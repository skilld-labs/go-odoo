package odoo

// SaleMassCancelOrders represents sale.mass.cancel.orders model.
type SaleMassCancelOrders struct {
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	HasConfirmedOrder *Bool     `xmlrpc:"has_confirmed_order,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	SaleOrderIds      *Relation `xmlrpc:"sale_order_ids,omitempty"`
	SaleOrdersCount   *Int      `xmlrpc:"sale_orders_count,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaleMassCancelOrderss represents array of sale.mass.cancel.orders model.
type SaleMassCancelOrderss []SaleMassCancelOrders

// SaleMassCancelOrdersModel is the odoo model name.
const SaleMassCancelOrdersModel = "sale.mass.cancel.orders"

// Many2One convert SaleMassCancelOrders to *Many2One.
func (smco *SaleMassCancelOrders) Many2One() *Many2One {
	return NewMany2One(smco.Id.Get(), "")
}

// CreateSaleMassCancelOrders creates a new sale.mass.cancel.orders model and returns its id.
func (c *Client) CreateSaleMassCancelOrders(smco *SaleMassCancelOrders) (int64, error) {
	ids, err := c.CreateSaleMassCancelOrderss([]*SaleMassCancelOrders{smco})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleMassCancelOrders creates a new sale.mass.cancel.orders model and returns its id.
func (c *Client) CreateSaleMassCancelOrderss(smcos []*SaleMassCancelOrders) ([]int64, error) {
	var vv []interface{}
	for _, v := range smcos {
		vv = append(vv, v)
	}
	return c.Create(SaleMassCancelOrdersModel, vv, nil)
}

// UpdateSaleMassCancelOrders updates an existing sale.mass.cancel.orders record.
func (c *Client) UpdateSaleMassCancelOrders(smco *SaleMassCancelOrders) error {
	return c.UpdateSaleMassCancelOrderss([]int64{smco.Id.Get()}, smco)
}

// UpdateSaleMassCancelOrderss updates existing sale.mass.cancel.orders records.
// All records (represented by ids) will be updated by smco values.
func (c *Client) UpdateSaleMassCancelOrderss(ids []int64, smco *SaleMassCancelOrders) error {
	return c.Update(SaleMassCancelOrdersModel, ids, smco, nil)
}

// DeleteSaleMassCancelOrders deletes an existing sale.mass.cancel.orders record.
func (c *Client) DeleteSaleMassCancelOrders(id int64) error {
	return c.DeleteSaleMassCancelOrderss([]int64{id})
}

// DeleteSaleMassCancelOrderss deletes existing sale.mass.cancel.orders records.
func (c *Client) DeleteSaleMassCancelOrderss(ids []int64) error {
	return c.Delete(SaleMassCancelOrdersModel, ids)
}

// GetSaleMassCancelOrders gets sale.mass.cancel.orders existing record.
func (c *Client) GetSaleMassCancelOrders(id int64) (*SaleMassCancelOrders, error) {
	smcos, err := c.GetSaleMassCancelOrderss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*smcos)[0]), nil
}

// GetSaleMassCancelOrderss gets sale.mass.cancel.orders existing records.
func (c *Client) GetSaleMassCancelOrderss(ids []int64) (*SaleMassCancelOrderss, error) {
	smcos := &SaleMassCancelOrderss{}
	if err := c.Read(SaleMassCancelOrdersModel, ids, nil, smcos); err != nil {
		return nil, err
	}
	return smcos, nil
}

// FindSaleMassCancelOrders finds sale.mass.cancel.orders record by querying it with criteria.
func (c *Client) FindSaleMassCancelOrders(criteria *Criteria) (*SaleMassCancelOrders, error) {
	smcos := &SaleMassCancelOrderss{}
	if err := c.SearchRead(SaleMassCancelOrdersModel, criteria, NewOptions().Limit(1), smcos); err != nil {
		return nil, err
	}
	return &((*smcos)[0]), nil
}

// FindSaleMassCancelOrderss finds sale.mass.cancel.orders records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleMassCancelOrderss(criteria *Criteria, options *Options) (*SaleMassCancelOrderss, error) {
	smcos := &SaleMassCancelOrderss{}
	if err := c.SearchRead(SaleMassCancelOrdersModel, criteria, options, smcos); err != nil {
		return nil, err
	}
	return smcos, nil
}

// FindSaleMassCancelOrdersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleMassCancelOrdersIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleMassCancelOrdersModel, criteria, options)
}

// FindSaleMassCancelOrdersId finds record id by querying it with criteria.
func (c *Client) FindSaleMassCancelOrdersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleMassCancelOrdersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
