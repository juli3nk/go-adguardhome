package adguardhome

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Status struct {
	Version                    string   `json:"domain"`
	Language                   string   `json:"answer"`
	DnsAddresses               []string `json:"dns_addresses"`
	DnsPort                    int      `json:"dns_port"`
	HttpPort                   string   `json:"http_port"`
	ProtectionDisabledDuration int      `json:"protection_disabled_duration"`
	ProtectionEnabled          bool     `json:"protection_enabled"`
	DhcpAvailable              bool     `json:"dhcp_available"`
	Running                    bool     `json:"running"`
}

func (c *DnsCfg) Status() (*Status, error) {
	url := c.url + "/control/status"
	mimeTypeJson := "application/json"

	tr := &http.Transport{
		IdleConnTimeout: time.Second * 5,
	}
	client := &http.Client{
		Transport: tr,
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err.Error())
	}
	req.Header.Add("Accept", mimeTypeJson)
	req.SetBasicAuth(c.username, c.password)

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d. Please contact an administrator if the problem persists", response.StatusCode)
	}

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err.Error())
	}

	var status Status
	if err := json.Unmarshal(respBody, &status); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %s", err.Error())
	}

	return &status, nil
}
