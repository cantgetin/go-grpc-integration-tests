package tests

import (
	"context"
	"go-integration-tests/pkg/users"
)

func (s *IntegrationSuite) TestUserGetUsers() {
	client := users.NewUserServiceClient(s.grpcConn)

	s.Run("valid requests", func() {
		tt := []struct {
			name string
			in   *users.GetUserRequest
			out  *users.GetUserResponse
		}{
			{
				name: "User that should be in DB #1",
				in: &users.GetUserRequest{
					Id: 1,
				},
				out: &users.GetUserResponse{
					User: &users.User{
						Id:           1,
						Name:         "John",
						Username:     "john12",
						PasswordHash: "2fbe00f6a2f5ca35a3d49adbdc33ce23",
					},
				},
			},
			{
				name: "User that should be in DB #2",
				in: &users.GetUserRequest{
					Id: 2,
				},
				out: &users.GetUserResponse{
					User: &users.User{
						Id:           2,
						Name:         "Ross",
						Username:     "frosty",
						PasswordHash: "11340a040bf55d21ab9e0cb1323ffa7c",
					},
				},
			},
		}

		for _, tc := range tt {
			s.Run(tc.name, func() {
				out, err := client.GetUser(context.Background(), tc.in)
				s.Require().NoError(err)
				s.Require().NotNil(out)

				want := tc.out.User
				got := out.User

				s.Assert().Equal(want.Id, got.Id)
				s.Assert().Equal(want.Name, got.Name)
				s.Assert().Equal(want.Username, got.Username)
				s.Assert().Equal(want.PasswordHash, got.PasswordHash)
			})
		}
	})
}

func (s *IntegrationSuite) TestUserListUsers() {
	client := users.NewUserServiceClient(s.grpcConn)

	s.Run("valid requests", func() {
		tt := []struct {
			name string
			in   *users.ListUsersRequest
			out  *users.ListUsersResponse
		}{
			{
				name: "Users that should be in DB",
				in:   &users.ListUsersRequest{},
				out: &users.ListUsersResponse{
					Users: []*users.User{
						{
							Id:           1,
							Name:         "John",
							Username:     "john12",
							PasswordHash: "2fbe00f6a2f5ca35a3d49adbdc33ce23",
						},
						{
							Id:           2,
							Name:         "Ross",
							Username:     "frosty",
							PasswordHash: "11340a040bf55d21ab9e0cb1323ffa7c",
						},
					},
				},
			},
		}

		for _, tc := range tt {
			s.Run(tc.name, func() {
				out, err := client.ListUsers(context.Background(), tc.in)
				s.Require().NoError(err)
				s.Require().NotNil(out)

				for i, b := range out.Users {
					s.Assert().Equal(tc.out.Users[i].Id, b.Id)
					s.Assert().Equal(tc.out.Users[i].Name, b.Name)
					s.Assert().Equal(tc.out.Users[i].Username, b.Username)
					s.Assert().Equal(tc.out.Users[i].PasswordHash, b.PasswordHash)
				}
			})
		}
	})
}
