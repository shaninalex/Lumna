package realtime

import "testing"

// A second connection of the same identity must keep receiving after the
// first one detaches (reconnects, several tabs, React StrictMode remounts).
func TestHub_DetachKeepsOtherConnections(t *testing.T) {
	h := NewHub(1)
	_, detachA := h.Attach(1)
	b, detachB := h.Attach(1)
	defer detachB()

	detachA()
	h.Send([]int{1}, Envelope{Name: "x"})

	select {
	case e := <-b.C():
		if e.Name != "x" {
			t.Fatalf("unexpected envelope: %+v", e)
		}
	default:
		t.Fatal("second subscriber got nothing after the first detached")
	}
}

func TestHub_SendsToEveryConnection(t *testing.T) {
	h := NewHub(1)
	a, _ := h.Attach(1)
	b, _ := h.Attach(1)

	h.Send([]int{1}, Envelope{Name: "x"})

	for i, s := range []*Subscriber{a, b} {
		select {
		case <-s.C():
		default:
			t.Fatalf("subscriber %d got nothing", i)
		}
	}
}
