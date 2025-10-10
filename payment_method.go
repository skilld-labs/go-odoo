package odoo

// PaymentMethod represents payment.method model.
type PaymentMethod struct {
	Active                 *Bool      `xmlrpc:"active,omitempty"`
	BrandIds               *Relation  `xmlrpc:"brand_ids,omitempty"`
	Code                   *String    `xmlrpc:"code,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	Image                  *String    `xmlrpc:"image,omitempty"`
	ImagePaymentForm       *String    `xmlrpc:"image_payment_form,omitempty"`
	IsPrimary              *Bool      `xmlrpc:"is_primary,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	PrimaryPaymentMethodId *Many2One  `xmlrpc:"primary_payment_method_id,omitempty"`
	ProviderIds            *Relation  `xmlrpc:"provider_ids,omitempty"`
	Sequence               *Int       `xmlrpc:"sequence,omitempty"`
	SupportExpressCheckout *Bool      `xmlrpc:"support_express_checkout,omitempty"`
	SupportRefund          *Selection `xmlrpc:"support_refund,omitempty"`
	SupportTokenization    *Bool      `xmlrpc:"support_tokenization,omitempty"`
	SupportedCountryIds    *Relation  `xmlrpc:"supported_country_ids,omitempty"`
	SupportedCurrencyIds   *Relation  `xmlrpc:"supported_currency_ids,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PaymentMethods represents array of payment.method model.
type PaymentMethods []PaymentMethod

// PaymentMethodModel is the odoo model name.
const PaymentMethodModel = "payment.method"

// Many2One convert PaymentMethod to *Many2One.
func (pm *PaymentMethod) Many2One() *Many2One {
	return NewMany2One(pm.Id.Get(), "")
}

// CreatePaymentMethod creates a new payment.method model and returns its id.
func (c *Client) CreatePaymentMethod(pm *PaymentMethod) (int64, error) {
	ids, err := c.CreatePaymentMethods([]*PaymentMethod{pm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePaymentMethod creates a new payment.method model and returns its id.
func (c *Client) CreatePaymentMethods(pms []*PaymentMethod) ([]int64, error) {
	var vv []interface{}
	for _, v := range pms {
		vv = append(vv, v)
	}
	return c.Create(PaymentMethodModel, vv, nil)
}

// UpdatePaymentMethod updates an existing payment.method record.
func (c *Client) UpdatePaymentMethod(pm *PaymentMethod) error {
	return c.UpdatePaymentMethods([]int64{pm.Id.Get()}, pm)
}

// UpdatePaymentMethods updates existing payment.method records.
// All records (represented by ids) will be updated by pm values.
func (c *Client) UpdatePaymentMethods(ids []int64, pm *PaymentMethod) error {
	return c.Update(PaymentMethodModel, ids, pm, nil)
}

// DeletePaymentMethod deletes an existing payment.method record.
func (c *Client) DeletePaymentMethod(id int64) error {
	return c.DeletePaymentMethods([]int64{id})
}

// DeletePaymentMethods deletes existing payment.method records.
func (c *Client) DeletePaymentMethods(ids []int64) error {
	return c.Delete(PaymentMethodModel, ids)
}

// GetPaymentMethod gets payment.method existing record.
func (c *Client) GetPaymentMethod(id int64) (*PaymentMethod, error) {
	pms, err := c.GetPaymentMethods([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pms)[0]), nil
}

// GetPaymentMethods gets payment.method existing records.
func (c *Client) GetPaymentMethods(ids []int64) (*PaymentMethods, error) {
	pms := &PaymentMethods{}
	if err := c.Read(PaymentMethodModel, ids, nil, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindPaymentMethod finds payment.method record by querying it with criteria.
func (c *Client) FindPaymentMethod(criteria *Criteria) (*PaymentMethod, error) {
	pms := &PaymentMethods{}
	if err := c.SearchRead(PaymentMethodModel, criteria, NewOptions().Limit(1), pms); err != nil {
		return nil, err
	}
	return &((*pms)[0]), nil
}

// FindPaymentMethods finds payment.method records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentMethods(criteria *Criteria, options *Options) (*PaymentMethods, error) {
	pms := &PaymentMethods{}
	if err := c.SearchRead(PaymentMethodModel, criteria, options, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindPaymentMethodIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentMethodIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PaymentMethodModel, criteria, options)
}

// FindPaymentMethodId finds record id by querying it with criteria.
func (c *Client) FindPaymentMethodId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentMethodModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
