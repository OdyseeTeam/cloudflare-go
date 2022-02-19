package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type StreamInputVideo struct {
	UID                   string                 `json:"uid"`
	Thumbnail             string                 `json:"thumbnail"`
	ThumbnailTimestampPct float64                `json:"thumbnailTimestampPct"`
	ReadyToStream         bool                   `json:"readyToStream"`
	Status                StreamInputVideoStatus `json:"status,omitempty"`
	Meta                  map[string]string      `json:"meta,omitempty"`
	Created               time.Time              `json:"created"`
	Modified              time.Time              `json:"modified"`
	Size                  int64                  `json:"size"`
	Preview               string                 `json:"preview"`
	AllowedOrigins        []string               `json:"allowedOrigins"`
	RequireSignedURLs     bool                   `json:"requireSignedURLs"`
	Uploaded              time.Time              `json:"uploaded"`
	UploadExpiry          time.Time              `json:"uploadExpiry"`
	MaxSizeBytes          int64                  `json:"maxSizeBytes"`
	MaxDurationSeconds    int64                  `json:"maxDurationSeconds"`
	Duration              float64                `json:"duration"`
	Input                 Input                  `json:"input,omitempty"`
	Playback              Playback               `json:"playback,omitempty"`
	Watermark             Watermark              `json:"watermark,omitempty"`
	LiveInput             string                 `json:"liveInput"`
	NFT                   NFT                    `json:"nft,omitempty"`
}

type Input struct {
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

type StreamInputVideoStatus struct {
	State           string `json:"state"`
	PercentComplete string `json:"pctComplete"`
	ErrorReasonCode string `json:"errorReasonCode"`
	ErrorReasonTxt  string `json:"errorReasonText"`
}

type Playback struct {
	HLS  string `json:"hls"`
	Dash string `json:"dash"`
}

type Watermark struct {
	UID            string    `json:"uid"`
	Size           int64     `json:"size"`
	Height         int64     `json:"height"`
	Width          int64     `json:"width"`
	Created        time.Time `json:"created"`
	DownloadedFrom string    `json:"downloadedFrom"`
	Name           string    `json:"name"`
	Opacity        float64   `json:"opacity"`
	Padding        float64   `json:"padding"`
	Scale          float64   `json:"scale"`
	Position       string    `json:"position"`
}

type NFT struct {
	Contract string `json:"contract"`
	Token    int64  `json:"token"`
}

// loadBalancerResponse represents the response from the Stream Input endpoints.
type streamInputVideoResponse struct {
	Response
	Result StreamInputVideo `json:"result"`
}

// streamInputListResponse represents the response from the List Stream Inputs endpoint.
type streamInputVideoListResponse struct {
	Response
	Result     []StreamInputVideo `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// CreateStreamInput creates a new stream input.
//
// API reference: ?
func (api *API) ListStreamInputVideos(ctx context.Context, accountID, liveInputID string) ([]StreamInputVideo, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/live_inputs/%s/videos", accountID, liveInputID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []StreamInputVideo{}, err
	}
	var r streamInputVideoListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []StreamInputVideo{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// StreamInputDetails retrieves the details of a stream input.
//
// API reference: ?
func (api *API) StreamInputVideoDetails(ctx context.Context, accountID, streamInputID, videoID string) (StreamInputVideo, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/live_inputs/%s/videos/%s", accountID, streamInputID, videoID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return StreamInputVideo{}, err
	}
	var r streamInputVideoResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return StreamInputVideo{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
