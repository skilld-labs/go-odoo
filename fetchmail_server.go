package odoo

import (
	"fmt"
)

// FetchmailServer represents fetchmail.server model.
type FetchmailServer struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	ActionId      *Many2One  `xmlrpc:"action_id,omptempty"`
	Active        *Bool      `xmlrpc:"active,omptempty"`
	Attach        *Bool      `xmlrpc:"attach,omptempty"`
	Configuration *String    `xmlrpc:"configuration,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date          *Time      `xmlrpc:"date,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	IsSsl         *Bool      `xmlrpc:"is_ssl,omptempty"`
	MessageIds    *Relation  `xmlrpc:"message_ids,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	ObjectId      *Many2One  `xmlrpc:"object_id,omptempty"`
	Original      *Bool      `xmlrpc:"original,omptempty"`
	Password      *String    `xmlrpc:"password,omptempty"`
	Port          *Int       `xmlrpc:"port,omptempty"`
	Priority      *Int       `xmlrpc:"priority,omptempty"`
	Script        *String    `xmlrpc:"script,omptempty"`
	Server        *String    `xmlrpc:"server,omptempty"`
	State         *Selection `xmlrpc:"state,omptempty"`
	Type          *Selection `xmlrpc:"type,omptempty"`
	User          *String    `xmlrpc:"user,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// FetchmailServers represents array of fetchmail.server model.
type FetchmailServers []FetchmailServer

// FetchmailServerModel is the odoo model name.
const FetchmailServerModel = "fetchmail.server"

// Many2One convert FetchmailServer to *Many2One.
func (fs *FetchmailServer) Many2One() *Many2One {
	return NewMany2One(fs.Id.Get(), "")
}

// CreateFetchmailServer creates a new fetchmail.server model and returns its id.
func (c *Client) CreateFetchmailServer(fs *FetchmailServer) (int64, error) {
	return c.Create(FetchmailServerModel, fs)
}

// UpdateFetchmailServer updates an existing fetchmail.server record.
func (c *Client) UpdateFetchmailServer(fs *FetchmailServer) error {
	return c.UpdateFetchmailServers([]int64{fs.Id.Get()}, fs)
}

// UpdateFetchmailServers updates existing fetchmail.server records.
// All records (represented by ids) will be updated by fs values.
func (c *Client) UpdateFetchmailServers(ids []int64, fs *FetchmailServer) error {
	return c.Update(FetchmailServerModel, ids, fs)
}

// DeleteFetchmailServer deletes an existing fetchmail.server record.
func (c *Client) DeleteFetchmailServer(id int64) error {
	return c.DeleteFetchmailServers([]int64{id})
}

// DeleteFetchmailServers deletes existing fetchmail.server records.
func (c *Client) DeleteFetchmailServers(ids []int64) error {
	return c.Delete(FetchmailServerModel, ids)
}

// GetFetchmailServer gets fetchmail.server existing record.
func (c *Client) GetFetchmailServer(id int64) (*FetchmailServer, error) {
	fss, err := c.GetFetchmailServers([]int64{id})
	if err != nil {
		return nil, err
	}
	if fss != nil && len(*fss) > 0 {
		return &((*fss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of fetchmail.server not found", id)
}

// GetFetchmailServers gets fetchmail.server existing records.
func (c *Client) GetFetchmailServers(ids []int64) (*FetchmailServers, error) {
	fss := &FetchmailServers{}
	if err := c.Read(FetchmailServerModel, ids, nil, fss); err != nil {
		return nil, err
	}
	return fss, nil
}

// FindFetchmailServer finds fetchmail.server record by querying it with criteria.
func (c *Client) FindFetchmailServer(criteria *Criteria) (*FetchmailServer, error) {
	fss := &FetchmailServers{}
	if err := c.SearchRead(FetchmailServerModel, criteria, NewOptions().Limit(1), fss); err != nil {
		return nil, err
	}
	if fss != nil && len(*fss) > 0 {
		return &((*fss)[0]), nil
	}
	return nil, fmt.Errorf("fetchmail.server was not found with criteria %v", criteria)
}

// FindFetchmailServers finds fetchmail.server records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFetchmailServers(criteria *Criteria, options *Options) (*FetchmailServers, error) {
	fss := &FetchmailServers{}
	if err := c.SearchRead(FetchmailServerModel, criteria, options, fss); err != nil {
		return nil, err
	}
	return fss, nil
}

// FindFetchmailServerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindFetchmailServerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(FetchmailServerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFetchmailServerId finds record id by querying it with criteria.
func (c *Client) FindFetchmailServerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FetchmailServerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("fetchmail.server was not found with criteria %v and options %v", criteria, options)
}
