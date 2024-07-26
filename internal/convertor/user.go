package convertor

import (
	repository "github.com/lookandhate/course_auth/internal/repository/model"
	"github.com/lookandhate/course_auth/internal/service/model"
	"github.com/lookandhate/course_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateUserFromProto converts data from protobuf generated struct to service CreateUserModel.
func CreateUserFromProto(user *auth_v1.CreateRequest) *model.CreateUserModel {
	if user == nil {
		return nil
	}

	return &model.CreateUserModel{
		Name:            user.GetName(),
		Email:           user.GetEmail(),
		Password:        user.GetPassword(),
		PasswordConfirm: user.GetPasswordConfirm(),
		Role:            model.UserRole(user.GetRole()),
	}
}

// CreateUserModelToRepo converts model from service layer to repo layer.
func CreateUserModelToRepo(user *model.CreateUserModel) *repository.CreateUserModel {
	if user == nil {
		return nil
	}

	return &repository.CreateUserModel{
		Name:  user.Name,
		Email: user.Email,
		// TODO hashing
		Password: user.Password,
		Role:     repository.UserRole(user.Role),
	}
}

// UserModelToGetResponseProto converts from UserModel to proto response.
func UserModelToGetResponseProto(user *model.UserModel) *auth_v1.GetResponse {
	if user == nil {
		return nil
	}

	return &auth_v1.GetResponse{
		Id:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

// UserUpdateFromProto converts from proto UpdateRequest to UpdateUserModel.
func UserUpdateFromProto(user *auth_v1.UpdateRequest) *model.UpdateUserModel {
	if user == nil {
		return nil
	}
	var name, email, password *string
	if user.Name != nil {
		name = &user.Name.Value
	}
	if user.Email != nil {
		email = &user.Email.Value
	}
	if user.Password != nil {
		password = &user.Password.Value
	}

	return &model.UpdateUserModel{
		Name:     name,
		Email:    email,
		Role:     int(user.GetRole()),
		Password: password,
		ID:       int(user.GetId()),
	}
}
