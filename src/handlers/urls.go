package handlers

import "github.com/labstack/echo/v4"

func (h *Handler) Urls() map[string]map[string]echo.HandlerFunc {
	return map[string]map[string]echo.HandlerFunc{
		"save/:user_id": {
			"POST": h.SaveMoney,
		},
		"goal/:user_id": {
			"POST": h.SetGoal,
		},
	}
}
