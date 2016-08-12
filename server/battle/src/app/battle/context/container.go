package context

import (
	battleAPI "app/typhenapi/type/submarine/battle"
)

type container struct {
	actors             []Actor
	actorsByID         map[int64]Actor
	submarinesByUserID map[int64]Actor
	users              []*User
}

func newContainer() *container {
	c := &container{
		actors:             make([]Actor, 0),
		actorsByID:         make(map[int64]Actor),
		submarinesByUserID: make(map[int64]Actor),
		users:              make([]*User, 0),
	}
	return c
}

func (c *container) addActor(actor Actor) {
	c.actorsByID[actor.ID()] = actor
	c.actors = append(c.actors, actor)
	if actor.Type() == battleAPI.ActorType_Submarine {
		c.submarinesByUserID[actor.User().ID] = actor
		c.users = append(c.users, actor.User())
	}
}

func (c *container) removeActor(rawActor Actor) Actor {
	actor := c.actorsByID[rawActor.ID()]
	if actor == nil {
		return nil
	}

	delete(c.actorsByID, actor.ID())
	actors := c.actors[:0]
	for _, a := range c.actors {
		if a != actor {
			actors = append(actors, a)
		}
	}
	c.actors = actors
	if actor.Type() == battleAPI.ActorType_Submarine {
		delete(c.submarinesByUserID, actor.User().ID)

		users := c.users[:0]
		for _, u := range c.users {
			if u != actor.User() {
				users = append(users, u)
			}
		}
		c.users = users
	}
	return actor
}
