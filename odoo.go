// Package odoo contains client code of library
//
//go:generate ./generator/generator -u $ODOO_ADMIN -p $ODOO_PASSWORD -d $ODOO_DATABASE --url $ODOO_URL -o $ODOO_REPO_PATH --models $ODOO_MODELS -t generator/cmd/tmpl/model.tmpl
package odoo

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/kolo/xmlrpc"
)

var (
	ErrClientConfigurationInvalid = errors.New("client configuration is invalid")
	ErrClientNotAuthenticate      = errors.New("client is not authenticate")
	ErrClientAuthentication       = errors.New("client authentication error: please verify client configuration")
	ErrNotFound                   = errors.New("not found")
	ErrPartiallyFound             = errors.New("partially found")
)

// ClientConfig is the configuration to create a new *Client by givin connection infomations.
type ClientConfig struct {
	Database          string
	Admin             string
	Password          string
	URL               string
	SyncWriteRequests bool
}

func (c *ClientConfig) valid() bool {
	return c.Database != "" &&
		c.Admin != "" &&
		c.Password != "" &&
		c.URL != ""
}

// Client provides high and low level functions to interact with odoo
type Client struct {
	common            *xmlrpc.Client
	object            *xmlrpc.Client
	cfg               *ClientConfig
	uid               int64
	auth              bool
	syncWriteRequests bool
	writeSyncer       *sync.Mutex
}

// NewClient creates a new *Client.
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
	if c.cfg.SyncWriteRequests {
		c.writeSyncer = &sync.Mutex{}
	}
	if err := c.authenticate(); err != nil {
		return nil, err
	}
	return c, nil
}

// Close closes all opened client connections.
func (c *Client) Close() {
	defer func() {
		if r := recover(); r != nil {
			log.Print("connections were already closed")
		}
	}()
	if c.common != nil {
		c.common.Close()
	}
	if c.object != nil {
		c.object.Close()
	}
}

// Version get informations about your odoo instance version.
func (c *Client) Version() (Version, error) {
	v := Version{}
	reply, err := c.commonCall("version", nil)
	if err != nil {
		return Version{}, err
	}
	convertFromDynamicToStatic(reply, &v)
	return v, nil
}

type combinedCriterionOperator string

const (
	and combinedCriterionOperator = "&"
	or                            = "|"
	not                           = "!"
)

type Criterion []interface{}

func (c *Criterion) ToInterface() []interface{} {
	if c != nil {
		return ([]interface{})(*c)
	}
	return []interface{}{}
}

type combinedCriterions struct {
	combinedCriterionOperator
	Criterions []*Criterion
}

/*
Criteria is a set of Criterion, each Criterion is a triple (field_name, operator, value).
It allow you to search models.
see documentation: https://www.odoo.com/documentation/13.0/reference/orm.html#reference-orm-domains
*/
type Criteria []interface{}

// NewCriteria creates a new *Criteria.
func NewCriteria() *Criteria {
	return &Criteria{}
}

func NewCriterion(field, operator string, value interface{}) *Criterion {
	c := Criterion(newTuple(field, operator, value))
	return &c
}

// Add a new Criterion to a *Criteria.
func (c *Criteria) Add(field, operator string, value interface{}) *Criteria {
	*c = append(*c, NewCriterion(field, operator, value).ToInterface())
	return c
}

// Add a new Criterion to a *Criteria.
func (c *Criteria) AddCriterion(cri *Criterion) *Criteria {
	*c = append(*c, cri.ToInterface())
	return c
}

func (c *Criteria) And(c1, c2 *Criterion) *Criteria {
	return c.combinedCriterions(and, c1, c2)
}

func (c *Criteria) Or(c1, c2 *Criterion) *Criteria {
	return c.combinedCriterions(or, c1, c2)
}

func (c *Criteria) Not(cri *Criterion) *Criteria {
	return c.combinedCriterions(not, cri)
}

func (c *Criteria) combinedCriterions(op combinedCriterionOperator, cc ...*Criterion) *Criteria {
	*c = append(*c, op)
	for _, cri := range cc {
		*c = append(*c, cri.ToInterface())
	}
	return c
}

// Options allow you to filter search results.
type Options map[string]interface{}

// NewOptions creates a new *Options
func NewOptions() *Options {
	opt := Options(make(map[string]interface{}))
	return &opt
}

// Offset adds the offset options.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#pagination
func (o *Options) Offset(offset int) *Options {
	return o.Add("offset", offset)
}

// Limit adds the limit options.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#pagination
func (o *Options) Limit(limit int) *Options {
	return o.Add("limit", limit)
}

// FetchFields allow you to precise the model fields you want odoo to return.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#search-and-read
func (o *Options) FetchFields(fields ...string) *Options {
	ff := []string{}
	for _, f := range fields {
		ff = append(ff, f)
	}
	return o.Add("fields", ff)
}

// AllFields is useful for FieldsGet function. It represents the fields to document
// you want odoo to return.
// https://www.odoo.com/documentation/13.0/reference/orm.html#fields-views
func (o *Options) AllFields(fields ...string) *Options {
	ff := []string{}
	for _, f := range fields {
		ff = append(ff, f)
	}
	return o.Add("allfields", ff)
}

// Attributes is useful for FieldsGet function. It represents the attributes to document
// you want odoo to return.
// https://www.odoo.com/documentation/13.0/reference/orm.html#fields-views
func (o *Options) Attributes(attributes ...string) *Options {
	aa := []string{}
	for _, a := range attributes {
		aa = append(aa, a)
	}
	return o.Add("attributes", aa)
}

