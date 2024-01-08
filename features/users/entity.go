package users

import "github.com/labstack/echo/v4"

type Contoh struct {
	Name  string `json:"nama"`
	Age   int
	Email string
	Skill Skill
}
type Skill struct {
	NamaSkill string
}


type Handler interface {
	AddUser() echo.HandlerFunc
}

type Repository interface{
	AddUser(newUser Contoh) (Contoh, error)
}

type Service interface {
	AddUser(newUser Contoh) (Contoh, error)
}
