package util

import (
	"context"
	"fmt"
)

type contextKey string

const userIDContextKey contextKey = "userId"

func SetUserID(c context.Context, userID int32) context.Context {
	return context.WithValue(c, userIDContextKey, userID)
}

func GetUserID(c context.Context) (int32, error) {
	v := c.Value(userIDContextKey)

	userID, ok := v.(int32)
	if !ok {
		return 0, fmt.Errorf("userId not found in context")
	}

	return userID, nil
}
