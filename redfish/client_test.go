package redfish

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	var err error
	var systems []System
	c := NewClient(ClientConfig{URL: "http://localhost:8000"})
	systems, err = c.GetSystems()
	require.NoError(t, err)
	for _, s := range systems {
		if s.Name == "live-iso" {
			err = s.Reset(resetTypeForceOff)
			require.NoError(t, err)
			err = s.InsertMedia(mediaTypeCD, "http://localhost/media.iso")
			require.NoError(t, err)
			err = s.SetBootSourceOverride(BootSourceCD, BootSourceEnabledContinuous)
			require.NoError(t, err)
			//err = s.SetBootSourceOverride(BootSourcePXE, BootSourceEnabledContinuous)
			err = s.Reset(resetTypeForceOn)
			require.NoError(t, err)
		}
		log.Println(s.String())
	}
}
