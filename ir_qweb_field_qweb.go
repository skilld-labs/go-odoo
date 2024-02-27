package odoo

// IrQwebFieldQweb represents ir.qweb.field.qweb model.
type IrQwebFieldQweb struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldQwebs represents array of ir.qweb.field.qweb model.
type IrQwebFieldQwebs []IrQwebFieldQweb

// IrQwebFieldQwebModel is the odoo model name.
const IrQwebFieldQwebModel = "ir.qweb.field.qweb"

// Many2One convert IrQwebFieldQweb to *Many2One.
func (iqfq *IrQwebFieldQweb) Many2One() *Many2One {
	return NewMany2One(iqfq.Id.Get(), "")
}

// CreateIrQwebFieldQweb creates a new ir.qweb.field.qweb model and returns its id.
func (c *Client) CreateIrQwebFieldQweb(iqfq *IrQwebFieldQweb) (int64, error) {
	ids, err := c.CreateIrQwebFieldQwebs([]*IrQwebFieldQweb{iqfq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldQweb creates a new ir.qweb.field.qweb model and returns its id.
func (c *Client) CreateIrQwebFieldQwebs(iqfqs []*IrQwebFieldQweb) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfqs {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldQwebModel, vv, nil)
}

// UpdateIrQwebFieldQweb updates an existing ir.qweb.field.qweb record.
func (c *Client) UpdateIrQwebFieldQweb(iqfq *IrQwebFieldQweb) error {
	return c.UpdateIrQwebFieldQwebs([]int64{iqfq.Id.Get()}, iqfq)
}

// UpdateIrQwebFieldQwebs updates existing ir.qweb.field.qweb records.
// All records (represented by ids) will be updated by iqfq values.
func (c *Client) UpdateIrQwebFieldQwebs(ids []int64, iqfq *IrQwebFieldQweb) error {
	return c.Update(IrQwebFieldQwebModel, ids, iqfq, nil)
}

// DeleteIrQwebFieldQweb deletes an existing ir.qweb.field.qweb record.
func (c *Client) DeleteIrQwebFieldQweb(id int64) error {
	return c.DeleteIrQwebFieldQwebs([]int64{id})
}

// DeleteIrQwebFieldQwebs deletes existing ir.qweb.field.qweb records.
func (c *Client) DeleteIrQwebFieldQwebs(ids []int64) error {
	return c.Delete(IrQwebFieldQwebModel, ids)
}

// GetIrQwebFieldQweb gets ir.qweb.field.qweb existing record.
func (c *Client) GetIrQwebFieldQweb(id int64) (*IrQwebFieldQweb, error) {
	iqfqs, err := c.GetIrQwebFieldQwebs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqfqs)[0]), nil
}

// GetIrQwebFieldQwebs gets ir.qweb.field.qweb existing records.
func (c *Client) GetIrQwebFieldQwebs(ids []int64) (*IrQwebFieldQwebs, error) {
	iqfqs := &IrQwebFieldQwebs{}
	if err := c.Read(IrQwebFieldQwebModel, ids, nil, iqfqs); err != nil {
		return nil, err
	}
	return iqfqs, nil
}

// FindIrQwebFieldQweb finds ir.qweb.field.qweb record by querying it with criteria.
func (c *Client) FindIrQwebFieldQweb(criteria *Criteria) (*IrQwebFieldQweb, error) {
	iqfqs := &IrQwebFieldQwebs{}
	if err := c.SearchRead(IrQwebFieldQwebModel, criteria, NewOptions().Limit(1), iqfqs); err != nil {
		return nil, err
	}
	return &((*iqfqs)[0]), nil
}

// FindIrQwebFieldQwebs finds ir.qweb.field.qweb records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldQwebs(criteria *Criteria, options *Options) (*IrQwebFieldQwebs, error) {
	iqfqs := &IrQwebFieldQwebs{}
	if err := c.SearchRead(IrQwebFieldQwebModel, criteria, options, iqfqs); err != nil {
		return nil, err
	}
	return iqfqs, nil
}

// FindIrQwebFieldQwebIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldQwebIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebFieldQwebModel, criteria, options)
}

// FindIrQwebFieldQwebId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldQwebId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldQwebModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
