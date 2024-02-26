package odoo

// MailMessageSubtype represents mail.message.subtype model.
type MailMessageSubtype struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	Default       *Bool     `xmlrpc:"default,omptempty"`
	Description   *String   `xmlrpc:"description,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Hidden        *Bool     `xmlrpc:"hidden,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	Internal      *Bool     `xmlrpc:"internal,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	ParentId      *Many2One `xmlrpc:"parent_id,omptempty"`
	RelationField *String   `xmlrpc:"relation_field,omptempty"`
	ResModel      *String   `xmlrpc:"res_model,omptempty"`
	Sequence      *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailMessageSubtypes represents array of mail.message.subtype model.
type MailMessageSubtypes []MailMessageSubtype

// MailMessageSubtypeModel is the odoo model name.
const MailMessageSubtypeModel = "mail.message.subtype"

// Many2One convert MailMessageSubtype to *Many2One.
func (mms *MailMessageSubtype) Many2One() *Many2One {
	return NewMany2One(mms.Id.Get(), "")
}

// CreateMailMessageSubtype creates a new mail.message.subtype model and returns its id.
func (c *Client) CreateMailMessageSubtype(mms *MailMessageSubtype) (int64, error) {
	ids, err := c.CreateMailMessageSubtypes([]*MailMessageSubtype{mms})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMessageSubtype creates a new mail.message.subtype model and returns its id.
func (c *Client) CreateMailMessageSubtypes(mmss []*MailMessageSubtype) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmss {
		vv = append(vv, v)
	}
	return c.Create(MailMessageSubtypeModel, vv, nil)
}

// UpdateMailMessageSubtype updates an existing mail.message.subtype record.
func (c *Client) UpdateMailMessageSubtype(mms *MailMessageSubtype) error {
	return c.UpdateMailMessageSubtypes([]int64{mms.Id.Get()}, mms)
}

// UpdateMailMessageSubtypes updates existing mail.message.subtype records.
// All records (represented by ids) will be updated by mms values.
func (c *Client) UpdateMailMessageSubtypes(ids []int64, mms *MailMessageSubtype) error {
	return c.Update(MailMessageSubtypeModel, ids, mms, nil)
}

// DeleteMailMessageSubtype deletes an existing mail.message.subtype record.
func (c *Client) DeleteMailMessageSubtype(id int64) error {
	return c.DeleteMailMessageSubtypes([]int64{id})
}

// DeleteMailMessageSubtypes deletes existing mail.message.subtype records.
func (c *Client) DeleteMailMessageSubtypes(ids []int64) error {
	return c.Delete(MailMessageSubtypeModel, ids)
}

// GetMailMessageSubtype gets mail.message.subtype existing record.
func (c *Client) GetMailMessageSubtype(id int64) (*MailMessageSubtype, error) {
	mmss, err := c.GetMailMessageSubtypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmss)[0]), nil
}

// GetMailMessageSubtypes gets mail.message.subtype existing records.
func (c *Client) GetMailMessageSubtypes(ids []int64) (*MailMessageSubtypes, error) {
	mmss := &MailMessageSubtypes{}
	if err := c.Read(MailMessageSubtypeModel, ids, nil, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMessageSubtype finds mail.message.subtype record by querying it with criteria.
func (c *Client) FindMailMessageSubtype(criteria *Criteria) (*MailMessageSubtype, error) {
	mmss := &MailMessageSubtypes{}
	if err := c.SearchRead(MailMessageSubtypeModel, criteria, NewOptions().Limit(1), mmss); err != nil {
		return nil, err
	}
	return &((*mmss)[0]), nil
}

// FindMailMessageSubtypes finds mail.message.subtype records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageSubtypes(criteria *Criteria, options *Options) (*MailMessageSubtypes, error) {
	mmss := &MailMessageSubtypes{}
	if err := c.SearchRead(MailMessageSubtypeModel, criteria, options, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMessageSubtypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageSubtypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMessageSubtypeModel, criteria, options)
}

// FindMailMessageSubtypeId finds record id by querying it with criteria.
func (c *Client) FindMailMessageSubtypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageSubtypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
