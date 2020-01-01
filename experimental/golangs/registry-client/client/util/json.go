package util

import "github.com/mkideal/pkg/encoding/jsonx"

func JsonX2Map(s string) (map[string]interface{}, error) {

	var ret map[string]interface{}
	err := jsonx.Unmarshal([]byte(s), &ret, jsonx.WithExtraComma(), jsonx.WithComment())
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func JsonFormat(s string) {

}
