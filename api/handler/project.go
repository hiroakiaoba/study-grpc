package handler

import (
	srv "api/gen/service"
	"context"
)

type ProjectHandler struct {
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

func (h *ProjectHandler) Create(c context.Context, r *srv.ProjectCreateRequest) (*srv.ProjectCreateResponse, error) {
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
