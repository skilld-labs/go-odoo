package odoo

// MailGatewayAllowed represents mail.gateway.allowed model.
type MailGatewayAllowed struct {
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Email           *String   `xmlrpc:"email,omitempty"`
	EmailNormalized *String   `xmlrpc:"email_normalized,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailGatewayAlloweds represents array of mail.gateway.allowed model.
type MailGatewayAlloweds []MailGatewayAllowed

// MailGatewayAllowedModel is the odoo model name.
const MailGatewayAllowedModel = "mail.gateway.allowed"

// Many2One convert MailGatewayAllowed to *Many2One.
func (mga *MailGatewayAllowed) Many2One() *Many2One {
	return NewMany2One(mga.Id.Get(), "")
}

// CreateMailGatewayAllowed creates a new mail.gateway.allowed model and returns its id.
func (c *Client) CreateMailGatewayAllowed(mga *MailGatewayAllowed) (int64, error) {
	ids, err := c.CreateMailGatewayAlloweds([]*MailGatewayAllowed{mga})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailGatewayAllowed creates a new mail.gateway.allowed model and returns its id.
func (c *Client) CreateMailGatewayAlloweds(mgas []*MailGatewayAllowed) ([]int64, error) {
	var vv []interface{}
	for _, v := range mgas {
		vv = append(vv, v)
	}
	return c.Create(MailGatewayAllowedModel, vv, nil)
}

// UpdateMailGatewayAllowed updates an existing mail.gateway.allowed record.
func (c *Client) UpdateMailGatewayAllowed(mga *MailGatewayAllowed) error {
	return c.UpdateMailGatewayAlloweds([]int64{mga.Id.Get()}, mga)
}

// UpdateMailGatewayAlloweds updates existing mail.gateway.allowed records.
// All records (represented by ids) will be updated by mga values.
func (c *Client) UpdateMailGatewayAlloweds(ids []int64, mga *MailGatewayAllowed) error {
	return c.Update(MailGatewayAllowedModel, ids, mga, nil)
}

// DeleteMailGatewayAllowed deletes an existing mail.gateway.allowed record.
func (c *Client) DeleteMailGatewayAllowed(id int64) error {
	return c.DeleteMailGatewayAlloweds([]int64{id})
}

// DeleteMailGatewayAlloweds deletes existing mail.gateway.allowed records.
func (c *Client) DeleteMailGatewayAlloweds(ids []int64) error {
	return c.Delete(MailGatewayAllowedModel, ids)
}

// GetMailGatewayAllowed gets mail.gateway.allowed existing record.
func (c *Client) GetMailGatewayAllowed(id int64) (*MailGatewayAllowed, error) {
	mgas, err := c.GetMailGatewayAlloweds([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mgas)[0]), nil
}

// GetMailGatewayAlloweds gets mail.gateway.allowed existing records.
func (c *Client) GetMailGatewayAlloweds(ids []int64) (*MailGatewayAlloweds, error) {
	mgas := &MailGatewayAlloweds{}
	if err := c.Read(MailGatewayAllowedModel, ids, nil, mgas); err != nil {
		return nil, err
	}
	return mgas, nil
}

// FindMailGatewayAllowed finds mail.gateway.allowed record by querying it with criteria.
func (c *Client) FindMailGatewayAllowed(criteria *Criteria) (*MailGatewayAllowed, error) {
	mgas := &MailGatewayAlloweds{}
	if err := c.SearchRead(MailGatewayAllowedModel, criteria, NewOptions().Limit(1), mgas); err != nil {
		return nil, err
	}
	return &((*mgas)[0]), nil
}

// FindMailGatewayAlloweds finds mail.gateway.allowed records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailGatewayAlloweds(criteria *Criteria, options *Options) (*MailGatewayAlloweds, error) {
	mgas := &MailGatewayAlloweds{}
	if err := c.SearchRead(MailGatewayAllowedModel, criteria, options, mgas); err != nil {
		return nil, err
	}
	return mgas, nil
}

// FindMailGatewayAllowedIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailGatewayAllowedIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailGatewayAllowedModel, criteria, options)
}

// FindMailGatewayAllowedId finds record id by querying it with criteria.
func (c *Client) FindMailGatewayAllowedId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailGatewayAllowedModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
