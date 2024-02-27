package access

import (
	"context"
	descAccess "github.com/Murat993/auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (i Implementation) Check(ctx context.Context, req *descAccess.CheckRequest) (*emptypb.Empty, error) {
	empty, err := i.accessService.Check(ctx, req)

	if err != nil {
		return nil, err
	}

	log.Printf("GetAccessToken")

	return empty, nil
}
