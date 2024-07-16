package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
)

type Event struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Time     string `json:"time"`
}

func textToJSON(text string) ([]byte, error) {
	cmd := exec.Command("ai", "-server", "openai", "text", "-sys", "extract all calendar events as json with key 'events', containing a json array.  each entry has keys for name, date, location, and time. use null if not provided.  if year isn't provided, assume 2024.  provide date as yyyymmdd.", "-json", text)
	output, err := cmd.Output()
	if err != nil {
		errText := ""
		if err := err.(*exec.ExitError); err != nil {
			if len(err.Stderr) > 0 {
				errText += string(err.Stderr) + "\n"
			}
		}
		if len(output) > 0 {
			errText += string(output)
		}
		return nil, fmt.Errorf("gpt call failed: %v", errText)
	}
	// uncomment me to gather test input text
	// fmt.Println(string(output))
	return output, nil
}

func parseEventsJSON(text []byte) ([]Event, error) {
	var events struct {
		Events []Event `json:"events"`
	}
	if err := json.Unmarshal(text, &events); err != nil {
		return nil, fmt.Errorf("json parse %q failed: %w", text, err)
	}
	return events.Events, nil
}

func eventToURL(event Event) string {
	u := "https://www.google.com/calendar/render?action=TEMPLATE"
	u += "&text=" + url.QueryEscape(event.Name)
	u += "&dates=" + url.QueryEscape(event.Date)
	if event.Location != "" {
		u += "&location=" + url.QueryEscape(event.Location)
	}
	return u
}

func run() error {
	// read all of stdin to string
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	json, err := textToJSON(string(input))
	if err != nil {
		return err
	}

	events, err := parseEventsJSON(json)
	if err != nil {
		return err
	}

	for _, event := range events {
		fmt.Printf("event: %v\n", event)
		fmt.Println(eventToURL(event))
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
