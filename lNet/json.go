package lNet

import "encoding/json"

func JSON(p []byte, err error) (map[string]interface{}, error) {
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	err = json.Unmarshal(p, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
