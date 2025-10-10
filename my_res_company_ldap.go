package odoo

// MyResCompanyLdap represents my.res.company.ldap model.
type MyResCompanyLdap struct {
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

// MyResCompanyLdaps represents array of my.res.company.ldap model.
type MyResCompanyLdaps []MyResCompanyLdap

// MyResCompanyLdapModel is the odoo model name.
const MyResCompanyLdapModel = "my.res.company.ldap"

// Many2One convert MyResCompanyLdap to *Many2One.
func (mrcl *MyResCompanyLdap) Many2One() *Many2One {
	return NewMany2One(mrcl.Id.Get(), "")
}

// CreateMyResCompanyLdap creates a new my.res.company.ldap model and returns its id.
func (c *Client) CreateMyResCompanyLdap(mrcl *MyResCompanyLdap) (int64, error) {
	ids, err := c.CreateMyResCompanyLdaps([]*MyResCompanyLdap{mrcl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMyResCompanyLdap creates a new my.res.company.ldap model and returns its id.
func (c *Client) CreateMyResCompanyLdaps(mrcls []*MyResCompanyLdap) ([]int64, error) {
	var vv []interface{}
	for _, v := range mrcls {
		vv = append(vv, v)
	}
	return c.Create(MyResCompanyLdapModel, vv, nil)
}

// UpdateMyResCompanyLdap updates an existing my.res.company.ldap record.
func (c *Client) UpdateMyResCompanyLdap(mrcl *MyResCompanyLdap) error {
	return c.UpdateMyResCompanyLdaps([]int64{mrcl.Id.Get()}, mrcl)
}

// UpdateMyResCompanyLdaps updates existing my.res.company.ldap records.
// All records (represented by ids) will be updated by mrcl values.
func (c *Client) UpdateMyResCompanyLdaps(ids []int64, mrcl *MyResCompanyLdap) error {
	return c.Update(MyResCompanyLdapModel, ids, mrcl, nil)
}

// DeleteMyResCompanyLdap deletes an existing my.res.company.ldap record.
func (c *Client) DeleteMyResCompanyLdap(id int64) error {
	return c.DeleteMyResCompanyLdaps([]int64{id})
}

// DeleteMyResCompanyLdaps deletes existing my.res.company.ldap records.
func (c *Client) DeleteMyResCompanyLdaps(ids []int64) error {
	return c.Delete(MyResCompanyLdapModel, ids)
}

// GetMyResCompanyLdap gets my.res.company.ldap existing record.
func (c *Client) GetMyResCompanyLdap(id int64) (*MyResCompanyLdap, error) {
	mrcls, err := c.GetMyResCompanyLdaps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mrcls)[0]), nil
}

// GetMyResCompanyLdaps gets my.res.company.ldap existing records.
func (c *Client) GetMyResCompanyLdaps(ids []int64) (*MyResCompanyLdaps, error) {
	mrcls := &MyResCompanyLdaps{}
	if err := c.Read(MyResCompanyLdapModel, ids, nil, mrcls); err != nil {
		return nil, err
	}
	return mrcls, nil
}

// FindMyResCompanyLdap finds my.res.company.ldap record by querying it with criteria.
func (c *Client) FindMyResCompanyLdap(criteria *Criteria) (*MyResCompanyLdap, error) {
	mrcls := &MyResCompanyLdaps{}
	if err := c.SearchRead(MyResCompanyLdapModel, criteria, NewOptions().Limit(1), mrcls); err != nil {
		return nil, err
	}
	return &((*mrcls)[0]), nil
}

// FindMyResCompanyLdaps finds my.res.company.ldap records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMyResCompanyLdaps(criteria *Criteria, options *Options) (*MyResCompanyLdaps, error) {
	mrcls := &MyResCompanyLdaps{}
	if err := c.SearchRead(MyResCompanyLdapModel, criteria, options, mrcls); err != nil {
		return nil, err
	}
	return mrcls, nil
}

// FindMyResCompanyLdapIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMyResCompanyLdapIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MyResCompanyLdapModel, criteria, options)
}

// FindMyResCompanyLdapId finds record id by querying it with criteria.
func (c *Client) FindMyResCompanyLdapId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MyResCompanyLdapModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
