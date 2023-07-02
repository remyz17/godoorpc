package godoorpc

import "fmt"

type OdooService struct {
	URL      string
	Db       string
	Username string
	Password string
	client   RpcClient
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
