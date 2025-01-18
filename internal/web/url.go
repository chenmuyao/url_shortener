package web

import "github.com/gofiber/fiber/v2"

type UrlShortenerHdl struct{}

func NewUrlShortenerHdl() *UrlShortenerHdl {
	return &UrlShortenerHdl{}
}

func (u *UrlShortenerHdl) RegisterHandlers(s *fiber.App) {
	s.Post("/", u.SetUrl)
	s.Get("/", u.GetShort)
}

func (u *UrlShortenerHdl) SetUrl(c *fiber.Ctx) error {
	c.SendString("POST")
	return nil
}

func (u *UrlShortenerHdl) GetShort(c *fiber.Ctx) error {
	c.SendString("GET")
	return nil
}
