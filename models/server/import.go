package server

type Import struct {
	Name         string           `valid:"required"`
	Description  string           `json:"omitempty"`
	GroupId      string           `valid:"required"`
	PrimaryDns   string           `json:"omitempty"`
	SecondaryDns string           `json:"omitempty"`
	NetworkId    string           `json:"omitempty"`
	RootPassword string           `valid:"required"`
	Cpu          int64            `valid:"required"`
	MemoryGB     int64            `valid:"required"`
	Type         string           `valid:"required"`
	StorageType  string           `json:"omitempty"`
	CustomFields []CustomFieldDef `json:"omitempty"`
	OvfId        string           `valid:"required"`
	OvfOsType    string           `valid:"required"`
}