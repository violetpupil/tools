package dispatcher

import "olive/engine/dispatch/enum"

type Dispatcher interface {
	Dispatch(event *Event) error
	DispatcherType() enum.DispatcherTypeID
	DispatchTypes() []enum.EventTypeID
}
