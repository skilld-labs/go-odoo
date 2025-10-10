package odoo

// MailIceServer represents mail.ice.server model.
type MailIceServer struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Credential  *String    `xmlrpc:"credential,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	ServerType  *Selection `xmlrpc:"server_type,omitempty"`
	Uri         *String    `xmlrpc:"uri,omitempty"`
	Username    *String    `xmlrpc:"username,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailIceServers represents array of mail.ice.server model.
type MailIceServers []MailIceServer

// MailIceServerModel is the odoo model name.
const MailIceServerModel = "mail.ice.server"

// Many2One convert MailIceServer to *Many2One.
func (mis *MailIceServer) Many2One() *Many2One {
	return NewMany2One(mis.Id.Get(), "")
}

// CreateMailIceServer creates a new mail.ice.server model and returns its id.
func (c *Client) CreateMailIceServer(mis *MailIceServer) (int64, error) {
	ids, err := c.CreateMailIceServers([]*MailIceServer{mis})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailIceServer creates a new mail.ice.server model and returns its id.
func (c *Client) CreateMailIceServers(miss []*MailIceServer) ([]int64, error) {
	var vv []interface{}
	for _, v := range miss {
		vv = append(vv, v)
	}
	return c.Create(MailIceServerModel, vv, nil)
}

// UpdateMailIceServer updates an existing mail.ice.server record.
func (c *Client) UpdateMailIceServer(mis *MailIceServer) error {
	return c.UpdateMailIceServers([]int64{mis.Id.Get()}, mis)
}

// UpdateMailIceServers updates existing mail.ice.server records.
// All records (represented by ids) will be updated by mis values.
func (c *Client) UpdateMailIceServers(ids []int64, mis *MailIceServer) error {
	return c.Update(MailIceServerModel, ids, mis, nil)
}

// DeleteMailIceServer deletes an existing mail.ice.server record.
func (c *Client) DeleteMailIceServer(id int64) error {
	return c.DeleteMailIceServers([]int64{id})
}

// DeleteMailIceServers deletes existing mail.ice.server records.
func (c *Client) DeleteMailIceServers(ids []int64) error {
	return c.Delete(MailIceServerModel, ids)
}

// GetMailIceServer gets mail.ice.server existing record.
func (c *Client) GetMailIceServer(id int64) (*MailIceServer, error) {
	miss, err := c.GetMailIceServers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*miss)[0]), nil
}

// GetMailIceServers gets mail.ice.server existing records.
func (c *Client) GetMailIceServers(ids []int64) (*MailIceServers, error) {
	miss := &MailIceServers{}
	if err := c.Read(MailIceServerModel, ids, nil, miss); err != nil {
		return nil, err
	}
	return miss, nil
}

// FindMailIceServer finds mail.ice.server record by querying it with criteria.
func (c *Client) FindMailIceServer(criteria *Criteria) (*MailIceServer, error) {
	miss := &MailIceServers{}
	if err := c.SearchRead(MailIceServerModel, criteria, NewOptions().Limit(1), miss); err != nil {
		return nil, err
	}
	return &((*miss)[0]), nil
}

// FindMailIceServers finds mail.ice.server records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailIceServers(criteria *Criteria, options *Options) (*MailIceServers, error) {
	miss := &MailIceServers{}
	if err := c.SearchRead(MailIceServerModel, criteria, options, miss); err != nil {
		return nil, err
	}
	return miss, nil
}

// FindMailIceServerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailIceServerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailIceServerModel, criteria, options)
}

// FindMailIceServerId finds record id by querying it with criteria.
func (c *Client) FindMailIceServerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailIceServerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
