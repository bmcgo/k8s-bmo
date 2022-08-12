package redfish

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestNewClient(t *testing.T) {
	var err error
	var systems []System
	l, _ := logr.FromContext(context.Background())
	c := NewClient(ClientConfig{URL: "http://localhost:8000"}, l)
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
