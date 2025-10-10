package odoo

// PaymentCaptureWizard represents payment.capture.wizard model.
type PaymentCaptureWizard struct {
	AmountToCapture        *Float    `xmlrpc:"amount_to_capture,omitempty"`
	AuthorizedAmount       *Float    `xmlrpc:"authorized_amount,omitempty"`
	AvailableAmount        *Float    `xmlrpc:"available_amount,omitempty"`
	CapturedAmount         *Float    `xmlrpc:"captured_amount,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId             *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	HasDraftChildren       *Bool     `xmlrpc:"has_draft_children,omitempty"`
	HasRemainingAmount     *Bool     `xmlrpc:"has_remaining_amount,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	IsAmountToCaptureValid *Bool     `xmlrpc:"is_amount_to_capture_valid,omitempty"`
	SupportPartialCapture  *Bool     `xmlrpc:"support_partial_capture,omitempty"`
	TransactionIds         *Relation `xmlrpc:"transaction_ids,omitempty"`
	VoidRemainingAmount    *Bool     `xmlrpc:"void_remaining_amount,omitempty"`
	VoidedAmount           *Float    `xmlrpc:"voided_amount,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PaymentCaptureWizards represents array of payment.capture.wizard model.
type PaymentCaptureWizards []PaymentCaptureWizard

// PaymentCaptureWizardModel is the odoo model name.
const PaymentCaptureWizardModel = "payment.capture.wizard"

// Many2One convert PaymentCaptureWizard to *Many2One.
func (pcw *PaymentCaptureWizard) Many2One() *Many2One {
	return NewMany2One(pcw.Id.Get(), "")
}

// CreatePaymentCaptureWizard creates a new payment.capture.wizard model and returns its id.
func (c *Client) CreatePaymentCaptureWizard(pcw *PaymentCaptureWizard) (int64, error) {
	ids, err := c.CreatePaymentCaptureWizards([]*PaymentCaptureWizard{pcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePaymentCaptureWizard creates a new payment.capture.wizard model and returns its id.
func (c *Client) CreatePaymentCaptureWizards(pcws []*PaymentCaptureWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcws {
		vv = append(vv, v)
	}
	return c.Create(PaymentCaptureWizardModel, vv, nil)
}

// UpdatePaymentCaptureWizard updates an existing payment.capture.wizard record.
func (c *Client) UpdatePaymentCaptureWizard(pcw *PaymentCaptureWizard) error {
	return c.UpdatePaymentCaptureWizards([]int64{pcw.Id.Get()}, pcw)
}

// UpdatePaymentCaptureWizards updates existing payment.capture.wizard records.
// All records (represented by ids) will be updated by pcw values.
func (c *Client) UpdatePaymentCaptureWizards(ids []int64, pcw *PaymentCaptureWizard) error {
	return c.Update(PaymentCaptureWizardModel, ids, pcw, nil)
}

// DeletePaymentCaptureWizard deletes an existing payment.capture.wizard record.
func (c *Client) DeletePaymentCaptureWizard(id int64) error {
	return c.DeletePaymentCaptureWizards([]int64{id})
}

// DeletePaymentCaptureWizards deletes existing payment.capture.wizard records.
func (c *Client) DeletePaymentCaptureWizards(ids []int64) error {
	return c.Delete(PaymentCaptureWizardModel, ids)
}

// GetPaymentCaptureWizard gets payment.capture.wizard existing record.
func (c *Client) GetPaymentCaptureWizard(id int64) (*PaymentCaptureWizard, error) {
	pcws, err := c.GetPaymentCaptureWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcws)[0]), nil
}

// GetPaymentCaptureWizards gets payment.capture.wizard existing records.
func (c *Client) GetPaymentCaptureWizards(ids []int64) (*PaymentCaptureWizards, error) {
	pcws := &PaymentCaptureWizards{}
	if err := c.Read(PaymentCaptureWizardModel, ids, nil, pcws); err != nil {
		return nil, err
	}
	return pcws, nil
}

// FindPaymentCaptureWizard finds payment.capture.wizard record by querying it with criteria.
func (c *Client) FindPaymentCaptureWizard(criteria *Criteria) (*PaymentCaptureWizard, error) {
	pcws := &PaymentCaptureWizards{}
	if err := c.SearchRead(PaymentCaptureWizardModel, criteria, NewOptions().Limit(1), pcws); err != nil {
		return nil, err
	}
	return &((*pcws)[0]), nil
}

// FindPaymentCaptureWizards finds payment.capture.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentCaptureWizards(criteria *Criteria, options *Options) (*PaymentCaptureWizards, error) {
	pcws := &PaymentCaptureWizards{}
	if err := c.SearchRead(PaymentCaptureWizardModel, criteria, options, pcws); err != nil {
		return nil, err
	}
	return pcws, nil
}

// FindPaymentCaptureWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentCaptureWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PaymentCaptureWizardModel, criteria, options)
}

// FindPaymentCaptureWizardId finds record id by querying it with criteria.
func (c *Client) FindPaymentCaptureWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentCaptureWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
