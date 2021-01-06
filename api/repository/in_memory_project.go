package repository

import "api/model"

type InMemoryProjectRepository struct{}

var _ IProjectRepository = (*InMemoryProjectRepository)(nil)

func NewInMemoryProjectRepository() *InMemoryProjectRepository {
	return &InMemoryProjectRepository{}
}

func (r *InMemoryProjectRepository) Create(p *model.Project) error {
	return nil
}
