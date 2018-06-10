package v1

import (
	"authentication-server/app/domains/entity"
	"authentication-server/app/controllers"
	"github.com/satori/go.uuid"
	"github.com/revel/revel"
	"gopkg.in/go-playground/validator.v9"
	"authentication-server/app/domains/service"
)

type UserController struct {
	controllers.BaseApiController
	validate *validator.Validate
}

var validate = validator.New()
var useService service.UserService

func (c UserController) PostUser(users entity.Users) revel.Result {

	users.Uuid = uuid.Must(uuid.NewV4()).String()
	users.Password = useService.BcryptPw(users.Password)

	err := validate.Struct(users)
	if err != nil {
		return c.BadRequest("001")
	}

	userData := useService.GetUserByEmail(users.Email).Response
	if userData == "server error" {
		return c.InternalServer("002")
	}

	if userData != nil {
		return c.UnprocessableEntity("003")
	}

	if useService.InsertUser(users).Response == nil {
		return c.InternalServer("004")
	}

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	return c.RenderJSON(success)
}
