package t3

import (
	"errors"
	"sync"
)

type DBPool struct {
	connections      []*Connection
	takenConnections map[*Connection]struct{}
	cond             sync.Cond
}

func NewDBPool(size int) *DBPool {
	pool := DBPool{
		connections:      []*Connection{},
		takenConnections: map[*Connection]struct{}{},
		cond:             *sync.NewCond(&sync.Mutex{}),
	}

	for connID := range size {
		pool.connections = append(pool.connections, createAMockConnection(connID))
	}
	return &pool
}

func (p *DBPool) Get() (*Connection, error) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()
	if len(p.takenConnections) == len(p.connections) {
		p.cond.Wait()
	}

	// здесь явно требуется какой-нибудь рефактор...
	for _, connptr := range p.connections {
		_, taken := p.takenConnections[connptr]
		if !taken {
			p.takenConnections[connptr] = struct{}{}

			return connptr, nil
		}
	}
	return nil, errors.New("i should not be here...")
}

func (p *DBPool) Release(connptrptr *Connection) {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	delete(p.takenConnections, connptrptr)
	p.cond.Signal()
}
