package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const collectTimeout = 5 * time.Second

var (
	infoLength    = 11
	infoRegexp    = regexp.MustCompile(`\|\s*([0-9]*)\s*(\S*)\s*\|\s*(\S*)\s*\|\s*(\S*)\s*(\S*|\S*\s/\s\S*)\s*\|\s\|\s*([0-9]*)\s*(\S*)\s*\|\s*(\S*)\s*\|\s*(\S*)\s*(\S*|\S*\s/\s\S*)\s*\|`)
	versionRegexp = regexp.MustCompile(`.*Version: (?P<version>\S*)\s*`)
	memRegexp     = regexp.MustCompile(`([0-9]+)\s/\s([0-9]+)`)
)

type NpuCollector struct{}

type Metrics struct {
	Version string
	Devices []*NpuDevice
}

type NpuDevice struct {
	NpuID         string
	Name          string
	Health        string
	PowerUsage    string
	Temperature   string
	ChipID        string
	Device        string
	BusID         string
	AICore        string
	MemoryUsageMB string
}

func NewNpuCollector() *NpuCollector {
	return &NpuCollector{}
}

// TODO: read info from npu driver kernel api by importing dsmi_common_interface.h
func (nc *NpuCollector) GetNPUInfo() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), collectTimeout)
	defer cancel()

	return ExecBashShellWithCtx(ctx, "npu-smi info")
}

func collectMetrics() (*Metrics, error) {
	cl := NewNpuCollector()
	result, err := cl.GetNPUInfo()
	if err != nil {
		return nil, err
	}

	return &Metrics{
		Version: parseNpuVersion(result),
		Devices: parseNpuDevices(result),
	}, nil
}

func parseNpuVersion(str string) string {
	var version string

	ver := versionRegexp.FindStringSubmatch(str)
	if len(ver) == 2 {
		version = ver[1]
	}
	return version
}

func parseNpuDevices(str string) []*NpuDevice {
	devices := make([]*NpuDevice, 0)
	allInfos := infoRegexp.FindAllStringSubmatch(str, -1)
	for _, info := range allInfos {
		if len(info) != infoLength {
			continue
		}
		allStr := ""
		device := &NpuDevice{}
		params := []struct {
			attr *string
		}{
			{&allStr},
			{&device.NpuID},
			{&device.Name},
			{&device.Health},
			{&device.PowerUsage},
			{&device.Temperature},
			{&device.ChipID},
			{&device.Device},
			{&device.BusID},
			{&device.AICore},
			{&device.MemoryUsageMB},
		}
		for k, v := range info {
			*params[k].attr = v
		}
		devices = append(devices, device)
	}
	return devices
}

func parseValueFloat(val string) float64 {
	valFloat, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Printf("parse value float error:%v, value:%s", err, val)
	}
	return valFloat
}

func parseMemoryByte(s string) (float64, float64) {
	var usage, total float64
	mem := memRegexp.FindStringSubmatch(s)
	if len(mem) == 3 {
		usage = parseValueFloat(mem[1]) * MiB
		total = parseValueFloat(mem[2]) * MiB
	}

	return usage, total
}
