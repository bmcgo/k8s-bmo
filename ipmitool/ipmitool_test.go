package ipmitool

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

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
*/

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

/*
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
*/
