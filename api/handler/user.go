package handler

import (
	"context"
	"encoding/base64"

	srv "api/gen/service"
	"api/model"
	"api/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userRepo repository.IUserRepository
}

func NewUserHandler(userRepo repository.IUserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h *UserHandler) SignUp(c context.Context, r *srv.SignUpRequest) (*srv.SignUpResponse, error) {
	// 入力チェック
	if r.LoginName == "" || r.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "LoginNameとPasswordを入力してくだい")
	}

	// バリデーション
	// 名前の重複確認
	user, err := h.userRepo.FindByLoginName(r.LoginName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	if user != nil {
		return nil, status.Errorf(codes.InvalidArgument, "その名前のユーザーはすでに存在しています")
	}
	// 名前とパスワード文字数などバリデーションは一旦省略

	// userの保存
	newUser := &model.User{
		LoginName: r.LoginName,
		Password:  r.Password,
	}
	err = h.userRepo.Create(newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	// token作成
	// めんどいので一旦base64で <LoginName-Passowrd>
	data := user.LoginName + "-" + user.Password
	enc := base64.StdEncoding.EncodeToString([]byte(data))

	return &srv.SignUpResponse{
		Token: enc,
	}, nil
}

func (h *UserHandler) SingIn(c context.Context, r *srv.SingInRequest) (*srv.SingInResponse, error) {
	// 入力チェック
	if r.LoginName == "" || r.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "LoginNameとPasswordを入力してくだい")
	}

	// login_nameでuserを取得
	// パスワードが正しいか確認
	// token作成

	return nil, nil
}

func (h *UserHandler) List(c context.Context, r *srv.ListRequest) (*srv.ListResponse, error) {
	return nil, nil
}
