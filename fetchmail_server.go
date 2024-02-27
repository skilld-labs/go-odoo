package odoo

// FetchmailServer represents fetchmail.server model.
type FetchmailServer struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	ActionId      *Many2One  `xmlrpc:"action_id,omitempty"`
	Active        *Bool      `xmlrpc:"active,omitempty"`
	Attach        *Bool      `xmlrpc:"attach,omitempty"`
	Configuration *String    `xmlrpc:"configuration,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date          *Time      `xmlrpc:"date,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	IsSsl         *Bool      `xmlrpc:"is_ssl,omitempty"`
	MessageIds    *Relation  `xmlrpc:"message_ids,omitempty"`
	Name          *String    `xmlrpc:"name,omitempty"`
	ObjectId      *Many2One  `xmlrpc:"object_id,omitempty"`
	Original      *Bool      `xmlrpc:"original,omitempty"`
	Password      *String    `xmlrpc:"password,omitempty"`
	Port          *Int       `xmlrpc:"port,omitempty"`
	Priority      *Int       `xmlrpc:"priority,omitempty"`
	Script        *String    `xmlrpc:"script,omitempty"`
	Server        *String    `xmlrpc:"server,omitempty"`
	State         *Selection `xmlrpc:"state,omitempty"`
	Type          *Selection `xmlrpc:"type,omitempty"`
	User          *String    `xmlrpc:"user,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateFetchmailServers([]*FetchmailServer{fs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateFetchmailServers creates a new fetchmail.server model and returns its id.
func (c *Client) CreateFetchmailServers(fss []*FetchmailServer) ([]int64, error) {
	var vv []interface{}
	for _, v := range fss {
		vv = append(vv, v)
	}
	return c.Create(FetchmailServerModel, vv, nil)
}

// UpdateFetchmailServer updates an existing fetchmail.server record.
func (c *Client) UpdateFetchmailServer(fs *FetchmailServer) error {
	return c.UpdateFetchmailServers([]int64{fs.Id.Get()}, fs)
}

// UpdateFetchmailServers updates existing fetchmail.server records.
// All records (represented by ids) will be updated by fs values.
func (c *Client) UpdateFetchmailServers(ids []int64, fs *FetchmailServer) error {
	return c.Update(FetchmailServerModel, ids, fs, nil)
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
	return &((*fss)[0]), nil
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
	return &((*fss)[0]), nil
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
	return c.Search(FetchmailServerModel, criteria, options)
}

// FindFetchmailServerId finds record id by querying it with criteria.
func (c *Client) FindFetchmailServerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FetchmailServerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
