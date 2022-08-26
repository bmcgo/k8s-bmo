package ipmitool

import (
	"errors"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type ChassisStatusSystemPower string
type ChassisStatusDriveFault string
type ChassisPower string
type BootDev string

const BootDevPxe = BootDev("pxe")
const BootDevDisk = BootDev("disk")
const ChassisStatusPowerOn = ChassisStatusSystemPower("on")
const ChassisStatusPowerOff = ChassisStatusSystemPower("off")
const ChassisStatusDriveFaultTrue = ChassisStatusDriveFault("true")
const ChassisStatusDriveFaultFalse = ChassisStatusDriveFault("false")
const ChassisPowerOn = ChassisPower("on")
const ChassisPowerCycle = ChassisPower("cycle")
const ChassisPowerSoft = ChassisPower("soft")
const ChassisPowerOff = ChassisPower("off")

type ChassisStatus struct {
	SystemPower ChassisStatusSystemPower
	DriveFault  ChassisStatusDriveFault
}

func (i IpmiTool) Power(power ChassisPower) error {
	_, err := i.execAndGetCombinedOutputFunc("chassis", "power", string(power))
	return err
}

func (i IpmiTool) SetBootDev(dev BootDev) error {
	_, err := i.execAndGetCombinedOutputFunc("chassis", "bootdev", string(dev))
	return err
}

func (i IpmiTool) GetBootDev() (BootDev, error) {
	output, err := i.execAndGetCombinedOutputFunc("chassis", "bootparam", "get", "5")
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(output, "\n") {
		bits := strings.Split(line, ":")
		if len(bits) == 2 {
			name := strings.Trim(bits[0], " -")
			if name == "Boot Device Selector" {
				if strings.Contains(bits[1], "PXE") {
					return BootDevPxe, nil
				}
				return BootDev(strings.Trim(bits[1], " ")), nil
			}
		}
	}
	return "Unknown", nil
}

func (i IpmiTool) GetChassisStatus() (status ChassisStatus, err error) {
	output, err := i.execAndGetCombinedOutputFunc("chassis", "status")
	if err != nil {
		return
	}
	for _, line := range strings.Split(output, "\n") {
		bits := strings.Split(line, ":")
		if len(bits) == 2 {
			value := strings.Trim(bits[1], " ")
			switch strings.Trim(bits[0], " ") {
			case "System Power":
				status.SystemPower = ChassisStatusSystemPower(value)
			case "Drive Fault":
				status.DriveFault = ChassisStatusDriveFault(value)
			}
		}
	}
	return
}

func (i IpmiTool) GetBMCGUID() (string, error) {
	output, err := i.execAndGetCombinedOutputFunc("bmc", "guid")
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(output, "\n") {
		bits := strings.Split(line, ":")
		if len(bits) == 2 {
			value := strings.Trim(bits[1], " ")
			if strings.Trim(bits[0], " ") == "System GUID" {
				return value, nil
			}
		}
	}
	log.Println("ipmitool output:")
	log.Println(output)
	return "", errors.New("failed to find System GUID in ipmitool output")
}

func (i IpmiTool) execAndGetCombinedOutput(args ...string) (string, error) {
	cmd := exec.Command(i.cmd, append(i.args, args...)...)
	out, err := cmd.CombinedOutput()
	//TODO: return own error with both original exec.Command error and combined output included
	return string(out), err
}

type IpmiTool struct {
	cmd                          string
	args                         []string
	execAndGetCombinedOutputFunc func(...string) (string, error)
}

func New(host string, portInt int, username string, password string) IpmiTool {
	port := strconv.Itoa(portInt)
	it := IpmiTool{
		cmd:  "ipmitool",
		args: []string{"-I", "lanplus", "-H", host, "-p", port, "-U", username, "-P", password},
	}
	it.execAndGetCombinedOutputFunc = it.execAndGetCombinedOutput
	return it
}
