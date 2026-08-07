package odoo

// IrBinary represents ir.binary model.
type IrBinary struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrBinarys represents array of ir.binary model.
type IrBinarys []IrBinary

// IrBinaryModel is the odoo model name.
const IrBinaryModel = "ir.binary"

// Many2One convert IrBinary to *Many2One.
func (ib *IrBinary) Many2One() *Many2One {
	return NewMany2One(ib.Id.Get(), "")
}

// CreateIrBinary creates a new ir.binary model and returns its id.
func (c *Client) CreateIrBinary(ib *IrBinary) (int64, error) {
	ids, err := c.CreateIrBinarys([]*IrBinary{ib})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrBinary creates a new ir.binary model and returns its id.
func (c *Client) CreateIrBinarys(ibs []*IrBinary) ([]int64, error) {
	var vv []interface{}
	for _, v := range ibs {
		vv = append(vv, v)
	}
	return c.Create(IrBinaryModel, vv, nil)
}

// UpdateIrBinary updates an existing ir.binary record.
func (c *Client) UpdateIrBinary(ib *IrBinary) error {
	return c.UpdateIrBinarys([]int64{ib.Id.Get()}, ib)
}

// UpdateIrBinarys updates existing ir.binary records.
// All records (represented by ids) will be updated by ib values.
func (c *Client) UpdateIrBinarys(ids []int64, ib *IrBinary) error {
	return c.Update(IrBinaryModel, ids, ib, nil)
}

// DeleteIrBinary deletes an existing ir.binary record.
func (c *Client) DeleteIrBinary(id int64) error {
	return c.DeleteIrBinarys([]int64{id})
}

// DeleteIrBinarys deletes existing ir.binary records.
func (c *Client) DeleteIrBinarys(ids []int64) error {
	return c.Delete(IrBinaryModel, ids)
}

// GetIrBinary gets ir.binary existing record.
func (c *Client) GetIrBinary(id int64) (*IrBinary, error) {
	ibs, err := c.GetIrBinarys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ibs)[0]), nil
}

// GetIrBinarys gets ir.binary existing records.
func (c *Client) GetIrBinarys(ids []int64) (*IrBinarys, error) {
	ibs := &IrBinarys{}
	if err := c.Read(IrBinaryModel, ids, nil, ibs); err != nil {
		return nil, err
	}
	return ibs, nil
}

// FindIrBinary finds ir.binary record by querying it with criteria.
func (c *Client) FindIrBinary(criteria *Criteria) (*IrBinary, error) {
	ibs := &IrBinarys{}
	if err := c.SearchRead(IrBinaryModel, criteria, NewOptions().Limit(1), ibs); err != nil {
		return nil, err
	}
	return &((*ibs)[0]), nil
}

// FindIrBinarys finds ir.binary records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrBinarys(criteria *Criteria, options *Options) (*IrBinarys, error) {
	ibs := &IrBinarys{}
	if err := c.SearchRead(IrBinaryModel, criteria, options, ibs); err != nil {
		return nil, err
	}
	return ibs, nil
}

// FindIrBinaryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrBinaryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrBinaryModel, criteria, options)
}

// FindIrBinaryId finds record id by querying it with criteria.
func (c *Client) FindIrBinaryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrBinaryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
