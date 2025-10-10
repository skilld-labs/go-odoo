package odoo

// SaleOrderDiscount represents sale.order.discount model.
type SaleOrderDiscount struct {
	CompanyId          *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId         *Many2One  `xmlrpc:"currency_id,omitempty"`
	DiscountAmount     *Float     `xmlrpc:"discount_amount,omitempty"`
	DiscountPercentage *Float     `xmlrpc:"discount_percentage,omitempty"`
	DiscountType       *Selection `xmlrpc:"discount_type,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	SaleOrderId        *Many2One  `xmlrpc:"sale_order_id,omitempty"`
	TaxIds             *Relation  `xmlrpc:"tax_ids,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SaleOrderDiscounts represents array of sale.order.discount model.
type SaleOrderDiscounts []SaleOrderDiscount

// SaleOrderDiscountModel is the odoo model name.
const SaleOrderDiscountModel = "sale.order.discount"

// Many2One convert SaleOrderDiscount to *Many2One.
func (sod *SaleOrderDiscount) Many2One() *Many2One {
	return NewMany2One(sod.Id.Get(), "")
}

// CreateSaleOrderDiscount creates a new sale.order.discount model and returns its id.
func (c *Client) CreateSaleOrderDiscount(sod *SaleOrderDiscount) (int64, error) {
	ids, err := c.CreateSaleOrderDiscounts([]*SaleOrderDiscount{sod})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderDiscount creates a new sale.order.discount model and returns its id.
func (c *Client) CreateSaleOrderDiscounts(sods []*SaleOrderDiscount) ([]int64, error) {
	var vv []interface{}
	for _, v := range sods {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderDiscountModel, vv, nil)
}

// UpdateSaleOrderDiscount updates an existing sale.order.discount record.
func (c *Client) UpdateSaleOrderDiscount(sod *SaleOrderDiscount) error {
	return c.UpdateSaleOrderDiscounts([]int64{sod.Id.Get()}, sod)
}

// UpdateSaleOrderDiscounts updates existing sale.order.discount records.
// All records (represented by ids) will be updated by sod values.
func (c *Client) UpdateSaleOrderDiscounts(ids []int64, sod *SaleOrderDiscount) error {
	return c.Update(SaleOrderDiscountModel, ids, sod, nil)
}

// DeleteSaleOrderDiscount deletes an existing sale.order.discount record.
func (c *Client) DeleteSaleOrderDiscount(id int64) error {
	return c.DeleteSaleOrderDiscounts([]int64{id})
}

// DeleteSaleOrderDiscounts deletes existing sale.order.discount records.
func (c *Client) DeleteSaleOrderDiscounts(ids []int64) error {
	return c.Delete(SaleOrderDiscountModel, ids)
}

// GetSaleOrderDiscount gets sale.order.discount existing record.
func (c *Client) GetSaleOrderDiscount(id int64) (*SaleOrderDiscount, error) {
	sods, err := c.GetSaleOrderDiscounts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sods)[0]), nil
}

// GetSaleOrderDiscounts gets sale.order.discount existing records.
func (c *Client) GetSaleOrderDiscounts(ids []int64) (*SaleOrderDiscounts, error) {
	sods := &SaleOrderDiscounts{}
	if err := c.Read(SaleOrderDiscountModel, ids, nil, sods); err != nil {
		return nil, err
	}
	return sods, nil
}

// FindSaleOrderDiscount finds sale.order.discount record by querying it with criteria.
func (c *Client) FindSaleOrderDiscount(criteria *Criteria) (*SaleOrderDiscount, error) {
	sods := &SaleOrderDiscounts{}
	if err := c.SearchRead(SaleOrderDiscountModel, criteria, NewOptions().Limit(1), sods); err != nil {
		return nil, err
	}
	return &((*sods)[0]), nil
}

// FindSaleOrderDiscounts finds sale.order.discount records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderDiscounts(criteria *Criteria, options *Options) (*SaleOrderDiscounts, error) {
	sods := &SaleOrderDiscounts{}
	if err := c.SearchRead(SaleOrderDiscountModel, criteria, options, sods); err != nil {
		return nil, err
	}
	return sods, nil
}

// FindSaleOrderDiscountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderDiscountIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleOrderDiscountModel, criteria, options)
}

// FindSaleOrderDiscountId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderDiscountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderDiscountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
