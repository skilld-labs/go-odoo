package odoo

// SalePaymentProviderOnboardingWizard represents sale.payment.provider.onboarding.wizard model.
type SalePaymentProviderOnboardingWizard struct {
	DataFetched        *Bool      `xmlrpc:"_data_fetched,omitempty"`
	AccNumber          *String    `xmlrpc:"acc_number,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	JournalName        *String    `xmlrpc:"journal_name,omitempty"`
	ManualName         *String    `xmlrpc:"manual_name,omitempty"`
	ManualPostMsg      *String    `xmlrpc:"manual_post_msg,omitempty"`
	PaymentMethod      *Selection `xmlrpc:"payment_method,omitempty"`
	PaypalEmailAccount *String    `xmlrpc:"paypal_email_account,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SalePaymentProviderOnboardingWizards represents array of sale.payment.provider.onboarding.wizard model.
type SalePaymentProviderOnboardingWizards []SalePaymentProviderOnboardingWizard

// SalePaymentProviderOnboardingWizardModel is the odoo model name.
const SalePaymentProviderOnboardingWizardModel = "sale.payment.provider.onboarding.wizard"

// Many2One convert SalePaymentProviderOnboardingWizard to *Many2One.
func (sppow *SalePaymentProviderOnboardingWizard) Many2One() *Many2One {
	return NewMany2One(sppow.Id.Get(), "")
}

// CreateSalePaymentProviderOnboardingWizard creates a new sale.payment.provider.onboarding.wizard model and returns its id.
func (c *Client) CreateSalePaymentProviderOnboardingWizard(sppow *SalePaymentProviderOnboardingWizard) (int64, error) {
	ids, err := c.CreateSalePaymentProviderOnboardingWizards([]*SalePaymentProviderOnboardingWizard{sppow})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSalePaymentProviderOnboardingWizard creates a new sale.payment.provider.onboarding.wizard model and returns its id.
func (c *Client) CreateSalePaymentProviderOnboardingWizards(sppows []*SalePaymentProviderOnboardingWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range sppows {
		vv = append(vv, v)
	}
	return c.Create(SalePaymentProviderOnboardingWizardModel, vv, nil)
}

// UpdateSalePaymentProviderOnboardingWizard updates an existing sale.payment.provider.onboarding.wizard record.
func (c *Client) UpdateSalePaymentProviderOnboardingWizard(sppow *SalePaymentProviderOnboardingWizard) error {
	return c.UpdateSalePaymentProviderOnboardingWizards([]int64{sppow.Id.Get()}, sppow)
}

// UpdateSalePaymentProviderOnboardingWizards updates existing sale.payment.provider.onboarding.wizard records.
// All records (represented by ids) will be updated by sppow values.
func (c *Client) UpdateSalePaymentProviderOnboardingWizards(ids []int64, sppow *SalePaymentProviderOnboardingWizard) error {
	return c.Update(SalePaymentProviderOnboardingWizardModel, ids, sppow, nil)
}

// DeleteSalePaymentProviderOnboardingWizard deletes an existing sale.payment.provider.onboarding.wizard record.
func (c *Client) DeleteSalePaymentProviderOnboardingWizard(id int64) error {
	return c.DeleteSalePaymentProviderOnboardingWizards([]int64{id})
}

// DeleteSalePaymentProviderOnboardingWizards deletes existing sale.payment.provider.onboarding.wizard records.
func (c *Client) DeleteSalePaymentProviderOnboardingWizards(ids []int64) error {
	return c.Delete(SalePaymentProviderOnboardingWizardModel, ids)
}

// GetSalePaymentProviderOnboardingWizard gets sale.payment.provider.onboarding.wizard existing record.
func (c *Client) GetSalePaymentProviderOnboardingWizard(id int64) (*SalePaymentProviderOnboardingWizard, error) {
	sppows, err := c.GetSalePaymentProviderOnboardingWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sppows)[0]), nil
}

// GetSalePaymentProviderOnboardingWizards gets sale.payment.provider.onboarding.wizard existing records.
func (c *Client) GetSalePaymentProviderOnboardingWizards(ids []int64) (*SalePaymentProviderOnboardingWizards, error) {
	sppows := &SalePaymentProviderOnboardingWizards{}
	if err := c.Read(SalePaymentProviderOnboardingWizardModel, ids, nil, sppows); err != nil {
		return nil, err
	}
	return sppows, nil
}

// FindSalePaymentProviderOnboardingWizard finds sale.payment.provider.onboarding.wizard record by querying it with criteria.
func (c *Client) FindSalePaymentProviderOnboardingWizard(criteria *Criteria) (*SalePaymentProviderOnboardingWizard, error) {
	sppows := &SalePaymentProviderOnboardingWizards{}
	if err := c.SearchRead(SalePaymentProviderOnboardingWizardModel, criteria, NewOptions().Limit(1), sppows); err != nil {
		return nil, err
	}
	return &((*sppows)[0]), nil
}

// FindSalePaymentProviderOnboardingWizards finds sale.payment.provider.onboarding.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSalePaymentProviderOnboardingWizards(criteria *Criteria, options *Options) (*SalePaymentProviderOnboardingWizards, error) {
	sppows := &SalePaymentProviderOnboardingWizards{}
	if err := c.SearchRead(SalePaymentProviderOnboardingWizardModel, criteria, options, sppows); err != nil {
		return nil, err
	}
	return sppows, nil
}

// FindSalePaymentProviderOnboardingWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSalePaymentProviderOnboardingWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SalePaymentProviderOnboardingWizardModel, criteria, options)
}

// FindSalePaymentProviderOnboardingWizardId finds record id by querying it with criteria.
func (c *Client) FindSalePaymentProviderOnboardingWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SalePaymentProviderOnboardingWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
