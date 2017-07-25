package api

import (
	"net/http"
	"reflect"

	"../types"
	"github.com/kolo/xmlrpc"
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

	SaleOrder              *SaleOrderService
	SaleOrderLine          *SaleOrderLineService
	ProjectProject         *ProjectProjectService
	ProjectTask            *ProjectTaskService
	AccountInvoice         *AccountInvoiceService
	AccountInvoiceLine     *AccountInvoiceLineService
	AccountAnalyticAccount *AccountAnalyticAccountService
	AccountAnalyticLine    *AccountAnalyticLineService
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
