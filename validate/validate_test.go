package validate_test

import (
	"fmt"
	"testing"

	"github.com/schrius/password-validator/validate"
	"gopkg.in/stretchr/testify.v1/assert"
)

func TestValidLength(t *testing.T) {
	testCases := []struct {
		name string
		password string
		expected bool
	} {
		{
			"Too short password",
			"123",
			false,
		},
		{
			"Too long password",
			"toolonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong"
			false,
		},
		{
			"Valid password length",
			"validtestpasswordlength",
			true,
		},
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testCase[%d] %s", i + 1, testCase.name), func(t *testing.T) {
			assert.True(t, testCase.expected, validate.ValidLength(testCase.password))
		})
	}
}

func TestValidLetter(t *testing.T) {
	testCases := []struct {
		name string
		password string
		expected bool
	} {
		{
			{
				"Valid Letter",
				"123",
				true,
			},
			{
				"Invalid Letter",
				"\u200b",
				false,
			},
		}
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testCase[%d] %s", i + 1, testCase.name), func(t *testing.T) {
			assert.True(t, testCase.expected, validate.ValidLetter(testCase.password))
		})
	}
}

func TestIsWeakPassword(t *testing.T) {
	testCases := []struct {
		name string
		password string
		path string
		expected bool
	} {
		{
			{
				"Too weak password",
				"password",
				"weak_password_list.txt"
				true,
			},
			{
				"Too weak password",
				"123456789",
				"weak_password_list.txt"
				true,
			},
		}
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testCase[%d] %s", i + 1, testCase.name), func(t *testing.T) {
			assert.True(t, testCase.expected, validate.IsWeakPassword(testCase.password, testCase.path))
		})
	}
}

func TestLoadWeakPasswordList(t *testing.T) {
	weakList := validate.LoadWeakPasswordList("weak_password_list.txt")
	assert.NotNil(t, weakList)
	assert.IsTypef(t, "map[string]bool", weakList)
}

func TestValidate(t *testing.T) {
	weakList := map[string]bool {
		"password": true
		"password1": true
	}
	testCases := []struct {
		name string
		password string
		expected string
	} {
		{
			{
				"Too weak password",
				"password",
				"password -> Error: Too Common"
			},
			{
				"Too weak password",
				"password1",
				"password1 -> Error: Too Common"
			},
		}
	}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testCase[%d] %s", i + 1, testCase.name), func(t *testing.T) {
			assert.EqualError(t, validate.Validate(testCase.password, weakList), testCase.expected)
		})
	}
}
