package notionapi

import (
	"encoding/json"
)

func (c *Client) LoadUserContent() (*ValueResponse, error) {

	req := struct{}{}

	apiURL := "/api/v3/loadUserContent"
	var rsp struct {
		RecordMap map[string]map[string]ValueResponse `json:"recordMap"`
	}
	var err error
	if _, err = doNotionAPI(c, apiURL, req, &rsp); err != nil {
		return nil, err
	}

	result := &ValueResponse{}

	for table, values := range rsp.RecordMap {
		for _, value := range values {
			var obj interface{}
			if table == TableUser {
				result.User = &User{}
				obj = result.User
			}
			if table == TableBlock {
				result.Block = &Block{}
				obj = result.Block
			}
			if table == TableSpace {
				result.Space = &Space{}
				obj = result.Space
			}
			if obj == nil {
				continue
			}
			if err := json.Unmarshal(value.Value, &obj); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}
