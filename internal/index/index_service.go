package index

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// IndexService ...
type IndexService interface {
	UpdateIndex() ([]*int, error)
}

type indexService struct {
	file []*int
}

// NewIndexService ...
func NewIndexService() IndexService {
	return &indexService{}
}

func ReadInputFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	return file, err
}

// UpdateIndex updates the index when called by the index monitor.
// The function reads the data file at the specified path and parses it into
// a slice of integers. The function returns the slice of integers, and an error
// if the file cannot be read.
func (is indexService) UpdateIndex() ([]*int, error) {
	filePath, err := filepath.Abs("data/input.txt")
	if err != nil {
		log.Error("Failed to get absolute path of file")
		return nil, err
	}

	log.Debugf("Reading file %s\n", filePath)

	// Load the file into a slice when the service starts
	positions, err := ReadInputFile(filePath)

	if err != nil {
		log.Errorf("Failed to read file %s: %v\n", filePath, err)
		return nil, err
	}

	log.Debugf("Read %d bytes from file %s\n", len(positions), filePath)

	// Split the string into a slice of integers
	strPositions := strings.Split(string(positions), "\n")

	log.Debugf("Split file into %d strings\n", len(strPositions))

	file := make([]*int, len(strPositions))
	for i, position := range strPositions {
		file[i] = new(int)
		*file[i], _ = strconv.Atoi(position)
	}

	return file, nil
}
