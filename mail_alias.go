package odoo

// MailAlias represents mail.alias model.
type MailAlias struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	AliasContact        *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults       *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain         *String    `xmlrpc:"alias_domain,omitempty"`
	AliasForceThreadId  *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasModelId        *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName           *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId  *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasUserId         *Many2One  `xmlrpc:"alias_user_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailAliass represents array of mail.alias model.
type MailAliass []MailAlias

// MailAliasModel is the odoo model name.
const MailAliasModel = "mail.alias"

// Many2One convert MailAlias to *Many2One.
func (ma *MailAlias) Many2One() *Many2One {
	return NewMany2One(ma.Id.Get(), "")
}

// CreateMailAlias creates a new mail.alias model and returns its id.
func (c *Client) CreateMailAlias(ma *MailAlias) (int64, error) {
	ids, err := c.CreateMailAliass([]*MailAlias{ma})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailAlias creates a new mail.alias model and returns its id.
func (c *Client) CreateMailAliass(mas []*MailAlias) ([]int64, error) {
	var vv []interface{}
	for _, v := range mas {
		vv = append(vv, v)
	}
	return c.Create(MailAliasModel, vv, nil)
}

// UpdateMailAlias updates an existing mail.alias record.
func (c *Client) UpdateMailAlias(ma *MailAlias) error {
	return c.UpdateMailAliass([]int64{ma.Id.Get()}, ma)
}

// UpdateMailAliass updates existing mail.alias records.
// All records (represented by ids) will be updated by ma values.
func (c *Client) UpdateMailAliass(ids []int64, ma *MailAlias) error {
	return c.Update(MailAliasModel, ids, ma, nil)
}

// DeleteMailAlias deletes an existing mail.alias record.
func (c *Client) DeleteMailAlias(id int64) error {
	return c.DeleteMailAliass([]int64{id})
}

// DeleteMailAliass deletes existing mail.alias records.
func (c *Client) DeleteMailAliass(ids []int64) error {
	return c.Delete(MailAliasModel, ids)
}

// GetMailAlias gets mail.alias existing record.
func (c *Client) GetMailAlias(id int64) (*MailAlias, error) {
	mas, err := c.GetMailAliass([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mas)[0]), nil
}

// GetMailAliass gets mail.alias existing records.
func (c *Client) GetMailAliass(ids []int64) (*MailAliass, error) {
	mas := &MailAliass{}
	if err := c.Read(MailAliasModel, ids, nil, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMailAlias finds mail.alias record by querying it with criteria.
func (c *Client) FindMailAlias(criteria *Criteria) (*MailAlias, error) {
	mas := &MailAliass{}
	if err := c.SearchRead(MailAliasModel, criteria, NewOptions().Limit(1), mas); err != nil {
		return nil, err
	}
	return &((*mas)[0]), nil
}

// FindMailAliass finds mail.alias records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAliass(criteria *Criteria, options *Options) (*MailAliass, error) {
	mas := &MailAliass{}
	if err := c.SearchRead(MailAliasModel, criteria, options, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMailAliasIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAliasIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailAliasModel, criteria, options)
}

// FindMailAliasId finds record id by querying it with criteria.
func (c *Client) FindMailAliasId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailAliasModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
