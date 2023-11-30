package dispatcher

import "olive/engine/dispatch/common"

type Event struct {
	Type   common.EventTypeID
	Object any
}

func NewEvent(typ common.EventTypeID, object any) *Event {
	return &Event{
		Type:   typ,
		Object: object,
	}
}
