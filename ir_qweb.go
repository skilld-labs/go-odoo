package odoo

// IrQweb represents ir.qweb model.
type IrQweb struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebs represents array of ir.qweb model.
type IrQwebs []IrQweb

// IrQwebModel is the odoo model name.
const IrQwebModel = "ir.qweb"

// Many2One convert IrQweb to *Many2One.
func (iq *IrQweb) Many2One() *Many2One {
	return NewMany2One(iq.Id.Get(), "")
}

// CreateIrQweb creates a new ir.qweb model and returns its id.
func (c *Client) CreateIrQweb(iq *IrQweb) (int64, error) {
	ids, err := c.CreateIrQwebs([]*IrQweb{iq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebs creates a new ir.qweb model and returns its id.
func (c *Client) CreateIrQwebs(iqs []*IrQweb) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqs {
		vv = append(vv, v)
	}
	return c.Create(IrQwebModel, vv, nil)
}

// UpdateIrQweb updates an existing ir.qweb record.
func (c *Client) UpdateIrQweb(iq *IrQweb) error {
	return c.UpdateIrQwebs([]int64{iq.Id.Get()}, iq)
}

// UpdateIrQwebs updates existing ir.qweb records.
// All records (represented by ids) will be updated by iq values.
func (c *Client) UpdateIrQwebs(ids []int64, iq *IrQweb) error {
	return c.Update(IrQwebModel, ids, iq, nil)
}

// DeleteIrQweb deletes an existing ir.qweb record.
func (c *Client) DeleteIrQweb(id int64) error {
	return c.DeleteIrQwebs([]int64{id})
}

// DeleteIrQwebs deletes existing ir.qweb records.
func (c *Client) DeleteIrQwebs(ids []int64) error {
	return c.Delete(IrQwebModel, ids)
}

// GetIrQweb gets ir.qweb existing record.
func (c *Client) GetIrQweb(id int64) (*IrQweb, error) {
	iqs, err := c.GetIrQwebs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqs)[0]), nil
}

// GetIrQwebs gets ir.qweb existing records.
func (c *Client) GetIrQwebs(ids []int64) (*IrQwebs, error) {
	iqs := &IrQwebs{}
	if err := c.Read(IrQwebModel, ids, nil, iqs); err != nil {
		return nil, err
	}
	return iqs, nil
}

// FindIrQweb finds ir.qweb record by querying it with criteria.
func (c *Client) FindIrQweb(criteria *Criteria) (*IrQweb, error) {
	iqs := &IrQwebs{}
	if err := c.SearchRead(IrQwebModel, criteria, NewOptions().Limit(1), iqs); err != nil {
		return nil, err
	}
	return &((*iqs)[0]), nil
}

// FindIrQwebs finds ir.qweb records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebs(criteria *Criteria, options *Options) (*IrQwebs, error) {
	iqs := &IrQwebs{}
	if err := c.SearchRead(IrQwebModel, criteria, options, iqs); err != nil {
		return nil, err
	}
	return iqs, nil
}

// FindIrQwebIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebModel, criteria, options)
}

// FindIrQwebId finds record id by querying it with criteria.
func (c *Client) FindIrQwebId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
