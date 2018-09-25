package service

import (
	"github.com/tomoyane/grant-n-z/domain/entity"
	"github.com/tomoyane/grant-n-z/domain/repository"
	"time"
	"github.com/satori/go.uuid"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/tomoyane/grant-n-z/infra"
	"github.com/tomoyane/grant-n-z/handler"
	"github.com/labstack/echo"
)

type TokenService struct {
	TokenRepository repository.TokenRepository
	UserRepository repository.UserRepository
}

func (t TokenService) GenerateJwt(username string, userUuid uuid.UUID) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["user_uuid"] = userUuid.String()
	claims["expires"] = time.Now().Add(time.Hour * 365).String()

	signedToken, err := token.SignedString([]byte(infra.Yaml.App.PrivateKey))
	if err != nil {
		handler.ErrorResponse{}.Print(http.StatusInternalServerError, "failed generate jwt", "")
		return ""
	}

	return signedToken
}

func (t TokenService) ParseJwt(token string) (map[string]string, bool) {
	resultMap := map[string]string{}

	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(infra.Yaml.App.PrivateKey), nil
	})

	if err != nil || !parseToken.Valid {
		return resultMap, false
	}

	claims := parseToken.Claims.(jwt.MapClaims)
	if _, ok := claims["username"].(string); !ok {
		return resultMap, false
	}

	if _, ok := claims["user_uuid"].(string); !ok {
		return resultMap, false
	}

	if _, ok := claims["expires"].(string); !ok {
		return resultMap, false
	}

	resultMap["username"] = claims["username"].(string)
	resultMap["user_uuid"] = claims["user_uuid"].(string)
	resultMap["expires"] = claims["expires"].(string)

	return resultMap, true
}

func (t TokenService) GetTokenByUserUuid(userUuid string) *entity.Token {
	return t.TokenRepository.FindByUserUuid(userUuid)
}

func (t TokenService) InsertToken(userUuid uuid.UUID, token string, refreshToken string) *entity.Token {
	data := entity.Token{
		TokenType: "Bearer",
		Token: token,
		RefreshToken: refreshToken,
		UserUuid: userUuid,
	}
	return t.TokenRepository.Save(data)
}

func (t TokenService) VerifyToken(c echo.Context, token string) (*handler.ErrorResponse) {

	if token == "" {
		return handler.Unauthorized("")
	}

	resultMap, result := t.ParseJwt(token)
	if !result {
		return handler.Unauthorized("")
	}

	user := t.UserRepository.FindByUserNameAndUuid(resultMap["username"], resultMap["user_uuid"])
	if user == nil {
		return handler.InternalServerError("")
	}

	if len(user.Email) == 0 {
		return handler.Unauthorized("")
	}

	//role := roleService.GetRoleByUserUuid(user.Uuid.String())
	//if role == nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError(""))
	//}

	//if len(role.UserUuid) == 0 {
	//	return echo.NewHTTPError(http.StatusForbidden, handler.Forbidden("019"))
	//}
	//
	//if role.Type != "user" && role.Type != "admin" {
	//	return echo.NewHTTPError(http.StatusForbidden, handler.Forbidden("020"))
	//}

	return nil
}