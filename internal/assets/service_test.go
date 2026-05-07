package assets

import (
	"testing"
)

func TestFilterByType(t *testing.T) {
	var assetData = []Asset{
		{"1", "server-1", "server", "eu-west"}}

	s := NewService(assetData)
	testCorrectType := "server"
	testWrongType := "database"
	if len(s.FilterByType(testCorrectType)) != 1 {
		t.Errorf("FilterByType did not find Type (%s)", testCorrectType)
	}
	if len(s.FilterByType(testWrongType)) != 0 {
		t.Errorf("FilterByType found nonExisting Type (%s)", testWrongType)
	}
}

func TestGetAll(t *testing.T) {
	var assetData = []Asset{
		{"1", "name-1", "type-1", "location-1"},
		{"2", "name-2", "type-2", "location-2"},
	}
	s := NewService(assetData)

	if len(s.GetAll()) != len(assetData) {
		t.Error("GetAll does not have same length than assetData.")
	}
}
