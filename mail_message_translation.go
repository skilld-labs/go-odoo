package odoo

// MailMessageTranslation represents mail.message.translation model.
type MailMessageTranslation struct {
	Body        *String   `xmlrpc:"body,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	MessageId   *Many2One `xmlrpc:"message_id,omitempty"`
	SourceLang  *String   `xmlrpc:"source_lang,omitempty"`
	TargetLang  *String   `xmlrpc:"target_lang,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailMessageTranslations represents array of mail.message.translation model.
type MailMessageTranslations []MailMessageTranslation

// MailMessageTranslationModel is the odoo model name.
const MailMessageTranslationModel = "mail.message.translation"

// Many2One convert MailMessageTranslation to *Many2One.
func (mmt *MailMessageTranslation) Many2One() *Many2One {
	return NewMany2One(mmt.Id.Get(), "")
}

// CreateMailMessageTranslation creates a new mail.message.translation model and returns its id.
func (c *Client) CreateMailMessageTranslation(mmt *MailMessageTranslation) (int64, error) {
	ids, err := c.CreateMailMessageTranslations([]*MailMessageTranslation{mmt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMessageTranslation creates a new mail.message.translation model and returns its id.
func (c *Client) CreateMailMessageTranslations(mmts []*MailMessageTranslation) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmts {
		vv = append(vv, v)
	}
	return c.Create(MailMessageTranslationModel, vv, nil)
}

// UpdateMailMessageTranslation updates an existing mail.message.translation record.
func (c *Client) UpdateMailMessageTranslation(mmt *MailMessageTranslation) error {
	return c.UpdateMailMessageTranslations([]int64{mmt.Id.Get()}, mmt)
}

// UpdateMailMessageTranslations updates existing mail.message.translation records.
// All records (represented by ids) will be updated by mmt values.
func (c *Client) UpdateMailMessageTranslations(ids []int64, mmt *MailMessageTranslation) error {
	return c.Update(MailMessageTranslationModel, ids, mmt, nil)
}

// DeleteMailMessageTranslation deletes an existing mail.message.translation record.
func (c *Client) DeleteMailMessageTranslation(id int64) error {
	return c.DeleteMailMessageTranslations([]int64{id})
}

// DeleteMailMessageTranslations deletes existing mail.message.translation records.
func (c *Client) DeleteMailMessageTranslations(ids []int64) error {
	return c.Delete(MailMessageTranslationModel, ids)
}

// GetMailMessageTranslation gets mail.message.translation existing record.
func (c *Client) GetMailMessageTranslation(id int64) (*MailMessageTranslation, error) {
	mmts, err := c.GetMailMessageTranslations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmts)[0]), nil
}

// GetMailMessageTranslations gets mail.message.translation existing records.
func (c *Client) GetMailMessageTranslations(ids []int64) (*MailMessageTranslations, error) {
	mmts := &MailMessageTranslations{}
	if err := c.Read(MailMessageTranslationModel, ids, nil, mmts); err != nil {
		return nil, err
	}
	return mmts, nil
}

// FindMailMessageTranslation finds mail.message.translation record by querying it with criteria.
func (c *Client) FindMailMessageTranslation(criteria *Criteria) (*MailMessageTranslation, error) {
	mmts := &MailMessageTranslations{}
	if err := c.SearchRead(MailMessageTranslationModel, criteria, NewOptions().Limit(1), mmts); err != nil {
		return nil, err
	}
	return &((*mmts)[0]), nil
}

// FindMailMessageTranslations finds mail.message.translation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageTranslations(criteria *Criteria, options *Options) (*MailMessageTranslations, error) {
	mmts := &MailMessageTranslations{}
	if err := c.SearchRead(MailMessageTranslationModel, criteria, options, mmts); err != nil {
		return nil, err
	}
	return mmts, nil
}

// FindMailMessageTranslationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageTranslationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMessageTranslationModel, criteria, options)
}

// FindMailMessageTranslationId finds record id by querying it with criteria.
func (c *Client) FindMailMessageTranslationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageTranslationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
