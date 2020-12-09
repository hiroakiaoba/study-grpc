package model

type User struct {
	ID        int32
	LoginName string // ユニーク
	Password  string
}
