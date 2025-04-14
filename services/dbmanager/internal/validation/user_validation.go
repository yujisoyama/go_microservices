package validation

import (
	"errors"

	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
)

func ValidateUpsertUserRequest(req *dbmanager.UpsertUserRequest) error {
	if len(req.User.Email) == 0 {
		return errors.New("email is required")
	}
	return nil
}
