package ipmitool

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const chassisStatusOutput = `System Power         : on
Power Overload       : false
Power Interlock      : inactive
Main Power Fault     : false
Power Control Fault  : false
Power Restore Policy : always-off
Last Power Event     :
Chassis Intrusion    : inactive
Front-Panel Lockout  : inactive
Drive Fault          : false
Cooling/Fan Fault    : false`

const bmcGUIDOutput = `Running Get PICMG Properties my_addr 0x20, transit 0, target 0x20
Error response 0xc1 from Get PICMG Properities
Running Get VSO Capabilities my_addr 0x20, transit 0, target 0x20
Invalid completion code received: Invalid command
Discovered IPMB address 0x0
System GUID  : 13a68dcb-e9cc-4d45-a155-01acdf360003
Timestamp    : 06/12/1980 18:26:19`

func TestIpmiTool_GetChassisStatus(t *testing.T) {
	it := New("1.2.3.4", 1234, "admin", "secret")
	it.execAndGetCombinedOutputFunc = func(args ...string) (string, error) {
		return chassisStatusOutput, nil
	}
	status, err := it.GetChassisStatus()
	require.NoError(t, err)
	assert.Equal(t, ChassisStatusPowerOn, status.SystemPower)
	assert.Equal(t, ChassisStatusDriveFaultFalse, status.DriveFault)
}

func TestIpmiTool_GetBMCGUID(t *testing.T) {
	it := New("1.2.3.4", 1234, "admin", "secret")
	it.execAndGetCombinedOutputFunc = func(args ...string) (string, error) {
		return bmcGUIDOutput, nil
	}
	guid, err := it.GetBMCGUID()
	require.NoError(t, err)
	assert.Equal(t, "13a68dcb-e9cc-4d45-a155-01acdf360003", guid)
}

/*
chassis status sample output by virtualbmc
ipmitool version 1.8.18
vbmc 2.2.2
# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password chassis status
System Power         : on
Power Overload       : false
Power Interlock      : inactive
Main Power Fault     : false
Power Control Fault  : false
Power Restore Policy : always-off
Last Power Event     :
Chassis Intrusion    : inactive
Front-Panel Lockout  : inactive
Drive Fault          : false
Cooling/Fan Fault    : false

chassis bootparam sample output by virtualbmc
ipmitool version 1.8.18
vbmc 2.2.2
# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password chassis bootparam get 5
Boot parameter version: 1
Boot parameter 5 is valid/unlocked
Boot parameter data: 8014000000
 Boot Flags :
   - Boot Flag Valid
   - Options apply to only next boot
   - BIOS PC Compatible (legacy) boot
   - Boot Device Selector : Force Boot from CD/DVD
   - Console Redirection control : System Default
   - BIOS verbosity : Console redirection occurs per BIOS configuration setting (default)
   - BIOS Mux Control Override : BIOS uses recommended setting of the mux at the end of POST

# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password chassis bootdev pxe
Set Boot Device to pxe
root@bs:~# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password chassis bootparam get 5
Boot parameter version: 1
Boot parameter 5 is valid/unlocked
Boot parameter data: 8004000000
 Boot Flags :
   - Boot Flag Valid
   - Options apply to only next boot
   - BIOS PC Compatible (legacy) boot
   - Boot Device Selector : Force PXE
   - Console Redirection control : System Default
   - BIOS verbosity : Console redirection occurs per BIOS configuration setting (default)
   - BIOS Mux Control Override : BIOS uses recommended setting of the mux at the end of POST
# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password chassis bootdev pxe
Set Boot Device to pxe
# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password chassis bootdev disk
Set Boot Device to disk
# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password chassis bootparam get 5
Boot parameter version: 1
Boot parameter 5 is valid/unlocked
Boot parameter data: 8008000000
 Boot Flags :
   - Boot Flag Valid
   - Options apply to only next boot
   - BIOS PC Compatible (legacy) boot
   - Boot Device Selector : Force Boot from default Hard-Drive
   - Console Redirection control : System Default
   - BIOS verbosity : Console redirection occurs per BIOS configuration setting (default)
   - BIOS Mux Control Override : BIOS uses recommended setting of the mux at the end of POST

sample ipmitool bmc guid (custom patched virtualbmc https://storyboard.openstack.org/#!/story/2010241)
ipmitool version 1.8.18
# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password -v bmc guid
Running Get PICMG Properties my_addr 0x20, transit 0, target 0x20
Error response 0xc1 from Get PICMG Properities
Running Get VSO Capabilities my_addr 0x20, transit 0, target 0x20
Invalid completion code received: Invalid command
Discovered IPMB address 0x0
System GUID  : 13a68dcb-e9cc-4d45-a155-01acdf360003
Timestamp    : 06/12/1980 18:26:19

# ipmitool -I lanplus -H 10.7.0.1 -p 3133 -U admin -P password -v raw 0x06 0x37
Running Get PICMG Properties my_addr 0x20, transit 0, target 0x20
Error response 0xc1 from Get PICMG Properities
Running Get VSO Capabilities my_addr 0x20, transit 0, target 0x20
Invalid completion code received: Invalid command
Discovered IPMB address 0x0
RAW REQ (channel=0x0 netfn=0x6 lun=0x0 cmd=0x37 data_len=0)
RAW RSP (16 bytes)
 cb 8d a6 13 cc e9 45 4d a1 55 01 ac df 36 00 03
*/
