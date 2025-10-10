package odoo

// ResCompanyLdap represents res.company.ldap model.
type ResCompanyLdap struct {
	Company        *Many2One `xmlrpc:"company,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	CreateUser     *Bool     `xmlrpc:"create_user,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	LdapBase       *String   `xmlrpc:"ldap_base,omitempty"`
	LdapBinddn     *String   `xmlrpc:"ldap_binddn,omitempty"`
	LdapFilter     *String   `xmlrpc:"ldap_filter,omitempty"`
	LdapPassword   *String   `xmlrpc:"ldap_password,omitempty"`
	LdapServer     *String   `xmlrpc:"ldap_server,omitempty"`
	LdapServerPort *Int      `xmlrpc:"ldap_server_port,omitempty"`
	LdapTls        *Bool     `xmlrpc:"ldap_tls,omitempty"`
	MailAttribute  *String   `xmlrpc:"mail_attribute,omitempty"`
	NameAttribute  *String   `xmlrpc:"name_attribute,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	User           *Many2One `xmlrpc:"user,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResCompanyLdaps represents array of res.company.ldap model.
type ResCompanyLdaps []ResCompanyLdap

// ResCompanyLdapModel is the odoo model name.
const ResCompanyLdapModel = "res.company.ldap"

// Many2One convert ResCompanyLdap to *Many2One.
func (rcl *ResCompanyLdap) Many2One() *Many2One {
	return NewMany2One(rcl.Id.Get(), "")
}

// CreateResCompanyLdap creates a new res.company.ldap model and returns its id.
func (c *Client) CreateResCompanyLdap(rcl *ResCompanyLdap) (int64, error) {
	ids, err := c.CreateResCompanyLdaps([]*ResCompanyLdap{rcl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCompanyLdap creates a new res.company.ldap model and returns its id.
func (c *Client) CreateResCompanyLdaps(rcls []*ResCompanyLdap) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcls {
		vv = append(vv, v)
	}
	return c.Create(ResCompanyLdapModel, vv, nil)
}

// UpdateResCompanyLdap updates an existing res.company.ldap record.
func (c *Client) UpdateResCompanyLdap(rcl *ResCompanyLdap) error {
	return c.UpdateResCompanyLdaps([]int64{rcl.Id.Get()}, rcl)
}

// UpdateResCompanyLdaps updates existing res.company.ldap records.
// All records (represented by ids) will be updated by rcl values.
func (c *Client) UpdateResCompanyLdaps(ids []int64, rcl *ResCompanyLdap) error {
	return c.Update(ResCompanyLdapModel, ids, rcl, nil)
}

// DeleteResCompanyLdap deletes an existing res.company.ldap record.
func (c *Client) DeleteResCompanyLdap(id int64) error {
	return c.DeleteResCompanyLdaps([]int64{id})
}

// DeleteResCompanyLdaps deletes existing res.company.ldap records.
func (c *Client) DeleteResCompanyLdaps(ids []int64) error {
	return c.Delete(ResCompanyLdapModel, ids)
}

// GetResCompanyLdap gets res.company.ldap existing record.
func (c *Client) GetResCompanyLdap(id int64) (*ResCompanyLdap, error) {
	rcls, err := c.GetResCompanyLdaps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcls)[0]), nil
}

// GetResCompanyLdaps gets res.company.ldap existing records.
func (c *Client) GetResCompanyLdaps(ids []int64) (*ResCompanyLdaps, error) {
	rcls := &ResCompanyLdaps{}
	if err := c.Read(ResCompanyLdapModel, ids, nil, rcls); err != nil {
		return nil, err
	}
	return rcls, nil
}

// FindResCompanyLdap finds res.company.ldap record by querying it with criteria.
func (c *Client) FindResCompanyLdap(criteria *Criteria) (*ResCompanyLdap, error) {
	rcls := &ResCompanyLdaps{}
	if err := c.SearchRead(ResCompanyLdapModel, criteria, NewOptions().Limit(1), rcls); err != nil {
		return nil, err
	}
	return &((*rcls)[0]), nil
}

// FindResCompanyLdaps finds res.company.ldap records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyLdaps(criteria *Criteria, options *Options) (*ResCompanyLdaps, error) {
	rcls := &ResCompanyLdaps{}
	if err := c.SearchRead(ResCompanyLdapModel, criteria, options, rcls); err != nil {
		return nil, err
	}
	return rcls, nil
}

// FindResCompanyLdapIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyLdapIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResCompanyLdapModel, criteria, options)
}

// FindResCompanyLdapId finds record id by querying it with criteria.
func (c *Client) FindResCompanyLdapId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyLdapModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
