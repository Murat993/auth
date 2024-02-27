package access

import (
	"context"
	"github.com/Murat993/auth/internal/config/env"
	"github.com/Murat993/auth/internal/utils"
	descCheck "github.com/Murat993/auth/pkg/access_v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"strings"
)

func (s *serverAccess) Check(ctx context.Context, req *descCheck.CheckRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], env.AuthPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], env.AuthPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(os.Getenv(env.AccessTokenSecretKey)))
	if err != nil {
		return nil, errors.New("access token is invalid")
	}

	accessibleMap, err := s.accessRepository.Check(ctx)
	if err != nil {
		return nil, errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[req.GetEndpointAddress()]
	if !ok {
		return &emptypb.Empty{}, nil
	}

	if role == claims.Role {
		return &emptypb.Empty{}, nil
	}

	return nil, errors.New("access denied")
}
