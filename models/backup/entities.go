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

type ServerPolicy struct {
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
