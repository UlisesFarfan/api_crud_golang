package user_controller_test

import (
	"fmt"
	"testing"
	"time"

	m "github.com/api-rest-go/models"
	user_repository "github.com/api-rest-go/repository/user.repository"
)

func TestCreate(t *testing.T) {

	user := m.User{
		Name:      "Ulises",
		Email:     "ulises@gmail.com",
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err := user_repository.Create(user)

	if err != nil {
		fmt.Println(err)
		t.Error("Create test failed.")
		t.Fail()
	} else {
		t.Log("Create test was successfuly.")
	}
}

func TestRead(t *testing.T) {

	users, err := user_repository.Read()

	if err != nil {
		t.Error("Something went worng.")
		t.Fail()
	}

	if len(users) == 0 {
		t.Error("The request not has retorned data.")
		t.Fail()
	} else {
		t.Log("Road test was soccessfuly.")
	}

}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
