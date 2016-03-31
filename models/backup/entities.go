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
