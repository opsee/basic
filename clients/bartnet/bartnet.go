// Package bartnet provides a client interface to bartnet
package bartnet

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/opsee/basic/schema"
	"github.com/opsee/basic/service"
	"io"
	"io/ioutil"
	"net/http"
)

type Client interface {
	ListChecks(user *schema.User) ([]*schema.Check, error)
	CreateCheck(user *schema.User, check *schema.Check) (*schema.Check, error)
}

type client struct {
	client   *http.Client
	endpoint string
}

// An endpoint is the address of the bartnet service.
func New(endpoint string) Client {
	return &client{
		client:   &http.Client{},
		endpoint: endpoint,
	}
}

// ListChecks lists the checks + assertions for a user's customer account, without the results
func (c *client) ListChecks(user *schema.User) ([]*schema.Check, error) {
	body, err := c.do(user, "GET", "application/x-protobuf", "/gql/checks", nil)
	if err != nil {
		return nil, err
	}

	checks := &service.CheckResourceRequest{}
	err = proto.Unmarshal(body, checks)
	if err != nil {
		return nil, err
	}

	return checks.Checks, nil
}

func (c *client) CreateCheck(user *schema.User, check *schema.Check) (*schema.Check, error) {
	marshaler := jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(check)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(user, "POST", "application/json", "/checks", bytes.NewBufferString(jsondata))
	if err != nil {
		return nil, err
	}

	checkResp := &schema.Check{}
	err = jsonpb.UnmarshalString(string(resp), checkResp)
	if err != nil {
		return nil, err
	}

	return checkResp, nil
}

func (c *client) do(user *schema.User, method, accept, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, c.endpoint+path, body)
	if err != nil {
		return nil, err
	}

	toke, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(toke)))
	req.Header.Set("Accept", accept)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("bartnet responded with error status: %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
