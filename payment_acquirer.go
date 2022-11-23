package odoo

import (
	"fmt"
)

// PaymentAcquirer represents payment.acquirer model.
type PaymentAcquirer struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	AuthorizeImplemented       *Bool      `xmlrpc:"authorize_implemented,omptempty"`
	CancelMsg                  *String    `xmlrpc:"cancel_msg,omptempty"`
	CaptureManually            *Bool      `xmlrpc:"capture_manually,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryIds                 *Relation  `xmlrpc:"country_ids,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description                *String    `xmlrpc:"description,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	DoneMsg                    *String    `xmlrpc:"done_msg,omptempty"`
	Environment                *Selection `xmlrpc:"environment,omptempty"`
	ErrorMsg                   *String    `xmlrpc:"error_msg,omptempty"`
	FeesActive                 *Bool      `xmlrpc:"fees_active,omptempty"`
	FeesDomFixed               *Float     `xmlrpc:"fees_dom_fixed,omptempty"`
	FeesDomVar                 *Float     `xmlrpc:"fees_dom_var,omptempty"`
	FeesImplemented            *Bool      `xmlrpc:"fees_implemented,omptempty"`
	FeesIntFixed               *Float     `xmlrpc:"fees_int_fixed,omptempty"`
	FeesIntVar                 *Float     `xmlrpc:"fees_int_var,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	Image                      *String    `xmlrpc:"image,omptempty"`
	ImageMedium                *String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall                 *String    `xmlrpc:"image_small,omptempty"`
	JournalId                  *Many2One  `xmlrpc:"journal_id,omptempty"`
	ModuleId                   *Many2One  `xmlrpc:"module_id,omptempty"`
	ModuleState                *Selection `xmlrpc:"module_state,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	PaymentFlow                *Selection `xmlrpc:"payment_flow,omptempty"`
	PaymentIconIds             *Relation  `xmlrpc:"payment_icon_ids,omptempty"`
	PendingMsg                 *String    `xmlrpc:"pending_msg,omptempty"`
	PostMsg                    *String    `xmlrpc:"post_msg,omptempty"`
	PreMsg                     *String    `xmlrpc:"pre_msg,omptempty"`
	Provider                   *Selection `xmlrpc:"provider,omptempty"`
	RegistrationViewTemplateId *Many2One  `xmlrpc:"registration_view_template_id,omptempty"`
	SaveToken                  *Selection `xmlrpc:"save_token,omptempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omptempty"`
	SpecificCountries          *Bool      `xmlrpc:"specific_countries,omptempty"`
	TokenImplemented           *Bool      `xmlrpc:"token_implemented,omptempty"`
	ViewTemplateId             *Many2One  `xmlrpc:"view_template_id,omptempty"`
	WebsitePublished           *Bool      `xmlrpc:"website_published,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PaymentAcquirers represents array of payment.acquirer model.
type PaymentAcquirers []PaymentAcquirer

// PaymentAcquirerModel is the odoo model name.
const PaymentAcquirerModel = "payment.acquirer"

// Many2One convert PaymentAcquirer to *Many2One.
func (pa *PaymentAcquirer) Many2One() *Many2One {
	return NewMany2One(pa.Id.Get(), "")
}

// CreatePaymentAcquirer creates a new payment.acquirer model and returns its id.
func (c *Client) CreatePaymentAcquirer(pa *PaymentAcquirer) (int64, error) {
	return c.Create(PaymentAcquirerModel, pa)
}

// UpdatePaymentAcquirer updates an existing payment.acquirer record.
func (c *Client) UpdatePaymentAcquirer(pa *PaymentAcquirer) error {
	return c.UpdatePaymentAcquirers([]int64{pa.Id.Get()}, pa)
}

// UpdatePaymentAcquirers updates existing payment.acquirer records.
// All records (represented by ids) will be updated by pa values.
func (c *Client) UpdatePaymentAcquirers(ids []int64, pa *PaymentAcquirer) error {
	return c.Update(PaymentAcquirerModel, ids, pa)
}

// DeletePaymentAcquirer deletes an existing payment.acquirer record.
func (c *Client) DeletePaymentAcquirer(id int64) error {
	return c.DeletePaymentAcquirers([]int64{id})
}

// DeletePaymentAcquirers deletes existing payment.acquirer records.
func (c *Client) DeletePaymentAcquirers(ids []int64) error {
	return c.Delete(PaymentAcquirerModel, ids)
}

// GetPaymentAcquirer gets payment.acquirer existing record.
func (c *Client) GetPaymentAcquirer(id int64) (*PaymentAcquirer, error) {
	pas, err := c.GetPaymentAcquirers([]int64{id})
	if err != nil {
		return nil, err
	}
	if pas != nil && len(*pas) > 0 {
		return &((*pas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.acquirer not found", id)
}

// GetPaymentAcquirers gets payment.acquirer existing records.
func (c *Client) GetPaymentAcquirers(ids []int64) (*PaymentAcquirers, error) {
	pas := &PaymentAcquirers{}
	if err := c.Read(PaymentAcquirerModel, ids, nil, pas); err != nil {
		return nil, err
	}
	return pas, nil
}

// FindPaymentAcquirer finds payment.acquirer record by querying it with criteria.
func (c *Client) FindPaymentAcquirer(criteria *Criteria) (*PaymentAcquirer, error) {
	pas := &PaymentAcquirers{}
	if err := c.SearchRead(PaymentAcquirerModel, criteria, NewOptions().Limit(1), pas); err != nil {
		return nil, err
	}
	if pas != nil && len(*pas) > 0 {
		return &((*pas)[0]), nil
	}
	return nil, fmt.Errorf("payment.acquirer was not found with criteria %v", criteria)
}

// FindPaymentAcquirers finds payment.acquirer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentAcquirers(criteria *Criteria, options *Options) (*PaymentAcquirers, error) {
	pas := &PaymentAcquirers{}
	if err := c.SearchRead(PaymentAcquirerModel, criteria, options, pas); err != nil {
		return nil, err
	}
	return pas, nil
}

// FindPaymentAcquirerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentAcquirerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentAcquirerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentAcquirerId finds record id by querying it with criteria.
func (c *Client) FindPaymentAcquirerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentAcquirerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("payment.acquirer was not found with criteria %v and options %v", criteria, options)
}
