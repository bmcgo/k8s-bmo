package redfish

type ResetType string

const (
	ResetForceRestart ResetType = "ForceRestart"
	ResetForceOn      ResetType = "ForceOn"
	ResetForceOff     ResetType = "ForceOff"
)

type ResetBody struct {
	ResetType ResetType `json:"ResetType"`
}

func (s *System) Reset(rt ResetType) error {
	//TODO: check allowable values
	return s.client.Post(s.url+"/Actions/ComputerSystem.Reset", ResetBody{ResetType: rt})
}
