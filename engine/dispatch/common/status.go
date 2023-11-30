package common

type StatusID = uint32

var Status = struct {
	Starting StatusID
	Pending  StatusID
	Running  StatusID
	Stopping StatusID
}{
	Starting: 0,
	Pending:  0,
	Running:  0,
	Stopping: 0,
}
