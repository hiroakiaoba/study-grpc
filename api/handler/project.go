package handler

import (
	srv "api/gen/service"
	"api/util"
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectHandler struct {
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

func (h *ProjectHandler) Create(c context.Context, r *srv.ProjectCreateRequest) (*srv.ProjectCreateResponse, error) {
	fmt.Println("receivce project create request")

	userID, err := util.GetUserID(c)
	if err != nil {
		log.Println("GetUserID err:", err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	log.Println(strings.Repeat("#", 100))
	log.Println("userID:", userID)
	log.Println("title:", r.Title)

	return nil, nil
}

func (h *ProjectHandler) Delete(context.Context, *srv.ProjectDeleteRequest) (*srv.ProjectDeleteResponse, error) {
	return nil, nil
}

func (h *ProjectHandler) List(context.Context, *srv.ProjectListRequest) (*srv.ProjectListResponse, error) {
	return nil, nil
}

func (h *ProjectHandler) Invite(context.Context, *srv.ProjectInviteRequest) (*srv.ProjectInviteResponse, error) {
	return nil, nil
}
