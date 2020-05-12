package util

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateUUID() string {

	uuid.New()
	UUID := strings.ReplaceAll(uuid.New().String(), "-", "")
	return UUID
}
