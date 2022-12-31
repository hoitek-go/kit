package body

import (
	"encoding/json"
	"io"
	"net/http"
)

func Bind(r *http.Request, result any) error {
	if r.PostForm == nil {
		r.ParseMultipartForm(32 << 20)
		r.ParseForm()
	}
	parsedData := map[string]interface{}{}
	bodyRaw, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if len(bodyRaw) > 0 {
		err = json.Unmarshal([]byte(bodyRaw), &parsedData)
		if err != nil {
			return err
		}
	} else {
		data, err := json.Marshal(r.Form)
		if err != nil {
			return err
		}
		dataMap := make(map[string][]interface{})
		err = json.Unmarshal(data, &dataMap)
		if err != nil {
			return err
		}
		for key, value := range dataMap {
			parsedData[key] = value[0]
		}
	}
	byteArray, err := json.Marshal(parsedData)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(byteArray, result); err != nil {
		return err
	}
	return nil
}
