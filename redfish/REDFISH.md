# Redfish client

    import "redfish"
    c := redfish.NewClient(redfish.ClientConfig{URL: "http://endpoint-1.example.net""})
    systems, _ := c.GetSystems()
    system := systems[0]
    system.InsertVirtualMedia(redfish.MediaTypeCD, "https://iso.example.net/image.iso")
    system.SetBootSourceOverride(redfish.BootSourceCD, redfish.BootSourceEnabledContinuous)
    system.Reset(redfish.ResetForceRestart)