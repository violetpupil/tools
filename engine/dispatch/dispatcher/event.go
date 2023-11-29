package dispatcher

import "olive/engine/dispatch/common"

type Event struct {
	Type   common.EventTypeID
	Object any
}
