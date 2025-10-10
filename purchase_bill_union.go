package odoo

// PurchaseBillUnion represents purchase.bill.union model.
type PurchaseBillUnion struct {
	Amount          *Float    `xmlrpc:"amount,omitempty"`
	CompanyId       *Many2One `xmlrpc:"company_id,omitempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omitempty"`
	Date            *Time     `xmlrpc:"date,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	PartnerId       *Many2One `xmlrpc:"partner_id,omitempty"`
	PurchaseOrderId *Many2One `xmlrpc:"purchase_order_id,omitempty"`
	Reference       *String   `xmlrpc:"reference,omitempty"`
	VendorBillId    *Many2One `xmlrpc:"vendor_bill_id,omitempty"`
}

// PurchaseBillUnions represents array of purchase.bill.union model.
type PurchaseBillUnions []PurchaseBillUnion

// PurchaseBillUnionModel is the odoo model name.
const PurchaseBillUnionModel = "purchase.bill.union"

// Many2One convert PurchaseBillUnion to *Many2One.
func (pbu *PurchaseBillUnion) Many2One() *Many2One {
	return NewMany2One(pbu.Id.Get(), "")
}

// CreatePurchaseBillUnion creates a new purchase.bill.union model and returns its id.
func (c *Client) CreatePurchaseBillUnion(pbu *PurchaseBillUnion) (int64, error) {
	ids, err := c.CreatePurchaseBillUnions([]*PurchaseBillUnion{pbu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePurchaseBillUnion creates a new purchase.bill.union model and returns its id.
func (c *Client) CreatePurchaseBillUnions(pbus []*PurchaseBillUnion) ([]int64, error) {
	var vv []interface{}
	for _, v := range pbus {
		vv = append(vv, v)
	}
	return c.Create(PurchaseBillUnionModel, vv, nil)
}

// UpdatePurchaseBillUnion updates an existing purchase.bill.union record.
func (c *Client) UpdatePurchaseBillUnion(pbu *PurchaseBillUnion) error {
	return c.UpdatePurchaseBillUnions([]int64{pbu.Id.Get()}, pbu)
}

// UpdatePurchaseBillUnions updates existing purchase.bill.union records.
// All records (represented by ids) will be updated by pbu values.
func (c *Client) UpdatePurchaseBillUnions(ids []int64, pbu *PurchaseBillUnion) error {
	return c.Update(PurchaseBillUnionModel, ids, pbu, nil)
}

// DeletePurchaseBillUnion deletes an existing purchase.bill.union record.
func (c *Client) DeletePurchaseBillUnion(id int64) error {
	return c.DeletePurchaseBillUnions([]int64{id})
}

// DeletePurchaseBillUnions deletes existing purchase.bill.union records.
func (c *Client) DeletePurchaseBillUnions(ids []int64) error {
	return c.Delete(PurchaseBillUnionModel, ids)
}

// GetPurchaseBillUnion gets purchase.bill.union existing record.
func (c *Client) GetPurchaseBillUnion(id int64) (*PurchaseBillUnion, error) {
	pbus, err := c.GetPurchaseBillUnions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pbus)[0]), nil
}

// GetPurchaseBillUnions gets purchase.bill.union existing records.
func (c *Client) GetPurchaseBillUnions(ids []int64) (*PurchaseBillUnions, error) {
	pbus := &PurchaseBillUnions{}
	if err := c.Read(PurchaseBillUnionModel, ids, nil, pbus); err != nil {
		return nil, err
	}
	return pbus, nil
}

// FindPurchaseBillUnion finds purchase.bill.union record by querying it with criteria.
func (c *Client) FindPurchaseBillUnion(criteria *Criteria) (*PurchaseBillUnion, error) {
	pbus := &PurchaseBillUnions{}
	if err := c.SearchRead(PurchaseBillUnionModel, criteria, NewOptions().Limit(1), pbus); err != nil {
		return nil, err
	}
	return &((*pbus)[0]), nil
}

// FindPurchaseBillUnions finds purchase.bill.union records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseBillUnions(criteria *Criteria, options *Options) (*PurchaseBillUnions, error) {
	pbus := &PurchaseBillUnions{}
	if err := c.SearchRead(PurchaseBillUnionModel, criteria, options, pbus); err != nil {
		return nil, err
	}
	return pbus, nil
}

// FindPurchaseBillUnionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseBillUnionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PurchaseBillUnionModel, criteria, options)
}

// FindPurchaseBillUnionId finds record id by querying it with criteria.
func (c *Client) FindPurchaseBillUnionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseBillUnionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
