package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateStreamInput(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": {
				"uid": "256e7e9c7273eebc30b2f9871c337a45",
				"rtmps": {
					"url": "rtmps://live.cloudflare.com:443/live/",
					"streamKey": "2f7e7ab7b386266db11e355a91a79de7k256e7e9c7273eebc30b2f9871c337a45"
				},
				"created": "2022-01-10T22:32:45.479097Z",
				"modified": "2022-01-10T22:32:45.479097Z",
				"meta": {},
				"status": null,
				"recording": {
					"mode": "off",
					"requireSignedURLs": true,
					"allowedOrigins": null
				}
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/stream/live_inputs", handler)

	actual, err := client.CreateStreamInput(context.Background(), testAccountID, StreamInput{})

	if assert.NoError(t, err) {
		assert.Equal(t, actual.RMTPInfo.URL, "rtmps://live.cloudflare.com:443/live/")
		assert.Equal(t, actual.RMTPInfo.StreamKey, "2f7e7ab7b386266db11e355a91a79de7k256e7e9c7273eebc30b2f9871c337a45")
		assert.Equal(t, actual.UID, "256e7e9c7273eebc30b2f9871c337a45")
		assert.Equal(t, actual.Recording.RequireSignedURLs, true)
	}
}

func TestListStreamInput(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": [
				{
					"uid": "256e7e9c7273eebc30b2f9871c337a45",
					"created": "2022-01-10T22:32:45.479097Z",
					"modified": "2022-01-10T22:39:39.291433Z",
					"meta": {}
				}
			],
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/stream/live_inputs", handler)

	actual, err := client.ListStreamInputs(context.Background(), testAccountID)

	if assert.NoError(t, err) {
		assert.Equal(t, actual[0].UID, "256e7e9c7273eebc30b2f9871c337a45")
	}
}
