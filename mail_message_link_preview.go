package odoo

// MailMessageLinkPreview represents mail.message.link.preview model.
type MailMessageLinkPreview struct {
	AuthorId      *Many2One `xmlrpc:"author_id,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	IsHidden      *Bool     `xmlrpc:"is_hidden,omitempty"`
	LinkPreviewId *Many2One `xmlrpc:"link_preview_id,omitempty"`
	MessageId     *Many2One `xmlrpc:"message_id,omitempty"`
	Sequence      *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailMessageLinkPreviews represents array of mail.message.link.preview model.
type MailMessageLinkPreviews []MailMessageLinkPreview

// MailMessageLinkPreviewModel is the odoo model name.
const MailMessageLinkPreviewModel = "mail.message.link.preview"

// Many2One convert MailMessageLinkPreview to *Many2One.
func (mmlp *MailMessageLinkPreview) Many2One() *Many2One {
	return NewMany2One(mmlp.Id.Get(), "")
}

// CreateMailMessageLinkPreview creates a new mail.message.link.preview model and returns its id.
func (c *Client) CreateMailMessageLinkPreview(mmlp *MailMessageLinkPreview) (int64, error) {
	ids, err := c.CreateMailMessageLinkPreviews([]*MailMessageLinkPreview{mmlp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMessageLinkPreview creates a new mail.message.link.preview model and returns its id.
func (c *Client) CreateMailMessageLinkPreviews(mmlps []*MailMessageLinkPreview) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmlps {
		vv = append(vv, v)
	}
	return c.Create(MailMessageLinkPreviewModel, vv, nil)
}

// UpdateMailMessageLinkPreview updates an existing mail.message.link.preview record.
func (c *Client) UpdateMailMessageLinkPreview(mmlp *MailMessageLinkPreview) error {
	return c.UpdateMailMessageLinkPreviews([]int64{mmlp.Id.Get()}, mmlp)
}

// UpdateMailMessageLinkPreviews updates existing mail.message.link.preview records.
// All records (represented by ids) will be updated by mmlp values.
func (c *Client) UpdateMailMessageLinkPreviews(ids []int64, mmlp *MailMessageLinkPreview) error {
	return c.Update(MailMessageLinkPreviewModel, ids, mmlp, nil)
}

// DeleteMailMessageLinkPreview deletes an existing mail.message.link.preview record.
func (c *Client) DeleteMailMessageLinkPreview(id int64) error {
	return c.DeleteMailMessageLinkPreviews([]int64{id})
}

// DeleteMailMessageLinkPreviews deletes existing mail.message.link.preview records.
func (c *Client) DeleteMailMessageLinkPreviews(ids []int64) error {
	return c.Delete(MailMessageLinkPreviewModel, ids)
}

// GetMailMessageLinkPreview gets mail.message.link.preview existing record.
func (c *Client) GetMailMessageLinkPreview(id int64) (*MailMessageLinkPreview, error) {
	mmlps, err := c.GetMailMessageLinkPreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmlps)[0]), nil
}

// GetMailMessageLinkPreviews gets mail.message.link.preview existing records.
func (c *Client) GetMailMessageLinkPreviews(ids []int64) (*MailMessageLinkPreviews, error) {
	mmlps := &MailMessageLinkPreviews{}
	if err := c.Read(MailMessageLinkPreviewModel, ids, nil, mmlps); err != nil {
		return nil, err
	}
	return mmlps, nil
}

// FindMailMessageLinkPreview finds mail.message.link.preview record by querying it with criteria.
func (c *Client) FindMailMessageLinkPreview(criteria *Criteria) (*MailMessageLinkPreview, error) {
	mmlps := &MailMessageLinkPreviews{}
	if err := c.SearchRead(MailMessageLinkPreviewModel, criteria, NewOptions().Limit(1), mmlps); err != nil {
		return nil, err
	}
	return &((*mmlps)[0]), nil
}

// FindMailMessageLinkPreviews finds mail.message.link.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageLinkPreviews(criteria *Criteria, options *Options) (*MailMessageLinkPreviews, error) {
	mmlps := &MailMessageLinkPreviews{}
	if err := c.SearchRead(MailMessageLinkPreviewModel, criteria, options, mmlps); err != nil {
		return nil, err
	}
	return mmlps, nil
}

// FindMailMessageLinkPreviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageLinkPreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMessageLinkPreviewModel, criteria, options)
}

// FindMailMessageLinkPreviewId finds record id by querying it with criteria.
func (c *Client) FindMailMessageLinkPreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageLinkPreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
