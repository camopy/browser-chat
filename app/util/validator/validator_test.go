package validator_test

import (
	"testing"

	"github.com/camopy/browser-chat/app/util/validator"
	"github.com/stretchr/testify/assert"
)

type formValidatorTestCase struct {
	name     string
	input    interface{}
	expected string
}

var formValidatorTests = []*formValidatorTestCase{
	{
		name: `empty user name`,
		input: struct {
			UserName string `json:"user_name" form:"required"`
		}{},
		expected: "UserName is required",
	},
	{
		name: `empty password`,
		input: struct {
			Password string `json:"password" form:"required"`
		}{
			Password: "",
		},
		expected: "Password is required",
	},
}

func TestToErrResponse(t *testing.T) {
	vr := validator.New()

	for _, tc := range formValidatorTests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := vr.Struct(tc.input)
			if errResp := validator.ToErrResponse(err); errResp == nil || len(errResp.Errors) != 1 {
				assert.Equal(t, tc.expected, "")
			} else {
				assert.Equal(t, tc.expected, errResp.Errors[0])
			}
		})
	}
}
