package assets

import (
	"testing"
)

var testAssets = []Asset{
	{"1", "name-1", "type-1", "location-1"},
	{"2", "name-2", "type-2", "location-2"},
}
var testService = NewService(testAssets)

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

func TestGetAll(t *testing.T) {

	if len(testService.GetAll()) != len(testAssets) {
		t.Errorf("GetAll does not have same length than assets.")
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

			if len(testService.FilterBy(a, testValue)) != len(testService.FilterByName(testValue)) {
				t.Errorf("FilterBy(name,_) not behaving the same as FilterByName")
			}
		case "location":
			if len(testService.FilterBy(a, testValue)) != len(testService.FilterByLocation(testValue)) {
				t.Errorf("FilterBy(name,_) not behaving the same as FilterByName")
			}
		case "type":
			if len(testService.FilterBy(a, testValue)) != len(testService.FilterByType(testValue)) {
				t.Errorf("FilterBy(name,_) not behaving the same as FilterByName")
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
