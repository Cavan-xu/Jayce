package mapping

import "encoding/json"

func UnmarshalJsonByte(content []byte, v interface{}) error {
	return json.Unmarshal(content, v)
}
