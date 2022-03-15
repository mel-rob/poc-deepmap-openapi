package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type UserAccounts struct {
}

func NewUserAccount() *UserAccounts {
	return &UserAccounts{}
}

func (u *UserAccounts) DeleteAccount(ctx echo.Context, uuid string) error {
	return nil
}

func (u *UserAccounts) GetUserAccount(ctx echo.Context, uuid string) error {
	return nil
}

func (u *UserAccounts) ModifyUserAccount(ctx echo.Context, uuid string) error {
	return nil
}

func (u *UserAccounts) VerifyAccount(ctx echo.Context, uuid string) error {
	return nil
}

// Email has a type mismatch in model and openapi component
func (u *UserAccounts) GetUserAccounts(ctx echo.Context, params GetUserAccountsParams) error {
	fmt.Println("reached GetUserAccounts")
	uuid := "16673520-0b73-4425-aa33-3bee939369f8"
	// emails := &[]string{
	// 	"melrob@supercool.com",
	// 	"openapi@swagger.com",
	// }
	firstname := "mel"
	lastname := "rob"
	result := UserAccount{
		UUID: &uuid,
		// Emails:    emails,
		FirstName: &firstname,
		LastName:  &lastname,
	}
	ctx.JSON(200, result)

	return nil
}

func (u *UserAccounts) CreateAccount(ctx echo.Context) error {
	return nil
}
