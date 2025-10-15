package json

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	// Marshal is exported by json-iterator/go package.
	Marshal = json.Marshal
	// Unmarshal is exported by json-iterator/go package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by json-iterator/go package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by json-iterator/go package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by json-iterator/go package.
	NewEncoder = json.NewEncoder
)
