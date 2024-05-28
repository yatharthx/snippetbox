package main

import (
	"net/http"
	"testing"

	"snippetbox.yatharthx.com/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name         string
		urlPath      string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid ID",
			urlPath:      "/snippet/view/1",
			expectedCode: http.StatusOK,
			expectedBody: "An old silent pond...",
		},
		{
			name:         "Non-existent ID",
			urlPath:      "/snippet/view/2",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Negative ID",
			urlPath:      "/snippet/view/-1",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Decimal ID",
			urlPath:      "/snippet/view/1.23",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "String ID",
			urlPath:      "/snippet/view/foo",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Empty ID",
			urlPath:      "/snippet/view/",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			assert.Equal(t, code, tt.expectedCode)

			if tt.expectedBody != "" {
				assert.StringContains(t, body, tt.expectedBody)
			}
		})
	}
}
