package odoo

// SaleAdvancePaymentInv represents sale.advance.payment.inv model.
type SaleAdvancePaymentInv struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omitempty"`
	AdvancePaymentMethod *Selection `xmlrpc:"advance_payment_method,omitempty"`
	Amount               *Float     `xmlrpc:"amount,omitempty"`
	Count                *Int       `xmlrpc:"count,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DepositAccountId     *Many2One  `xmlrpc:"deposit_account_id,omitempty"`
	DepositTaxesId       *Relation  `xmlrpc:"deposit_taxes_id,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	ProductId            *Many2One  `xmlrpc:"product_id,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SaleAdvancePaymentInvs represents array of sale.advance.payment.inv model.
type SaleAdvancePaymentInvs []SaleAdvancePaymentInv

// SaleAdvancePaymentInvModel is the odoo model name.
const SaleAdvancePaymentInvModel = "sale.advance.payment.inv"

// Many2One convert SaleAdvancePaymentInv to *Many2One.
func (sapi *SaleAdvancePaymentInv) Many2One() *Many2One {
	return NewMany2One(sapi.Id.Get(), "")
}

// CreateSaleAdvancePaymentInv creates a new sale.advance.payment.inv model and returns its id.
func (c *Client) CreateSaleAdvancePaymentInv(sapi *SaleAdvancePaymentInv) (int64, error) {
	ids, err := c.CreateSaleAdvancePaymentInvs([]*SaleAdvancePaymentInv{sapi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleAdvancePaymentInv creates a new sale.advance.payment.inv model and returns its id.
func (c *Client) CreateSaleAdvancePaymentInvs(sapis []*SaleAdvancePaymentInv) ([]int64, error) {
	var vv []interface{}
	for _, v := range sapis {
		vv = append(vv, v)
	}
	return c.Create(SaleAdvancePaymentInvModel, vv, nil)
}

// UpdateSaleAdvancePaymentInv updates an existing sale.advance.payment.inv record.
func (c *Client) UpdateSaleAdvancePaymentInv(sapi *SaleAdvancePaymentInv) error {
	return c.UpdateSaleAdvancePaymentInvs([]int64{sapi.Id.Get()}, sapi)
}

// UpdateSaleAdvancePaymentInvs updates existing sale.advance.payment.inv records.
// All records (represented by ids) will be updated by sapi values.
func (c *Client) UpdateSaleAdvancePaymentInvs(ids []int64, sapi *SaleAdvancePaymentInv) error {
	return c.Update(SaleAdvancePaymentInvModel, ids, sapi, nil)
}

// DeleteSaleAdvancePaymentInv deletes an existing sale.advance.payment.inv record.
func (c *Client) DeleteSaleAdvancePaymentInv(id int64) error {
	return c.DeleteSaleAdvancePaymentInvs([]int64{id})
}

// DeleteSaleAdvancePaymentInvs deletes existing sale.advance.payment.inv records.
func (c *Client) DeleteSaleAdvancePaymentInvs(ids []int64) error {
	return c.Delete(SaleAdvancePaymentInvModel, ids)
}

// GetSaleAdvancePaymentInv gets sale.advance.payment.inv existing record.
func (c *Client) GetSaleAdvancePaymentInv(id int64) (*SaleAdvancePaymentInv, error) {
	sapis, err := c.GetSaleAdvancePaymentInvs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sapis)[0]), nil
}

// GetSaleAdvancePaymentInvs gets sale.advance.payment.inv existing records.
func (c *Client) GetSaleAdvancePaymentInvs(ids []int64) (*SaleAdvancePaymentInvs, error) {
	sapis := &SaleAdvancePaymentInvs{}
	if err := c.Read(SaleAdvancePaymentInvModel, ids, nil, sapis); err != nil {
		return nil, err
	}
	return sapis, nil
}

// FindSaleAdvancePaymentInv finds sale.advance.payment.inv record by querying it with criteria.
func (c *Client) FindSaleAdvancePaymentInv(criteria *Criteria) (*SaleAdvancePaymentInv, error) {
	sapis := &SaleAdvancePaymentInvs{}
	if err := c.SearchRead(SaleAdvancePaymentInvModel, criteria, NewOptions().Limit(1), sapis); err != nil {
		return nil, err
	}
	return &((*sapis)[0]), nil
}

// FindSaleAdvancePaymentInvs finds sale.advance.payment.inv records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleAdvancePaymentInvs(criteria *Criteria, options *Options) (*SaleAdvancePaymentInvs, error) {
	sapis := &SaleAdvancePaymentInvs{}
	if err := c.SearchRead(SaleAdvancePaymentInvModel, criteria, options, sapis); err != nil {
		return nil, err
	}
	return sapis, nil
}

// FindSaleAdvancePaymentInvIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleAdvancePaymentInvIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleAdvancePaymentInvModel, criteria, options)
}

// FindSaleAdvancePaymentInvId finds record id by querying it with criteria.
func (c *Client) FindSaleAdvancePaymentInvId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleAdvancePaymentInvModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
