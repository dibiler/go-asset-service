package cli

import (
	"testing"
	"time"

	"github.com/dibiler/go-asset-service/internal/assets"
)

var createdAt time.Time = (func() time.Time {
	date, _ := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
	return date
})()
var updatedAt time.Time = (func() time.Time {
	date, _ := time.Parse(time.RFC3339, "2024-02-01T00:00:00Z")
	return date
})()

var testAssets = []assets.Asset{
	{ID: "1", Name: "name-1", Type: "type-1", Location: "location-1", Environment: "development", Status: "active", CreatedAt: createdAt, UpdatedAt: updatedAt},
	{ID: "2", Name: "name-2", Type: "type-2", Location: "location-2", Environment: "staging", Status: "inactive", CreatedAt: createdAt, UpdatedAt: updatedAt},
}
var testService = assets.NewService(testAssets)

func FakeRequest(method, filter, value string) Request {
	return Request{method: method, filter: filter, value: value}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name          string
		req           Request
		wantFirstLine string
		wantLen       int
	}{
		{name: "list", req: Request{method: "list"}, wantFirstLine: "All Assets:", wantLen: 3},
		{name: "count empty", req: Request{method: "countbytype"}, wantFirstLine: "can't count an empty type. Please add value parameter", wantLen: 1},
		{name: "bad filter", req: Request{method: "filterby", filter: "bad", value: "x"}, wantFirstLine: "filter parameter [bad] not in: name, location, type", wantLen: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := assets.NewService(testAssets) // fresh service per case
			got := Run(tc.req, svc)

			if len(got) != tc.wantLen {
				t.Fatalf("len got %d want %d", len(got), tc.wantLen)
			}
			if tc.wantFirstLine != "" && got[0] != tc.wantFirstLine {
				t.Fatalf("first line got %q want %q", got[0], tc.wantFirstLine)
			}
		})
	}
}
