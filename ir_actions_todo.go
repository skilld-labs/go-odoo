package odoo

// IrActionsTodo represents ir.actions.todo model.
type IrActionsTodo struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	ActionId    *Many2One  `xmlrpc:"action_id,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	Sequence    *Int       `xmlrpc:"sequence,omitempty"`
	State       *Selection `xmlrpc:"state,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrActionsTodos represents array of ir.actions.todo model.
type IrActionsTodos []IrActionsTodo

// IrActionsTodoModel is the odoo model name.
const IrActionsTodoModel = "ir.actions.todo"

// Many2One convert IrActionsTodo to *Many2One.
func (iat *IrActionsTodo) Many2One() *Many2One {
	return NewMany2One(iat.Id.Get(), "")
}

// CreateIrActionsTodo creates a new ir.actions.todo model and returns its id.
func (c *Client) CreateIrActionsTodo(iat *IrActionsTodo) (int64, error) {
	ids, err := c.CreateIrActionsTodos([]*IrActionsTodo{iat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrActionsTodo creates a new ir.actions.todo model and returns its id.
func (c *Client) CreateIrActionsTodos(iats []*IrActionsTodo) ([]int64, error) {
	var vv []interface{}
	for _, v := range iats {
		vv = append(vv, v)
	}
	return c.Create(IrActionsTodoModel, vv, nil)
}

// UpdateIrActionsTodo updates an existing ir.actions.todo record.
func (c *Client) UpdateIrActionsTodo(iat *IrActionsTodo) error {
	return c.UpdateIrActionsTodos([]int64{iat.Id.Get()}, iat)
}

// UpdateIrActionsTodos updates existing ir.actions.todo records.
// All records (represented by ids) will be updated by iat values.
func (c *Client) UpdateIrActionsTodos(ids []int64, iat *IrActionsTodo) error {
	return c.Update(IrActionsTodoModel, ids, iat, nil)
}

// DeleteIrActionsTodo deletes an existing ir.actions.todo record.
func (c *Client) DeleteIrActionsTodo(id int64) error {
	return c.DeleteIrActionsTodos([]int64{id})
}

// DeleteIrActionsTodos deletes existing ir.actions.todo records.
func (c *Client) DeleteIrActionsTodos(ids []int64) error {
	return c.Delete(IrActionsTodoModel, ids)
}

// GetIrActionsTodo gets ir.actions.todo existing record.
func (c *Client) GetIrActionsTodo(id int64) (*IrActionsTodo, error) {
	iats, err := c.GetIrActionsTodos([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iats)[0]), nil
}

// GetIrActionsTodos gets ir.actions.todo existing records.
func (c *Client) GetIrActionsTodos(ids []int64) (*IrActionsTodos, error) {
	iats := &IrActionsTodos{}
	if err := c.Read(IrActionsTodoModel, ids, nil, iats); err != nil {
		return nil, err
	}
	return iats, nil
}

// FindIrActionsTodo finds ir.actions.todo record by querying it with criteria.
func (c *Client) FindIrActionsTodo(criteria *Criteria) (*IrActionsTodo, error) {
	iats := &IrActionsTodos{}
	if err := c.SearchRead(IrActionsTodoModel, criteria, NewOptions().Limit(1), iats); err != nil {
		return nil, err
	}
	return &((*iats)[0]), nil
}

// FindIrActionsTodos finds ir.actions.todo records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsTodos(criteria *Criteria, options *Options) (*IrActionsTodos, error) {
	iats := &IrActionsTodos{}
	if err := c.SearchRead(IrActionsTodoModel, criteria, options, iats); err != nil {
		return nil, err
	}
	return iats, nil
}

// FindIrActionsTodoIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsTodoIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrActionsTodoModel, criteria, options)
}

// FindIrActionsTodoId finds record id by querying it with criteria.
func (c *Client) FindIrActionsTodoId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsTodoModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
