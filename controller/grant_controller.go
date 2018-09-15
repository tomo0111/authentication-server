package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/tomoyane/grant-n-z/di"
	"github.com/tomoyane/grant-n-z/domain"
)

func GrantToken(c echo.Context) (err error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.Unauthorized("014"))
	}

	resultMap, result := di.ProviderTokenService.ParseJwt(token)
	if !result {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.Unauthorized("015"))
	}

	user := di.ProviderUserService.GetUserByUuid(resultMap["username"], resultMap["user_uuid"])
	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.InternalServerError("016"))
	}

	if len(user.Email) == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, domain.Unauthorized("017"))
	}

	role := di.ProviderRoleService.GetRoleByUserUuid(user.Uuid.String())
	if role == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.InternalServerError("018"))
	}

	if len(role.UserUuid) == 0 {
		return echo.NewHTTPError(http.StatusForbidden, domain.Forbidden("019"))
	}

	if role.Type != "user" && role.Type != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, domain.Forbidden("020"))
	}

	success := map[string]bool {
		"authority": true,
	}

	return c.JSON(http.StatusOK, success)
}
