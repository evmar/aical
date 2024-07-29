package main

import "testing"

const JSON_SAMPLE = `{
    "events": [
        {
            "name": "Finding Dory (in celebration of Disability Pride Month!)",
            "date": "20240725",
            "location": "Westlake Park",
            "time": null
        },
        {
            "name": "Kung Fu Panda 4",
            "date": "20240801",
            "location": "Millennium Plaza Park",
            "time": null
        },
		{
			"name": "Potluck Style",
			"date": "20240822",
			"location": "1234 Johnson Road",
			"time": "1800"
		}
    ]
}`

func TestJSONToURL(t *testing.T) {
	events, err := parseEventsJSON([]byte(JSON_SAMPLE))
	if err != nil {
		t.Fatal(err)
	}

	urls := []string{
		"https://www.google.com/calendar/render?action=TEMPLATE&text=Finding+Dory+%28in+celebration+of+Disability+Pride+Month%21%29&dates=20240725T000000%2F20240725T000000&location=Westlake+Park",
		"https://www.google.com/calendar/render?action=TEMPLATE&text=Kung+Fu+Panda+4&dates=20240801T000000%2F20240801T000000&location=Millennium+Plaza+Park",
		"https://www.google.com/calendar/render?action=TEMPLATE&text=Potluck+Style&dates=20240822T180000%2F20240822T180000&location=1234+Johnson+Road",
	}
	if len(events) != len(urls) {
		t.Fatalf("expected %d events, got %d", len(urls), len(events))
	}
	for i, event := range events {
		if got := eventToURL(event); got != urls[i] {
			t.Fatalf("expected %q, got %q", urls[i], got)
		}
	}
}
