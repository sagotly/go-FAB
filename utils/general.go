package utils

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
)

func GenerateUniqueFileName(filename string) string {
	fileName := fmt.Sprintf("id_%s_%s", uuid.New().String()[:8], filename)
	fullPath := filepath.Join("data", fileName)
	return fullPath
}
