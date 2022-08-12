package redfish

type BootSourceOverrideTarget string
type BootSourceOverrideEnabled string

const (
	BootSourceCD                BootSourceOverrideTarget  = "Cd"
	BootSourcePXE               BootSourceOverrideTarget  = "Pxe"
	BootSourceEnabledContinuous BootSourceOverrideEnabled = "Continuous"
	BootSourceEnabledDisabled   BootSourceOverrideEnabled = "Disabled"
)

type patchBootSourceOverride struct {
	Boot bootStructOverrideBoot `json:"Boot"`
}

type bootStructOverrideBoot struct {
	BootSourceOverrideTarget  BootSourceOverrideTarget  `json:"BootSourceOverrideTarget"`
	BootSourceOverrideEnabled BootSourceOverrideEnabled `json:"BootSourceOverrideEnabled"`
}

func (s *System) SetBootSourceOverride(target BootSourceOverrideTarget, enabled BootSourceOverrideEnabled) error {
	//TODO: check allowable values
	return s.client.Patch(s.url,
		patchBootSourceOverride{
			Boot: bootStructOverrideBoot{
				BootSourceOverrideTarget:  target,
				BootSourceOverrideEnabled: enabled,
			},
		})
}
