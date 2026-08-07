package odoo

// AccountMoveSendWizard represents account.move.send.wizard model.
type AccountMoveSendWizard struct {
	Alerts                   interface{} `xmlrpc:"alerts,omitempty"`
	AttachmentsNotSupported  interface{} `xmlrpc:"attachments_not_supported,omitempty"`
	AvailablePdfReportIds    *Relation   `xmlrpc:"available_pdf_report_ids,omitempty"`
	Body                     *String     `xmlrpc:"body,omitempty"`
	BodyHasTemplateValue     *Bool       `xmlrpc:"body_has_template_value,omitempty"`
	CanEditBody              *Bool       `xmlrpc:"can_edit_body,omitempty"`
	CompanyId                *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayAttachmentsWidget *Bool       `xmlrpc:"display_attachments_widget,omitempty"`
	DisplayName              *String     `xmlrpc:"display_name,omitempty"`
	DisplayPdfReportId       *Bool       `xmlrpc:"display_pdf_report_id,omitempty"`
	ExtraEdiCheckboxes       interface{} `xmlrpc:"extra_edi_checkboxes,omitempty"`
	ExtraEdis                interface{} `xmlrpc:"extra_edis,omitempty"`
	Id                       *Int        `xmlrpc:"id,omitempty"`
	InvoiceEdiFormat         *Selection  `xmlrpc:"invoice_edi_format,omitempty"`
	IsMailTemplateEditor     *Bool       `xmlrpc:"is_mail_template_editor,omitempty"`
	Lang                     *String     `xmlrpc:"lang,omitempty"`
	MailAttachmentsWidget    interface{} `xmlrpc:"mail_attachments_widget,omitempty"`
	MailPartnerIds           *Relation   `xmlrpc:"mail_partner_ids,omitempty"`
	Model                    *String     `xmlrpc:"model,omitempty"`
	MoveId                   *Many2One   `xmlrpc:"move_id,omitempty"`
	PdfReportId              *Many2One   `xmlrpc:"pdf_report_id,omitempty"`
	RenderModel              *String     `xmlrpc:"render_model,omitempty"`
	ResIds                   *String     `xmlrpc:"res_ids,omitempty"`
	SendingMethodCheckboxes  interface{} `xmlrpc:"sending_method_checkboxes,omitempty"`
	SendingMethods           interface{} `xmlrpc:"sending_methods,omitempty"`
	Subject                  *String     `xmlrpc:"subject,omitempty"`
	TemplateId               *Many2One   `xmlrpc:"template_id,omitempty"`
	TemplateName             *String     `xmlrpc:"template_name,omitempty"`
	WriteDate                *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountMoveSendWizards represents array of account.move.send.wizard model.
type AccountMoveSendWizards []AccountMoveSendWizard

// AccountMoveSendWizardModel is the odoo model name.
const AccountMoveSendWizardModel = "account.move.send.wizard"

// Many2One convert AccountMoveSendWizard to *Many2One.
func (amsw *AccountMoveSendWizard) Many2One() *Many2One {
	return NewMany2One(amsw.Id.Get(), "")
}

// CreateAccountMoveSendWizard creates a new account.move.send.wizard model and returns its id.
func (c *Client) CreateAccountMoveSendWizard(amsw *AccountMoveSendWizard) (int64, error) {
	ids, err := c.CreateAccountMoveSendWizards([]*AccountMoveSendWizard{amsw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMoveSendWizard creates a new account.move.send.wizard model and returns its id.
func (c *Client) CreateAccountMoveSendWizards(amsws []*AccountMoveSendWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range amsws {
		vv = append(vv, v)
	}
	return c.Create(AccountMoveSendWizardModel, vv, nil)
}

// UpdateAccountMoveSendWizard updates an existing account.move.send.wizard record.
func (c *Client) UpdateAccountMoveSendWizard(amsw *AccountMoveSendWizard) error {
	return c.UpdateAccountMoveSendWizards([]int64{amsw.Id.Get()}, amsw)
}

// UpdateAccountMoveSendWizards updates existing account.move.send.wizard records.
// All records (represented by ids) will be updated by amsw values.
func (c *Client) UpdateAccountMoveSendWizards(ids []int64, amsw *AccountMoveSendWizard) error {
	return c.Update(AccountMoveSendWizardModel, ids, amsw, nil)
}

// DeleteAccountMoveSendWizard deletes an existing account.move.send.wizard record.
func (c *Client) DeleteAccountMoveSendWizard(id int64) error {
	return c.DeleteAccountMoveSendWizards([]int64{id})
}

// DeleteAccountMoveSendWizards deletes existing account.move.send.wizard records.
func (c *Client) DeleteAccountMoveSendWizards(ids []int64) error {
	return c.Delete(AccountMoveSendWizardModel, ids)
}

// GetAccountMoveSendWizard gets account.move.send.wizard existing record.
func (c *Client) GetAccountMoveSendWizard(id int64) (*AccountMoveSendWizard, error) {
	amsws, err := c.GetAccountMoveSendWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amsws)[0]), nil
}

// GetAccountMoveSendWizards gets account.move.send.wizard existing records.
func (c *Client) GetAccountMoveSendWizards(ids []int64) (*AccountMoveSendWizards, error) {
	amsws := &AccountMoveSendWizards{}
	if err := c.Read(AccountMoveSendWizardModel, ids, nil, amsws); err != nil {
		return nil, err
	}
	return amsws, nil
}

// FindAccountMoveSendWizard finds account.move.send.wizard record by querying it with criteria.
func (c *Client) FindAccountMoveSendWizard(criteria *Criteria) (*AccountMoveSendWizard, error) {
	amsws := &AccountMoveSendWizards{}
	if err := c.SearchRead(AccountMoveSendWizardModel, criteria, NewOptions().Limit(1), amsws); err != nil {
		return nil, err
	}
	return &((*amsws)[0]), nil
}

// FindAccountMoveSendWizards finds account.move.send.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSendWizards(criteria *Criteria, options *Options) (*AccountMoveSendWizards, error) {
	amsws := &AccountMoveSendWizards{}
	if err := c.SearchRead(AccountMoveSendWizardModel, criteria, options, amsws); err != nil {
		return nil, err
	}
	return amsws, nil
}

// FindAccountMoveSendWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMoveSendWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountMoveSendWizardModel, criteria, options)
}

// FindAccountMoveSendWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountMoveSendWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMoveSendWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
