package helper

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetUserDetailFromContext(ctx echo.Context) (*int, *int, error) {
	userID := ctx.Get("user_id")
	conv, err := strconv.Atoi(fmt.Sprint(userID))
	if err != nil {
		return nil, nil, err
	}

	level := ctx.Get("level")
	convert, err := strconv.Atoi(fmt.Sprint(level))
	if err != nil {
		return nil, nil, err
	}
	return &conv, &convert, nil
}
