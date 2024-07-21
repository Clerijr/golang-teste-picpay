package user_test

import (
	"errors"
	"testing"

	"github.com/clerijr/teste-picpay-go/entities/user"
)

func TestUser_NewUser(t *testing.T) {
	type testcase struct {
		test        string
		newUser     *user.NewUser
		expectedErr error
	}

	testCases := []testcase{
		{
			test: "Invalid name",
			newUser: &user.NewUser{
				Name:     "",
				Lastname: "",
				UType:    "fisica",
				Document: "12312312312",
				Email:    "john@doe.com",
				Password: "123123",
			},
			expectedErr: user.ErrInvalidName,
		},
		{
			test: "Invalid user type",
			newUser: &user.NewUser{
				Name:     "John",
				Lastname: "",
				UType:    "",
				Document: "12312312312",
				Email:    "john@doe.com",
				Password: "123123",
			},
			expectedErr: user.ErrInvalidUserType,
		},
		{
			test: "Invalid document",
			newUser: &user.NewUser{
				Name:     "John",
				Lastname: "Doe",
				UType:    "fisica",
				Document: "",
				Email:    "john@doe.com",
				Password: "123123",
			},
			expectedErr: user.ErrInvalidDocument,
		},
		{
			test: "Invalid email",
			newUser: &user.NewUser{
				Name:     "John",
				Lastname: "",
				UType:    "fisica",
				Document: "12312312312",
				Email:    "",
				Password: "123123",
			},
			expectedErr: user.ErrInvalidEmail,
		},
		{
			test: "Invalid Password",
			newUser: &user.NewUser{
				Name:     "John",
				Lastname: "",
				UType:    "fisica",
				Document: "12312312312",
				Email:    "john@doe.com",
				Password: "",
			},
			expectedErr: user.ErrInvalidPassword,
		},
		{
			test: "Success",
			newUser: &user.NewUser{
				Name:     "John",
				Lastname: "Doe",
				UType:    "fisica",
				Document: "12312312312",
				Email:    "john@doe.com",
				Password: "123123",
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := user.New(*tc.newUser)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
