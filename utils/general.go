package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUniqueFileName(filename string) string {
	fileName := fmt.Sprintf("id_%s_%s", uuid.New().String()[:8], filename)
	fullPath := fmt.Sprintf("data/%s", fileName)
	return fullPath
}
