package odoo

// MailWizardInvite represents mail.wizard.invite model.
type MailWizardInvite struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	ChannelIds  *Relation `xmlrpc:"channel_ids,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Message     *String   `xmlrpc:"message,omitempty"`
	PartnerIds  *Relation `xmlrpc:"partner_ids,omitempty"`
	ResId       *Int      `xmlrpc:"res_id,omitempty"`
	ResModel    *String   `xmlrpc:"res_model,omitempty"`
	SendMail    *Bool     `xmlrpc:"send_mail,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailWizardInvites represents array of mail.wizard.invite model.
type MailWizardInvites []MailWizardInvite

// MailWizardInviteModel is the odoo model name.
const MailWizardInviteModel = "mail.wizard.invite"

// Many2One convert MailWizardInvite to *Many2One.
func (mwi *MailWizardInvite) Many2One() *Many2One {
	return NewMany2One(mwi.Id.Get(), "")
}

// CreateMailWizardInvite creates a new mail.wizard.invite model and returns its id.
func (c *Client) CreateMailWizardInvite(mwi *MailWizardInvite) (int64, error) {
	ids, err := c.CreateMailWizardInvites([]*MailWizardInvite{mwi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailWizardInvite creates a new mail.wizard.invite model and returns its id.
func (c *Client) CreateMailWizardInvites(mwis []*MailWizardInvite) ([]int64, error) {
	var vv []interface{}
	for _, v := range mwis {
		vv = append(vv, v)
	}
	return c.Create(MailWizardInviteModel, vv, nil)
}

// UpdateMailWizardInvite updates an existing mail.wizard.invite record.
func (c *Client) UpdateMailWizardInvite(mwi *MailWizardInvite) error {
	return c.UpdateMailWizardInvites([]int64{mwi.Id.Get()}, mwi)
}

// UpdateMailWizardInvites updates existing mail.wizard.invite records.
// All records (represented by ids) will be updated by mwi values.
func (c *Client) UpdateMailWizardInvites(ids []int64, mwi *MailWizardInvite) error {
	return c.Update(MailWizardInviteModel, ids, mwi, nil)
}

// DeleteMailWizardInvite deletes an existing mail.wizard.invite record.
func (c *Client) DeleteMailWizardInvite(id int64) error {
	return c.DeleteMailWizardInvites([]int64{id})
}

// DeleteMailWizardInvites deletes existing mail.wizard.invite records.
func (c *Client) DeleteMailWizardInvites(ids []int64) error {
	return c.Delete(MailWizardInviteModel, ids)
}

// GetMailWizardInvite gets mail.wizard.invite existing record.
func (c *Client) GetMailWizardInvite(id int64) (*MailWizardInvite, error) {
	mwis, err := c.GetMailWizardInvites([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mwis)[0]), nil
}

// GetMailWizardInvites gets mail.wizard.invite existing records.
func (c *Client) GetMailWizardInvites(ids []int64) (*MailWizardInvites, error) {
	mwis := &MailWizardInvites{}
	if err := c.Read(MailWizardInviteModel, ids, nil, mwis); err != nil {
		return nil, err
	}
	return mwis, nil
}

// FindMailWizardInvite finds mail.wizard.invite record by querying it with criteria.
func (c *Client) FindMailWizardInvite(criteria *Criteria) (*MailWizardInvite, error) {
	mwis := &MailWizardInvites{}
	if err := c.SearchRead(MailWizardInviteModel, criteria, NewOptions().Limit(1), mwis); err != nil {
		return nil, err
	}
	return &((*mwis)[0]), nil
}

// FindMailWizardInvites finds mail.wizard.invite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailWizardInvites(criteria *Criteria, options *Options) (*MailWizardInvites, error) {
	mwis := &MailWizardInvites{}
	if err := c.SearchRead(MailWizardInviteModel, criteria, options, mwis); err != nil {
		return nil, err
	}
	return mwis, nil
}

// FindMailWizardInviteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailWizardInviteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailWizardInviteModel, criteria, options)
}

// FindMailWizardInviteId finds record id by querying it with criteria.
func (c *Client) FindMailWizardInviteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailWizardInviteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
