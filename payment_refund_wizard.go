package odoo

// PaymentRefundWizard represents payment.refund.wizard model.
type PaymentRefundWizard struct {
	AmountAvailableForRefund *Float     `xmlrpc:"amount_available_for_refund,omitempty"`
	AmountToRefund           *Float     `xmlrpc:"amount_to_refund,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	HasPendingRefund         *Bool      `xmlrpc:"has_pending_refund,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	PaymentAmount            *Float     `xmlrpc:"payment_amount,omitempty"`
	PaymentId                *Many2One  `xmlrpc:"payment_id,omitempty"`
	RefundedAmount           *Float     `xmlrpc:"refunded_amount,omitempty"`
	SupportRefund            *Selection `xmlrpc:"support_refund,omitempty"`
	TransactionId            *Many2One  `xmlrpc:"transaction_id,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PaymentRefundWizards represents array of payment.refund.wizard model.
type PaymentRefundWizards []PaymentRefundWizard

// PaymentRefundWizardModel is the odoo model name.
const PaymentRefundWizardModel = "payment.refund.wizard"

// Many2One convert PaymentRefundWizard to *Many2One.
func (prw *PaymentRefundWizard) Many2One() *Many2One {
	return NewMany2One(prw.Id.Get(), "")
}

// CreatePaymentRefundWizard creates a new payment.refund.wizard model and returns its id.
func (c *Client) CreatePaymentRefundWizard(prw *PaymentRefundWizard) (int64, error) {
	ids, err := c.CreatePaymentRefundWizards([]*PaymentRefundWizard{prw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePaymentRefundWizard creates a new payment.refund.wizard model and returns its id.
func (c *Client) CreatePaymentRefundWizards(prws []*PaymentRefundWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range prws {
		vv = append(vv, v)
	}
	return c.Create(PaymentRefundWizardModel, vv, nil)
}

// UpdatePaymentRefundWizard updates an existing payment.refund.wizard record.
func (c *Client) UpdatePaymentRefundWizard(prw *PaymentRefundWizard) error {
	return c.UpdatePaymentRefundWizards([]int64{prw.Id.Get()}, prw)
}

// UpdatePaymentRefundWizards updates existing payment.refund.wizard records.
// All records (represented by ids) will be updated by prw values.
func (c *Client) UpdatePaymentRefundWizards(ids []int64, prw *PaymentRefundWizard) error {
	return c.Update(PaymentRefundWizardModel, ids, prw, nil)
}

// DeletePaymentRefundWizard deletes an existing payment.refund.wizard record.
func (c *Client) DeletePaymentRefundWizard(id int64) error {
	return c.DeletePaymentRefundWizards([]int64{id})
}

// DeletePaymentRefundWizards deletes existing payment.refund.wizard records.
func (c *Client) DeletePaymentRefundWizards(ids []int64) error {
	return c.Delete(PaymentRefundWizardModel, ids)
}

// GetPaymentRefundWizard gets payment.refund.wizard existing record.
func (c *Client) GetPaymentRefundWizard(id int64) (*PaymentRefundWizard, error) {
	prws, err := c.GetPaymentRefundWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prws)[0]), nil
}

// GetPaymentRefundWizards gets payment.refund.wizard existing records.
func (c *Client) GetPaymentRefundWizards(ids []int64) (*PaymentRefundWizards, error) {
	prws := &PaymentRefundWizards{}
	if err := c.Read(PaymentRefundWizardModel, ids, nil, prws); err != nil {
		return nil, err
	}
	return prws, nil
}

// FindPaymentRefundWizard finds payment.refund.wizard record by querying it with criteria.
func (c *Client) FindPaymentRefundWizard(criteria *Criteria) (*PaymentRefundWizard, error) {
	prws := &PaymentRefundWizards{}
	if err := c.SearchRead(PaymentRefundWizardModel, criteria, NewOptions().Limit(1), prws); err != nil {
		return nil, err
	}
	return &((*prws)[0]), nil
}

// FindPaymentRefundWizards finds payment.refund.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentRefundWizards(criteria *Criteria, options *Options) (*PaymentRefundWizards, error) {
	prws := &PaymentRefundWizards{}
	if err := c.SearchRead(PaymentRefundWizardModel, criteria, options, prws); err != nil {
		return nil, err
	}
	return prws, nil
}

// FindPaymentRefundWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentRefundWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PaymentRefundWizardModel, criteria, options)
}

// FindPaymentRefundWizardId finds record id by querying it with criteria.
func (c *Client) FindPaymentRefundWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentRefundWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
