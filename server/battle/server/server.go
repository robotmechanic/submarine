package server

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	webapi "github.com/shiwano/submarine/server/battle/lib/typhenapi/web/submarine"
	"github.com/shiwano/submarine/server/battle/server/config"
	"github.com/shiwano/submarine/server/battle/server/logger"
)

// Server represents a battle server.
type Server struct {
	*gin.Engine
	logWriter   *io.PipeWriter
	roomManager *RoomManager
	webAPI      *webapi.WebAPI
}

// NewServer creates a Server.
func NewServer() *Server {
	server := &Server{
		Engine:      gin.New(),
		logWriter:   logger.Log.Writer(),
		roomManager: newRoomManager(),
		webAPI:      NewWebAPI("http://localhost:3000"),
	}
	server.Use(gin.Recovery(), gin.LoggerWithWriter(server.logWriter))

	server.GET("/rooms/:id", server.roomsGET)

	return server
}

func (s *Server) roomsGET(c *gin.Context) {
	roomID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || roomID <= 0 {
		logger.Log.Error(err)
		c.String(http.StatusForbidden, "Invalid room id")
		return
	}

	room, err := s.roomManager.fetchRoom(roomID)
	if err != nil {
		logger.Log.Error(err)
		c.String(http.StatusForbidden, "Failed getting the room")
		return
	}

	res, err := s.webAPI.Battle.FindRoomMember(c.Query("room_key"))
	if err != nil {
		logger.Log.Error(err)
		c.String(http.StatusInternalServerError, "Failed authenticating the room key")
		return
	}
	if res.RoomMember == nil {
		c.String(http.StatusForbidden, "Invalid room key")
		return
	}

	session := newSession(res.RoomMember, roomID)
	if err := session.Connect(c.Writer, c.Request); err != nil {
		logger.Log.Error(err)
		c.String(http.StatusForbidden, "Failed upgrading the connection to the Web Socket Protocol")
		return
	}

	if room.isClosed.IsSet() {
		logger.Log.Infof("%v is already closed", room)
		session.close()
		return
	}

	room.joinCh <- session
	logger.Log.Infof("%v is created", session)
}

func init() {
	switch config.Env {
	case "test":
		gin.SetMode(gin.TestMode)
	case "development":
		gin.SetMode(gin.DebugMode)
	}
}