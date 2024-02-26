package odoo

// MailShortcode represents mail.shortcode model.
type MailShortcode struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description   *String    `xmlrpc:"description,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	ShortcodeType *Selection `xmlrpc:"shortcode_type,omptempty"`
	Source        *String    `xmlrpc:"source,omptempty"`
	Substitution  *String    `xmlrpc:"substitution,omptempty"`
	UnicodeSource *String    `xmlrpc:"unicode_source,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailShortcodes represents array of mail.shortcode model.
type MailShortcodes []MailShortcode

// MailShortcodeModel is the odoo model name.
const MailShortcodeModel = "mail.shortcode"

// Many2One convert MailShortcode to *Many2One.
func (ms *MailShortcode) Many2One() *Many2One {
	return NewMany2One(ms.Id.Get(), "")
}

// CreateMailShortcode creates a new mail.shortcode model and returns its id.
func (c *Client) CreateMailShortcode(ms *MailShortcode) (int64, error) {
	ids, err := c.CreateMailShortcodes([]*MailShortcode{ms})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailShortcode creates a new mail.shortcode model and returns its id.
func (c *Client) CreateMailShortcodes(mss []*MailShortcode) ([]int64, error) {
	var vv []interface{}
	for _, v := range mss {
		vv = append(vv, v)
	}
	return c.Create(MailShortcodeModel, vv, nil)
}

// UpdateMailShortcode updates an existing mail.shortcode record.
func (c *Client) UpdateMailShortcode(ms *MailShortcode) error {
	return c.UpdateMailShortcodes([]int64{ms.Id.Get()}, ms)
}

// UpdateMailShortcodes updates existing mail.shortcode records.
// All records (represented by ids) will be updated by ms values.
func (c *Client) UpdateMailShortcodes(ids []int64, ms *MailShortcode) error {
	return c.Update(MailShortcodeModel, ids, ms, nil)
}

// DeleteMailShortcode deletes an existing mail.shortcode record.
func (c *Client) DeleteMailShortcode(id int64) error {
	return c.DeleteMailShortcodes([]int64{id})
}

// DeleteMailShortcodes deletes existing mail.shortcode records.
func (c *Client) DeleteMailShortcodes(ids []int64) error {
	return c.Delete(MailShortcodeModel, ids)
}

// GetMailShortcode gets mail.shortcode existing record.
func (c *Client) GetMailShortcode(id int64) (*MailShortcode, error) {
	mss, err := c.GetMailShortcodes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mss)[0]), nil
}

// GetMailShortcodes gets mail.shortcode existing records.
func (c *Client) GetMailShortcodes(ids []int64) (*MailShortcodes, error) {
	mss := &MailShortcodes{}
	if err := c.Read(MailShortcodeModel, ids, nil, mss); err != nil {
		return nil, err
	}
	return mss, nil
}

// FindMailShortcode finds mail.shortcode record by querying it with criteria.
func (c *Client) FindMailShortcode(criteria *Criteria) (*MailShortcode, error) {
	mss := &MailShortcodes{}
	if err := c.SearchRead(MailShortcodeModel, criteria, NewOptions().Limit(1), mss); err != nil {
		return nil, err
	}
	return &((*mss)[0]), nil
}

// FindMailShortcodes finds mail.shortcode records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailShortcodes(criteria *Criteria, options *Options) (*MailShortcodes, error) {
	mss := &MailShortcodes{}
	if err := c.SearchRead(MailShortcodeModel, criteria, options, mss); err != nil {
		return nil, err
	}
	return mss, nil
}

// FindMailShortcodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailShortcodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailShortcodeModel, criteria, options)
}

// FindMailShortcodeId finds record id by querying it with criteria.
func (c *Client) FindMailShortcodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailShortcodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
