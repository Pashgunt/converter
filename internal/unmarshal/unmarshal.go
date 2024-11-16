package unmarshal

import "encoding/json"

func Decode(data []byte) (map[string]interface{}, error) {
	var rawData map[string]interface{}

	if err := json.Unmarshal(data, &rawData); err != nil {
		return nil, err
	}

	return rawData, nil
}
