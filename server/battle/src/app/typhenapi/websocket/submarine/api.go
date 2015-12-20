// This file was generated by typhen-api

package submarine

import (
	"app/typhenapi/core"
	_submarine_battle "app/typhenapi/websocket/submarine/battle"
)

// WebSocketAPI sends messages, and dispatches message events.
type WebSocketAPI struct {
	session      typhenapi.Session
	serializer   *typhenapi.Serializer
	errorHandler func([]byte, error)
	Battle       *_submarine_battle.WebSocketAPI
}

// New creates a WebSocketAPI.
func New(session typhenapi.Session, serializer *typhenapi.Serializer, errorHandler func([]byte, error)) *WebSocketAPI {
	api := &WebSocketAPI{}
	api.session = session
	api.serializer = serializer
	api.errorHandler = errorHandler
	api.Battle = _submarine_battle.New(session, serializer, errorHandler)
	return api
}

// DispatchMessageEvent dispatches a binary message.
func (api *WebSocketAPI) DispatchMessageEvent(data []byte) {
	api.Battle.DispatchMessageEvent(data)

}
