package db

type ActionLogEntry struct {
	Timestamp int64
	Message   string
	Details   string
	User      string
}

type Promotion struct {
	Id   int64
	Code string
}

type DataCenterInfo struct {
	DataCenter   string
	FriendlyName string
	Active       bool
}

type Destination struct {
	Id              int64
	DestinationType string
	Location        string
	Notifications   []Notification
	Verified        bool
}

type Notification struct {
	Id               int64
	NotificationType string
}
