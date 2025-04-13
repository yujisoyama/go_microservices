package dbmanagerdto

import (
	"github.com/yujisoyama/go_microservices/pkg/pb/dbmanager"
	"github.com/yujisoyama/go_microservices/services/dbmanager/internal/entity"
)

func InputDtoUserToEntity(req *dbmanager.UpsertUserRequest) *entity.UserEntity {
	return &entity.UserEntity{
		OauthId:       req.User.OauthId,
		OauthType:     req.User.OauthType,
		Email:         req.User.Email,
		VerifiedEmail: req.User.VerifiedEmail,
		FirstName:     req.User.FirstName,
		LastName:      req.User.LastName,
		Picture:       req.User.Picture,
	}
}
