package host

import (
	"github.com/Ullaakut/nmap/v2"
	"github.com/linkinyy/CycleScan/pkg/logger"
)

type Target struct {
	Ip     string
	Url    string
	Status bool
	Ports  []Port
	Os     []string
}

type Port struct {
	Id        uint16
	Service   string
	Protocol  string
	State     string
	Product   string
	Version   string
	ExtraInfo string
}

func (t *Target) Scan(ports []string) {
	t.ScanWithNmap(ports)
}

func (t *Target) ScanWithNmap(ports []string) {
	options := []nmap.Option{
		nmap.WithTargets(t.Ip),
		nmap.WithDisabledDNSResolution(),
		nmap.WithReason(),
		nmap.WithServiceInfo(),
		nmap.WithScripts("smb-os-discovery"),
		nmap.WithOSScanGuess(),
		nmap.WithOSDetection(),
	}
	if ports != nil {
		options = append(options, nmap.WithPorts(ports...))
	}
	scanner, err := nmap.NewScanner(options...)
	if err != nil {
		logger.Error(err)
		return
	}
	result, warnings, err := scanner.Run()
	if err != nil {
		logger.Error(err)
		return
	}
	if warnings != nil {
		logger.Warn(warnings)
	}
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		t.Status = true

		for _, match := range host.OS.Matches {
			t.Os = append(t.Os, match.Name)
		}

		for _, port := range host.Ports {
			p := Port{
				Id:        port.ID,
				Service:   port.Service.Name,
				Product:   port.Service.Product,
				Version:   port.Service.Version,
				ExtraInfo: port.Service.ExtraInfo,
				Protocol:  port.Protocol,
				State:     port.State.State,
			}
			t.Ports = append(t.Ports, p)
		}
	}
}

func (t *Target) IsAlive() bool {
	return t.Status
}

func (t *Target) OpenPorts() []Port {
	return t.Ports
}

//
//func checkPortsService(ports []int) ([]Id, error) {
//}
