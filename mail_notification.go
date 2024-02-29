package odoo

// MailNotification represents mail.notification model.
type MailNotification struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	EmailStatus   *Selection `xmlrpc:"email_status,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	IsEmail       *Bool      `xmlrpc:"is_email,omitempty"`
	IsRead        *Bool      `xmlrpc:"is_read,omitempty"`
	MailMessageId *Many2One  `xmlrpc:"mail_message_id,omitempty"`
	ResPartnerId  *Many2One  `xmlrpc:"res_partner_id,omitempty"`
}

// MailNotifications represents array of mail.notification model.
type MailNotifications []MailNotification

// MailNotificationModel is the odoo model name.
const MailNotificationModel = "mail.notification"

// Many2One convert MailNotification to *Many2One.
func (mn *MailNotification) Many2One() *Many2One {
	return NewMany2One(mn.Id.Get(), "")
}

// CreateMailNotification creates a new mail.notification model and returns its id.
func (c *Client) CreateMailNotification(mn *MailNotification) (int64, error) {
	ids, err := c.CreateMailNotifications([]*MailNotification{mn})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailNotifications creates a new mail.notification model and returns its id.
func (c *Client) CreateMailNotifications(mns []*MailNotification) ([]int64, error) {
	var vv []interface{}
	for _, v := range mns {
		vv = append(vv, v)
	}
	return c.Create(MailNotificationModel, vv, nil)
}

// UpdateMailNotification updates an existing mail.notification record.
func (c *Client) UpdateMailNotification(mn *MailNotification) error {
	return c.UpdateMailNotifications([]int64{mn.Id.Get()}, mn)
}

// UpdateMailNotifications updates existing mail.notification records.
// All records (represented by ids) will be updated by mn values.
func (c *Client) UpdateMailNotifications(ids []int64, mn *MailNotification) error {
	return c.Update(MailNotificationModel, ids, mn, nil)
}

// DeleteMailNotification deletes an existing mail.notification record.
func (c *Client) DeleteMailNotification(id int64) error {
	return c.DeleteMailNotifications([]int64{id})
}

// DeleteMailNotifications deletes existing mail.notification records.
func (c *Client) DeleteMailNotifications(ids []int64) error {
	return c.Delete(MailNotificationModel, ids)
}

// GetMailNotification gets mail.notification existing record.
func (c *Client) GetMailNotification(id int64) (*MailNotification, error) {
	mns, err := c.GetMailNotifications([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mns)[0]), nil
}

// GetMailNotifications gets mail.notification existing records.
func (c *Client) GetMailNotifications(ids []int64) (*MailNotifications, error) {
	mns := &MailNotifications{}
	if err := c.Read(MailNotificationModel, ids, nil, mns); err != nil {
		return nil, err
	}
	return mns, nil
}

// FindMailNotification finds mail.notification record by querying it with criteria.
func (c *Client) FindMailNotification(criteria *Criteria) (*MailNotification, error) {
	mns := &MailNotifications{}
	if err := c.SearchRead(MailNotificationModel, criteria, NewOptions().Limit(1), mns); err != nil {
		return nil, err
	}
	return &((*mns)[0]), nil
}

// FindMailNotifications finds mail.notification records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailNotifications(criteria *Criteria, options *Options) (*MailNotifications, error) {
	mns := &MailNotifications{}
	if err := c.SearchRead(MailNotificationModel, criteria, options, mns); err != nil {
		return nil, err
	}
	return mns, nil
}

// FindMailNotificationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailNotificationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailNotificationModel, criteria, options)
}

// FindMailNotificationId finds record id by querying it with criteria.
func (c *Client) FindMailNotificationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailNotificationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
