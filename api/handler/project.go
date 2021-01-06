package handler

import (
	srv "api/gen/service"
	"api/model"
	"api/repository"
	"api/util"
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectHandler struct {
	projectRepo repository.IProjectRepository
}

func NewProjectHandler(projectRepo repository.IProjectRepository) *ProjectHandler {
	return &ProjectHandler{
		projectRepo: projectRepo,
	}
}

func (h *ProjectHandler) Create(c context.Context, r *srv.ProjectCreateRequest) (*srv.ProjectCreateResponse, error) {
	fmt.Println("receivce project create request")

	userID, err := util.GetUserID(c)
	if err != nil {
		log.Println("GetUserID err:", err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	newProject := model.NewProject(r.Title, userID)

	// Createを実装する
	if err := h.projectRepo.Create(newProject); err != nil {
		log.Println("failed to create project err:", err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	// userを取得して渡すようにする
	return &srv.ProjectCreateResponse{
		Project: serializeNewProject(newProject, &model.User{}),
	}, nil
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

func serializeNewProject(p *model.Project, author *model.User) *srv.Project {
	authorUser := &srv.User{
		Id:        author.ID,
		LoginName: author.LoginName,
	}

	return &srv.Project{
		Id:     p.ID,
		Title:  p.Title,
		Users:  []*srv.User{authorUser},
		Author: authorUser,
		CreatedAt: &timestamp.Timestamp{
			Seconds: p.CreatedAt.Unix(),
			Nanos:   int32(p.CreatedAt.Nanosecond()),
		},
	}
}
