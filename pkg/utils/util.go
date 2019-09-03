package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	testConfigurationFileName = "test_setup.json"
)

// EdgexTest - general Edgex client test structure
type S3xClientTestConfig struct {
	Mockup  int    `json:"mockup"`
	Url     string `json:"url"`
	Authkey string `json:"authkey"`
	Secret  string `json:"secret"`
	Bucket  string `json:"bucket"`
	Object  string `json:"object"`
	//Debug   int    `json:"debug"`
}

func GetTestConfig() (*S3xClientTestConfig, error) {

	buf, err := ioutil.ReadFile(testConfigurationFileName)
	if err != nil {
		return nil, err
	}

	config := &S3xClientTestConfig{}
	err = json.Unmarshal(buf, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetBucketPath(bucket string) (string, error) {
	bucket = strings.TrimSpace(bucket)
	if len(bucket) == 0 {
		return "", fmt.Errorf("Invalid bucket name `%s`", bucket)
	}
	return bucket, nil
}

func GetObjectPath(bucket, object string) (string, error) {

	bucket = strings.TrimSpace(bucket)
	if len(bucket) == 0 {
		return "", fmt.Errorf("Invalid bucket name `%s`", bucket)
	}

	object = strings.TrimSpace(object)
	if len(object) == 0 {
		return "", fmt.Errorf("Invalid object name: `%s`", object)
	}
	return fmt.Sprintf("%s/%s", bucket, object), nil
}

// ArrToJSON - convert k/v pairs to json
func ArrToJSON(arr ...string) string {
	var b bytes.Buffer

	b.WriteString("{")
	n := 0
	for i := 0; i < len(arr); i += 2 {
		if n > 0 {
			b.WriteString(", ")
		}
		b.WriteString(" \"")
		b.WriteString(arr[i])
		b.WriteString("\": \"")
		b.WriteString(arr[i+1])
		b.WriteString("\"")
		n++
	}
	b.WriteString("}")

	return b.String()
}

// ArrToCVS - convert k/v pairs to cvs
func ArrToCVS(arr ...string) string {
	var b bytes.Buffer

	n := 0
	for i := 0; i < len(arr); i += 2 {
		if n > 0 {
			b.WriteString("\n")
		}
		b.WriteString(arr[i])
		b.WriteString(";")
		b.WriteString(arr[i+1])
		n++
	}

	return b.String()
}
