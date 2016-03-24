package db

type DeleteBackup struct {
	SubscriptionId string `json:"-" URIParam:"yes" valid:"required"`
	BackupId       string `json:"-" URIParam:"yes" valid:"required"`
}
