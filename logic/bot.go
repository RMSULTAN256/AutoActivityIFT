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

type StartBotRequest struct {
	Name       string   `json:"name" binding:"required"`
	Mode       string   `json:"mode"`
	Activities []string `json:"activities,omitempty"`
	Interval   string   `json:"interval"`
	TTL        string   `json:"ttl"`
}

func (e *ApiResp) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.Data)
}

var (
	url      = `http://localhost:5544/api/v1/adduser`
	urlStart = `http://localhost:5544/api/v1/startbot`
	urlStop  = `http://localhost:5544/api/v1/stopbot`
	urlReset = `http://localhost:5544/api/v1/resetlogs`
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
		return fmt.Errorf("send Failed")
	}

	return nil
}

func BrowserIdle(name, schedule, action, remain string) error {
	var ListActivites []string

	if action != "idle" {
		ListActivites = []string{action}
	} else {
		ListActivites = []string{}
	}

	formattedInterval := ""
	if schedule != "" {
		formattedInterval = schedule + "m"
	}

	formattedTTL := ""
	if remain != "" {
		formattedTTL = remain + "h"
	}

	data := StartBotRequest{
		Name:       name,
		Mode:       action,
		Activities: ListActivites,
		Interval:   formattedInterval,
		TTL:        formattedTTL,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	resp, err := http.Post(urlStart, "application/json", bytes.NewBuffer(jsonData))
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
		return fmt.Errorf("send Failed")
	}

	return nil
}

func BrowserClose(name string) error {
	Data := StartBotRequest{
		Name: name,
	}

	jsonData, err := json.Marshal(Data)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	resp, err := http.Post(urlStop, "application/json", bytes.NewBuffer(jsonData))
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

	if resp.Status == "fail" {
		if res.Data != "" {
			return fmt.Errorf(res.Data)
		}
		return fmt.Errorf("send Failed")
	}

	return nil
}

func ResetLogsData() error {
	req, err := http.NewRequest(http.MethodDelete, urlReset, nil)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("server error (%d)", resp.StatusCode)
	}

	return nil
}
