package task

// task type
const (
	CompressTaskType = iota
	DecompressTaskType
	TransferTaskType
	ImportTaskType
	RecycleTaskType
)

// task status
const (
	Queued = iota
	Processing
	Error
	Canceled
	Complete
)

// task progress
const (
	PendingProgress = iota
	CompressingProgress
	DecompressingProgress
	DownloadingProgress
	TransferringProgress
	ListingProgress
	InsertingProgress
)

type Job interface {
	Type() int
	Creator() uint
	Props() string
	Model() *Task //  对应的数据库模型
	SetStatus(status int)
	Do()
	SetError(*JobError)
	GetError() *JobError
}

type JobError struct {
	Msg string `json:"msg,omitempty"`
	Err string `json:"error,omitempty"`
}
