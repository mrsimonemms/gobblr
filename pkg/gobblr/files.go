package gobblr

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/relvacode/iso8601"
)

type FileType string

const (
	FileTypeJSON FileType = "json"
	FileTypeJS   FileType = "js"
)

var AllowedFileTypes = map[FileType]struct{}{
	FileTypeJSON: {},
	FileTypeJS:   {},
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

// Iterate through all keys searching for ISO8601 dates and convert to Golang time.Date
func detectAndConvertDates(data map[string]interface{}) map[string]interface{} {
	for k, v := range data {
		mapValue, ok := data[k].(map[string]interface{})
		if ok {
			v = detectAndConvertDates(mapValue)
		}

		f, err := iso8601.ParseString(fmt.Sprintf("%v", v))
		if err == nil {
			v = f
		}

		data[k] = v
	}

	return data
}

func loadJS(input []byte) ([]byte, error) {
	vm := goja.New()
	_, err := vm.RunString(string(input))
	if err != nil {
		return nil, err
	}

	var fn func() map[string]interface{}
	err = vm.ExportTo(vm.Get("data"), &fn)
	if err != nil {
		return nil, err
	}

	// Convert the response to a byte array
	jsonData, err := json.Marshal(fn())
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func (f *File) LoadFile() (jsonData *FileData, err error) {
	var fileData []byte

	switch f.Type {
	case FileTypeJSON, FileTypeJS:
		fileData, err = os.ReadFile(f.Path)
		if err != nil {
			return nil, err
		}

		if f.Type == FileTypeJS {
			fileData, err = loadJS(fileData)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("unknown file type: %s", f.Type)
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
		for k, v := range jsonData.Data {
			now := time.Now()

			// Iterate and detect dates
			jsonData.Data[k] = detectAndConvertDates(v)

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
