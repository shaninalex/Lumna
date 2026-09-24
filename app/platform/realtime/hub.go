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
	m := make(map[*Subscriber]any)
	m[subscriber] = struct{}{}
	h.subs[identityID] = m

	return subscriber, func() {
		h.mu.Lock()
		defer h.mu.Unlock()
		delete(h.subs, identityID)
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
	h.mu.RLock()
	defer h.mu.RUnlock()
	for k := range h.subs {
		delete(h.subs, k)
	}
}
