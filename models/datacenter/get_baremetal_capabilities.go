package datacenter

type GetBMCapReq struct {
	DataCenter string `valid:"required" URIParam:"yes"`
}

type GetBMCapRes struct {
	SKUs             []SKU             `json:"skus"`
	OperatingSystems []OperatingSystem `json:"operatingSystems"`
}

type SKU struct {
	ID           string    `json:"id"`
	HourlyRate   float32   `json:"hourlyRate"`
	Availability string    `json:"availability"`
	Memory       []Memory  `json:"memory"`
	Processor    Processor `json:"processor"`
	Storage      []Storage `json:"storage"`
}

type Memory struct {
	CapacityInGB int `json:"capacityGB"`
}

type Processor struct {
	Sockets        int    `json:"sockets"`
	CoresPerSocket int    `json:"coresPerSocket"`
	Description    string `json:"description"`
}

type Storage struct {
	Type         string `json:"type"`
	CapacityInGB int    `json:"capacityGB"`
	SpeedInRPM   int    `json:"speedRpm"`
}

type OperatingSystem struct {
	Type                string  `json:"type"`
	Description         string  `json:"description"`
	HourlyRatePerSocket float32 `json:"hourlyRatePerSocket"`
}
