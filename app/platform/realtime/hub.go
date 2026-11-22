package realtime

import "sync"

type Envelope struct {
	Name    string
	Payload any
}

type Subscriber struct {
	ch chan Envelope
}

func (s *Subscriber) C() <-chan Envelope {
	return s.ch
}

type Hub struct {
	mu sync.RWMutex

	// identityID - connections
	subs map[int]map[*Subscriber]any
	buf  int
}

func NewHub(buf int) *Hub {
	return &Hub{
		mu:   sync.RWMutex{},
		subs: make(map[int]map[*Subscriber]any),
		buf:  buf,
	}
}

func (h *Hub) Attach(identityID int) (*Subscriber, func()) {
	h.mu.Lock()
	defer h.mu.Unlock()

	subscriber := &Subscriber{ch: make(chan Envelope, h.buf)}
	if h.subs[identityID] == nil {
		h.subs[identityID] = make(map[*Subscriber]any)
	}
	h.subs[identityID][subscriber] = struct{}{}

	return subscriber, func() {
		h.mu.Lock()
		defer h.mu.Unlock()
		delete(h.subs[identityID], subscriber)
		if len(h.subs[identityID]) == 0 {
			delete(h.subs, identityID)
		}
	}
}

func (h *Hub) Send(identityIDs []int, e Envelope) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for _, id := range identityIDs {
		for s := range h.subs[id] {
			select {
			case s.ch <- e:
			default:
			}
		}
	}
}

// Close closes all subscriber channels on shutdown
func (h *Hub) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()
	for id, subs := range h.subs {
		for s := range subs {
			close(s.ch)
		}
		delete(h.subs, id)
	}
}
