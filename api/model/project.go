package model

import "time"

type Project struct {
	ID        int32
	Title     string
	UserIDs   []int32
	AuthorID  int32
	CreatedAt time.Time
}

func NewProject(title string, authorID int32) *Project {
	return &Project{
		Title:     title,
		UserIDs:   []int32{},
		AuthorID:  authorID,
		CreatedAt: time.Now(),
	}
}