// Add on option by providing option name and value.
func (o *Options) Add(opt string, v interface{}) *Options {
	(*o)[opt] = v
	return o
}

func getValuesFromInterface(v interface{}) map[string]interface{} {
	switch sv := v.(type) {
	case map[string]interface{}:
		return sv
	default:
		return convertFromStaticToDynamic(sv)
	}
}

// Create new model instances.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#create-records
func (c *Client) Create(model string, values []interface{}, options *Options) ([]int64, error) {
	var args []interface{}
	if len(values) == 0 {
		return nil, nil
	} else if len(values) == 1 { // ensure it works with odoo 11 that doesn't allow to create multiple instances with one API call
		args = []interface{}{getValuesFromInterface(values[0])}
	} else {
		vv := make([]interface{}, len(values))
		for i := 0; i < len(values); i++ {
			vv[i] = getValuesFromInterface(values[i])
		}
		args = []interface{}{vv}
	}
	resp, err := c.ExecuteKw("create", model, args, options)
	if err != nil {
		return nil, err
	}
	if len(values) == 1 {
		return []int64{resp.(int64)}, nil
	}
	r := resp.([]interface{})
	ids := make([]int64, len(r))
	for i := 0; i < len(ids); i++ {
		ids[i] = r[i].(int64)
	}
	return ids, nil
}

// Update existing model row(s).
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#update-records
func (c *Client) Update(model string, ids []int64, values interface{}, options *Options) error {
	v := getValuesFromInterface(values)
	_, err := c.ExecuteKw("write", model, []interface{}{ids, v}, options)
	if err != nil {
		return err
	}
	return nil
}

// Delete existing model row(s).
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#delete-records
func (c *Client) Delete(model string, ids []int64) error {
	_, err := c.ExecuteKw("unlink", model, []interface{}{ids}, nil)
	if err != nil {
		return err
	}
	return nil
}

// SearchRead search model records matching with *Criteria and read it.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#search-and-read
func (c *Client) SearchRead(model string, criteria *Criteria, options *Options, elem interface{}) error {
	resp, err := c.ExecuteKw("search_read", model, argsFromCriteria(criteria), options)
	if err != nil {
		return err
	}
	respLen := len(resp.([]interface{}))
	if respLen == 0 {
		return fmt.Errorf("%s model was %w with criteria %v and options %v", model, ErrNotFound, criteria, options)
	}
	if err := convertFromDynamicToStatic(resp, elem); err != nil {
		return err
	}
	return nil
}

// Read model records matching with ids.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#read-records
func (c *Client) Read(model string, ids []int64, options *Options, elem interface{}) error {
	resp, err := c.ExecuteKw("read", model, []interface{}{ids}, options)
	if err != nil {
		return err
	}
	respLen := len(resp.([]interface{}))
	if respLen == 0 {
		return fmt.Errorf("%s ids %v was %w with options %v", model, ids, ErrNotFound, options)
	}
	if err := convertFromDynamicToStatic(resp, elem); err != nil {
		return err
	}
	if respLen != len(ids) {
		return fmt.Errorf("%s ids %v was %w with options %v", model, ids, ErrPartiallyFound, options)
	}
	return nil
}

// Count model records matching with *Criteria.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#count-records
func (c *Client) Count(model string, criteria *Criteria, options *Options) (int64, error) {
	resp, err := c.ExecuteKw("search_count", model, argsFromCriteria(criteria), options)
	if err != nil {
		return -1, err
	}
	return resp.(int64), nil
}

// Search model record ids matching with *Criteria.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#list-records
func (c *Client) Search(model string, criteria *Criteria, options *Options) ([]int64, error) {
	resp, err := c.ExecuteKw("search", model, argsFromCriteria(criteria), options)
	if err != nil {
		return []int64{}, err
	}
	respLen := len(resp.([]interface{}))
	if respLen == 0 {
		return []int64{}, fmt.Errorf("%s model was %w with criteria %v and options %v", model, ErrNotFound, criteria, options)
	}
	return sliceInterfaceToInt64Slice(resp.([]interface{})), nil
}

// FieldsGet inspect model fields.
// https://www.odoo.com/documentation/13.0/webservices/odoo.html#listing-record-fields
func (c *Client) FieldsGet(model string, options *Options) (map[string]interface{}, error) {
	resp, err := c.ExecuteKw("fields_get", model, nil, options)
	if err != nil {
		return nil, nil
	}
	return resp.(map[string]interface{}), nil
}

// ExecuteKw is a RPC function. The lowest library function. It is use for all
// function related to "xmlrpc/2/object" endpoint.
func (c *Client) ExecuteKw(method, model string, args []interface{}, options *Options) (interface{}, error) {
	if err := c.checkForAuthentication(); err != nil {
		return nil, err
	}
	if c.cfg.SyncWriteRequests && isWriteMethod(method) {
		c.writeSyncer.Lock()
		defer c.writeSyncer.Unlock()
	}

	resp, err := c.objectCall("execute_kw", []interface{}{c.cfg.Database, c.uid, c.cfg.Password, model, method, args, options})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) authenticate() error {
	if !c.isAuthenticate() {
		resp, err := c.commonCall("authenticate", []interface{}{c.cfg.Database, c.cfg.Admin, c.cfg.Password, ""})
		if err != nil {
			return err
		}
		if _, ok := resp.(bool); ok {
			return ErrClientAuthentication
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
		newClient, err := xmlrpc.NewClient(c.cfg.URL+path, nil)
		if err != nil {
			return err
		}
		*x = *newClient
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

func isWriteMethod(method string) bool {
	switch method {
	case "create", "write", "unlink":
		return true
	default:
		return false
	}
}
