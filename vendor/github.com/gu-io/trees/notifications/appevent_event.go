package notifications

import "sync"

// AppEventSubscriber defines a interface that which is used to subscribe specifically for
// events  AppEvent type.
type AppEventSubscriber interface {
	Receive(AppEvent)
}

//=========================================================================================================

// AppEventHandler defines a structure type which implements the
// AppEventSubscriber interface and the EventDistributor interface.
type AppEventHandler struct {
	handle func(AppEvent)
}

// NewAppEventHandler returns a new instance of a AppEventHandler.
func NewAppEventHandler(fn func(AppEvent)) *AppEventHandler {
	return &AppEventHandler{
		handle: fn,
	}
}

// Receive takes the giving value and execute it against the underline handler.
func (sn *AppEventHandler) Receive(elem AppEvent) {
	sn.handle(elem)
}

// Handle takes the giving value and asserts the expected value to match the
// AppEvent type then passes it to the Receive method.
func (sn *AppEventHandler) Handle(receive interface{}) {
	if elem, ok := receive.(AppEvent); ok {
		sn.Receive(elem)
	}
}

//=========================================================================================================

// AppEventNotification defines a structure type which must be used to
// receive AppEvent type has a event.
type AppEventNotification struct {
	sml        sync.Mutex
	subs       []AppEventSubscriber
	validation func(AppEvent) bool
	register   map[AppEventSubscriber]int
}

// NewAppEventNotificationWith returns a new instance of AppEventNotification.
func NewAppEventNotificationWith(validation func(AppEvent) bool) *AppEventNotification {
	var elem AppEventNotification

	elem.validation = validation
	elem.register = make(map[AppEventSubscriber]int, 0)

	return &elem
}

// NewAppEventNotification returns a new instance of NewAppEventNotification.
func NewAppEventNotification() *AppEventNotification {
	var elem AppEventNotification
	elem.register = make(map[AppEventSubscriber]int, 0)

	return &elem
}

// UnNotify removes the given subscriber from the notification's list if found from future events.
func (sn *AppEventNotification) UnNotify(sub AppEventSubscriber) {
	sn.do(func() {
		index, ok := sn.register[sub]
		if !ok {
			return
		}

		sn.subs = append(sn.subs[:index], sn.subs[index+1:]...)
	})
}

// Notify adds the given subscriber into the notification list and will await an update of
// a new event of the given AppEvent type.
func (sn *AppEventNotification) Notify(sub AppEventSubscriber) {
	sn.do(func() {
		sn.register[sub] = len(sn.subs)
		sn.subs = append(sn.subs, sub)
	})
}

// Handle takes the giving value and asserts the expected value to be of
// the type and pass on to it's underline subscribers else ignoring the event.
func (sn *AppEventNotification) Handle(elem interface{}) {
	if elemEvent, ok := elem.(AppEvent); ok {
		if sn.validation != nil && sn.validation(elemEvent) {
			sn.do(func() {
				for _, sub := range sn.subs {
					sub.Receive(elemEvent)
				}
			})

			return
		}

		sn.do(func() {
			for _, sub := range sn.subs {
				sub.Receive(elemEvent)
			}
		})
	}
}

// do performs action with the mutex locked and unlocked appropriately, ensuring safe
// concurrent access.
func (sn *AppEventNotification) do(fn func()) {
	if fn == nil {
		return
	}

	sn.sml.Lock()
	defer sn.sml.Unlock()

	fn()
}
