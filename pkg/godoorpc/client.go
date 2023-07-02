package godoorpc

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kolo/xmlrpc"
)

type RpcClient interface {
	CommonCall(serviceMethod string, args interface{}) (interface{}, error)
	ObjectCall(serviceMethod string, args interface{}) (interface{}, error)
	DbCall(serviceMethod string, args interface{}) (interface{}, error)
}

type XMLRPCClient struct {
	common *xmlrpc.Client
	object *xmlrpc.Client
	db     *xmlrpc.Client
}

func NewXMLRPCClient(url string) (*XMLRPCClient, error) {
	commonClient, err := xmlrpc.NewClient(url+"/xmlrpc/2/common", nil)
	if err != nil {
		return nil, errors.New("failed to create common client: %w", err)
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

func (x *XMLRPCClient) call(c *xmlrpc.Client, serviceMethod string, args interface{}) (interface{}, error) {
	var reply interface{}
	if err := c.Call(serviceMethod, args, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}

func (x *XMLRPCClient) CommonCall(serviceMethod string, args interface{}) (interface{}, error) {
	resp, err := x.call(x.common, serviceMethod, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (x *XMLRPCClient) ObjectCall(serviceMethod string, args interface{}) (interface{}, error) {
	resp, err := x.call(x.object, serviceMethod, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (x *XMLRPCClient) DbCall(serviceMethod string, args interface{}) (interface{}, error) {
	resp, err := x.call(x.db, serviceMethod, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
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

func (j *JSONRPCClient) CommonCall(serviceMethod string, args interface{}) (interface{}, error) {
	return "", errors.New("Not Implemented")
}

func (j *JSONRPCClient) ObjectCall(serviceMethod string, args interface{}) (interface{}, error) {
	return "", errors.New("Not Implemented")
}

func (j *JSONRPCClient) DbCall(serviceMethod string, args interface{}) (interface{}, error) {
	return "", errors.New("Not Implemented")
}
