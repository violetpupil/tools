package dispatcher

import "olive/engine/dispatch/common"

type Dispatcher interface {
	Dispatch(event *Event) error
	DispatcherType() common.DispatcherTypeID
	DispatchTypes() []common.EventTypeID
}
