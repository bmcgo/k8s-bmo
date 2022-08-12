package redfish

import "fmt"

type System struct {
	Name string
	UUID string
	Boot bootStructOverrideBoot

	computerSystem computerSystem
	client         *Client
	url            string
}

func (s *System) String() string {
	return fmt.Sprintf("System<%s [%s] boot:%s>", s.Name, s.UUID, s.computerSystem.Boot.BootSourceOverrideTarget)
}

type systemsCollectionMember struct {
	Id string `json:"@odata.id"`
}

type systemsCollection struct {
	Members []systemsCollectionMember `json:"Members"`
}

type ComputerSystemBoot struct {
	BootSourceOverrideEnabled               string   `json:"BootSourceOverrideEnabled"`
	BootSourceOverrideTarget                string   `json:"BootSourceOverrideTarget"`
	BootSourceOverrideTargetAllowableValues []string `json:"BootSourceOverrideTarget@Redfish.AllowableValues"`
}

type ComputerSystemLinks struct {
	Chassis   []OdataId `json:"Chassis"`
	ManagedBy []OdataId `json:"ManagedBy"`
}

type OdataId struct {
	Id string `json:"@odata.id"`
}

type computerSystem struct {
	Id         string              `json:"Id"`
	Name       string              `json:"Name"`
	UUID       string              `json:"UUID"`
	PowerState string              `json:"PowerState"`
	Boot       ComputerSystemBoot  `json:"Boot"`
	Links      ComputerSystemLinks `json:"Links"`
}

func (s *Client) GetSystems() ([]System, error) {
	var err error
	var cs *System
	var member systemsCollectionMember
	var systems []System

	collection := systemsCollection{}
	err = s.Get(s.urlSystems, &collection)
	if err != nil {
		return nil, err
	}
	for _, member = range collection.Members {
		cs, err = s.GetSystem(s.endpoint + member.Id)
		if err != nil {
			return nil, err
		}
		systems = append(systems, *cs)
	}
	return systems, nil
}

func (s *Client) GetSystem(url string) (*System, error) {
	cs := &computerSystem{}
	err := s.Get(url, cs)
	if err != nil {
		return nil, err
	}
	return &System{
		Name:           cs.Name,
		UUID:           cs.UUID,
		url:            url,
		client:         s,
		computerSystem: *cs,
	}, nil
}
