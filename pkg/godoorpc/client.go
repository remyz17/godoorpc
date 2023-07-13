package godoorpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/kolo/xmlrpc"
)

type RpcClient interface {
	CommonCall(serviceMethod string, reply any, args interface{}) error
	ObjectCall(serviceMethod string, reply any, args interface{}) error
	DbCall(serviceMethod string, reply any, args interface{}) error
}

type XMLRPCClient struct {
	common *xmlrpc.Client
	object *xmlrpc.Client
	db     *xmlrpc.Client
}

func NewXMLRPCClient(url string) (*XMLRPCClient, error) {
	commonClient, err := xmlrpc.NewClient(url+"/xmlrpc/2/common", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create common client: %w", err)
	}

	objectClient, err := xmlrpc.NewClient(url+"/xmlrpc/2/object", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create object client: %w", err)
	}

	dbClient, err := xmlrpc.NewClient(url+"/xmlrpc/2/db", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %w", err)
	}

	return &XMLRPCClient{
		common: commonClient,
		object: objectClient,
		db:     dbClient,
	}, nil
}

func (x *XMLRPCClient) call(c *xmlrpc.Client, serviceMethod string, reply any, args interface{}) error {
	return c.Call(serviceMethod, args, reply)
}

func (x *XMLRPCClient) CommonCall(serviceMethod string, reply any, args interface{}) error {
	return x.call(x.common, serviceMethod, reply, args)
}

func (x *XMLRPCClient) ObjectCall(serviceMethod string, reply any, args interface{}) error {
	return x.call(x.object, serviceMethod, reply, args)
}

func (x *XMLRPCClient) DbCall(serviceMethod string, reply any, args interface{}) error {
	return x.call(x.db, serviceMethod, reply, args)
}

type JSONRPCClient struct {
	url    string
	client *http.Client
}

func NewJSONRPCClient(url string) (*JSONRPCClient, error) {
	return &JSONRPCClient{
		url:    url,
		client: &http.Client{},
	}, nil
}

func (j *JSONRPCClient) doCall(service, serviceMethod string, reply any, args interface{}) error {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "call",
		"params": map[string]interface{}{
			"service": service,
			"method":  serviceMethod,
			"args":    args,
		},
		"id": rand.Intn(1000000000),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", j.url+"/jsonrpc", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := j.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(reply)
	if err != nil {
		return err
	}

	return nil
}

func (j *JSONRPCClient) CommonCall(serviceMethod string, reply any, args interface{}) error {
	return j.doCall("common", serviceMethod, reply, args)
}

func (j *JSONRPCClient) ObjectCall(serviceMethod string, reply any, args interface{}) error {
	return j.doCall("object", serviceMethod, reply, args)
}

func (j *JSONRPCClient) DbCall(serviceMethod string, reply any, args interface{}) error {
	return j.doCall("db", serviceMethod, reply, args)
}
