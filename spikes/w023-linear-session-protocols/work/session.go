package session

// Zero-allocation phantom typestate representations for compile-time channel protocols.

// ProtoStep phantom witnesses (zero runtime size).
type (
	PingState struct{}
	PongState struct{}
	Closed    struct{}
)

// SessionChan represents a compile-time tracked channel cursor.
// In the emitted Go code, this lowers directly to `chan T`.
type SessionChan[P any, T any] struct {
	ch chan T
}

func NewSessionChan[P any, T any](buffer int) SessionChan[P, T] {
	return SessionChan[P, T]{ch: make(chan T, buffer)}
}

// BareChannel exposes the underlying raw Go channel for identity emission.
func (s SessionChan[P, T]) BareChannel() chan T {
	return s.ch
}
