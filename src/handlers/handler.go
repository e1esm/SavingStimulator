package handlers

import "github.com/labstack/echo/v4"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) SaveMoney(c echo.Context) error {
	return nil
}

func (h *Handler) SetGoal(c echo.Context) error {
	return nil
}
