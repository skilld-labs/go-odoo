package main

import (
	"github.com/kolo/xmlrpc"
	"net/http"
)

type Client struct {
	client *xmlrpc.Client
	Session struct {
		DbName   string
		Admin    string
		Password string
		UID      int
	}
	URI string
	Transport http.RoundTripper
}

func NewClient(uri string, transport http.RoundTripper) (*Client, error) {
	c, err := xmlrpc.NewClient(uri + "/xmlrpc/2/object", transport)
	if err != nil {
		return nil, err
	}
	return &Client{client: c, URI: uri, Transport: transport}, err
}

func (c *Client) Login (dbName string, admin string, password string) error {
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

func (c *Client) Create (model string, args []interface{}) (int, error) {
	var id int
	err := c.DoRequest("create", model, args, nil, &id)
	return id, err
}

func (c *Client) Update (model string, args []interface{}) error {
	return c.DoRequest("write", model, args, nil, nil)
}

func (c * Client) Delete (model string, args []interface{}) error {
	return c.DoRequest("unlink", model, args, nil, nil)
}

func (c *Client) Search (model string, args []interface{}, options interface{}) ([]int, error) {
	var ids []int
	err := c.DoRequest("search", model, args, options, &ids)
	return ids, err
}

func (c *Client) Read (model string, args []interface{},options interface{}, elem interface{}) error {
	err := c.DoRequest("read", model, args, options, elem)
	return err
}

func (c *Client) SearchRead (model string, args []interface{}, options interface{}, elem interface{}) error {
	err := c.DoRequest("search_read", model, args, options, elem)
	return err
}

func (c *Client) SearchCount (model string, args []interface{}) (int, error) {
	var n int
	err := c.DoRequest("search_count", model, args, nil, &n)
	return n, err
}

func (c *Client) DoRequest (method string, model string, args []interface{}, options interface{}, elem interface{}) error {
	var err error
	if options != nil {
		if elem != nil {
			err = c.client.Call("execute_kw", []interface{}{c.Session.DbName, c.Session.UID, c.Session.Password, model, method, args, options}, elem)
		} else {
			err = c.client.Call("execute_kw", []interface{}{c.Session.DbName, c.Session.UID, c.Session.Password, model, method, args, options}, nil)
		}
	} else {
		if elem != nil {
			err = c.client.Call("execute_kw", []interface{}{c.Session.DbName, c.Session.UID, c.Session.Password, model, method, args}, elem)
		} else {
			err = c.client.Call("execute_kw", []interface{}{c.Session.DbName, c.Session.UID, c.Session.Password, model, method, args}, nil)
		}
	}
	return err
}
