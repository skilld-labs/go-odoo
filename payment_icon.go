package odoo

// PaymentIcon represents payment.icon model.
type PaymentIcon struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omitempty"`
	AcquirerIds      *Relation `xmlrpc:"acquirer_ids,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	Image            *String   `xmlrpc:"image,omitempty"`
	ImagePaymentForm *String   `xmlrpc:"image_payment_form,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PaymentIcons represents array of payment.icon model.
type PaymentIcons []PaymentIcon

// PaymentIconModel is the odoo model name.
const PaymentIconModel = "payment.icon"

// Many2One convert PaymentIcon to *Many2One.
func (pi *PaymentIcon) Many2One() *Many2One {
	return NewMany2One(pi.Id.Get(), "")
}

// CreatePaymentIcon creates a new payment.icon model and returns its id.
func (c *Client) CreatePaymentIcon(pi *PaymentIcon) (int64, error) {
	ids, err := c.CreatePaymentIcons([]*PaymentIcon{pi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePaymentIcon creates a new payment.icon model and returns its id.
func (c *Client) CreatePaymentIcons(pis []*PaymentIcon) ([]int64, error) {
	var vv []interface{}
	for _, v := range pis {
		vv = append(vv, v)
	}
	return c.Create(PaymentIconModel, vv, nil)
}

// UpdatePaymentIcon updates an existing payment.icon record.
func (c *Client) UpdatePaymentIcon(pi *PaymentIcon) error {
	return c.UpdatePaymentIcons([]int64{pi.Id.Get()}, pi)
}

// UpdatePaymentIcons updates existing payment.icon records.
// All records (represented by ids) will be updated by pi values.
func (c *Client) UpdatePaymentIcons(ids []int64, pi *PaymentIcon) error {
	return c.Update(PaymentIconModel, ids, pi, nil)
}

// DeletePaymentIcon deletes an existing payment.icon record.
func (c *Client) DeletePaymentIcon(id int64) error {
	return c.DeletePaymentIcons([]int64{id})
}

// DeletePaymentIcons deletes existing payment.icon records.
func (c *Client) DeletePaymentIcons(ids []int64) error {
	return c.Delete(PaymentIconModel, ids)
}

// GetPaymentIcon gets payment.icon existing record.
func (c *Client) GetPaymentIcon(id int64) (*PaymentIcon, error) {
	pis, err := c.GetPaymentIcons([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pis)[0]), nil
}

// GetPaymentIcons gets payment.icon existing records.
func (c *Client) GetPaymentIcons(ids []int64) (*PaymentIcons, error) {
	pis := &PaymentIcons{}
	if err := c.Read(PaymentIconModel, ids, nil, pis); err != nil {
		return nil, err
	}
	return pis, nil
}

// FindPaymentIcon finds payment.icon record by querying it with criteria.
func (c *Client) FindPaymentIcon(criteria *Criteria) (*PaymentIcon, error) {
	pis := &PaymentIcons{}
	if err := c.SearchRead(PaymentIconModel, criteria, NewOptions().Limit(1), pis); err != nil {
		return nil, err
	}
	return &((*pis)[0]), nil
}

// FindPaymentIcons finds payment.icon records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentIcons(criteria *Criteria, options *Options) (*PaymentIcons, error) {
	pis := &PaymentIcons{}
	if err := c.SearchRead(PaymentIconModel, criteria, options, pis); err != nil {
		return nil, err
	}
	return pis, nil
}

// FindPaymentIconIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentIconIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PaymentIconModel, criteria, options)
}

// FindPaymentIconId finds record id by querying it with criteria.
func (c *Client) FindPaymentIconId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentIconModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
