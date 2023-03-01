package settings

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	// Configuration defines properties of the configuration loader
	Configuration struct {
		filename string
		format   FileFormat
	}

	// FileFormat defines supported configuration file formats
	FileFormat int
)

const (
	// FileFormatUndefined represents unset value
	FileFormatUndefined FileFormat = iota - 1
	// FileFormatJSON defines JSON syntax
	FileFormatJSON
)

const (
	// DefaultFileFormat is undefined value
	DefaultFileFormat = FileFormatUndefined

	// DefaultFilename is an empty string (undefined)
	DefaultFilename = ""
)

// NewConfiguration returns a reference to a new instance of "Configuration" type
func NewConfiguration() *Configuration {
	return &Configuration{
		filename: DefaultFilename,
		format:   DefaultFileFormat,
	}
}

// FromFile defines configuration filename and its data format
func (c *Configuration) FromFile(name string, format FileFormat) *Configuration {
	c.filename = name
	c.format = format
	return c
}

// Populate loads configuration object with values from the defined sources
func (c *Configuration) Populate(object any) error {
	if c.filename != "" && c.format != FileFormatUndefined {
		if err := c.populateFromFile(c.filename, object); err != nil {
			return fmt.Errorf("unable to load data: %w", err)
		}
	}
	return nil
}

func (c *Configuration) populateFromFile(name string, object any) error {
	if name == "" {
		return nil
	}

	data, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("can't read from a file: %w", err)
	}

	if err := json.Unmarshal(data, object); err != nil {
		return fmt.Errorf("can't unmarshal: %w", err)
	}

	return nil
}
