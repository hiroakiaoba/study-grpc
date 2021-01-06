package middleware

import (
	"api/repository"
	"api/util"
	"context"
	"errors"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type AuthMiddleware struct {
	userRepo repository.IUserRepository
	authUtil util.Auther
}

func NewAuthMiddleware(userRepo repository.IUserRepository, authUtil util.Auther) *AuthMiddleware {
	return &AuthMiddleware{
		userRepo: userRepo,
		authUtil: authUtil,
	}
}

func (a *AuthMiddleware) Authenticate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		log.Println("failed to load auth header err:", err)
		return nil, err
	}

	userID, err := a.authUtil.Verify(token)
	if err != nil {
		log.Println("failed to verify token err:", err)
		return nil, errors.New("unauthorized")
	}

	newCtx := util.SetUserID(ctx, userID)
	return newCtx, nil
}
