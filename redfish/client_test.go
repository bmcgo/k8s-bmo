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
			err = s.Reset(ResetForceOff)
			require.NoError(t, err)
			err = s.InsertVirtualMedia(MediaTypeCD, "http://localhost/media.iso")
			require.NoError(t, err)
			err = s.SetBootSourceOverride(BootSourceCD, BootSourceEnabledContinuous)
			require.NoError(t, err)
			//err = s.SetBootSourceOverride(BootSourcePXE, BootSourceEnabledContinuous)
			err = s.Reset(ResetForceOn)
			require.NoError(t, err)
		}
		log.Println(s.String())
	}
}
