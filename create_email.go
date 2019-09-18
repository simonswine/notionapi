package notionapi

import (
	"errors"
)

// CreateUser invites a new user through his email address
func (c *Client) CreateUser(email string) (*UserWithRole, error) {
	req := struct {
		Email string `json:"email"`
	}{
		Email: email,
	}

	// response is empty, as far as I can tell
	var rsp struct {
		UserID    string `json:"userId"`
		RecordMap struct {
			NotionUser map[string]UserWithRole `json:"notion_user"`
		} `json:"recordMap"`
	}

	apiURL := "/api/v3/createEmailUser"
	_, err := doNotionAPI(c, apiURL, req, &rsp)

	users, ok := rsp.RecordMap.NotionUser[rsp.UserID]
	if !ok {
		return nil, errors.New("error creating user")
	}

	return &users, err
}
