package common

type DispatcherTypeID uint32

var DispatcherType = struct {
	Monitor  DispatcherTypeID
	Recorder DispatcherTypeID
}{
	Monitor:  100,
	Recorder: 200,
}

func (dt DispatcherTypeID) String() string {
	switch dt {
	case DispatcherType.Monitor:
		return "monitor"
	case DispatcherType.Recorder:
		return "recorder"
	}
	return "undefined"
}
