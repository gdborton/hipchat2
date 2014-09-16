package hipchat2

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"strings"
	"strconv"
)

type privateClient struct {
	AuthToken string
	Host string
}

type ErrorResponse struct {
	Error struct {
		Code    int
		Type    string
		Message string
	}
}

func NewClient (AuthToken string) *privateClient {
	return &privateClient{Host: "api.hipchat.com", AuthToken: AuthToken}
}

func (client *privateClient) GetUsers () ([]User, error) {
	uri := fmt.Sprintf("https://%s/v2/user?auth_token=%s", client.Host, client.AuthToken)
	body, err := getJSONRoute(uri)
	if err != nil {
		return nil, err
	}

	usersResponse := &struct{ Items []User }{}

	if err := json.Unmarshal(body, usersResponse); err != nil {
		return nil, err
	}
	return usersResponse.Items, nil
}

func (client *privateClient) GetUser (emailMentionOrName string) (*User, error) {
	uri := fmt.Sprintf("https://%s/v2/user/%s?auth_token=%s", client.Host, emailMentionOrName, client.AuthToken)
	body, err := getJSONRoute(uri)
	if err != nil {
		return nil, err
	}

	userResponse := new(User)

	if err := json.Unmarshal(body, &userResponse); err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (client *privateClient) GetRooms () ([]Room, error) {
	uri := fmt.Sprintf("https://%s/v2/room?auth_token=%s", client.Host, client.AuthToken)

	body, err := getJSONRoute(uri)
	if err != nil {
		return nil, err
	}

	roomsResp := &struct{ Items []Room }{}

	if err := json.Unmarshal(body, roomsResp); err != nil {
		return nil, err
	}

	return roomsResp.Items, nil
}

func (client *privateClient) GetRoom (nameOrId string) (*Room, error) {
	uri := fmt.Sprintf("https://%s/v2/room/%s?auth_token=%s", client.Host, nameOrId, client.AuthToken)

	body, err := getJSONRoute(uri)
	if err != nil {
		return nil, err
	}
	roomResp := new(Room)

	if err := json.Unmarshal(body, &roomResp); err != nil {
		return nil, err
	}

	return roomResp, nil
}

func (client *privateClient) CreateRoom (topic string, guestAccess bool, name string, ownerUserId string, privacy string) (int, error) {
	uri := fmt.Sprintf("https://%s/v2/room?auth_token=%s", client.Host, client.AuthToken)

	payload := fmt.Sprintf("{\"topic\": %q, \"guest_access\": %s, \"name\": %q, \"owner_user_id\": %q, \"privacy\": %q}", topic, strconv.FormatBool(guestAccess), name, ownerUserId, privacy)
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", uri, strings.NewReader(payload))
	req.Header.Set("content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	createResponse := &struct{ Id int `json: "id"`}{}
	if err := json.Unmarshal(body, createResponse); err != nil {
		return 0, err
	}

	return createResponse.Id, nil
}

func getJSONRoute (route string) ([]byte, error) {
	response, err := http.Get(route)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, err
		}
		return nil, errors.New(errResp.Error.Message)
	}
	return body, err
}
