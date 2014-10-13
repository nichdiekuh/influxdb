package broker

import "errors"

var (
	// ErrPathRequired is returned when opening a broker without a path.
	ErrPathRequired = errors.New("path required")

	// ErrClosed is returned when closing a broker that's already closed.
	ErrClosed = errors.New("broker already closed")

	// ErrSubscribed is returned when a stream is already subscribed to a topic.
	ErrSubscribed = errors.New("already subscribed")

	// errReplicaUnavailable is returned when writing bytes to a replica when
	// there is no writer attached to the replica.
	errReplicaUnavailable = errors.New("replica unavailable")
)