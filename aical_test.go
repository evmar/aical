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
        }
    ]
}`

func TestJSONToURL(t *testing.T) {
	events, err := parseEventsJSON([]byte(JSON_SAMPLE))
	if err != nil {
		t.Fatal(err)
	}

	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}

	urls := []string{
		"https://www.google.com/calendar/render?action=TEMPLATE&text=Finding+Dory+%28in+celebration+of+Disability+Pride+Month%21%29&dates=20240725&location=Westlake+Park",
		"https://www.google.com/calendar/render?action=TEMPLATE&text=Kung+Fu+Panda+4&dates=20240801&location=Millennium+Plaza+Park",
	}
	for i, event := range events {
		if got := eventToURL(event); got != urls[i] {
			t.Fatalf("expected %q, got %q", urls[i], got)
		}
	}
}
