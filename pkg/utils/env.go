package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		err := fmt.Sprintf("Env %s not found", key)
		panic(errors.New(err))
	}

	return strings.TrimSpace(value)
}
