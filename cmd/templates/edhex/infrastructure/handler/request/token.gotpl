package request

import "github.com/labstack/echo/v4"

func GetUserID(c echo.Context) uint {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return 0
	}

	return userID
}
