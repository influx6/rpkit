package notifications

import (
	"sync"
)

// Remover defines an interface which exposes a remove method.
type Remover interface {
	Remove()
	Add(func())
}

// dispatch provides a default dispatcher for listening to events.
var dispatch = New()

// Unsubscribe adds a new listener to the dispatcher.
func Unsubscribe(dist EventDistributor) {
	dispatch.UnNotify(dist)
}

// Subscribe adds a new listener to the dispatcher.
func Subscribe(dist EventDistributor) {
	dispatch.Notify(dist)
}

// SubscribeWithRemover adds a new listener to the dispatcher and returns a common.Remover .
func SubscribeWithRemover(dist EventDistributor) Remover {
	dispatch.Notify(dist)

	return listenerRemover{
		root:    dispatch,
		handler: dist,
	}
}

// ListenerRemover defines a struct which implements the common.Remover interface.
type listenerRemover struct {
	handler EventDistributor
	root    *Notifications
	fn      []func()
}

// Adds adds a callback to be called when Remove is called.
func (l listenerRemover) Add(fn func()) {
	l.fn = append(l.fn, fn)
}

// Remove implements the common.Remover.
func (l listenerRemover) Remove() {
	l.root.UnNotify(l.handler)

	for _, fx := range l.fn {
		fx()
	}

	l.fn = nil
}

// Dispatch emits a event into the dispatch callback listeners.
func Dispatch(q interface{}) {
	dispatch.Handle(q)
}

// EventDistributor defines a interface that exposes a single method which
// will process a provided event received.
type EventDistributor interface {
	Handle(interface{})
}

// Notifications defines a central delivery pipe where all types of event notifications
// will pass through to be delivered to all EventDistributor listening.
type Notifications struct {
	ml       sync.Mutex
	sources  []EventDistributor
	register map[EventDistributor]int
}

// New returns a new instance of a Notification primitive.
func New() *Notifications {
	var nl Notifications
	nl.register = make(map[EventDistributor]int, 0)
	return &nl
}

// UnNotify removes the giving distributor from the notification system.
func (n *Notifications) UnNotify(source EventDistributor) {
	n.do(func() {
		index, ok := n.register[source]
		if !ok {
			return
		}

		n.sources[index] = nil
		n.sources = append(n.sources[:index], n.sources[index+1:]...)
	})
}

// Notify adds a giving EventDistributor into the notifications list.
func (n *Notifications) Notify(source EventDistributor) {
	n.do(func() {
		n.register[source] = len(n.sources)
		n.sources = append(n.sources, source)
	})
}

// Handle will publish giving type to all internal EventDistributor who are
// expected to convert the needed interface{} into expected type for consumption
// for their internal state or operations.
func (n *Notifications) Handle(item interface{}) {
	n.do(func() {
		for _, source := range n.sources {
			if source != nil {
				source.Handle(item)
			}
		}
	})
}

// do performs the needed function call guarded by a mutex call block.
func (n *Notifications) do(action func()) {
	if action == nil {
		return
	}

	n.ml.Lock()
	defer n.ml.Unlock()
	action()
}
