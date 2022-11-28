package gobblr

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type FileType string

const (
	FileTypeJSON FileType = "json"
)

var AllowedFileTypes = map[FileType]struct{}{
	FileTypeJSON: {},
}

type File struct {
	Path      string
	TableName string
	Type      FileType
}

type FileData struct {
	Meta MetaData                 `json:"meta"`
	Data []map[string]interface{} `json:"data"`
}

type MetaData struct {
	Created    bool   `json:"created"`
	CreatedKey string `json:"createdKey"`
	Updated    bool   `json:"updated"`
	UpdatedKey string `json:"updatedKey"`
}

func (f *File) LoadFile() (jsonData *FileData, err error) {
	var fileData []byte

	switch f.Type {
	case FileTypeJSON:
		fileData, err = os.ReadFile(f.Path)
		if err != nil {
			return nil, err
		}
	}

	// Set default parameters
	jsonData = &FileData{
		Meta: MetaData{
			Created:    true,
			CreatedKey: "createdAt",
			Updated:    true,
			UpdatedKey: "updatedAt",
		},
		Data: make([]map[string]interface{}, 0),
	}

	// First, check if it's in the data-only format
	var rawData []map[string]interface{}
	if err := json.Unmarshal([]byte(fileData), &rawData); err != nil {
		// Error parsing - treat JSON as in FileData format
		if err := json.Unmarshal([]byte(fileData), &jsonData); err != nil {
			// Data in unknown format
			return nil, err
		}
	} else {
		// JSON is just the data
		jsonData.Data = rawData
	}

	addCreated := jsonData.Meta.Created
	addUpdated := jsonData.Meta.Updated
	if addCreated || addUpdated {
		// Add the created/updated timestamp(s)
		for k := range jsonData.Data {
			now := time.Now()

			if addCreated {
				jsonData.Data[k][jsonData.Meta.CreatedKey] = now
			}
			if addUpdated {
				jsonData.Data[k][jsonData.Meta.UpdatedKey] = now
			}
		}
	}

	return jsonData, nil
}

func FindFiles(dataPath string) ([]File, error) {
	dataFiles := make([]File, 0)

	files, err := os.ReadDir(dataPath)
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile(`(?P<ID>\d+)-(?P<Table>\w+).(?P<Ext>\w+)`)
	if err != nil {
		return nil, err
	}
	regexNames := regex.SubexpNames()

	for _, file := range files {
		fileType := FileType(strings.Trim(filepath.Ext(file.Name()), "."))

		// Check file in legitimate format
		if !regex.MatchString(file.Name()) {
			// Invalid naming convention - ignore
			continue
		}

		matches := regex.FindAllStringSubmatch(file.Name(), -1)
		m := map[string]string{}
		for i, n := range matches[0] {
			m[regexNames[i]] = n
		}

		// Check file type is allowed
		if _, ok := AllowedFileTypes[fileType]; !ok {
			// Invalid file type - ignore
			continue
		}

		dataFiles = append(dataFiles, File{
			Path:      path.Join(dataPath, file.Name()),
			TableName: m["Table"],
			Type:      fileType,
		})
	}

	return dataFiles, nil
}
