package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FetchParam struct {
	Limit    uint64
	CursorID uint64
}

type StandarResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FetchParamsPaginationFromRequest(c *fiber.Ctx) FetchParam {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	if limit == 0 {
		limit = 10
	}

	cursorStr := c.Query("cursor")
	cursor, _ := strconv.Atoi(cursorStr)
	if cursor == 0 {
		cursor = 1
	}

	fetchParam := FetchParam{
		Limit:    uint64(limit),
		CursorID: uint64(cursor),
	}

	return fetchParam
}
