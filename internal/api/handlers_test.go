package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/XellarReps/metricscollector/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateHandler(t *testing.T) {
	s := &Server{
		Storage: storage.NewMemStorage(),
	}

	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name   string
		method string
		url    string
		want   want
	}{
		{
			name:   "positive test update gauge",
			method: http.MethodPost,
			url:    "/update/gauge/testik/123.123",
			want: want{
				code:        http.StatusOK,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "positive test update counter",
			method: http.MethodPost,
			url:    "/update/counter/testik/123",
			want: want{
				code:        http.StatusOK,
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "test error metric type",
			method: http.MethodPost,
			url:    "/update/counterrrr/testik/123",
			want: want{
				code:        http.StatusBadRequest,
				response:    "invalid metric type `counterrrr`\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "test error url",
			method: http.MethodPost,
			url:    "/update/gauge",
			want: want{
				code:        http.StatusNotFound,
				response:    "some of the request elements are missing\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "test error method",
			method: http.MethodGet,
			url:    "/update/gauge/abacaba/123.123",
			want: want{
				code:        http.StatusMethodNotAllowed,
				response:    "only POST method allowed\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "test error metric value",
			method: http.MethodPost,
			url:    "/update/counter/testik/xellar",
			want: want{
				code:        http.StatusBadRequest,
				response:    "invalid metric value type: strconv.ParseInt: parsing \"xellar\": invalid syntax\n",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()
			s.UpdateHandler(w, request)

			res := w.Result()

			assert.Equal(t, res.StatusCode, tt.want.code)

			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.Equal(t, tt.want.response, string(resBody))
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
