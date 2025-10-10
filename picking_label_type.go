package odoo

// PickingLabelType represents picking.label.type model.
type PickingLabelType struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	LabelType   *Selection `xmlrpc:"label_type,omitempty"`
	PickingIds  *Relation  `xmlrpc:"picking_ids,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PickingLabelTypes represents array of picking.label.type model.
type PickingLabelTypes []PickingLabelType

// PickingLabelTypeModel is the odoo model name.
const PickingLabelTypeModel = "picking.label.type"

// Many2One convert PickingLabelType to *Many2One.
func (plt *PickingLabelType) Many2One() *Many2One {
	return NewMany2One(plt.Id.Get(), "")
}

// CreatePickingLabelType creates a new picking.label.type model and returns its id.
func (c *Client) CreatePickingLabelType(plt *PickingLabelType) (int64, error) {
	ids, err := c.CreatePickingLabelTypes([]*PickingLabelType{plt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePickingLabelType creates a new picking.label.type model and returns its id.
func (c *Client) CreatePickingLabelTypes(plts []*PickingLabelType) ([]int64, error) {
	var vv []interface{}
	for _, v := range plts {
		vv = append(vv, v)
	}
	return c.Create(PickingLabelTypeModel, vv, nil)
}

// UpdatePickingLabelType updates an existing picking.label.type record.
func (c *Client) UpdatePickingLabelType(plt *PickingLabelType) error {
	return c.UpdatePickingLabelTypes([]int64{plt.Id.Get()}, plt)
}

// UpdatePickingLabelTypes updates existing picking.label.type records.
// All records (represented by ids) will be updated by plt values.
func (c *Client) UpdatePickingLabelTypes(ids []int64, plt *PickingLabelType) error {
	return c.Update(PickingLabelTypeModel, ids, plt, nil)
}

// DeletePickingLabelType deletes an existing picking.label.type record.
func (c *Client) DeletePickingLabelType(id int64) error {
	return c.DeletePickingLabelTypes([]int64{id})
}

// DeletePickingLabelTypes deletes existing picking.label.type records.
func (c *Client) DeletePickingLabelTypes(ids []int64) error {
	return c.Delete(PickingLabelTypeModel, ids)
}

// GetPickingLabelType gets picking.label.type existing record.
func (c *Client) GetPickingLabelType(id int64) (*PickingLabelType, error) {
	plts, err := c.GetPickingLabelTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*plts)[0]), nil
}

// GetPickingLabelTypes gets picking.label.type existing records.
func (c *Client) GetPickingLabelTypes(ids []int64) (*PickingLabelTypes, error) {
	plts := &PickingLabelTypes{}
	if err := c.Read(PickingLabelTypeModel, ids, nil, plts); err != nil {
		return nil, err
	}
	return plts, nil
}

// FindPickingLabelType finds picking.label.type record by querying it with criteria.
func (c *Client) FindPickingLabelType(criteria *Criteria) (*PickingLabelType, error) {
	plts := &PickingLabelTypes{}
	if err := c.SearchRead(PickingLabelTypeModel, criteria, NewOptions().Limit(1), plts); err != nil {
		return nil, err
	}
	return &((*plts)[0]), nil
}

// FindPickingLabelTypes finds picking.label.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPickingLabelTypes(criteria *Criteria, options *Options) (*PickingLabelTypes, error) {
	plts := &PickingLabelTypes{}
	if err := c.SearchRead(PickingLabelTypeModel, criteria, options, plts); err != nil {
		return nil, err
	}
	return plts, nil
}

// FindPickingLabelTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPickingLabelTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PickingLabelTypeModel, criteria, options)
}

// FindPickingLabelTypeId finds record id by querying it with criteria.
func (c *Client) FindPickingLabelTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PickingLabelTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
