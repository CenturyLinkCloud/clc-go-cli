package backup

type AccountPolicy struct {
	ClcAccountAlias        string
	BackupIntervalHours    int64
	ExcludedDirectoryPaths []string
	Name                   string
	OsType                 string
	Paths                  []string
	PolicyId               string
	RetentionDays          int64
	Status                 string
}

type Region struct {
	Messages    []string
	Name        string
	RegionLabel string
}

type ServerPolicyDetails struct {
	AccountPolicyId     string
	AccountPolicyStatus string
	BackupIntervalHours int64
	BackupProvider      string
	ClcAccountAlias     string
	EligibleForBackup   bool
	Name                string
	OsType              string
	Paths               []string
	RetentionDays       int64
	ServerId            string
	ServerPolicyId      string
	ServerPolicyStatus  string
	StorageRegion       string
}

type ServerPolicy struct {
	AccountPolicyId  string
	ClcAccountAlias  string
	ExpirationDate   int64
	ServerId         string
	ServerPolicyId   string
	Status           string
	StorageAccountId string
	StorageRegion    string
	UnsubscribedDate int64
}

type StoredData struct {
	BytesStored     string
	GigabytesStored string
}

type RestoreDetails struct {
	BackupFinishedDate            string
	BackupStartedDate             string
	BytesFailedToTransfer         int64
	BytesInStorageForItemsRemoved int64
	BytesTransferredToStorage     int64
	FilesFailedTransferToStorage  int64
	FilesRemovedFromDisk          int64
	FilesTransferredToStorage     int64
	NumberOfProtectedFiles        int64
	PolicyId                      string
	RestorePointCreationStatus    string
	RestorePointId                string
	RetentionDays                 int64
	RetentionExpiredDate          string
	UnchangedBytesInStorage       int64
	UnchangedFilesNotTransferred  int64
}
