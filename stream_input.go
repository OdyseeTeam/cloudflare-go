package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type StreamInput struct {
	UID       string            `json:"uid"`
	RMTPInfo  RMTPInfo          `json:"rtmps,omitempty"`
	Created   time.Time         `json:"created"`
	Modified  time.Time         `json:"modified"`
	Meta      map[string]string `json:"meta,omitempty"`
	Status    StreamInputStatus `json:"status,omitempty"`
	Recording RecordingState    `json:"recording,omitempty"`
}

type RecordingState struct {
	Mode              string   `json:"mode"`
	TimeoutSeconds    int64    `json:"timeoutSeconds"`
	RequireSignedURLs bool     `json:"requireSignedURLs"`
	AllowedOrigins    []string `json:"allowedOrigins"`
}

type StreamInputStatus struct {
	Current StreamInputState   `json:"current"`
	History []StreamInputState `json:"history"`
}

type StreamInputState struct {
	Reason          string    `json:"reason"`
	State           string    `json:"state"`
	StatusEnteredAt time.Time `json:"statusEnteredAt"`
	StatusLastSeen  time.Time `json:"statusLastSeen"`
}

type RMTPInfo struct {
	URL       string `json:"url"`
	StreamKey string `json:"streamKey"`
}

// loadBalancerResponse represents the response from the Stream Input endpoints.
type streamInputResponse struct {
	Response
	Result StreamInput `json:"result"`
}

// streamInputListResponse represents the response from the List Stream Inputs endpoint.
type streamInputListResponse struct {
	Response
	Result     []StreamInput `json:"result"`
	ResultInfo ResultInfo    `json:"result_info"`
}

// CreateStreamInput creates a new stream input.
//
// API reference: https://api.cloudflare.com/#stream-live-inputs-create-a-live-input
func (api *API) CreateStreamInput(ctx context.Context, accountID string, si StreamInput) (StreamInput, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/live_inputs", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, si)
	if err != nil {
		return StreamInput{}, err
	}
	var r streamInputResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return StreamInput{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// StreamInputDetails retrieves the details of a stream input.
//
// API reference: https://api.cloudflare.com/#stream-live-inputs-live-input-details
func (api *API) StreamInputDetails(ctx context.Context, streamInputID, accountID string) (StreamInput, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/live_inputs/%s", accountID, streamInputID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return StreamInput{}, err
	}
	var r streamInputResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return StreamInput{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListStreamInputs lists stream inputs created for an account.
//
// API reference: https://api.cloudflare.com/#stream-live-inputs-list-live-inputs
func (api *API) ListStreamInputs(ctx context.Context, accountID string) ([]StreamInput, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/live_inputs", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r streamInputListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
