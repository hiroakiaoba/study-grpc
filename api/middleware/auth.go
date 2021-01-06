package middleware

import (
	"api/repository"
	"api/util"
	"context"
	"log"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	log.Println("認証Middleware通ります！！")
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		log.Println("failed to load auth header err:", err)
		return nil, status.Errorf(codes.Unauthenticated, "Tokenを送信してください")
	}

	userID, err := a.authUtil.Verify(token)
	if err != nil {
		log.Println("failed to verify token err:", err)
		return nil, status.Errorf(codes.Unauthenticated, "Tokenが不正です")
	}

	newCtx := util.SetUserID(ctx, userID)
	return newCtx, nil
}
