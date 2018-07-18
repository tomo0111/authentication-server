package controllers

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/tomo0111/grant-n-z/app/controllers/v1"
	"github.com/tomo0111/grant-n-z/app/domains/entity"
	"github.com/tomo0111/grant-n-z/tests"
	"github.com/tomo0111/grant-n-z/app"
)

type UserControllerTest struct {
	tests.AppTest
}

var userController = v1.UserController{}

// Start MySQL
func (t UserControllerTest) Before() {

	_, mock, err := sqlmock.NewWithDSN("sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}

	db, err := gorm.Open("sqlmock", "sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}
	defer db.Close()

	db.DB()
	app.Db = db


	rs := sqlmock.NewRows([]string{"count"}).FromCSVString("1")
	mock.ExpectQuery(`SELECT * FROM "users"`).
		WithArgs(12345).
		WillReturnRows(rs)
}

// Stop MySQL
func (t UserControllerTest) After() {
	app.Db.Close()
}

func (t UserControllerTest) TestPostUserOk() {
	app.Db.CreateTable(&entity.Users{})

	users := entity.Users{
		Username: "test",
		Email: "test@gmail.com",
		Password: "testtest",
	}

	var response = userController.PostUser(users)

	success := map[string]string {
		"message": "user creation succeeded.",
	}

	t.AssertEqual(success, response)
}