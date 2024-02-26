package odoo

// IrMailServer represents ir.mail_server model.
type IrMailServer struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	Active         *Bool      `xmlrpc:"active,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	Sequence       *Int       `xmlrpc:"sequence,omptempty"`
	SmtpDebug      *Bool      `xmlrpc:"smtp_debug,omptempty"`
	SmtpEncryption *Selection `xmlrpc:"smtp_encryption,omptempty"`
	SmtpHost       *String    `xmlrpc:"smtp_host,omptempty"`
	SmtpPass       *String    `xmlrpc:"smtp_pass,omptempty"`
	SmtpPort       *Int       `xmlrpc:"smtp_port,omptempty"`
	SmtpUser       *String    `xmlrpc:"smtp_user,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrMailServers represents array of ir.mail_server model.
type IrMailServers []IrMailServer

// IrMailServerModel is the odoo model name.
const IrMailServerModel = "ir.mail_server"

// Many2One convert IrMailServer to *Many2One.
func (im *IrMailServer) Many2One() *Many2One {
	return NewMany2One(im.Id.Get(), "")
}

// CreateIrMailServer creates a new ir.mail_server model and returns its id.
func (c *Client) CreateIrMailServer(im *IrMailServer) (int64, error) {
	ids, err := c.CreateIrMailServers([]*IrMailServer{im})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrMailServer creates a new ir.mail_server model and returns its id.
func (c *Client) CreateIrMailServers(ims []*IrMailServer) ([]int64, error) {
	var vv []interface{}
	for _, v := range ims {
		vv = append(vv, v)
	}
	return c.Create(IrMailServerModel, vv, nil)
}

// UpdateIrMailServer updates an existing ir.mail_server record.
func (c *Client) UpdateIrMailServer(im *IrMailServer) error {
	return c.UpdateIrMailServers([]int64{im.Id.Get()}, im)
}

// UpdateIrMailServers updates existing ir.mail_server records.
// All records (represented by ids) will be updated by im values.
func (c *Client) UpdateIrMailServers(ids []int64, im *IrMailServer) error {
	return c.Update(IrMailServerModel, ids, im, nil)
}

// DeleteIrMailServer deletes an existing ir.mail_server record.
func (c *Client) DeleteIrMailServer(id int64) error {
	return c.DeleteIrMailServers([]int64{id})
}

// DeleteIrMailServers deletes existing ir.mail_server records.
func (c *Client) DeleteIrMailServers(ids []int64) error {
	return c.Delete(IrMailServerModel, ids)
}

// GetIrMailServer gets ir.mail_server existing record.
func (c *Client) GetIrMailServer(id int64) (*IrMailServer, error) {
	ims, err := c.GetIrMailServers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ims)[0]), nil
}

// GetIrMailServers gets ir.mail_server existing records.
func (c *Client) GetIrMailServers(ids []int64) (*IrMailServers, error) {
	ims := &IrMailServers{}
	if err := c.Read(IrMailServerModel, ids, nil, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindIrMailServer finds ir.mail_server record by querying it with criteria.
func (c *Client) FindIrMailServer(criteria *Criteria) (*IrMailServer, error) {
	ims := &IrMailServers{}
	if err := c.SearchRead(IrMailServerModel, criteria, NewOptions().Limit(1), ims); err != nil {
		return nil, err
	}
	return &((*ims)[0]), nil
}

// FindIrMailServers finds ir.mail_server records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrMailServers(criteria *Criteria, options *Options) (*IrMailServers, error) {
	ims := &IrMailServers{}
	if err := c.SearchRead(IrMailServerModel, criteria, options, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindIrMailServerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrMailServerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrMailServerModel, criteria, options)
}

// FindIrMailServerId finds record id by querying it with criteria.
func (c *Client) FindIrMailServerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrMailServerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
