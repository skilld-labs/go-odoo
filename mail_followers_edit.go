package odoo

// MailFollowersEdit represents mail.followers.edit model.
type MailFollowersEdit struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Message     *String    `xmlrpc:"message,omitempty"`
	Notify      *Bool      `xmlrpc:"notify,omitempty"`
	Operation   *Selection `xmlrpc:"operation,omitempty"`
	PartnerIds  *Relation  `xmlrpc:"partner_ids,omitempty"`
	ResIds      *String    `xmlrpc:"res_ids,omitempty"`
	ResModel    *String    `xmlrpc:"res_model,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailFollowersEdits represents array of mail.followers.edit model.
type MailFollowersEdits []MailFollowersEdit

// MailFollowersEditModel is the odoo model name.
const MailFollowersEditModel = "mail.followers.edit"

// Many2One convert MailFollowersEdit to *Many2One.
func (mfe *MailFollowersEdit) Many2One() *Many2One {
	return NewMany2One(mfe.Id.Get(), "")
}

// CreateMailFollowersEdit creates a new mail.followers.edit model and returns its id.
func (c *Client) CreateMailFollowersEdit(mfe *MailFollowersEdit) (int64, error) {
	ids, err := c.CreateMailFollowersEdits([]*MailFollowersEdit{mfe})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailFollowersEdit creates a new mail.followers.edit model and returns its id.
func (c *Client) CreateMailFollowersEdits(mfes []*MailFollowersEdit) ([]int64, error) {
	var vv []interface{}
	for _, v := range mfes {
		vv = append(vv, v)
	}
	return c.Create(MailFollowersEditModel, vv, nil)
}

// UpdateMailFollowersEdit updates an existing mail.followers.edit record.
func (c *Client) UpdateMailFollowersEdit(mfe *MailFollowersEdit) error {
	return c.UpdateMailFollowersEdits([]int64{mfe.Id.Get()}, mfe)
}

// UpdateMailFollowersEdits updates existing mail.followers.edit records.
// All records (represented by ids) will be updated by mfe values.
func (c *Client) UpdateMailFollowersEdits(ids []int64, mfe *MailFollowersEdit) error {
	return c.Update(MailFollowersEditModel, ids, mfe, nil)
}

// DeleteMailFollowersEdit deletes an existing mail.followers.edit record.
func (c *Client) DeleteMailFollowersEdit(id int64) error {
	return c.DeleteMailFollowersEdits([]int64{id})
}

// DeleteMailFollowersEdits deletes existing mail.followers.edit records.
func (c *Client) DeleteMailFollowersEdits(ids []int64) error {
	return c.Delete(MailFollowersEditModel, ids)
}

// GetMailFollowersEdit gets mail.followers.edit existing record.
func (c *Client) GetMailFollowersEdit(id int64) (*MailFollowersEdit, error) {
	mfes, err := c.GetMailFollowersEdits([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mfes)[0]), nil
}

// GetMailFollowersEdits gets mail.followers.edit existing records.
func (c *Client) GetMailFollowersEdits(ids []int64) (*MailFollowersEdits, error) {
	mfes := &MailFollowersEdits{}
	if err := c.Read(MailFollowersEditModel, ids, nil, mfes); err != nil {
		return nil, err
	}
	return mfes, nil
}

// FindMailFollowersEdit finds mail.followers.edit record by querying it with criteria.
func (c *Client) FindMailFollowersEdit(criteria *Criteria) (*MailFollowersEdit, error) {
	mfes := &MailFollowersEdits{}
	if err := c.SearchRead(MailFollowersEditModel, criteria, NewOptions().Limit(1), mfes); err != nil {
		return nil, err
	}
	return &((*mfes)[0]), nil
}

// FindMailFollowersEdits finds mail.followers.edit records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailFollowersEdits(criteria *Criteria, options *Options) (*MailFollowersEdits, error) {
	mfes := &MailFollowersEdits{}
	if err := c.SearchRead(MailFollowersEditModel, criteria, options, mfes); err != nil {
		return nil, err
	}
	return mfes, nil
}

// FindMailFollowersEditIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailFollowersEditIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailFollowersEditModel, criteria, options)
}

// FindMailFollowersEditId finds record id by querying it with criteria.
func (c *Client) FindMailFollowersEditId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailFollowersEditModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
