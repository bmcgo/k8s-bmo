package redfish

import (
	"fmt"
)

type jsonManager struct {
	VirtualMedia OdataId `json:"VirtualMedia"`
}

type Manager struct {
	manager jsonManager
	client  *Client
}

type jsonVirtualMediaTarget struct {
	Target string `json:"Target"`
}

type jsonVirtualMediaActions struct {
	Eject  jsonVirtualMediaTarget `json:"#VirtualMedia.EjectMedia"`
	Insert jsonVirtualMediaTarget `json:"#VirtualMedia.InsertMedia"`
}

type jsonVirtualMediaList struct {
	Members []OdataId `json:"Members"`
}

type jsonVirtualMedia struct {
	MediaTypes []MediaType             `json:"MediaTypes"`
	Actions    jsonVirtualMediaActions `json:"Actions"`
}

type jsonVirtualMediaActionInsertPatch struct {
	Image    string `json:"Image"`
	Inserted bool   `json:"Inserted"`
}

func (s *Manager) InsertVirtualMedia(mt MediaType, image string) (err error) {

	jsonVMList := jsonVirtualMediaList{}
	err = s.client.Get(s.client.endpoint+s.manager.VirtualMedia.Id, &jsonVMList)
	if err != nil {
		return
	}
	for _, cvm := range jsonVMList.Members {
		virtualMedia := jsonVirtualMedia{}
		err = s.client.Get(s.client.endpoint+cvm.Id, &virtualMedia)

		for _, cmt := range virtualMedia.MediaTypes {
			if cmt == mt {
				err = s.client.Post(s.client.endpoint+virtualMedia.Actions.Insert.Target, jsonVirtualMediaActionInsertPatch{
					Inserted: true,
					Image:    image,
				})
				return
			}
		}
	}

	return fmt.Errorf("virtual media of type %s is not supported by the system", mt)
}
