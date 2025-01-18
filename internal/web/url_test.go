package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/chenmuyao/url_shortener/internal/service"
	urlsvcmock "github.com/chenmuyao/url_shortener/internal/service/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccessLogAPI(t *testing.T) {
	testCases := []struct {
		Name string

		mock func(ctrl *gomock.Controller) service.UrlShortenerSvc

		// Inputs
		reqBuilder func(t *testing.T) *http.Request

		// Outputs
		wantCode int
		wantRes  ShortUrlRes
	}{
		{
			Name: "test ok",
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{
"url": "http://vinchent.xyz"
}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			mock: func(ctrl *gomock.Controller) service.UrlShortenerSvc {
				urlsvc := urlsvcmock.NewMockUrlShortenerSvc(ctrl)
				urlsvc.EXPECT().Shorten(gomock.Any(), "http://vinchent.xyz").Return("sdf123", nil)
				return urlsvc
			},
			wantCode: http.StatusOK,
			wantRes: ShortUrlRes{
				URL: "http://localhost:3000/sdf123",
			},
		},
		{
			Name: "wrong json format",
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{
"url": "http://vinchent.xyz",
}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			mock: func(ctrl *gomock.Controller) service.UrlShortenerSvc {
				urlsvc := urlsvcmock.NewMockUrlShortenerSvc(ctrl)
				return urlsvc
			},
			wantCode: http.StatusBadRequest,
		},
		{
			Name: "wrong url format",
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{
"url": "vinchent.xyz"
}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			mock: func(ctrl *gomock.Controller) service.UrlShortenerSvc {
				urlsvc := urlsvcmock.NewMockUrlShortenerSvc(ctrl)
				return urlsvc
			},
			wantCode: http.StatusBadRequest,
		},
		{
			Name: "internal error",
			reqBuilder: func(t *testing.T) *http.Request {
				req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{
"url": "http://vinchent.xyz"
}`)))
				req.Header.Set("Content-Type", "application/json")
				assert.NoError(t, err)
				return req
			},
			mock: func(ctrl *gomock.Controller) service.UrlShortenerSvc {
				urlsvc := urlsvcmock.NewMockUrlShortenerSvc(ctrl)
				urlsvc.EXPECT().
					Shorten(gomock.Any(), "http://vinchent.xyz").
					Return("", errors.New("some error"))
				return urlsvc
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			url := tc.mock(ctrl)
			v := validator.New(validator.WithRequiredStructEnabled())
			// create the test server
			hdl := NewUrlShortenerHdl(v, url)

			server := fiber.New()
			hdl.RegisterHandlers(server)

			req := tc.reqBuilder(t)

			resp, _ := server.Test(req)
			t.Log(resp)

			assert.Equal(t, tc.wantCode, resp.StatusCode)
			var res ShortUrlRes
			// Only check if there is a wanted result
			if tc.wantRes != res {
				err := json.NewDecoder(resp.Body).Decode(&res)
				assert.NoError(t, err)
				assert.Equal(t, tc.wantRes, res)
			}
		})
	}
}
