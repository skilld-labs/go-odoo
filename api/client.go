package api

import (
	"net/http"
	"reflect"

	"github.com/kolo/xmlrpc"
	"github.com/skilld-labs/go-odoo/types"
)

type Client struct {
	client    *xmlrpc.Client
	URI       string
	Transport http.RoundTripper
	Session   struct {
		DbName   string
		Admin    string
		Password string
		UID      int
	}
}

func NewClient(uri string, transport http.RoundTripper) (*Client, error) {
	c, err := xmlrpc.NewClient(uri+"/xmlrpc/2/object", transport)
	if err != nil {
		return nil, err
	}
	client := &Client{client: c, URI: uri, Transport: transport}
	return client, err
}

func (c *Client) Login(dbName string, admin string, password string) error {
	var uid int
	uriTemp := c.URI + "/xmlrpc/2/common"
	cTemp, err := xmlrpc.NewClient(uriTemp, c.Transport)
	if err != nil {
		return err
	}
	clientTemp := &Client{client: cTemp, URI: uriTemp, Transport: c.Transport}
	err = clientTemp.client.Call("authenticate", []interface{}{dbName, admin, password, ""}, &uid)
	if err != nil {
		return err
	}
	c.Session.DbName = dbName
	c.Session.Admin = admin
	c.Session.Password = password
	c.Session.UID = uid
	return err
}

func (c *Client) GetReport(model string, ids []int64) (map[string]interface{}, error) {
	client, err := xmlrpc.NewClient(c.URI+"/xmlrpc/2/report", c.Transport)
	if err != nil {
		return nil, err
	}
	var report map[string]interface{}
	reportService := NewIrActionsReportService(c)
	fields, err := reportService.GetByField("model", model)
	if err != nil {
		return nil, err
	}
	return report, client.Call("render_report", []interface{}{c.Session.DbName, c.Session.UID, c.Session.Password, (*fields)[0].ReportName, ids}, &report)
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
	return c.client.Call("execute_kw", []interface{}{c.Session.DbName, c.Session.UID, c.Session.Password, model, method, args, options}, elem)
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
