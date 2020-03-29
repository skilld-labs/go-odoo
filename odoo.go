//go:generate ./generator/generator -u $ODOO_ADMIN -p $ODOO_PASSWORD -d $ODOO_DATABASE --url $ODOO_URL -o $ODOO_REPO_PATH --models $ODOO_MODELS
package odoo

import (
	"errors"

	"github.com/kolo/xmlrpc"
)

var (
	ErrClientConfigurationInvalid = errors.New("client configuration is invalid")
	ErrClientNotAuthenticate      = errors.New("client is not authenticate")
)

type ClientConfig struct {
	Database string
	Admin    string
	Password string
	URL      string
}

func (c *ClientConfig) valid() bool {
	return c.Database != "" &&
		c.Admin != "" &&
		c.Password != "" &&
		c.URL != ""
}

type Client struct {
	common *xmlrpc.Client
	object *xmlrpc.Client
	cfg    *ClientConfig
	uid    int64
	auth   bool
}

func NewClient(cfg *ClientConfig) (*Client, error) {
	if !cfg.valid() {
		return nil, ErrClientConfigurationInvalid
	}
	c := &Client{
		cfg:    cfg,
		common: &xmlrpc.Client{},
		object: &xmlrpc.Client{},
		auth:   false,
	}
	if err := c.Authenticate(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) Close() {
	if c.common != nil {
		c.common.Close()
	}
	if c.object != nil {
		c.object.Close()
	}
}

func (c *Client) Version() (Version, error) {
	v := Version{}
	reply, err := c.commonCall("version", nil)
	if err != nil {
		return Version{}, err
	}
	ConvertFromDynamicToStatic(reply, &v)
	return v, nil
}

type Criterion []interface{}
type Criteria []*Criterion

func NewCriteria(criterions ...*Criterion) *Criteria {
	cc := Criteria{}
	for _, c := range criterions {
		cc = append(cc, c)
	}
	return &cc
}

func NewCriterion(field, operator string, value interface{}) *Criterion {
	c := Criterion(newTuple(field, operator, value))
	return &c
}

func (c *Criteria) Add(field, operator string, value interface{}) *Criteria {
	*c = append(*c, NewCriterion(field, operator, value))
	return c
}

type Options map[string]interface{}

func NewOptions() *Options {
	opt := Options(make(map[string]interface{}))
	return &opt
}

func (o *Options) Offset(offset int) *Options {
	return o.Add("offset", offset)
}

func (o *Options) Limit(limit int) *Options {
	return o.Add("limit", limit)
}

func (o *Options) FetchFields(fields ...string) *Options {
	ff := []string{}
	for _, f := range fields {
		ff = append(ff, f)
	}
	return o.Add("fields", ff)
}

func (o *Options) AllFields(fields ...string) *Options {
	ff := []string{}
	for _, f := range fields {
		ff = append(ff, f)
	}
	return o.Add("allfields", ff)
}

func (o *Options) Attributes(attributes ...string) *Options {
	aa := []string{}
	for _, a := range attributes {
		aa = append(aa, a)
	}
	return o.Add("attributes", aa)
}

func (o *Options) Add(opt string, v interface{}) *Options {
	(*o)[opt] = v
	return o
}

func getValuesFromInterface(v interface{}) map[string]interface{} {
	switch v.(type) {
	case map[string]interface{}:
		return v.(map[string]interface{})
	default:
		return ConvertFromStaticToDynamic(v)
	}
}

func (c *Client) Create(model string, values interface{}) (int64, error) {
	v := getValuesFromInterface(values)
	resp, err := c.ExecuteKw("create", model, []interface{}{v}, nil)
	if err != nil {
		return -1, err
	}
	return resp.(int64), nil
}

func (c *Client) Update(model string, ids []int64, values interface{}) error {
	v := getValuesFromInterface(values)
	_, err := c.ExecuteKw("write", model, []interface{}{ids, v}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Delete(model string, ids []int64) error {
	_, err := c.ExecuteKw("unlink", model, []interface{}{ids}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SearchRead(model string, criteria *Criteria, options *Options, elem interface{}) error {
	resp, err := c.ExecuteKw("search_read", model, argsFromCriteria(criteria), options)
	if err != nil {
		return err
	}
	if err := ConvertFromDynamicToStatic(resp, elem); err != nil {
		return err
	}
	return nil
}

func (c *Client) Read(model string, ids []int64, options *Options, elem interface{}) error {
	resp, err := c.ExecuteKw("read", model, []interface{}{ids}, options)
	if err != nil {
		return err
	}
	if err := ConvertFromDynamicToStatic(resp, elem); err != nil {
		return err
	}
	return nil
}

func (c *Client) Count(model string, criteria *Criteria, options *Options) (int64, error) {
	resp, err := c.ExecuteKw("search_count", model, argsFromCriteria(criteria), options)
	if err != nil {
		return -1, err
	}
	return resp.(int64), nil
}

func (c *Client) Search(model string, criteria *Criteria, options *Options) ([]int64, error) {
	resp, err := c.ExecuteKw("search", model, argsFromCriteria(criteria), options)
	if err != nil {
		return []int64{}, err
	}
	return SliceInterfaceToInt64Slice(resp.([]interface{})), nil
}

func (c *Client) FieldsGet(model string, options *Options) (map[string]interface{}, error) {
	resp, err := c.ExecuteKw("fields_get", model, nil, options)
	if err != nil {
		return nil, nil
	}
	return resp.(map[string]interface{}), nil
}

func (c *Client) ExecuteKw(method, model string, args []interface{}, options *Options) (interface{}, error) {
	if err := c.checkForAuthentication(); err != nil {
		return nil, err
	}
	resp, err := c.objectCall("execute_kw", []interface{}{c.cfg.Database, c.uid, c.cfg.Password, model, method, args, options})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Authenticate() error {
	if !c.isAuthenticate() {
		resp, err := c.commonCall("authenticate", []interface{}{c.cfg.Database, c.cfg.Admin, c.cfg.Password, ""})
		if err != nil {
			return err
		}
		c.uid = resp.(int64)
		c.auth = true
	}
	return nil
}

func (c *Client) commonCall(serviceMethod string, args interface{}) (interface{}, error) {
	if err := c.loadCommonClient(); err != nil {
		return nil, err
	}
	resp, err := c.call(c.common, serviceMethod, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) objectCall(serviceMethod string, args interface{}) (interface{}, error) {
	if err := c.loadObjectClient(); err != nil {
		return nil, err
	}
	resp, err := c.call(c.object, serviceMethod, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) call(x *xmlrpc.Client, serviceMethod string, args interface{}) (interface{}, error) {
	var reply interface{}
	if err := x.Call(serviceMethod, args, &reply); err != nil {
		return nil, err
	}
	return reply, nil
}

func (c *Client) loadCommonClient() error {
	return c.loadXmlrpcClient(c.common, "/xmlrpc/2/common")
}

func (c *Client) loadObjectClient() error {
	return c.loadXmlrpcClient(c.object, "/xmlrpc/2/object")
}

func (c *Client) loadXmlrpcClient(x *xmlrpc.Client, path string) error {
	if x.Client == nil {
		if newClient, err := xmlrpc.NewClient(c.cfg.URL+path, nil); err != nil {
			return err
		} else {
			*x = *newClient
		}
	}
	return nil
}

func (c *Client) checkForAuthentication() error {
	if !c.isAuthenticate() {
		return ErrClientNotAuthenticate
	}
	return nil
}

func (c *Client) isAuthenticate() bool {
	return c.auth
}

func newTuple(values ...interface{}) []interface{} {
	t := make([]interface{}, len(values))
	for i, v := range values {
		t[i] = v
	}
	return t
}

func argsFromCriteria(c *Criteria) []interface{} {
	if c != nil {
		return []interface{}{*c}
	}
	return []interface{}{}
}
