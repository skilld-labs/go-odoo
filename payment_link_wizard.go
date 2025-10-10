package odoo

// PaymentLinkWizard represents payment.link.wizard model.
type PaymentLinkWizard struct {
	Amount                  *Float      `xmlrpc:"amount,omitempty"`
	AmountMax               *Float      `xmlrpc:"amount_max,omitempty"`
	AmountPaid              *Float      `xmlrpc:"amount_paid,omitempty"`
	CompanyId               *Many2One   `xmlrpc:"company_id,omitempty"`
	ConfirmationMessage     *String     `xmlrpc:"confirmation_message,omitempty"`
	CreateDate              *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId              *Many2One   `xmlrpc:"currency_id,omitempty"`
	DiscountDate            *Time       `xmlrpc:"discount_date,omitempty"`
	DisplayName             *String     `xmlrpc:"display_name,omitempty"`
	DisplayOpenInstallments *Bool       `xmlrpc:"display_open_installments,omitempty"`
	EpdInfo                 *String     `xmlrpc:"epd_info,omitempty"`
	HasEligibleEpd          *Bool       `xmlrpc:"has_eligible_epd,omitempty"`
	Id                      *Int        `xmlrpc:"id,omitempty"`
	InvoiceAmountDue        *Float      `xmlrpc:"invoice_amount_due,omitempty"`
	Link                    *String     `xmlrpc:"link,omitempty"`
	OpenInstallments        interface{} `xmlrpc:"open_installments,omitempty"`
	OpenInstallmentsPreview *String     `xmlrpc:"open_installments_preview,omitempty"`
	PartnerEmail            *String     `xmlrpc:"partner_email,omitempty"`
	PartnerId               *Many2One   `xmlrpc:"partner_id,omitempty"`
	ResId                   *Int        `xmlrpc:"res_id,omitempty"`
	ResModel                *String     `xmlrpc:"res_model,omitempty"`
	WarningMessage          *String     `xmlrpc:"warning_message,omitempty"`
	WriteDate               *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// PaymentLinkWizards represents array of payment.link.wizard model.
type PaymentLinkWizards []PaymentLinkWizard

// PaymentLinkWizardModel is the odoo model name.
const PaymentLinkWizardModel = "payment.link.wizard"

// Many2One convert PaymentLinkWizard to *Many2One.
func (plw *PaymentLinkWizard) Many2One() *Many2One {
	return NewMany2One(plw.Id.Get(), "")
}

// CreatePaymentLinkWizard creates a new payment.link.wizard model and returns its id.
func (c *Client) CreatePaymentLinkWizard(plw *PaymentLinkWizard) (int64, error) {
	ids, err := c.CreatePaymentLinkWizards([]*PaymentLinkWizard{plw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePaymentLinkWizard creates a new payment.link.wizard model and returns its id.
func (c *Client) CreatePaymentLinkWizards(plws []*PaymentLinkWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range plws {
		vv = append(vv, v)
	}
	return c.Create(PaymentLinkWizardModel, vv, nil)
}

// UpdatePaymentLinkWizard updates an existing payment.link.wizard record.
func (c *Client) UpdatePaymentLinkWizard(plw *PaymentLinkWizard) error {
	return c.UpdatePaymentLinkWizards([]int64{plw.Id.Get()}, plw)
}

// UpdatePaymentLinkWizards updates existing payment.link.wizard records.
// All records (represented by ids) will be updated by plw values.
func (c *Client) UpdatePaymentLinkWizards(ids []int64, plw *PaymentLinkWizard) error {
	return c.Update(PaymentLinkWizardModel, ids, plw, nil)
}

// DeletePaymentLinkWizard deletes an existing payment.link.wizard record.
func (c *Client) DeletePaymentLinkWizard(id int64) error {
	return c.DeletePaymentLinkWizards([]int64{id})
}

// DeletePaymentLinkWizards deletes existing payment.link.wizard records.
func (c *Client) DeletePaymentLinkWizards(ids []int64) error {
	return c.Delete(PaymentLinkWizardModel, ids)
}

// GetPaymentLinkWizard gets payment.link.wizard existing record.
func (c *Client) GetPaymentLinkWizard(id int64) (*PaymentLinkWizard, error) {
	plws, err := c.GetPaymentLinkWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*plws)[0]), nil
}

// GetPaymentLinkWizards gets payment.link.wizard existing records.
func (c *Client) GetPaymentLinkWizards(ids []int64) (*PaymentLinkWizards, error) {
	plws := &PaymentLinkWizards{}
	if err := c.Read(PaymentLinkWizardModel, ids, nil, plws); err != nil {
		return nil, err
	}
	return plws, nil
}

// FindPaymentLinkWizard finds payment.link.wizard record by querying it with criteria.
func (c *Client) FindPaymentLinkWizard(criteria *Criteria) (*PaymentLinkWizard, error) {
	plws := &PaymentLinkWizards{}
	if err := c.SearchRead(PaymentLinkWizardModel, criteria, NewOptions().Limit(1), plws); err != nil {
		return nil, err
	}
	return &((*plws)[0]), nil
}

// FindPaymentLinkWizards finds payment.link.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentLinkWizards(criteria *Criteria, options *Options) (*PaymentLinkWizards, error) {
	plws := &PaymentLinkWizards{}
	if err := c.SearchRead(PaymentLinkWizardModel, criteria, options, plws); err != nil {
		return nil, err
	}
	return plws, nil
}

// FindPaymentLinkWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentLinkWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PaymentLinkWizardModel, criteria, options)
}

// FindPaymentLinkWizardId finds record id by querying it with criteria.
func (c *Client) FindPaymentLinkWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentLinkWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
