package web

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/chenmuyao/url_shortener/config"
	"github.com/chenmuyao/url_shortener/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FullUrlReq struct {
	URL string `json:"url" validate:"required,url"`
}

type ShortUrlRes struct {
	URL string `json:"url"`
}

type UrlShortenerHdl struct {
	validate *validator.Validate
	svc      service.UrlShortenerSvc
	baseURL  string
}

func NewUrlShortenerHdl(v *validator.Validate, svc service.UrlShortenerSvc) *UrlShortenerHdl {
	return &UrlShortenerHdl{validate: v, svc: svc, baseURL: config.App.BaseURL}
}

func (u *UrlShortenerHdl) RegisterHandlers(s *fiber.App) {
	s.Post("/", u.SetUrl)
	s.Get("/:short", u.GetFull)
	s.Get("/:short/count", u.GetCount)
}

func (u *UrlShortenerHdl) SetUrl(c *fiber.Ctx) error {
	c.Accepts("application/json")

	// Check the input
	var req FullUrlReq
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		slog.Error(
			"Failed to unmarshal",
			slog.Any("ip", c.IP()),
			slog.Any("body", c.Body()),
			slog.Any("err", err),
		)
		return c.SendStatus(http.StatusBadRequest)
	}

	err = u.validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			slog.Error(
				"Input error",
				slog.Any("ip", c.IP()),
				slog.Any("field", err.Field()),
				slog.Any("tag", err.Tag()),
				slog.Any("value", err.Value()),
			)
		}
		return c.SendStatus(http.StatusBadRequest)
	}

	// Do the biz
	short, err := u.svc.Shorten(c.Context(), req.URL)
	if err != nil {
		slog.Error("Failed to shorten the url", slog.Any("url", req.URL), slog.Any("err", err))
		return c.SendStatus(http.StatusInternalServerError)
	}

	// Return the result
	return c.Status(http.StatusOK).JSON(ShortUrlRes{
		URL: fmt.Sprintf("%s/%s", u.baseURL, short),
	})
}

func (u *UrlShortenerHdl) GetFull(c *fiber.Ctx) error {
	short := c.Params("short")

	url, err := u.svc.GetURL(c.Context(), short)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	slog.Debug("Got full url", slog.Any("url", url.Url))

	return c.Redirect(url.Url)
}

func (u *UrlShortenerHdl) GetCount(c *fiber.Ctx) error {
	short := c.Params("short")

	url, err := u.svc.GetURL(c.Context(), short)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	slog.Debug("Got full url", slog.Any("url", url.Url), slog.Any("count", url.Count))

	return c.Status(http.StatusOK).JSON(url)
}
