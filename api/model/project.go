package model

import "time"

type Project struct {
	ID        int32
	Title     string
	UserIDs   []int32
	AuthorID  int32
	CreatedAt time.Time
}
