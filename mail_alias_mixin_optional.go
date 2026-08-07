package odoo

// MailAliasMixinOptional represents mail.alias.mixin.optional model.
type MailAliasMixinOptional struct {
	AliasDefaults *String   `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain   *String   `xmlrpc:"alias_domain,omitempty"`
	AliasDomainId *Many2One `xmlrpc:"alias_domain_id,omitempty"`
	AliasEmail    *String   `xmlrpc:"alias_email,omitempty"`
	AliasId       *Many2One `xmlrpc:"alias_id,omitempty"`
	AliasName     *String   `xmlrpc:"alias_name,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
}

// MailAliasMixinOptionals represents array of mail.alias.mixin.optional model.
type MailAliasMixinOptionals []MailAliasMixinOptional

// MailAliasMixinOptionalModel is the odoo model name.
const MailAliasMixinOptionalModel = "mail.alias.mixin.optional"

// Many2One convert MailAliasMixinOptional to *Many2One.
func (mamo *MailAliasMixinOptional) Many2One() *Many2One {
	return NewMany2One(mamo.Id.Get(), "")
}

// CreateMailAliasMixinOptional creates a new mail.alias.mixin.optional model and returns its id.
func (c *Client) CreateMailAliasMixinOptional(mamo *MailAliasMixinOptional) (int64, error) {
	ids, err := c.CreateMailAliasMixinOptionals([]*MailAliasMixinOptional{mamo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailAliasMixinOptional creates a new mail.alias.mixin.optional model and returns its id.
func (c *Client) CreateMailAliasMixinOptionals(mamos []*MailAliasMixinOptional) ([]int64, error) {
	var vv []interface{}
	for _, v := range mamos {
		vv = append(vv, v)
	}
	return c.Create(MailAliasMixinOptionalModel, vv, nil)
}

// UpdateMailAliasMixinOptional updates an existing mail.alias.mixin.optional record.
func (c *Client) UpdateMailAliasMixinOptional(mamo *MailAliasMixinOptional) error {
	return c.UpdateMailAliasMixinOptionals([]int64{mamo.Id.Get()}, mamo)
}

// UpdateMailAliasMixinOptionals updates existing mail.alias.mixin.optional records.
// All records (represented by ids) will be updated by mamo values.
func (c *Client) UpdateMailAliasMixinOptionals(ids []int64, mamo *MailAliasMixinOptional) error {
	return c.Update(MailAliasMixinOptionalModel, ids, mamo, nil)
}

// DeleteMailAliasMixinOptional deletes an existing mail.alias.mixin.optional record.
func (c *Client) DeleteMailAliasMixinOptional(id int64) error {
	return c.DeleteMailAliasMixinOptionals([]int64{id})
}

// DeleteMailAliasMixinOptionals deletes existing mail.alias.mixin.optional records.
func (c *Client) DeleteMailAliasMixinOptionals(ids []int64) error {
	return c.Delete(MailAliasMixinOptionalModel, ids)
}

// GetMailAliasMixinOptional gets mail.alias.mixin.optional existing record.
func (c *Client) GetMailAliasMixinOptional(id int64) (*MailAliasMixinOptional, error) {
	mamos, err := c.GetMailAliasMixinOptionals([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mamos)[0]), nil
}

// GetMailAliasMixinOptionals gets mail.alias.mixin.optional existing records.
func (c *Client) GetMailAliasMixinOptionals(ids []int64) (*MailAliasMixinOptionals, error) {
	mamos := &MailAliasMixinOptionals{}
	if err := c.Read(MailAliasMixinOptionalModel, ids, nil, mamos); err != nil {
		return nil, err
	}
	return mamos, nil
}

// FindMailAliasMixinOptional finds mail.alias.mixin.optional record by querying it with criteria.
func (c *Client) FindMailAliasMixinOptional(criteria *Criteria) (*MailAliasMixinOptional, error) {
	mamos := &MailAliasMixinOptionals{}
	if err := c.SearchRead(MailAliasMixinOptionalModel, criteria, NewOptions().Limit(1), mamos); err != nil {
		return nil, err
	}
	return &((*mamos)[0]), nil
}

// FindMailAliasMixinOptionals finds mail.alias.mixin.optional records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAliasMixinOptionals(criteria *Criteria, options *Options) (*MailAliasMixinOptionals, error) {
	mamos := &MailAliasMixinOptionals{}
	if err := c.SearchRead(MailAliasMixinOptionalModel, criteria, options, mamos); err != nil {
		return nil, err
	}
	return mamos, nil
}

// FindMailAliasMixinOptionalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAliasMixinOptionalIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailAliasMixinOptionalModel, criteria, options)
}

// FindMailAliasMixinOptionalId finds record id by querying it with criteria.
func (c *Client) FindMailAliasMixinOptionalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailAliasMixinOptionalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
