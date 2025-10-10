package odoo

// MailAliasDomain represents mail.alias.domain model.
type MailAliasDomain struct {
	BounceAlias      *String   `xmlrpc:"bounce_alias,omitempty"`
	BounceEmail      *String   `xmlrpc:"bounce_email,omitempty"`
	CatchallAlias    *String   `xmlrpc:"catchall_alias,omitempty"`
	CatchallEmail    *String   `xmlrpc:"catchall_email,omitempty"`
	CompanyIds       *Relation `xmlrpc:"company_ids,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DefaultFrom      *String   `xmlrpc:"default_from,omitempty"`
	DefaultFromEmail *String   `xmlrpc:"default_from_email,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	Sequence         *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailAliasDomains represents array of mail.alias.domain model.
type MailAliasDomains []MailAliasDomain

// MailAliasDomainModel is the odoo model name.
const MailAliasDomainModel = "mail.alias.domain"

// Many2One convert MailAliasDomain to *Many2One.
func (mad *MailAliasDomain) Many2One() *Many2One {
	return NewMany2One(mad.Id.Get(), "")
}

// CreateMailAliasDomain creates a new mail.alias.domain model and returns its id.
func (c *Client) CreateMailAliasDomain(mad *MailAliasDomain) (int64, error) {
	ids, err := c.CreateMailAliasDomains([]*MailAliasDomain{mad})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailAliasDomain creates a new mail.alias.domain model and returns its id.
func (c *Client) CreateMailAliasDomains(mads []*MailAliasDomain) ([]int64, error) {
	var vv []interface{}
	for _, v := range mads {
		vv = append(vv, v)
	}
	return c.Create(MailAliasDomainModel, vv, nil)
}

// UpdateMailAliasDomain updates an existing mail.alias.domain record.
func (c *Client) UpdateMailAliasDomain(mad *MailAliasDomain) error {
	return c.UpdateMailAliasDomains([]int64{mad.Id.Get()}, mad)
}

// UpdateMailAliasDomains updates existing mail.alias.domain records.
// All records (represented by ids) will be updated by mad values.
func (c *Client) UpdateMailAliasDomains(ids []int64, mad *MailAliasDomain) error {
	return c.Update(MailAliasDomainModel, ids, mad, nil)
}

// DeleteMailAliasDomain deletes an existing mail.alias.domain record.
func (c *Client) DeleteMailAliasDomain(id int64) error {
	return c.DeleteMailAliasDomains([]int64{id})
}

// DeleteMailAliasDomains deletes existing mail.alias.domain records.
func (c *Client) DeleteMailAliasDomains(ids []int64) error {
	return c.Delete(MailAliasDomainModel, ids)
}

// GetMailAliasDomain gets mail.alias.domain existing record.
func (c *Client) GetMailAliasDomain(id int64) (*MailAliasDomain, error) {
	mads, err := c.GetMailAliasDomains([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mads)[0]), nil
}

// GetMailAliasDomains gets mail.alias.domain existing records.
func (c *Client) GetMailAliasDomains(ids []int64) (*MailAliasDomains, error) {
	mads := &MailAliasDomains{}
	if err := c.Read(MailAliasDomainModel, ids, nil, mads); err != nil {
		return nil, err
	}
	return mads, nil
}

// FindMailAliasDomain finds mail.alias.domain record by querying it with criteria.
func (c *Client) FindMailAliasDomain(criteria *Criteria) (*MailAliasDomain, error) {
	mads := &MailAliasDomains{}
	if err := c.SearchRead(MailAliasDomainModel, criteria, NewOptions().Limit(1), mads); err != nil {
		return nil, err
	}
	return &((*mads)[0]), nil
}

// FindMailAliasDomains finds mail.alias.domain records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAliasDomains(criteria *Criteria, options *Options) (*MailAliasDomains, error) {
	mads := &MailAliasDomains{}
	if err := c.SearchRead(MailAliasDomainModel, criteria, options, mads); err != nil {
		return nil, err
	}
	return mads, nil
}

// FindMailAliasDomainIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAliasDomainIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailAliasDomainModel, criteria, options)
}

// FindMailAliasDomainId finds record id by querying it with criteria.
func (c *Client) FindMailAliasDomainId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailAliasDomainModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
