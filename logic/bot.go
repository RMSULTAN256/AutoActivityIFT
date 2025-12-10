package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"masterc/models"
	"net/http"
)

type ApiResp struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func (e *ApiResp) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.Data)
}

var (
	url = `http://localhost:5544/api/v1/adduser`
)

func HandleCrendetials(name, user, pass, platform string) error {

	data := models.Users{
		Name:     name,
		Username: user,
		Password: pass,
		Platform: platform,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	var res ApiResp
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return fmt.Errorf("failed to parse response (status %d): %w", resp.StatusCode, err)
	}

	if resp.StatusCode >= 400 {
		if res.Data != "" {
			return fmt.Errorf(res.Data)
		}
		return fmt.Errorf("server error (%d)", resp.StatusCode)
	}

	if res.Status == "fail" {
		if res.Data != "" {
			return fmt.Errorf(res.Data)
		}
		return fmt.Errorf("Send Failed")
	}

	return nil
}

func HandleBrowser(name, action string) error {
	data := models.Action{
		Name:       name,
		Activities: []string{action},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	var res ApiResp
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return fmt.Errorf("failed to parse response (status %d): %w", resp.StatusCode, err)
	}

	if resp.StatusCode >= 400 {
		if res.Data != "" {
			return fmt.Errorf(res.Data)
		}
		return fmt.Errorf("server error (%d)", resp.StatusCode)
	}

	if res.Status == "fail" {
		if res.Data != "" {
			return fmt.Errorf(res.Data)
		}
		return fmt.Errorf("Send Failed")
	}
	return nil
}
