package redfish

import "fmt"

type MediaType string

const mediaTypeCD MediaType = "CD"

func (s *System) InsertMedia(mt MediaType, image string) error {
	if len(s.computerSystem.Links.ManagedBy) == 0 {
		return fmt.Errorf("no managers in system")
	}
	m, err := s.client.GetManagerById(s.computerSystem.Links.ManagedBy[0].Id)
	if err != nil {
		return err
	}
	return m.InsertMedia(mt, image)
}
