package event

const (
	// URLTask 从URL添加的任务
	URLTask = iota
	// TorrentTask 种子任务
	TorrentTask
)

const (
	// Ready 准备就绪
	Ready = iota
	// Downloading 下载中
	Downloading
	// Paused 暂停中
	Paused
	// Error 出错
	Error
	// Complete 完成
	Complete
	// Canceled 取消/停止
	Canceled
	// Unknown 未知状态
	Unknown
	// Seeding 做种中
	Seeding
)
