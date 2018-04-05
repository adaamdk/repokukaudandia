package endpoint

import (
	"context"
	f "fmt"

	svc "prjResto/user/server"
)

func (ue UserEndpoint) AddUserService(ctx context.Context, usr svc.User) error {
	_, err := ue.AddUserEndpoint(ctx, usr)
	return err
}

func (ue UserEndpoint) UpdateUserService(ctx context.Context, usr svc.User) error {
	_, err := ue.UpdateUserEndpoint(ctx, usr)
	if err != nil {
		f.Println("error pada endpoint:", err)
	}
	return err
}
