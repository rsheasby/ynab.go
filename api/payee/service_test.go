package payee_test

import (
	"net/http"
	"testing"

	"strconv"

	"github.com/stretchr/testify/assert"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/payee"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestService_GetPayees(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/payees"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "payees": [
      {
        "id": "34e88373-ef48-4386-9ab3-7f86c2a8988f",
        "name": "Supermarket",
        "transfer_account_id": null,
        "deleted": false
      }
		]
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	payees, err := client.Payee().GetPayees("aa248caa-eed7-4575-a990-717386438d2c")
	assert.NoError(t, err)

	expected := []*payee.Payee{
		{
			ID:      "34e88373-ef48-4386-9ab3-7f86c2a8988f",
			Name:    "Supermarket",
			Deleted: false,
		},
	}

	assert.Equal(t, expected, payees)
}

func TestService_GetPayee(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/payees/34e88373-ef48-4386-9ab3-7f86c2a8988f"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
		"payee": {
			"id": "34e88373-ef48-4386-9ab3-7f86c2a8988f",
			"name": "Supermarket",
			"transfer_account_id": null,
			"deleted": false
		}
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	p, err := client.Payee().GetPayee(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"34e88373-ef48-4386-9ab3-7f86c2a8988f",
	)
	assert.NoError(t, err)

	expected := &payee.Payee{
		ID:      "34e88373-ef48-4386-9ab3-7f86c2a8988f",
		Name:    "Supermarket",
		Deleted: false,
	}

	assert.Equal(t, expected, p)
}

func TestService_GetPayeeLocations(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/payee_locations"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "payee_locations": [
      {
        "id": "34fabc3-1234-4a11-8bcd-7f63756b7193",
        "payee_id": "34e88373-ef48-4386-9ab3-7f86c2a8988f",
        "latitude": "42.496348",
        "longitude": "23.3095594",
        "deleted": false
      }
		]
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	locations, err := client.Payee().GetPayeeLocations("aa248caa-eed7-4575-a990-717386438d2c")
	assert.NoError(t, err)

	latitude, err := strconv.ParseFloat("42.496348", 64)
	assert.NoError(t, err)

	longitude, err := strconv.ParseFloat("23.3095594", 64)
	assert.NoError(t, err)

	expected := []*payee.Location{
		{
			ID:        "34fabc3-1234-4a11-8bcd-7f63756b7193",
			PayeeID:   "34e88373-ef48-4386-9ab3-7f86c2a8988f",
			Latitude:  &latitude,
			Longitude: &longitude,
			Deleted:   false,
		},
	}

	assert.Equal(t, expected, locations)
}

func TestService_GetPayeeLocation(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/payee_locations/34fabc3-1234-4a11-8bcd-7f63756b7193"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "payee_location": {
			"id": "34fabc3-1234-4a11-8bcd-7f63756b7193",
			"payee_id": "34e88373-ef48-4386-9ab3-7f86c2a8988f",
			"latitude": "42.496348",
			"longitude": "23.3095594",
			"deleted": false
		}
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	location, err := client.Payee().GetPayeeLocation(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"34fabc3-1234-4a11-8bcd-7f63756b7193",
	)
	assert.NoError(t, err)

	latitude, err := strconv.ParseFloat("42.496348", 64)
	assert.NoError(t, err)

	longitude, err := strconv.ParseFloat("23.3095594", 64)
	assert.NoError(t, err)

	expected := &payee.Location{
		ID:        "34fabc3-1234-4a11-8bcd-7f63756b7193",
		PayeeID:   "34e88373-ef48-4386-9ab3-7f86c2a8988f",
		Latitude:  &latitude,
		Longitude: &longitude,
		Deleted:   false,
	}

	assert.Equal(t, expected, location)
}

func TestService_GetPayeeLocationsByPayee(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/payees/34e88373-ef48-4386-9ab3-7f86c2a8988f/payee_locations"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "payee_locations": [
      {
        "id": "34fabc3-1234-4a11-8bcd-7f63756b7193",
        "payee_id": "34e88373-ef48-4386-9ab3-7f86c2a8988f",
        "latitude": "42.496348",
        "longitude": "23.3095594",
        "deleted": false
      }
		]
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	locations, err := client.Payee().GetPayeeLocationsByPayee(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"34e88373-ef48-4386-9ab3-7f86c2a8988f",
	)
	assert.NoError(t, err)

	latitude, err := strconv.ParseFloat("42.496348", 64)
	assert.NoError(t, err)

	longitude, err := strconv.ParseFloat("23.3095594", 64)
	assert.NoError(t, err)

	expected := []*payee.Location{
		{
			ID:        "34fabc3-1234-4a11-8bcd-7f63756b7193",
			PayeeID:   "34e88373-ef48-4386-9ab3-7f86c2a8988f",
			Latitude:  &latitude,
			Longitude: &longitude,
			Deleted:   false,
		},
	}

	assert.Equal(t, expected, locations)
}
