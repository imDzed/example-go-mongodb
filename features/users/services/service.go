package services

import "gomongo/features/users"

type UserService struct {
	m users.Repository
}

func New(model users.Repository) users.Service {
	return &UserService{
		m: model,
	}
}


// AddUser implements users.Service.
func (au *UserService) AddUser(newUser users.Contoh) (users.Contoh, error) {
	result, err := au.m.AddUser(newUser)
	if err != nil {
		return users.Contoh{}, err
	}

	return result, nil
}
