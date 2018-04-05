package endpoint

import (
	"context"
	f "fmt"

	svc "prjResto/gudang/server"
)

func (ue GudangEndpoint) AddGudangService(ctx context.Context, usr svc.Gudang) error {
	_, err := ue.AddGudangEndpoint(ctx, usr)
	return err
}

func (ue GudangEndpoint) UpdateGudangService(ctx context.Context, usr svc.Gudang) error {
	_, err := ue.UpdateGudangEndpoint(ctx, usr)
	if err != nil {
		f.Println("error pada endpoint:", err)
	}
	return err
}
