package helper

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateCode() string {
	u := uuid.New()
	return strings.Replace(u.String(), "-", "", -1)[:8]
}
