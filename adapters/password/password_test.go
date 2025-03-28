package password

import (
	"testing"

	"github.com/ville-koskela/go-ldap-server/test"
)

func TestHashPassword(t *testing.T) {
	t.Run("should hash password successfully", func(t *testing.T) {
		password := "SecurePassword123"

		hash, err := PasswordTool.HashPassword(password)

		test.Assert(t, nil, err, "HashPassword failed")
		test.Assert(t, true, hash != "", "Hash should not be empty")
		test.Assert(t, true, hash != password, "Hash should not equal original password")
	})

	t.Run("should hash empty password", func(t *testing.T) {
		password := ""

		hash, err := PasswordTool.HashPassword(password)

		test.Assert(t, nil, err, "HashPassword failed with empty password")
		test.Assert(t, true, hash != "", "Hash should not be empty even with empty password")
	})
}

func TestComparePassword(t *testing.T) {
	t.Run("should return true for matching password", func(t *testing.T) {
		password := "SecurePassword123"
		hash, err := PasswordTool.HashPassword(password)

		test.Assert(t, nil, err, "HashPassword failed")

		match := PasswordTool.ComparePassword(hash, password)

		test.Assert(t, true, match, "Password comparison failed")
	})

	t.Run("should return false for non-matching password", func(t *testing.T) {
		password := "SecurePassword123"
		wrongPassword := "WrongPassword123"
		hash, err := PasswordTool.HashPassword(password)

		test.Assert(t, nil, err, "HashPassword failed")

		match := PasswordTool.ComparePassword(hash, wrongPassword)

		test.Assert(t, false, match, "Password comparison should return false for non-matching password")
	})

	t.Run("should return false for empty password", func(t *testing.T) {
		password := "SecurePassword123"
		hash, err := PasswordTool.HashPassword(password)

		test.Assert(t, nil, err, "HashPassword failed")

		match := PasswordTool.ComparePassword(hash, "")

		test.Assert(t, false, match, "Password comparison should return false for empty password")
	})

	t.Run("should return false for invalid hash", func(t *testing.T) {
		password := "SecurePassword123"

		match := PasswordTool.ComparePassword("invalid_hash", password)

		test.Assert(t, false, match, "Password comparison should return false for invalid hash")
	})
}
