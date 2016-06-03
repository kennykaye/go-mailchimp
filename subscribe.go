package mailchimp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AreaHQ/mailchimp/status"
)

// Subscribe ...
func (c *Client) Subscribe(email string, listID string, mergeFields map[string]interface{}) (*MemberResponse, error) {
	// Make request
	params := map[string]interface{}{
		"email_address": email,
		"status":        status.Subscribed,
		"merge_fields":  mergeFields,
	}
	resp, err := c.do(
		"POST",
		fmt.Sprintf("/lists/%s/members/", listID),
		&params,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// If the request failed
	if resp.StatusCode > 299 {
		errorResponse, err := extractError(data)
		if err != nil {
			return nil, err
		}
		return nil, errorResponse
	}

	// Unmarshal response into MemberResponse struct
	memberResponse := new(MemberResponse)
	if err := json.Unmarshal(data, memberResponse); err != nil {
		return nil, err
	}
	return memberResponse, nil
}
