package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

// streamVideoMP4Response represents the response from the Stream MP4 endpoints.
type streamVideoMP4Response struct {
	Response
	Result StreamVideoDownloads `json:"result"`
}

type StreamVideoDownloads struct {
	Default StreamVideoMP4 `json:"default"`
}

type StreamVideoMP4 struct {
	URL             string  `json:"url"`
	Status          string  `json:"status"`
	PercentComplete float64 `json:"percentComplete"`
}

// ListStreamInputVideoDownloads retrieves the list of downloads available for a video
//
// API reference: https://api.cloudflare.com/#stream-mp4-downloads-list-downloads
func (api *API) CreateVideoDownloads(ctx context.Context, accountID, videoID string) (StreamVideoDownloads, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/%s/downloads", accountID, videoID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return StreamVideoDownloads{}, err
	}
	var r streamVideoMP4Response
	if err := json.Unmarshal(res, &r); err != nil {
		return StreamVideoDownloads{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListStreamInputVideoDownloads retrieves the list of downloads available for a video
//
// API reference: https://api.cloudflare.com/#stream-mp4-downloads-list-downloads
func (api *API) ListVideoDownloads(ctx context.Context, accountID, videoID string) (StreamVideoDownloads, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/%s/downloads", accountID, videoID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return StreamVideoDownloads{}, err
	}
	var r streamVideoMP4Response
	if err := json.Unmarshal(res, &r); err != nil {
		return StreamVideoDownloads{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
