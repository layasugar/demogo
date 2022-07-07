package benchmark

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkJsonMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		JsonMarshal()
	}
}

func BenchmarkGobMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GobMarshal()
	}
}

func BenchmarkJsonIteratorMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		JsonIterMarshal()
	}
}

func JsonMarshal() {
	var data = map[string]interface{}{
		"name": "aaa",
		"id":   111.1,
	}

	json.Marshal(data)
}

func JsonIterMarshal() {
	var data = map[string]interface{}{
		"name": "aaa",
		"id":   111.1,
	}

	jsonIter.Marshal(data)
}

func GobMarshal() {
	var data = map[string]interface{}{
		"name": "aaa",
		"id":   111.1,
	}
	dec := gob.NewEncoder(bytes.NewBuffer(nil))
	dec.Encode(data)
}

func JsonUnmarshal() {
	var data map[string]interface{}
	var b = []byte(`{"name":"aaa","id":111.1}`)
	json.Unmarshal(b, &data)
}

func JsonIterUnmarshal() {
	var data map[string]interface{}
	var b = []byte(`{"name":"aaa","id":111.1}`)
	jsonIter.Unmarshal(b, &data)
}

func GobUnmarshal() {
	var data map[string]interface{}
	var b = []byte(`{"name":"aaa","id":111.1}`)
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&data)
	if err != nil {
		return
	}
}
