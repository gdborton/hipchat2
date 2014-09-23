/* Package hipchat2 is a client for the hipchat v2 api.

 https://www.hipchat.com/docs/apiv2
 */
package hipchat2

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"strings"
	"bytes"
)

var (
	AuthToken string
	Host string = "api.hipchat.com"
)

type ErrorResponse struct {
	Error struct {
		Code    int
		Type    string
		Message string
	}
}

func get (route string) ([]byte, error) {
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

func post (uri string, payload []byte) ([]byte, error) {
	httpClient := &http.Client{}
	req, err:= http.NewRequest("POST", uri, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode <= 200 || response.StatusCode > 300 {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, err
		}
		return nil, errors.New(errResp.Error.Message)
	}
	return body, nil
}

func delete (uri string) (error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("DELETE", uri, strings.NewReader(""))
	response, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return err
		}
		if err != nil {
			return err
		}
		return errors.New(errResp.Error.Message)
	}
	return nil
}

func put (uri string, payload []byte) ([]byte, error) {
	httpClient := &http.Client{}
	req, err:= http.NewRequest("PUT", uri, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode <= 200 || response.StatusCode > 300 {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, err
		}
		return nil, errors.New(errResp.Error.Message)
	}
	return body, nil
}
