package cloudflare

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListStreamInputVideos(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
  "result": [
    {
      "uid": "ee4bbe9e47e10a4bdfbe42c08a37b0cd",
      "thumbnail": "https://videodelivery.net/ee4bbe9e47e10a4bdfbe42c08a37b0cd/thumbnails/thumbnail.jpg",
      "thumbnailTimestampPct": 0,
      "readyToStream": false,
      "status": {
        "state": "live-inprogress",
        "errorReasonCode": "",
        "errorReasonText": ""
      },
      "meta": {
        "name": "13 Feb 22 17:27 UTC"
      },
      "created": "2022-02-13T17:27:34.572191Z",
      "modified": "2022-02-13T17:27:34.572191Z",
      "size": 0,
      "preview": "https://watch.videodelivery.net/ee4bbe9e47e10a4bdfbe42c08a37b0cd",
      "allowedOrigins": [],
      "requireSignedURLs": false,
      "uploaded": "2022-02-13T17:27:34.572179Z",
      "uploadExpiry": null,
      "maxSizeBytes": null,
      "maxDurationSeconds": null,
      "duration": -1,
      "input": {
        "width": -1,
        "height": -1
      },
      "playback": {
        "hls": "https://videodelivery.net/ee4bbe9e47e10a4bdfbe42c08a37b0cd/manifest/video.m3u8",
        "dash": "https://videodelivery.net/ee4bbe9e47e10a4bdfbe42c08a37b0cd/manifest/video.mpd"
      },
      "watermark": null,
      "liveInput": "6bb55c9389c27877baf1c185facca222"
    },
    {
      "uid": "52ecc71c930803874c5145006b2a6c5a",
      "thumbnail": "https://videodelivery.net/52ecc71c930803874c5145006b2a6c5a/thumbnails/thumbnail.jpg",
      "thumbnailTimestampPct": 0,
      "readyToStream": true,
      "status": {
        "state": "ready",
        "pctComplete": "100.000000",
        "errorReasonCode": "",
        "errorReasonText": ""
      },
      "meta": {
        "name": "13 Feb 22 17:26 UTC"
      },
      "created": "2022-02-13T17:26:33.865603Z",
      "modified": "2022-02-13T17:27:18.873315Z",
      "size": 0,
      "preview": "https://watch.videodelivery.net/52ecc71c930803874c5145006b2a6c5a",
      "allowedOrigins": [],
      "requireSignedURLs": false,
      "uploaded": "2022-02-13T17:26:33.865594Z",
      "uploadExpiry": null,
      "maxSizeBytes": null,
      "maxDurationSeconds": null,
      "duration": 14,
      "input": {
        "width": 1280,
        "height": 720
      },
      "playback": {
        "hls": "https://videodelivery.net/52ecc71c930803874c5145006b2a6c5a/manifest/video.m3u8",
        "dash": "https://videodelivery.net/52ecc71c930803874c5145006b2a6c5a/manifest/video.mpd"
      },
      "watermark": null,
      "liveInput": "6bb55c9389c27877baf1c185facca222"
    },
    {
      "uid": "40b953dfafbdae0014cc383630635d6a",
      "thumbnail": "https://videodelivery.net/40b953dfafbdae0014cc383630635d6a/thumbnails/thumbnail.jpg",
      "thumbnailTimestampPct": 0,
      "readyToStream": true,
      "status": {
        "state": "ready",
        "pctComplete": "100.000000",
        "errorReasonCode": "",
        "errorReasonText": ""
      },
      "meta": {
        "name": "13 Feb 22 17:23 UTC"
      },
      "created": "2022-02-13T17:23:54.070279Z",
      "modified": "2022-02-13T17:26:20.977092Z",
      "size": 0,
      "preview": "https://watch.videodelivery.net/40b953dfafbdae0014cc383630635d6a",
      "allowedOrigins": [],
      "requireSignedURLs": false,
      "uploaded": "2022-02-13T17:23:54.070263Z",
      "uploadExpiry": null,
      "maxSizeBytes": null,
      "maxDurationSeconds": null,
      "duration": 132,
      "input": {
        "width": 1280,
        "height": 720
      },
      "playback": {
        "hls": "https://videodelivery.net/40b953dfafbdae0014cc383630635d6a/manifest/video.m3u8",
        "dash": "https://videodelivery.net/40b953dfafbdae0014cc383630635d6a/manifest/video.mpd"
      },
      "watermark": null,
      "liveInput": "6bb55c9389c27877baf1c185facca222"
    },
    {
      "uid": "1d8329a30e312f99a7125608a8fbc784",
      "thumbnail": "https://videodelivery.net/1d8329a30e312f99a7125608a8fbc784/thumbnails/thumbnail.jpg",
      "thumbnailTimestampPct": 0,
      "readyToStream": true,
      "status": {
        "state": "ready",
        "pctComplete": "100.000000",
        "errorReasonCode": "",
        "errorReasonText": ""
      },
      "meta": {
        "name": "13 Feb 22 17:18 UTC"
      },
      "created": "2022-02-13T17:18:58.356017Z",
      "modified": "2022-02-13T17:23:07.233741Z",
      "size": 0,
      "preview": "https://watch.videodelivery.net/1d8329a30e312f99a7125608a8fbc784",
      "allowedOrigins": [],
      "requireSignedURLs": false,
      "uploaded": "2022-02-13T17:18:58.356003Z",
      "uploadExpiry": null,
      "maxSizeBytes": null,
      "maxDurationSeconds": null,
      "duration": 219,
      "input": {
        "width": 1280,
        "height": 720
      },
      "playback": {
        "hls": "https://videodelivery.net/1d8329a30e312f99a7125608a8fbc784/manifest/video.m3u8",
        "dash": "https://videodelivery.net/1d8329a30e312f99a7125608a8fbc784/manifest/video.mpd"
      },
      "watermark": null,
      "liveInput": "6bb55c9389c27877baf1c185facca222"
    },
    {
      "uid": "f2de09bb47997827e8a91edc8569698f",
      "thumbnail": "https://videodelivery.net/f2de09bb47997827e8a91edc8569698f/thumbnails/thumbnail.jpg",
      "thumbnailTimestampPct": 0,
      "readyToStream": true,
      "status": {
        "state": "ready",
        "pctComplete": "100.000000",
        "errorReasonCode": "",
        "errorReasonText": ""
      },
      "meta": {
        "name": "11 Feb 22 19:49 UTC"
      },
      "created": "2022-02-11T19:49:51.074112Z",
      "modified": "2022-02-11T19:50:42.853036Z",
      "size": 0,
      "preview": "https://watch.videodelivery.net/f2de09bb47997827e8a91edc8569698f",
      "allowedOrigins": [],
      "requireSignedURLs": false,
      "uploaded": "2022-02-11T19:49:51.074103Z",
      "uploadExpiry": null,
      "maxSizeBytes": null,
      "maxDurationSeconds": null,
      "duration": 37.18,
      "input": {
        "width": 1280,
        "height": 720
      },
      "playback": {
        "hls": "https://videodelivery.net/f2de09bb47997827e8a91edc8569698f/manifest/video.m3u8",
        "dash": "https://videodelivery.net/f2de09bb47997827e8a91edc8569698f/manifest/video.mpd"
      },
      "watermark": null,
      "liveInput": "6bb55c9389c27877baf1c185facca222"
    },
    {
      "uid": "769ddba4719bdf99f389dffeb9bb081a",
      "thumbnail": "https://videodelivery.net/769ddba4719bdf99f389dffeb9bb081a/thumbnails/thumbnail.jpg",
      "thumbnailTimestampPct": 0,
      "readyToStream": true,
      "status": {
        "state": "ready",
        "pctComplete": "100.000000",
        "errorReasonCode": "",
        "errorReasonText": ""
      },
      "meta": {
        "name": "11 Feb 22 19:29 UTC"
      },
      "created": "2022-02-11T19:29:48.32338Z",
      "modified": "2022-02-11T19:30:05.9377Z",
      "size": 0,
      "preview": "https://watch.videodelivery.net/769ddba4719bdf99f389dffeb9bb081a",
      "allowedOrigins": [],
      "requireSignedURLs": false,
      "uploaded": "2022-02-11T19:29:48.323372Z",
      "uploadExpiry": null,
      "maxSizeBytes": null,
      "maxDurationSeconds": null,
      "duration": 2,
      "input": {
        "width": 1280,
        "height": 720
      },
      "playback": {
        "hls": "https://videodelivery.net/769ddba4719bdf99f389dffeb9bb081a/manifest/video.m3u8",
        "dash": "https://videodelivery.net/769ddba4719bdf99f389dffeb9bb081a/manifest/video.mpd"
      },
      "watermark": null,
      "liveInput": "6bb55c9389c27877baf1c185facca222"
    },
    {
      "uid": "e3f3e64d420fc977828757ae14770949",
      "thumbnail": "https://videodelivery.net/e3f3e64d420fc977828757ae14770949/thumbnails/thumbnail.jpg",
      "thumbnailTimestampPct": 0,
      "readyToStream": true,
      "status": {
        "state": "ready",
        "pctComplete": "100.000000",
        "errorReasonCode": "",
        "errorReasonText": ""
      },
      "meta": {
        "name": "11 Feb 22 19:08 UTC"
      },
      "created": "2022-02-11T19:08:04.097633Z",
      "modified": "2022-02-11T19:27:29.872898Z",
      "size": 0,
      "preview": "https://watch.videodelivery.net/e3f3e64d420fc977828757ae14770949",
      "allowedOrigins": [],
      "requireSignedURLs": false,
      "uploaded": "2022-02-11T19:08:04.097619Z",
      "uploadExpiry": null,
      "maxSizeBytes": null,
      "maxDurationSeconds": null,
      "duration": 1151.3,
      "input": {
        "width": 1280,
        "height": 720
      },
      "playback": {
        "hls": "https://videodelivery.net/e3f3e64d420fc977828757ae14770949/manifest/video.m3u8",
        "dash": "https://videodelivery.net/e3f3e64d420fc977828757ae14770949/manifest/video.mpd"
      },
      "watermark": null,
      "liveInput": "6bb55c9389c27877baf1c185facca222"
    },
    {
      "uid": "cda0b7ddec959ce66ec0acbc993923ed",
      "thumbnail": "https://videodelivery.net/cda0b7ddec959ce66ec0acbc993923ed/thumbnails/thumbnail.jpg",
      "thumbnailTimestampPct": 0,
      "readyToStream": true,
      "status": {
        "state": "ready",
        "pctComplete": "100.000000",
        "errorReasonCode": "",
        "errorReasonText": ""
      },
      "meta": {
        "name": "11 Feb 22 17:40 UTC"
      },
      "created": "2022-02-11T17:40:45.100795Z",
      "modified": "2022-02-11T17:44:00.305488Z",
      "size": 0,
      "preview": "https://watch.videodelivery.net/cda0b7ddec959ce66ec0acbc993923ed",
      "allowedOrigins": [],
      "requireSignedURLs": false,
      "uploaded": "2022-02-11T17:40:45.100787Z",
      "uploadExpiry": null,
      "maxSizeBytes": null,
      "maxDurationSeconds": null,
      "duration": 180.67,
      "input": {
        "width": 1280,
        "height": 720
      },
      "playback": {
        "hls": "https://videodelivery.net/cda0b7ddec959ce66ec0acbc993923ed/manifest/video.m3u8",
        "dash": "https://videodelivery.net/cda0b7ddec959ce66ec0acbc993923ed/manifest/video.mpd"
      },
      "watermark": null,
      "liveInput": "6bb55c9389c27877baf1c185facca222"
    }
  ],
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

func TestStreamInputVideoDetails(t *testing.T) {
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
