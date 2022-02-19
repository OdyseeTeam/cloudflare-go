package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateVideoDownloads(t *testing.T) {
	setup()
	defer teardown()
	const testVideoID = "ac136ad7149fff3bdb1de1dbc388d547"

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
    "result": {
        "default": {
            "status": "ready",
            "url": "https://videodelivery.net/`+testVideoID+`/downloads/default.mp4",
            "percentComplete": 100
        }
    },
    "success": true,
    "errors": [],
    "messages": []
}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/stream/"+testVideoID+"/downloads", handler)

	actual, err := client.CreateVideoDownloads(context.Background(), testAccountID, testVideoID)

	if assert.NoError(t, err) {
		assert.Equal(t, actual.Default.URL, "https://videodelivery.net/"+testVideoID+"/downloads/default.mp4")
		assert.Equal(t, actual.Default.Status, "ready")
		assert.Equal(t, actual.Default.PercentComplete, 100.0)
	}
}
