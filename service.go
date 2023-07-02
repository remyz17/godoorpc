package godoorpc

type OdooService struct {
	URL      string
	Db       string
	Username string
	Password string
	client   RpcClient
}

func NewOdooService(protocol, url, db, username, password string) *OdooService {
	var client RpcClient
	switch protocol {
	case "xmlrpc":
		client = &XMLRPCClient{}
	case "jsonrpc":
		client = &JSONRPCClient{}
	}
	return &OdooService{URL: url, Db: db, Username: username, Password: password, client: client}
}
