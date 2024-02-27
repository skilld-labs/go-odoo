package odoo

// PaymentAcquirer represents payment.acquirer model.
type PaymentAcquirer struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omitempty"`
	AuthorizeImplemented       *Bool      `xmlrpc:"authorize_implemented,omitempty"`
	CancelMsg                  *String    `xmlrpc:"cancel_msg,omitempty"`
	CaptureManually            *Bool      `xmlrpc:"capture_manually,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryIds                 *Relation  `xmlrpc:"country_ids,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description                *String    `xmlrpc:"description,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	DoneMsg                    *String    `xmlrpc:"done_msg,omitempty"`
	Environment                *Selection `xmlrpc:"environment,omitempty"`
	ErrorMsg                   *String    `xmlrpc:"error_msg,omitempty"`
	FeesActive                 *Bool      `xmlrpc:"fees_active,omitempty"`
	FeesDomFixed               *Float     `xmlrpc:"fees_dom_fixed,omitempty"`
	FeesDomVar                 *Float     `xmlrpc:"fees_dom_var,omitempty"`
	FeesImplemented            *Bool      `xmlrpc:"fees_implemented,omitempty"`
	FeesIntFixed               *Float     `xmlrpc:"fees_int_fixed,omitempty"`
	FeesIntVar                 *Float     `xmlrpc:"fees_int_var,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	Image                      *String    `xmlrpc:"image,omitempty"`
	ImageMedium                *String    `xmlrpc:"image_medium,omitempty"`
	ImageSmall                 *String    `xmlrpc:"image_small,omitempty"`
	JournalId                  *Many2One  `xmlrpc:"journal_id,omitempty"`
	ModuleId                   *Many2One  `xmlrpc:"module_id,omitempty"`
	ModuleState                *Selection `xmlrpc:"module_state,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	PaymentFlow                *Selection `xmlrpc:"payment_flow,omitempty"`
	PaymentIconIds             *Relation  `xmlrpc:"payment_icon_ids,omitempty"`
	PendingMsg                 *String    `xmlrpc:"pending_msg,omitempty"`
	PostMsg                    *String    `xmlrpc:"post_msg,omitempty"`
	PreMsg                     *String    `xmlrpc:"pre_msg,omitempty"`
	Provider                   *Selection `xmlrpc:"provider,omitempty"`
	RegistrationViewTemplateId *Many2One  `xmlrpc:"registration_view_template_id,omitempty"`
	SaveToken                  *Selection `xmlrpc:"save_token,omitempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omitempty"`
	SpecificCountries          *Bool      `xmlrpc:"specific_countries,omitempty"`
	TokenImplemented           *Bool      `xmlrpc:"token_implemented,omitempty"`
	ViewTemplateId             *Many2One  `xmlrpc:"view_template_id,omitempty"`
	WebsitePublished           *Bool      `xmlrpc:"website_published,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreatePaymentAcquirers([]*PaymentAcquirer{pa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePaymentAcquirer creates a new payment.acquirer model and returns its id.
func (c *Client) CreatePaymentAcquirers(pas []*PaymentAcquirer) ([]int64, error) {
	var vv []interface{}
	for _, v := range pas {
		vv = append(vv, v)
	}
	return c.Create(PaymentAcquirerModel, vv, nil)
}

// UpdatePaymentAcquirer updates an existing payment.acquirer record.
func (c *Client) UpdatePaymentAcquirer(pa *PaymentAcquirer) error {
	return c.UpdatePaymentAcquirers([]int64{pa.Id.Get()}, pa)
}

// UpdatePaymentAcquirers updates existing payment.acquirer records.
// All records (represented by ids) will be updated by pa values.
func (c *Client) UpdatePaymentAcquirers(ids []int64, pa *PaymentAcquirer) error {
	return c.Update(PaymentAcquirerModel, ids, pa, nil)
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
	return &((*pas)[0]), nil
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
	return &((*pas)[0]), nil
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
	return c.Search(PaymentAcquirerModel, criteria, options)
}

// FindPaymentAcquirerId finds record id by querying it with criteria.
func (c *Client) FindPaymentAcquirerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentAcquirerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
