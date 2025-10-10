package odoo

// PaymentProvider represents payment.provider model.
type PaymentProvider struct {
	AllowExpressCheckout      *Bool      `xmlrpc:"allow_express_checkout,omitempty"`
	AllowTokenization         *Bool      `xmlrpc:"allow_tokenization,omitempty"`
	AuthMsg                   *String    `xmlrpc:"auth_msg,omitempty"`
	AvailableCountryIds       *Relation  `xmlrpc:"available_country_ids,omitempty"`
	AvailableCurrencyIds      *Relation  `xmlrpc:"available_currency_ids,omitempty"`
	CancelMsg                 *String    `xmlrpc:"cancel_msg,omitempty"`
	CaptureManually           *Bool      `xmlrpc:"capture_manually,omitempty"`
	Code                      *Selection `xmlrpc:"code,omitempty"`
	Color                     *Int       `xmlrpc:"color,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomMode                *Selection `xmlrpc:"custom_mode,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	DoneMsg                   *String    `xmlrpc:"done_msg,omitempty"`
	ExpressCheckoutFormViewId *Many2One  `xmlrpc:"express_checkout_form_view_id,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	Image128                  *String    `xmlrpc:"image_128,omitempty"`
	InlineFormViewId          *Many2One  `xmlrpc:"inline_form_view_id,omitempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omitempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omitempty"`
	MainCurrencyId            *Many2One  `xmlrpc:"main_currency_id,omitempty"`
	MaximumAmount             *Float     `xmlrpc:"maximum_amount,omitempty"`
	ModuleId                  *Many2One  `xmlrpc:"module_id,omitempty"`
	ModuleState               *Selection `xmlrpc:"module_state,omitempty"`
	ModuleToBuy               *Bool      `xmlrpc:"module_to_buy,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	PaymentMethodIds          *Relation  `xmlrpc:"payment_method_ids,omitempty"`
	PendingMsg                *String    `xmlrpc:"pending_msg,omitempty"`
	PreMsg                    *String    `xmlrpc:"pre_msg,omitempty"`
	QrCode                    *Bool      `xmlrpc:"qr_code,omitempty"`
	RedirectFormViewId        *Many2One  `xmlrpc:"redirect_form_view_id,omitempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omitempty"`
	SoReferenceType           *Selection `xmlrpc:"so_reference_type,omitempty"`
	State                     *Selection `xmlrpc:"state,omitempty"`
	SupportExpressCheckout    *Bool      `xmlrpc:"support_express_checkout,omitempty"`
	SupportManualCapture      *Selection `xmlrpc:"support_manual_capture,omitempty"`
	SupportRefund             *Selection `xmlrpc:"support_refund,omitempty"`
	SupportTokenization       *Bool      `xmlrpc:"support_tokenization,omitempty"`
	TokenInlineFormViewId     *Many2One  `xmlrpc:"token_inline_form_view_id,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PaymentProviders represents array of payment.provider model.
type PaymentProviders []PaymentProvider

// PaymentProviderModel is the odoo model name.
const PaymentProviderModel = "payment.provider"

// Many2One convert PaymentProvider to *Many2One.
func (pp *PaymentProvider) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreatePaymentProvider creates a new payment.provider model and returns its id.
func (c *Client) CreatePaymentProvider(pp *PaymentProvider) (int64, error) {
	ids, err := c.CreatePaymentProviders([]*PaymentProvider{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePaymentProvider creates a new payment.provider model and returns its id.
func (c *Client) CreatePaymentProviders(pps []*PaymentProvider) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(PaymentProviderModel, vv, nil)
}

// UpdatePaymentProvider updates an existing payment.provider record.
func (c *Client) UpdatePaymentProvider(pp *PaymentProvider) error {
	return c.UpdatePaymentProviders([]int64{pp.Id.Get()}, pp)
}

// UpdatePaymentProviders updates existing payment.provider records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdatePaymentProviders(ids []int64, pp *PaymentProvider) error {
	return c.Update(PaymentProviderModel, ids, pp, nil)
}

// DeletePaymentProvider deletes an existing payment.provider record.
func (c *Client) DeletePaymentProvider(id int64) error {
	return c.DeletePaymentProviders([]int64{id})
}

// DeletePaymentProviders deletes existing payment.provider records.
func (c *Client) DeletePaymentProviders(ids []int64) error {
	return c.Delete(PaymentProviderModel, ids)
}

// GetPaymentProvider gets payment.provider existing record.
func (c *Client) GetPaymentProvider(id int64) (*PaymentProvider, error) {
	pps, err := c.GetPaymentProviders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// GetPaymentProviders gets payment.provider existing records.
func (c *Client) GetPaymentProviders(ids []int64) (*PaymentProviders, error) {
	pps := &PaymentProviders{}
	if err := c.Read(PaymentProviderModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPaymentProvider finds payment.provider record by querying it with criteria.
func (c *Client) FindPaymentProvider(criteria *Criteria) (*PaymentProvider, error) {
	pps := &PaymentProviders{}
	if err := c.SearchRead(PaymentProviderModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	return &((*pps)[0]), nil
}

// FindPaymentProviders finds payment.provider records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentProviders(criteria *Criteria, options *Options) (*PaymentProviders, error) {
	pps := &PaymentProviders{}
	if err := c.SearchRead(PaymentProviderModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPaymentProviderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentProviderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PaymentProviderModel, criteria, options)
}

// FindPaymentProviderId finds record id by querying it with criteria.
func (c *Client) FindPaymentProviderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentProviderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
