package db

import "fmt"

type Update struct {
	SubscriptionId      string         `json:"-" URIParam:"yes" valid:"required"`
	MachineConfig       MachineConfig  `json:"-"`
	MachineConfigValue  *MachineConfig `json:"machineConfig,omitempty" argument:"ignore"`
	BackupRetentionDays *int64         `json:"backupRetentionDays,omitempty"`
	BackupTime          BackupTime     `json:"-"`
	BackupTimeValue     *BackupTime    `json:"backupTime,omitempty" argument:"ignore"`
}

func (u *Update) Validate() error {
	emptyConfig := MachineConfig{}
	emptyBackup := BackupTime{}
	if (u.BackupRetentionDays == nil || *u.BackupRetentionDays == 0) && u.MachineConfig == emptyConfig && u.BackupTime == emptyBackup {
		return fmt.Errorf("At least one of --machine-config, --backup-retention-days, and --backup-time must be set and non empty")
	}

	if u.BackupTime != emptyBackup {
		b := u.BackupTime
		if b.Hour == nil || b.Minute == nil {
			return fmt.Errorf("Both hour and minute have to be specified")
		}
		u.BackupTimeValue = &b
	}

	if u.MachineConfig != emptyConfig {
		c := u.MachineConfig
		if c.Memory == nil || c.Cpu == nil || c.Storage == nil {
			return fmt.Errorf("All of the cpu, memory, and storage have to be specified")
		}
		u.MachineConfigValue = &c
	}
	return nil
}
