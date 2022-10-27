package odoo

import (
	"fmt"
)

// SalePaymentAcquirerOnboardingWizard represents sale.payment.acquirer.onboarding.wizard model.
type SalePaymentAcquirerOnboardingWizard struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	DataFetched          *Bool      `xmlrpc:"_data_fetched,omitempty"`
	AccNumber            *String    `xmlrpc:"acc_number,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	JournalName          *String    `xmlrpc:"journal_name,omitempty"`
	ManualName           *String    `xmlrpc:"manual_name,omitempty"`
	ManualPostMsg        *String    `xmlrpc:"manual_post_msg,omitempty"`
	PaymentMethod        *Selection `xmlrpc:"payment_method,omitempty"`
	PaypalEmailAccount   *String    `xmlrpc:"paypal_email_account,omitempty"`
	PaypalPdtToken       *String    `xmlrpc:"paypal_pdt_token,omitempty"`
	PaypalSellerAccount  *String    `xmlrpc:"paypal_seller_account,omitempty"`
	PaypalUserType       *Selection `xmlrpc:"paypal_user_type,omitempty"`
	StripePublishableKey *String    `xmlrpc:"stripe_publishable_key,omitempty"`
	StripeSecretKey      *String    `xmlrpc:"stripe_secret_key,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SalePaymentAcquirerOnboardingWizards represents array of sale.payment.acquirer.onboarding.wizard model.
type SalePaymentAcquirerOnboardingWizards []SalePaymentAcquirerOnboardingWizard

// SalePaymentAcquirerOnboardingWizardModel is the odoo model name.
const SalePaymentAcquirerOnboardingWizardModel = "sale.payment.acquirer.onboarding.wizard"

// Many2One convert SalePaymentAcquirerOnboardingWizard to *Many2One.
func (spaow *SalePaymentAcquirerOnboardingWizard) Many2One() *Many2One {
	return NewMany2One(spaow.Id.Get(), "")
}

// CreateSalePaymentAcquirerOnboardingWizard creates a new sale.payment.acquirer.onboarding.wizard model and returns its id.
func (c *Client) CreateSalePaymentAcquirerOnboardingWizard(spaow *SalePaymentAcquirerOnboardingWizard) (int64, error) {
	return c.Create(SalePaymentAcquirerOnboardingWizardModel, spaow)
}

// UpdateSalePaymentAcquirerOnboardingWizard updates an existing sale.payment.acquirer.onboarding.wizard record.
func (c *Client) UpdateSalePaymentAcquirerOnboardingWizard(spaow *SalePaymentAcquirerOnboardingWizard) error {
	return c.UpdateSalePaymentAcquirerOnboardingWizards([]int64{spaow.Id.Get()}, spaow)
}

// UpdateSalePaymentAcquirerOnboardingWizards updates existing sale.payment.acquirer.onboarding.wizard records.
// All records (represented by ids) will be updated by spaow values.
func (c *Client) UpdateSalePaymentAcquirerOnboardingWizards(ids []int64, spaow *SalePaymentAcquirerOnboardingWizard) error {
	return c.Update(SalePaymentAcquirerOnboardingWizardModel, ids, spaow)
}

// DeleteSalePaymentAcquirerOnboardingWizard deletes an existing sale.payment.acquirer.onboarding.wizard record.
func (c *Client) DeleteSalePaymentAcquirerOnboardingWizard(id int64) error {
	return c.DeleteSalePaymentAcquirerOnboardingWizards([]int64{id})
}

// DeleteSalePaymentAcquirerOnboardingWizards deletes existing sale.payment.acquirer.onboarding.wizard records.
func (c *Client) DeleteSalePaymentAcquirerOnboardingWizards(ids []int64) error {
	return c.Delete(SalePaymentAcquirerOnboardingWizardModel, ids)
}

// GetSalePaymentAcquirerOnboardingWizard gets sale.payment.acquirer.onboarding.wizard existing record.
func (c *Client) GetSalePaymentAcquirerOnboardingWizard(id int64) (*SalePaymentAcquirerOnboardingWizard, error) {
	spaows, err := c.GetSalePaymentAcquirerOnboardingWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if spaows != nil && len(*spaows) > 0 {
		return &((*spaows)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.payment.acquirer.onboarding.wizard not found", id)
}

// GetSalePaymentAcquirerOnboardingWizards gets sale.payment.acquirer.onboarding.wizard existing records.
func (c *Client) GetSalePaymentAcquirerOnboardingWizards(ids []int64) (*SalePaymentAcquirerOnboardingWizards, error) {
	spaows := &SalePaymentAcquirerOnboardingWizards{}
	if err := c.Read(SalePaymentAcquirerOnboardingWizardModel, ids, nil, spaows); err != nil {
		return nil, err
	}
	return spaows, nil
}

// FindSalePaymentAcquirerOnboardingWizard finds sale.payment.acquirer.onboarding.wizard record by querying it with criteria.
func (c *Client) FindSalePaymentAcquirerOnboardingWizard(criteria *Criteria) (*SalePaymentAcquirerOnboardingWizard, error) {
	spaows := &SalePaymentAcquirerOnboardingWizards{}
	if err := c.SearchRead(SalePaymentAcquirerOnboardingWizardModel, criteria, NewOptions().Limit(1), spaows); err != nil {
		return nil, err
	}
	if spaows != nil && len(*spaows) > 0 {
		return &((*spaows)[0]), nil
	}
	return nil, fmt.Errorf("sale.payment.acquirer.onboarding.wizard was not found")
}

// FindSalePaymentAcquirerOnboardingWizards finds sale.payment.acquirer.onboarding.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSalePaymentAcquirerOnboardingWizards(criteria *Criteria, options *Options) (*SalePaymentAcquirerOnboardingWizards, error) {
	spaows := &SalePaymentAcquirerOnboardingWizards{}
	if err := c.SearchRead(SalePaymentAcquirerOnboardingWizardModel, criteria, options, spaows); err != nil {
		return nil, err
	}
	return spaows, nil
}

// FindSalePaymentAcquirerOnboardingWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSalePaymentAcquirerOnboardingWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SalePaymentAcquirerOnboardingWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSalePaymentAcquirerOnboardingWizardId finds record id by querying it with criteria.
func (c *Client) FindSalePaymentAcquirerOnboardingWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SalePaymentAcquirerOnboardingWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.payment.acquirer.onboarding.wizard was not found")
}
