package server

import (
	"context"
	"fmt"
)

// 5.
type user struct {
	writer ReadWriter // diambil dari interface di service.go
}

func NewUser(writer ReadWriter) UserService {
	return &user{writer: writer}
}

//Methode pada interface UserService di service.go
func (u *user) AddUserService(ctx context.Context, user User) error {
	fmt.Println("Input data berhasil, silahkan periksa DB Anda, gan!")
	err := u.writer.AddUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *user) UpdateUserService(ctx context.Context, usr User) error {
	err := u.writer.UpdateUser(usr)
	if err != nil {
		return err
	}
	return nil
}
