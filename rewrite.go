package adguardhome

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Record struct {
	Domain string `json:"domain"`
	Answer string `json:"answer"`
}

func (c *DnsCfg) RewriteList() ([]Record, error) {
	url := c.url + "/control/rewrite/list"
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

	var records []Record
	if err := json.Unmarshal(respBody, &records); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %s", err.Error())
	}

	return records, nil
}

func (c *DnsCfg) RewriteAdd(payload *Record) error {
	url := c.url + "/control/rewrite/add"
	mimeTypeJson := "application/json"

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	tr := &http.Transport{
		IdleConnTimeout: time.Second * 5,
	}
	client := &http.Client{
		Transport: tr,
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err.Error())
	}
	req.Header.Add("Content-Type", mimeTypeJson)
	req.SetBasicAuth(c.username, c.password)

	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d. Please contact an administrator if the problem persists", response.StatusCode)
	}

	return nil
}

func (c *DnsCfg) RewriteDelete(payload *Record) error {
	url := c.url + "/control/rewrite/delete"
	mimeTypeJson := "application/json"

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	tr := &http.Transport{
		IdleConnTimeout: time.Second * 5,
	}
	client := &http.Client{
		Transport: tr,
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err.Error())
	}
	req.Header.Add("Content-Type", mimeTypeJson)
	req.SetBasicAuth(c.username, c.password)

	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d. Please contact an administrator if the problem persists", response.StatusCode)
	}

	return nil
}
