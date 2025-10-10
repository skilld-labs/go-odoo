package odoo

// MailResendPartner represents mail.resend.partner model.
type MailResendPartner struct {
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Email           *String   `xmlrpc:"email,omitempty"`
	FailureReason   *String   `xmlrpc:"failure_reason,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Message         *String   `xmlrpc:"message,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	NotificationId  *Many2One `xmlrpc:"notification_id,omitempty"`
	PartnerId       *Many2One `xmlrpc:"partner_id,omitempty"`
	PartnerReadonly *Bool     `xmlrpc:"partner_readonly,omitempty"`
	Resend          *Bool     `xmlrpc:"resend,omitempty"`
	ResendWizardId  *Many2One `xmlrpc:"resend_wizard_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailResendPartners represents array of mail.resend.partner model.
type MailResendPartners []MailResendPartner

// MailResendPartnerModel is the odoo model name.
const MailResendPartnerModel = "mail.resend.partner"

// Many2One convert MailResendPartner to *Many2One.
func (mrp *MailResendPartner) Many2One() *Many2One {
	return NewMany2One(mrp.Id.Get(), "")
}

// CreateMailResendPartner creates a new mail.resend.partner model and returns its id.
func (c *Client) CreateMailResendPartner(mrp *MailResendPartner) (int64, error) {
	ids, err := c.CreateMailResendPartners([]*MailResendPartner{mrp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailResendPartner creates a new mail.resend.partner model and returns its id.
func (c *Client) CreateMailResendPartners(mrps []*MailResendPartner) ([]int64, error) {
	var vv []interface{}
	for _, v := range mrps {
		vv = append(vv, v)
	}
	return c.Create(MailResendPartnerModel, vv, nil)
}

// UpdateMailResendPartner updates an existing mail.resend.partner record.
func (c *Client) UpdateMailResendPartner(mrp *MailResendPartner) error {
	return c.UpdateMailResendPartners([]int64{mrp.Id.Get()}, mrp)
}

// UpdateMailResendPartners updates existing mail.resend.partner records.
// All records (represented by ids) will be updated by mrp values.
func (c *Client) UpdateMailResendPartners(ids []int64, mrp *MailResendPartner) error {
	return c.Update(MailResendPartnerModel, ids, mrp, nil)
}

// DeleteMailResendPartner deletes an existing mail.resend.partner record.
func (c *Client) DeleteMailResendPartner(id int64) error {
	return c.DeleteMailResendPartners([]int64{id})
}

// DeleteMailResendPartners deletes existing mail.resend.partner records.
func (c *Client) DeleteMailResendPartners(ids []int64) error {
	return c.Delete(MailResendPartnerModel, ids)
}

// GetMailResendPartner gets mail.resend.partner existing record.
func (c *Client) GetMailResendPartner(id int64) (*MailResendPartner, error) {
	mrps, err := c.GetMailResendPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mrps)[0]), nil
}

// GetMailResendPartners gets mail.resend.partner existing records.
func (c *Client) GetMailResendPartners(ids []int64) (*MailResendPartners, error) {
	mrps := &MailResendPartners{}
	if err := c.Read(MailResendPartnerModel, ids, nil, mrps); err != nil {
		return nil, err
	}
	return mrps, nil
}

// FindMailResendPartner finds mail.resend.partner record by querying it with criteria.
func (c *Client) FindMailResendPartner(criteria *Criteria) (*MailResendPartner, error) {
	mrps := &MailResendPartners{}
	if err := c.SearchRead(MailResendPartnerModel, criteria, NewOptions().Limit(1), mrps); err != nil {
		return nil, err
	}
	return &((*mrps)[0]), nil
}

// FindMailResendPartners finds mail.resend.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailResendPartners(criteria *Criteria, options *Options) (*MailResendPartners, error) {
	mrps := &MailResendPartners{}
	if err := c.SearchRead(MailResendPartnerModel, criteria, options, mrps); err != nil {
		return nil, err
	}
	return mrps, nil
}

// FindMailResendPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailResendPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailResendPartnerModel, criteria, options)
}

// FindMailResendPartnerId finds record id by querying it with criteria.
func (c *Client) FindMailResendPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailResendPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
