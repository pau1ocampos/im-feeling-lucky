package lucky

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubFilesystem struct {
	Error error
	Bytes []byte
}

func TestParseFromFile_ReturnsError(t *testing.T) {
	// Arrange
	stub := StubFilesystem{
		Error: errors.New("An error occurred"),
	}
	opts := Options{
		FileSystem: stub,
	}
	// Act
	r, err := opts.FileSystem.ReadFile("path")
	// Assert
	assert.Nil(t, r)
	assert.Error(t, err)
}

func TestParseFromFile_ValidFile_ReturnsExpectedStruct(t *testing.T) {
	// Arrange
	content := `[
		{
			"numbers":[7,8,24,25,47],
			"stars":[8,9]},
		{
			"numbers":[3,4,27,29,37],
			"stars":[5,6]
		},
		{
			"numbers":[15,19,22,46,49],
			"stars":[2,9]
		}
	]`

	expected := &Draws{
		Draw{
			Numbers: []int{7, 8, 24, 25, 47},
			Starts:  []int{8, 9},
		},
		Draw{
			Numbers: []int{3, 4, 27, 29, 37},
			Starts:  []int{5, 6},
		},
		Draw{
			Numbers: []int{15, 19, 22, 46, 49},
			Starts:  []int{2, 9},
		},
	}

	stub := StubFilesystem{
		Error: nil,
		Bytes: []byte(content),
	}

	opts := Options{
		FileSystem: stub,
	}
	// Act
	r, err := opts.ParseFromFile("file")
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, r)
}

func (fs StubFilesystem) ReadFile(path string) ([]byte, error) {
	return fs.Bytes, fs.Error
}
