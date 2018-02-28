package api

import (
	"net/http"
	"reflect"

	"../types"
	"github.com/kolo/xmlrpc"
)

type Config struct {
	DbName    string
	Admin     string
	Password  string
	URI       string
	Transport http.RoundTripper
}

type Client struct {
	CommonClient *xmlrpc.Client
	ObjectClient *xmlrpc.Client
	DbClient     *xmlrpc.Client
	ReportClient *xmlrpc.Client
	Session      *Session
}

type Session struct {
	DbName   string
	Admin    string
	Password string
	UID      int
}

func (config *Config) NewClient() (*Client, error) {
	commonClient, err := GetCommonClient(config.URI, config.Transport)
	if err != nil {
		return nil, err
	}
	objectClient, err := GetObjectClient(config.URI, config.Transport)
	if err != nil {
		return nil, err
	}
	dbClient, err := GetDbClient(config.URI, config.Transport)
	if err != nil {
		return nil, err
	}
	reportClient, err := GetObjectClient(config.URI, config.Transport)
	if err != nil {
		return nil, err
	}
	return &Client{
		CommonClient: commonClient,
		ObjectClient: objectClient,
		DbClient:     dbClient,
		ReportClient: reportClient,
		Session: &Session{
			Admin:    config.Admin,
			Password: config.Password,
			DbName:   config.DbName,
		},
	}, err
}

func (c *Client) CompleteSession() error {
	var uid int
	err := c.CommonClient.Call("authenticate", []interface{}{c.Session.DbName, c.Session.Admin, c.Session.Password, ""}, &uid)
	if err != nil {
		return err
	}
	c.Session.UID = uid
	return nil
}

func GetObjectClient(uri string, transport http.RoundTripper) (*xmlrpc.Client, error) {
	return xmlrpc.NewClient(uri+"/xmlrpc/2/object", transport)
}

func GetCommonClient(uri string, transport http.RoundTripper) (*xmlrpc.Client, error) {
	return xmlrpc.NewClient(uri+"/xmlrpc/2/common", transport)
}

func GetDbClient(uri string, transport http.RoundTripper) (*xmlrpc.Client, error) {
	return xmlrpc.NewClient(uri+"/xmlrpc/2/db", transport)
}

func GetReportClient(uri string, transport http.RoundTripper) (*xmlrpc.Client, error) {
	return xmlrpc.NewClient(uri+"/xmlrpc/2/report", transport)
}

func (c *Client) Create(model string, args []interface{}, elem interface{}) error {
	return c.DoRequest("create", model, args, nil, elem)
}

func (c *Client) Update(model string, args []interface{}) error {
	return c.DoRequest("write", model, args, nil, nil)
}

func (c *Client) Delete(model string, args []interface{}) error {
	return c.DoRequest("unlink", model, args, nil, nil)
}

func (c *Client) Search(model string, args []interface{}, options interface{}, elem interface{}) error {
	return c.DoRequest("search", model, args, options, elem)
}

func (c *Client) Read(model string, args []interface{}, options interface{}, elem interface{}) error {
	ne := elem.(types.Type).NilableType_()
	err := c.DoRequest("read", model, args, options, ne)
	if err == nil {
		reflect.ValueOf(elem).Elem().Set(reflect.ValueOf(ne.(types.NilableType).Type_()).Elem())
	}
	return err
}

func (c *Client) SearchRead(model string, args []interface{}, options interface{}, elem interface{}) error {
	ne := elem.(types.Type).NilableType_()
	err := c.DoRequest("search_read", model, args, options, ne)
	if err == nil {
		reflect.ValueOf(elem).Elem().Set(reflect.ValueOf(ne.(types.NilableType).Type_()).Elem())
	}
	return err
}

func (c *Client) SearchCount(model string, args []interface{}, elem interface{}) error {
	return c.DoRequest("search_count", model, args, nil, elem)
}

func (c *Client) DoRequest(method string, model string, args []interface{}, options interface{}, elem interface{}) error {
	return c.ObjectClient.Call("execute_kw", []interface{}{c.Session.DbName, c.Session.UID, c.Session.Password, model, method, args, options}, elem)
}

func (c *Client) getIdsByName(model string, name string) ([]int64, error) {
	var ids []int64
	err := c.Search(model, []interface{}{[]string{"name", "=", name}}, nil, &ids)
	return ids, err
}

func (c *Client) getByIds(model string, ids []int64, elem interface{}) error {
	err := c.Read(model, []interface{}{ids}, nil, elem)
	return err
}

func (c *Client) getByName(model string, name string, elem interface{}) error {
	err := c.SearchRead(model, []interface{}{[]interface{}{[]string{"name", "=", name}}}, nil, elem)
	return err
}

func (c *Client) getByField(model string, field string, value string, elem interface{}) error {
	err := c.SearchRead(model, []interface{}{[]interface{}{[]string{field, "=", value}}}, nil, elem)
	return err
}

func (c *Client) getAll(model string, elem interface{}) error {
	err := c.SearchRead(model, []interface{}{[]interface{}{}}, nil, elem)
	return err
}

func (c *Client) create(model string, fields map[string]interface{}, relation *types.Relations) (int64, error) {
	var id int64
	if relation != nil {
		types.HandleRelations(&fields, relation)
	}
	err := c.Create(model, []interface{}{fields}, &id)
	return id, err
}

func (c *Client) update(model string, ids []int64, fields map[string]interface{}, relation *types.Relations) error {
	if relation != nil {
		types.HandleRelations(&fields, relation)
	}
	err := c.Update(model, []interface{}{ids, fields})
	return err
}

func (c *Client) delete(model string, ids []int64) error {
	return c.Delete(model, []interface{}{ids})
}

func (c *Client) GetAllModels() ([]string, error) {
	var content []map[string]interface{}
	err := c.DoRequest("search_read", "ir.model", []interface{}{[]interface{}{}}, nil, &content)
	if err != nil {
		return []string{}, err
	}
	models := make([]string, len(content))
	for i, modelFields := range content {
		for field, model := range modelFields {
			if field == "model" {
				models[i] = model.(string)
			}
		}
	}
	return models, err
}
