package controllers

import (
	"fmt"

	"github.com/MiniProject/go-oauth2/internal/config"
	"github.com/MiniProject/go-oauth2/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	cfg     *config.Config
	Service usecase.UserService
}

func New(UserService usecase.UserService, cfg *config.Config, h *fiber.App) {
	r := &UserController{
		Service: UserService,
		cfg:     cfg,
	}

	h.Get("/login", r.GoogleLogin)
	h.Get("/google_callback", r.GoogleFeedBack)
}

func (u *UserController) GoogleLogin(c *fiber.Ctx) error {
	link := u.Service.GoogleLogins()

	c.Status(fiber.StatusSeeOther)
	c.Redirect(link)
	return c.SendString("JADI GAESS")
}

func (u *UserController) GoogleFeedBack(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != u.cfg.AuthCodeUrl {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	resp, err := u.Service.GoogleLogin(code)

	if err != nil {
		fmt.Println(err)
		return c.SendString("code dosn't catch correctly")
	}

	return c.SendString(resp)

}
