package assets

import (
	"testing"
	"time"
)

func fakeAsset(fakeField string) Asset {

	asset := Asset{
		ID:          "1",
		Name:        "name-1",
		Type:        "server",
		Location:    "eu-west",
		Environment: "production",
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	switch fakeField {
	case "ID":
		asset.ID = ""
	case "Name":
		asset.Name = ""
	case "Location":
		asset.Location = ""
	case "Type":
		asset.Type = "fake-type"
	case "Environment":
		asset.Environment = "fake-environment"
	case "Status":
		asset.Status = "fake-status"
	case "CreatedAt":
		asset.CreatedAt = time.Time{}
	case "UpdatedAt":
		asset.UpdatedAt = time.Time{}
	}

	return asset
}
func TestValidateAsset(t *testing.T) {
	asset := fakeAsset("ok")

	err := ValidateAsset(asset)

	if err != nil {
		t.Errorf("expected no validation error.")
	}
}
func TestEmptyId(t *testing.T) {
	asset := fakeAsset("ID")

	err := ValidateAsset(asset)

	if err.Error() != "asset ID cannot be empty" {
		t.Error("expected id validation error")
	}
}
func TestEmptyName(t *testing.T) {
	asset := fakeAsset("Name")

	err := ValidateAsset(asset)

	if err.Error() != "asset name cannot be empty" {
		t.Error("expected name validation error")
	}
}
func TestEmptyLocation(t *testing.T) {
	asset := fakeAsset("Location")

	err := ValidateAsset(asset)

	if err != nil && err.Error() != "location cannot be empty" {
		t.Error("expected location validation error")
	}
}

func TestInvalidType(t *testing.T) {
	asset := fakeAsset("Type")

	err := ValidateAsset(asset)

	if err == nil {
		t.Error("expected asset type error")
	}
}
func TestInvalidEnvironment(t *testing.T) {
	asset := fakeAsset("Environment")

	err := ValidateAsset(asset)

	if err == nil {
		t.Error("expected invalid environment error")
	}
}
func TestInvalidStatus(t *testing.T) {
	asset := fakeAsset("Status")

	err := ValidateAsset(asset)

	if err == nil {
		t.Error("expected invalid status error")
	}
}
func TestInvalidCreatedAt(t *testing.T) {
	asset := fakeAsset("CreatedAt")

	err := ValidateAsset(asset)

	if err.Error() != "created_at cannot be empty" {
		t.Error("expected invalid created_at error")
	}
}
func TestInvalidUpdatedAt(t *testing.T) {
	asset := fakeAsset("UpdatedAt")

	err := ValidateAsset(asset)

	if err.Error() != "updated_at cannot be empty" {
		t.Error("expected invalid updated_at error")
	}
}
