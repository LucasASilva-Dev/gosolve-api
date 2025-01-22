package index

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputFileValidFilename(t *testing.T) {
	// Create a test file
	tmpFile, err := os.CreateTemp("", "test-input-file")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Write some content to the file
	_, err = tmpFile.Write([]byte("Hello, World!"))
	assert.NoError(t, err)

	// Read the file using the function under test
	data, err := ReadInputFile(tmpFile.Name())
	assert.NoError(t, err)

	// Check that the data was read correctly
	assert.Equal(t, []byte("Hello, World!"), data)
}

func TestReadInputFileInvalidFilename(t *testing.T) {
	// Pass an invalid filename to the function under test
	data, err := ReadInputFile("")
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestReadInputFileNonExistentFilename(t *testing.T) {
	// Pass a non-existent filename to the function under test
	data, err := ReadInputFile("non-existent-file.txt")
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestNonExistingReadFile(t *testing.T) {
	_, err := ReadInputFile("non-existing-file")
	assert.Error(t, err)
}

func TestReadFile(t *testing.T) {
	_, err := ReadInputFile("../../data/input.txt")
	assert.NoError(t, err)
}
