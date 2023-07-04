package godoorpc

import (
	"fmt"
)

type OdooService struct {
	URL      string
	Db       string
	Username string
	Password string
	client   RpcClient
}

type Version struct {
	ServerVersion     string      `xmlrpc:"server_version"`
	ServerVersionInfo interface{} `xmlrpc:"server_version_info"`
	ServerSerie       string      `xmlrpc:"server_serie"`
	ProtocolVersion   int         `xmlrpc:"protocol_version"`
}

func (s *OdooService) Version() (*Version, error) {
	var err error
	reply := Version{}
	err = s.client.CommonCall("version", &reply, nil)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func (s *OdooService) ListDatabase() (*[]string, error) {
	var dbs []string
	err := s.client.DbCall("list", &dbs, nil)
	if err != nil {
		return nil, err
	}
	return &dbs, nil
}

func NewOdooService(protocol, url, db, username, password string) (*OdooService, error) {
	var client RpcClient
	var err error
	switch protocol {
	case "xmlrpc":
		client, err = NewXMLRPCClient(url)
	case "jsonrpc":
		client, err = NewJSONRPCClient(url)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create %s client: %w", protocol, err)
	}
	return &OdooService{URL: url, Db: db, Username: username, Password: password, client: client}, nil
}
