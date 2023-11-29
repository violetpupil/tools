package dispatcher

import "olive/engine/dispatch/enum"

type Event struct {
	Type   enum.EventTypeID
	Object any
}
