package battle

import (
	"time"
)

// Battle represents a battle.
type Battle struct {
	Gateway   *Gateway
	context   *Context
	createdAt time.Time
	startedAt time.Time
	timeLimit time.Duration
	IsStarted bool
}

// New creates a new battle.
func New(timeLimit time.Duration) *Battle {
	return &Battle{
		Gateway:   newGateway(),
		context:   newContext(),
		createdAt: time.Now(),
		timeLimit: timeLimit,
	}
}

// CreateSubmarineUnlessExists creates the user's submarine unless it exists.
func (b *Battle) CreateSubmarineUnlessExists(userID int64) {
	if s := b.context.container.getSubmarineByUserID(userID); s == nil {
		b.context.container.createSubmarine(userID)
	}
}

// Start the battle.
func (b *Battle) Start() {
	if !b.IsStarted {
		b.IsStarted = true
		go b.run()
	}
}

func (b *Battle) run() {
	ticker := time.Tick(time.Second / 30)
	b.startedAt = time.Now()
	b.Gateway.start(b.startedAt)

loop:
	for {
		select {
		case now := <-ticker:
			b.context.now = now
			if b.context.now.After(b.startedAt.Add(b.timeLimit)) {
				break loop
			}
		case <-b.Gateway.Close:
			break loop
		}
	}

	b.Gateway.finish(true, b.context.now)
}