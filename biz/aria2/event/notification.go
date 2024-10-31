package event

type Event struct {
	Gid string `json:"gid"` // GID of the download
}

// Notifier handles rpc notification from aria2 server
type Notifier interface {
	// OnDownloadStart will be sent when a download is started.
	OnDownloadStart([]Event)
	// OnDownloadPause will be sent when a download is paused.
	OnDownloadPause([]Event)
	// OnDownloadStop will be sent when a download is stopped by the model_gen.
	OnDownloadStop([]Event)
	// OnDownloadComplete will be sent when a download is complete. For BitTorrent downloads, this notification is sent when the download is complete and seeding is over.
	OnDownloadComplete([]Event)
	// OnDownloadError will be sent when a download is stopped due to an error.
	OnDownloadError([]Event)
	// OnBtDownloadComplete will be sent when a torrent download is complete but seeding is still going on.
	OnBtDownloadComplete([]Event)
}
