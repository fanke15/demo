package pkg

import (
	"bytes"

	iter "github.com/json-iterator/go"
)

var json = iter.ConfigCompatibleWithStandardLibrary

func Marshal(data interface{}) []byte {
	v, _ := json.Marshal(data)
	return v
}

func MarshalToString(data interface{}) string {
	var (
		bf          = bytes.NewBuffer([]byte{})
		jsonEncoder = json.NewEncoder(bf)
	)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(data)
	return bf.String()
}

func UnMarshal(data []byte, v interface{}) {
	_ = json.Unmarshal(data, v)
}
