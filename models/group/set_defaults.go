package group

import "fmt"

type SetDefaults struct {
	Group        `URIParam:"GroupId" argument:"composed" json:"-"`
	Cpu          int64  `json:",omitempty"`
	MemoryGb     int64  `json:",omitempty"`
	NetworkId    string `json:",omitempty"`
	PrimaryDns   string `json:",omitempty"`
	SecondaryDns string `json:",omitempty"`
	TemplateName string `json:",omitempty"`
}

func (s *SetDefaults) Validate() error {
	if err := s.Group.Validate(); err != nil {
		return err
	}

	if s.Cpu == 0 &&
		s.MemoryGb == 0 &&
		s.NetworkId == "" &&
		s.PrimaryDns == "" &&
		s.SecondaryDns == "" &&
		s.TemplateName == "" {
		return fmt.Errorf("A non-zero value for at least one of the --cpu, --memory-gb, --network-id, --primary-dns, --secondary-dns, and --template-name must be specified")
	}
	return nil
}
