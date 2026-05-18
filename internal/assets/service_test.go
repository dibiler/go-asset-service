package assets

import (
	"slices"
	"testing"
	"time"
)

var createdAt time.Time = (func() time.Time {
	date, _ := time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")
	return date
})()
var updatedAt time.Time = (func() time.Time {
	date, _ := time.Parse(time.RFC3339, "2024-02-01T00:00:00Z")
	return date
})()

var testAssets = []Asset{
	{"1", "name-1", "type-1", "location-1", "development", "active", createdAt, updatedAt},
	{"2", "name-2", "type-2", "location-2", "staging", "inactive", createdAt, updatedAt},
}
var testService = NewService(testAssets)

func TestGetAll(t *testing.T) {

	if len(testService.GetAll()) != len(testAssets) {
		t.Errorf("GetAll does not have same length than assets.")
	}
}

func TestFilterByType(t *testing.T) {
	testCorrectType := "type-1"
	testWrongType := "type-fake"
	if len(testService.FilterByType(testCorrectType)) != 1 {
		t.Errorf("FilterByType did not find Type (%s)", testCorrectType)
	}
	if len(testService.FilterByType(testWrongType)) != 0 {
		t.Errorf("FilterByType found nonExisting Type (%s)", testWrongType)
	}
}

func TestFilterByLocation(t *testing.T) {

	filterLoc1 := testService.FilterByLocation("location-1")
	if len(filterLoc1) != 1 {
		t.Errorf("FilterByLocation did not find asset.")
	}
	filterLocFake := testService.FilterByLocation("location-fake")
	if len(filterLocFake) != 0 {
		t.Errorf("FilterByLocation found wrong asset.")
	}
}
func TestFilterByName(t *testing.T) {

	filterName1 := testService.FilterByName("name-1")
	if len(filterName1) != 1 {
		t.Errorf("FilterByName did not find asset.")
	}
	filterNameFake := testService.FilterByName("name-fake")
	if len(filterNameFake) != 0 {
		t.Errorf("FilterByName found wrong asset.")
	}
}
func TestFilterBy(t *testing.T) {

	var types = []string{"name", "location", "type"}

	for _, a := range types {
		testValue := a + "-1"
		switch a {
		case "name":
			filterByNameValue := testService.FilterByName(testValue)
			filterByValue := testService.FilterBy(func(asset Asset) bool {
				return asset.Name == testValue
			})
			if !slices.Equal(filterByValue, filterByNameValue) {
				t.Errorf("FilterBy(name,_) not behaving the same as FilterByName")
			}
		case "location":
			filterByLocationValue := testService.FilterByLocation(testValue)
			filterByValue := testService.FilterBy(func(asset Asset) bool {
				return asset.Location == testValue
			})
			if !slices.Equal(filterByValue, filterByLocationValue) {
				t.Errorf("FilterBy(location,_) not behaving the same as FilterByLocation")
			}
		case "type":
			filterByTypeValue := testService.FilterByType(testValue)
			filterByValue := testService.FilterBy(func(asset Asset) bool {
				return asset.Type == testValue
			})
			if !slices.Equal(filterByValue, filterByTypeValue) {
				t.Errorf("FilterBy(type,_) not behaving the same as FilterByType")
			}
		default:
			t.Errorf("Wrong filter using TestFilterBy (%s)", a)
		}
	}
	filterName1 := testService.FilterByName("name-1")
	if len(filterName1) != 1 {
		t.Errorf("FilterByName did not find asset.")
	}
	filterNameFake := testService.FilterByName("name-fake")
	if len(filterNameFake) != 0 {
		t.Errorf("FilterByName found wrong asset.")
	}
}

func TestCountByType(t *testing.T) {
	testType := "type-1"

	assetsFound := testService.CountByType(testType)

	if assetsFound != 1 {
		t.Errorf("CountByType can't find types correctly. Found %d assets instead of 1.", assetsFound)
	}
}
