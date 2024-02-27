package odoo

// IrDefault represents ir.default model.
type IrDefault struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	Condition   *String   `xmlrpc:"condition,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	FieldId     *Many2One `xmlrpc:"field_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	JsonValue   *String   `xmlrpc:"json_value,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrDefaults represents array of ir.default model.
type IrDefaults []IrDefault

// IrDefaultModel is the odoo model name.
const IrDefaultModel = "ir.default"

// Many2One convert IrDefault to *Many2One.
func (ID *IrDefault) Many2One() *Many2One {
	return NewMany2One(ID.Id.Get(), "")
}

// CreateIrDefault creates a new ir.default model and returns its id.
func (c *Client) CreateIrDefault(ID *IrDefault) (int64, error) {
	ids, err := c.CreateIrDefaults([]*IrDefault{ID})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrDefault creates a new ir.default model and returns its id.
func (c *Client) CreateIrDefaults(IDs []*IrDefault) ([]int64, error) {
	var vv []interface{}
	for _, v := range IDs {
		vv = append(vv, v)
	}
	return c.Create(IrDefaultModel, vv, nil)
}

// UpdateIrDefault updates an existing ir.default record.
func (c *Client) UpdateIrDefault(ID *IrDefault) error {
	return c.UpdateIrDefaults([]int64{ID.Id.Get()}, ID)
}

// UpdateIrDefaults updates existing ir.default records.
// All records (represented by ids) will be updated by ID values.
func (c *Client) UpdateIrDefaults(ids []int64, ID *IrDefault) error {
	return c.Update(IrDefaultModel, ids, ID, nil)
}

// DeleteIrDefault deletes an existing ir.default record.
func (c *Client) DeleteIrDefault(id int64) error {
	return c.DeleteIrDefaults([]int64{id})
}

// DeleteIrDefaults deletes existing ir.default records.
func (c *Client) DeleteIrDefaults(ids []int64) error {
	return c.Delete(IrDefaultModel, ids)
}

// GetIrDefault gets ir.default existing record.
func (c *Client) GetIrDefault(id int64) (*IrDefault, error) {
	IDs, err := c.GetIrDefaults([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*IDs)[0]), nil
}

// GetIrDefaults gets ir.default existing records.
func (c *Client) GetIrDefaults(ids []int64) (*IrDefaults, error) {
	IDs := &IrDefaults{}
	if err := c.Read(IrDefaultModel, ids, nil, IDs); err != nil {
		return nil, err
	}
	return IDs, nil
}

// FindIrDefault finds ir.default record by querying it with criteria.
func (c *Client) FindIrDefault(criteria *Criteria) (*IrDefault, error) {
	IDs := &IrDefaults{}
	if err := c.SearchRead(IrDefaultModel, criteria, NewOptions().Limit(1), IDs); err != nil {
		return nil, err
	}
	return &((*IDs)[0]), nil
}

// FindIrDefaults finds ir.default records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDefaults(criteria *Criteria, options *Options) (*IrDefaults, error) {
	IDs := &IrDefaults{}
	if err := c.SearchRead(IrDefaultModel, criteria, options, IDs); err != nil {
		return nil, err
	}
	return IDs, nil
}

// FindIrDefaultIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDefaultIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrDefaultModel, criteria, options)
}

// FindIrDefaultId finds record id by querying it with criteria.
func (c *Client) FindIrDefaultId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrDefaultModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
