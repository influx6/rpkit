package notifications

// AppEvent defines a struct to contain a event which occurs to be delivered to
// a giving AppNotification instance.
//
//@notification:event
type AppEvent struct {
	UUID  string
	Event interface{}
}

// Deliver delivers the giving value has a AppEvent.
func (sn *AppEventNotification) Deliver(uuid string, item interface{}) {
	sn.Handle(AppEvent{UUID: uuid, Event: item})
}

// AppNotification defines a structure which provides a local notification
// framework for the pubsub.
func AppNotification(uid string) *AppEventNotification {
	app := NewAppEventNotificationWith(func(ev AppEvent) bool {
		return ev.UUID == uid
	})

	Subscribe(app)

	return app
}
