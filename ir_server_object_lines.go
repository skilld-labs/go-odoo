package odoo

// IrServerObjectLines represents ir.server.object.lines model.
type IrServerObjectLines struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Col1        *Many2One  `xmlrpc:"col1,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	ServerId    *Many2One  `xmlrpc:"server_id,omptempty"`
	Type        *Selection `xmlrpc:"type,omptempty"`
	Value       *String    `xmlrpc:"value,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrServerObjectLiness represents array of ir.server.object.lines model.
type IrServerObjectLiness []IrServerObjectLines

// IrServerObjectLinesModel is the odoo model name.
const IrServerObjectLinesModel = "ir.server.object.lines"

// Many2One convert IrServerObjectLines to *Many2One.
func (isol *IrServerObjectLines) Many2One() *Many2One {
	return NewMany2One(isol.Id.Get(), "")
}

// CreateIrServerObjectLines creates a new ir.server.object.lines model and returns its id.
func (c *Client) CreateIrServerObjectLines(isol *IrServerObjectLines) (int64, error) {
	ids, err := c.CreateIrServerObjectLiness([]*IrServerObjectLines{isol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrServerObjectLines creates a new ir.server.object.lines model and returns its id.
func (c *Client) CreateIrServerObjectLiness(isols []*IrServerObjectLines) ([]int64, error) {
	var vv []interface{}
	for _, v := range isols {
		vv = append(vv, v)
	}
	return c.Create(IrServerObjectLinesModel, vv, nil)
}

// UpdateIrServerObjectLines updates an existing ir.server.object.lines record.
func (c *Client) UpdateIrServerObjectLines(isol *IrServerObjectLines) error {
	return c.UpdateIrServerObjectLiness([]int64{isol.Id.Get()}, isol)
}

// UpdateIrServerObjectLiness updates existing ir.server.object.lines records.
// All records (represented by ids) will be updated by isol values.
func (c *Client) UpdateIrServerObjectLiness(ids []int64, isol *IrServerObjectLines) error {
	return c.Update(IrServerObjectLinesModel, ids, isol, nil)
}

// DeleteIrServerObjectLines deletes an existing ir.server.object.lines record.
func (c *Client) DeleteIrServerObjectLines(id int64) error {
	return c.DeleteIrServerObjectLiness([]int64{id})
}

// DeleteIrServerObjectLiness deletes existing ir.server.object.lines records.
func (c *Client) DeleteIrServerObjectLiness(ids []int64) error {
	return c.Delete(IrServerObjectLinesModel, ids)
}

// GetIrServerObjectLines gets ir.server.object.lines existing record.
func (c *Client) GetIrServerObjectLines(id int64) (*IrServerObjectLines, error) {
	isols, err := c.GetIrServerObjectLiness([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*isols)[0]), nil
}

// GetIrServerObjectLiness gets ir.server.object.lines existing records.
func (c *Client) GetIrServerObjectLiness(ids []int64) (*IrServerObjectLiness, error) {
	isols := &IrServerObjectLiness{}
	if err := c.Read(IrServerObjectLinesModel, ids, nil, isols); err != nil {
		return nil, err
	}
	return isols, nil
}

// FindIrServerObjectLines finds ir.server.object.lines record by querying it with criteria.
func (c *Client) FindIrServerObjectLines(criteria *Criteria) (*IrServerObjectLines, error) {
	isols := &IrServerObjectLiness{}
	if err := c.SearchRead(IrServerObjectLinesModel, criteria, NewOptions().Limit(1), isols); err != nil {
		return nil, err
	}
	return &((*isols)[0]), nil
}

// FindIrServerObjectLiness finds ir.server.object.lines records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrServerObjectLiness(criteria *Criteria, options *Options) (*IrServerObjectLiness, error) {
	isols := &IrServerObjectLiness{}
	if err := c.SearchRead(IrServerObjectLinesModel, criteria, options, isols); err != nil {
		return nil, err
	}
	return isols, nil
}

// FindIrServerObjectLinesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrServerObjectLinesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrServerObjectLinesModel, criteria, options)
}

// FindIrServerObjectLinesId finds record id by querying it with criteria.
func (c *Client) FindIrServerObjectLinesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrServerObjectLinesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
