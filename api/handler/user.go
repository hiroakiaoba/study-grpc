package handler

import (
	"context"
	"log"

	srv "api/gen/service"
	"api/model"
	"api/repository"
	"api/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var InternalErrMsg = "Internal Server Error"
var InvalidInputErrMsg = "入力値を確認してください"

type UserHandler struct {
	userRepo repository.IUserRepository
	auth     util.Auther
}

func NewUserHandler(userRepo repository.IUserRepository, auth util.Auther) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
		auth:     auth,
	}
}

func (h *UserHandler) SignUp(c context.Context, r *srv.SignUpRequest) (*srv.SignUpResponse, error) {
	log.Println("reveived sign up request!!")
	// 入力チェック
	if r.LoginName == "" || r.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "LoginNameとPasswordを入力してくだい")
	}

	// バリデーション
	// 名前の重複確認
	user, err := h.userRepo.FindByLoginName(r.LoginName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, InternalErrMsg)
	}
	if user != nil {
		return nil, status.Errorf(codes.InvalidArgument, "その名前のユーザーはすでに存在しています")
	}
	// 名前とパスワード文字数などバリデーションは一旦省略

	// userの保存(一旦パスワードは平で保存)
	newUser := &model.User{
		LoginName: r.LoginName,
		Password:  r.Password,
	}
	err = h.userRepo.Create(newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, InternalErrMsg)
	}

	// token作成
	token, err := h.auth.GenToken(newUser.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, InternalErrMsg)
	}

	return &srv.SignUpResponse{
		Token: token,
	}, nil
}

func (h *UserHandler) SignIn(c context.Context, r *srv.SingInRequest) (*srv.SingInResponse, error) {
	// 入力チェック
	if r.LoginName == "" || r.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "LoginNameとPasswordを入力してくだい")
	}

	// login_nameでuserを取得
	user, err := h.userRepo.FindByLoginName(r.LoginName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, InternalErrMsg)
	}
	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, InvalidInputErrMsg)
	}

	// パスワードが正しいか確認
	if user.Password != r.Password {
		return nil, status.Errorf(codes.InvalidArgument, InvalidInputErrMsg)
	}

	// token作成
	token, err := h.auth.GenToken(user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, InternalErrMsg)
	}

	return &srv.SingInResponse{
		Token: token,
	}, nil
}

func (h *UserHandler) List(c context.Context, r *srv.ListRequest) (*srv.ListResponse, error) {
	users, err := h.userRepo.List()
	if err != nil {
		return nil, status.Errorf(codes.Internal, InternalErrMsg)
	}

	return &srv.ListResponse{
		Users: SerializeUsers(users),
	}, nil
}

func SerializeUsers(users []*model.User) []*srv.User {
	serializedUsers := make([]*srv.User, 0)
	for _, u := range users {
		user := &srv.User{
			Id:        u.ID,
			LoginName: u.LoginName,
		}
		serializedUsers = append(serializedUsers, user)
	}
	return serializedUsers
}
