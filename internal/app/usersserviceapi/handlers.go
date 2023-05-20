package userserviceapi

import (
	"context"
	"go-integration-tests/internal/database/repository/userrepository"
	"go-integration-tests/pkg/users"
)

func GetUserRepositoryToResponse(
	in *userrepository.User) *users.GetUserResponse {
	out := &users.GetUserResponse{}

	out.User = &users.User{
		Id:           in.Id,
		Name:         in.Name,
		Username:     in.Username,
		PasswordHash: in.PasswordHash,
	}
	return out
}

func ListUsersRepositoryToResponse(
	in []*userrepository.User) *users.ListUsersResponse {
	out := &users.ListUsersResponse{}

	out.Users = make([]*users.User, len(in))
	for i, v := range in {
		out.Users[i] = &users.User{
			Id:           v.Id,
			Name:         v.Name,
			Username:     v.Username,
			PasswordHash: v.PasswordHash,
		}
	}

	return out
}
func (s *ServiceImpl) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	user, err := s.serviceRepository.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return GetUserRepositoryToResponse(user), nil
}

func (s *ServiceImpl) ListUsers(ctx context.Context, req *users.ListUsersRequest) (*users.ListUsersResponse, error) {
	listUsers, err := s.serviceRepository.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return ListUsersRepositoryToResponse(listUsers), nil
}
