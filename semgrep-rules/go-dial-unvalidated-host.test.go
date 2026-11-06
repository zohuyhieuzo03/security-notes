// Synthetic test fixtures for go-dial-unvalidated-host.yaml.
// These are invented examples illustrating a vulnerability CLASS, not code
// from, or details of, any specific project or report.
package example

import (
	"encoding/json"
	"net/http"
)

type DeviceConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type DeviceClient struct{}

func (c *DeviceClient) Dial(host string, port int, opts ...any) error {
	return nil
}

func guardPublicOnly(host string, port int) error {
	// pretend this resolves host, then checks the resolved IP against
	// loopback/private/link-local/metadata ranges.
	return nil
}

// A handler that lets a caller update a device's connection settings and
// immediately reconnects using the device's already-stored credentials,
// with no check on where it is now connecting to.
func UpdateDeviceHandlerVulnerable(w http.ResponseWriter, r *http.Request) {
	var cfg DeviceConfig
	json.NewDecoder(r.Body).Decode(&cfg)

	client := &DeviceClient{}
	// ruleid: go-dial-unvalidated-host
	client.Dial(cfg.Host, cfg.Port)
}

// Same shape, but the destination is validated before dialing.
func UpdateDeviceHandlerSafe(w http.ResponseWriter, r *http.Request) {
	var cfg DeviceConfig
	json.NewDecoder(r.Body).Decode(&cfg)

	if err := guardPublicOnly(cfg.Host, cfg.Port); err != nil {
		http.Error(w, "destination not allowed", http.StatusForbidden)
		return
	}

	client := &DeviceClient{}
	// ok: go-dial-unvalidated-host
	client.Dial(cfg.Host, cfg.Port)
}
