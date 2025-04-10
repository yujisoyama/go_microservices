package validation

import (
	"errors"

	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
)

func ValidateUpsertUserRequest(req *dbmanager.UpsertUserRequest) error {
	if len(req.Name) == 0 {
		return errors.New("name is required")
	}
	if len(req.Email) == 0 {
		return errors.New("email is required")
	}
	return nil
}
