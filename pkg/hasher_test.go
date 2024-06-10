package pkg_test

import (
	"errors"
	"testing"

	"github.com/clerijr/teste-picpay-go/pkg"
)

func TestHasher_NewHash(t *testing.T) {
	type testcase struct {
		test        string
		password    string
		expectedErr error
	}

	testcases := []testcase{
		{
			test:        "Empty string given",
			password:    "",
			expectedErr: pkg.ErrEmptyString,
		},
		{
			test:        "Success",
			password:    "123123",
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			hasher := pkg.NewHasher()
			_, err := hasher.Hash(tc.password)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
