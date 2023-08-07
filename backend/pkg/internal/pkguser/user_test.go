package pkguser_test

import (
	"testing"

	"github.com/fahmifan/commurz/pkg/internal/pkguser"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	const testPassword = "test1234"

	type input struct {
		email string
	}

	type expect struct {
		shouldErr bool
	}

	type testCase struct {
		input  input
		expect expect
	}

	testCases := []testCase{
		{
			input:  input{email: "ha@email.com"},
			expect: expect{shouldErr: false},
		},
		{
			input:  input{email: "haemail.com"},
			expect: expect{shouldErr: true},
		},
		{
			input:  input{email: "@email.com"},
			expect: expect{shouldErr: true},
		},
	}

	for _, tc := range testCases {
		_, err := pkguser.NewUser(tc.input.email, testPassword)
		if tc.expect.shouldErr {
			require.Error(t, err)
			continue
		}

		require.NoError(t, err)
	}
}
