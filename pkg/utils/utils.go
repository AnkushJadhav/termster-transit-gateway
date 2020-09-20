package utils

import (
	"github.com/google/uuid"
)

// GenerateUID generates an unique identifier using Googles UUID system
func GenerateUID() string {
	return uuid.New().String()
}
