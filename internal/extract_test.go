package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractContent(t *testing.T) {
	type testCase struct {
		name        string
		input       string
		inputType   string
		want        string
		expectError bool
	}
	testCases := []testCase{
		{
			name:        "valid input",
			input:       "```json{\"name\": \"John\", \"age\": 25}```",
			inputType:   "json",
			want:        "{\"name\": \"John\", \"age\": 25}",
			expectError: false,
		},
		{
			name:        "invalid input",
			input:       "json\n{\"name\": \"John\", \"age\": 25\n",
			inputType:   "json",
			want:        "",
			expectError: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ExtractContent(tc.input, tc.inputType)
			assert.Equal(t, tc.want, got)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
