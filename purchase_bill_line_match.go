package odoo

// PurchaseBillLineMatch represents purchase.bill.line.match model.
type PurchaseBillLineMatch struct {
	AccountMoveId         *Many2One `xmlrpc:"account_move_id,omitempty"`
	AmlId                 *Many2One `xmlrpc:"aml_id,omitempty"`
	BilledAmountUntaxed   *Float    `xmlrpc:"billed_amount_untaxed,omitempty"`
	CompanyId             *Many2One `xmlrpc:"company_id,omitempty"`
	CurrencyId            *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	LineAmountUntaxed     *Float    `xmlrpc:"line_amount_untaxed,omitempty"`
	LineQty               *Float    `xmlrpc:"line_qty,omitempty"`
	LineUomId             *Many2One `xmlrpc:"line_uom_id,omitempty"`
	PartnerId             *Many2One `xmlrpc:"partner_id,omitempty"`
	PolId                 *Many2One `xmlrpc:"pol_id,omitempty"`
	ProductId             *Many2One `xmlrpc:"product_id,omitempty"`
	ProductUomId          *Many2One `xmlrpc:"product_uom_id,omitempty"`
	ProductUomPrice       *Float    `xmlrpc:"product_uom_price,omitempty"`
	ProductUomQty         *Float    `xmlrpc:"product_uom_qty,omitempty"`
	PurchaseAmountUntaxed *Float    `xmlrpc:"purchase_amount_untaxed,omitempty"`
	PurchaseOrderId       *Many2One `xmlrpc:"purchase_order_id,omitempty"`
	QtyInvoiced           *Float    `xmlrpc:"qty_invoiced,omitempty"`
	Reference             *String   `xmlrpc:"reference,omitempty"`
	State                 *String   `xmlrpc:"state,omitempty"`
}

// PurchaseBillLineMatchs represents array of purchase.bill.line.match model.
type PurchaseBillLineMatchs []PurchaseBillLineMatch

// PurchaseBillLineMatchModel is the odoo model name.
const PurchaseBillLineMatchModel = "purchase.bill.line.match"

// Many2One convert PurchaseBillLineMatch to *Many2One.
func (pblm *PurchaseBillLineMatch) Many2One() *Many2One {
	return NewMany2One(pblm.Id.Get(), "")
}

// CreatePurchaseBillLineMatch creates a new purchase.bill.line.match model and returns its id.
func (c *Client) CreatePurchaseBillLineMatch(pblm *PurchaseBillLineMatch) (int64, error) {
	ids, err := c.CreatePurchaseBillLineMatchs([]*PurchaseBillLineMatch{pblm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePurchaseBillLineMatch creates a new purchase.bill.line.match model and returns its id.
func (c *Client) CreatePurchaseBillLineMatchs(pblms []*PurchaseBillLineMatch) ([]int64, error) {
	var vv []interface{}
	for _, v := range pblms {
		vv = append(vv, v)
	}
	return c.Create(PurchaseBillLineMatchModel, vv, nil)
}

// UpdatePurchaseBillLineMatch updates an existing purchase.bill.line.match record.
func (c *Client) UpdatePurchaseBillLineMatch(pblm *PurchaseBillLineMatch) error {
	return c.UpdatePurchaseBillLineMatchs([]int64{pblm.Id.Get()}, pblm)
}

// UpdatePurchaseBillLineMatchs updates existing purchase.bill.line.match records.
// All records (represented by ids) will be updated by pblm values.
func (c *Client) UpdatePurchaseBillLineMatchs(ids []int64, pblm *PurchaseBillLineMatch) error {
	return c.Update(PurchaseBillLineMatchModel, ids, pblm, nil)
}

// DeletePurchaseBillLineMatch deletes an existing purchase.bill.line.match record.
func (c *Client) DeletePurchaseBillLineMatch(id int64) error {
	return c.DeletePurchaseBillLineMatchs([]int64{id})
}

// DeletePurchaseBillLineMatchs deletes existing purchase.bill.line.match records.
func (c *Client) DeletePurchaseBillLineMatchs(ids []int64) error {
	return c.Delete(PurchaseBillLineMatchModel, ids)
}

// GetPurchaseBillLineMatch gets purchase.bill.line.match existing record.
func (c *Client) GetPurchaseBillLineMatch(id int64) (*PurchaseBillLineMatch, error) {
	pblms, err := c.GetPurchaseBillLineMatchs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pblms)[0]), nil
}

// GetPurchaseBillLineMatchs gets purchase.bill.line.match existing records.
func (c *Client) GetPurchaseBillLineMatchs(ids []int64) (*PurchaseBillLineMatchs, error) {
	pblms := &PurchaseBillLineMatchs{}
	if err := c.Read(PurchaseBillLineMatchModel, ids, nil, pblms); err != nil {
		return nil, err
	}
	return pblms, nil
}

// FindPurchaseBillLineMatch finds purchase.bill.line.match record by querying it with criteria.
func (c *Client) FindPurchaseBillLineMatch(criteria *Criteria) (*PurchaseBillLineMatch, error) {
	pblms := &PurchaseBillLineMatchs{}
	if err := c.SearchRead(PurchaseBillLineMatchModel, criteria, NewOptions().Limit(1), pblms); err != nil {
		return nil, err
	}
	return &((*pblms)[0]), nil
}

// FindPurchaseBillLineMatchs finds purchase.bill.line.match records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseBillLineMatchs(criteria *Criteria, options *Options) (*PurchaseBillLineMatchs, error) {
	pblms := &PurchaseBillLineMatchs{}
	if err := c.SearchRead(PurchaseBillLineMatchModel, criteria, options, pblms); err != nil {
		return nil, err
	}
	return pblms, nil
}

// FindPurchaseBillLineMatchIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseBillLineMatchIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PurchaseBillLineMatchModel, criteria, options)
}

// FindPurchaseBillLineMatchId finds record id by querying it with criteria.
func (c *Client) FindPurchaseBillLineMatchId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseBillLineMatchModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
