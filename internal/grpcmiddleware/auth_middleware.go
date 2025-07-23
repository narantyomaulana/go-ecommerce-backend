package grpcmiddleware

import (
	"context"

	jwtentity "github.com/narantyomaulana/go-grpc-ercommerce-be/internal/entity/jwt"
	"github.com/narantyomaulana/go-grpc-ercommerce-be/internal/utils"
	gocache "github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
)

type authMiddleware struct {
	cacheService *gocache.Cache
}

func (am *authMiddleware) Middleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if info.FullMethod == "/auth.AuthService/Register" || info.FullMethod == "/auth.AuthService/Login" {
		return handler(ctx, req)
	}

	tokenStr, err := jwtentity.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	_, ok := am.cacheService.Get(tokenStr)
	if ok {
		return nil, utils.UnauthenticatedResponse()
	}

	claims, err := jwtentity.GetClaimsFromToken(tokenStr)
	if err != nil {
		return nil, err
	}

	ctx = claims.SetToContext(ctx)

	res, err := handler(ctx, req)
	return res, err
}

func NewAuthMiddleware(cacheService *gocache.Cache) *authMiddleware {
	return &authMiddleware{
		cacheService: cacheService,
	}
}
