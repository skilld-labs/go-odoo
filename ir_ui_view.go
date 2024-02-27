package odoo

// IrUiView represents ir.ui.view model.
type IrUiView struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	Active             *Bool      `xmlrpc:"active,omitempty"`
	Arch               *String    `xmlrpc:"arch,omitempty"`
	ArchBase           *String    `xmlrpc:"arch_base,omitempty"`
	ArchDb             *String    `xmlrpc:"arch_db,omitempty"`
	ArchFs             *String    `xmlrpc:"arch_fs,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	FieldParent        *String    `xmlrpc:"field_parent,omitempty"`
	GroupsId           *Relation  `xmlrpc:"groups_id,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	InheritChildrenIds *Relation  `xmlrpc:"inherit_children_ids,omitempty"`
	InheritId          *Many2One  `xmlrpc:"inherit_id,omitempty"`
	Key                *String    `xmlrpc:"key,omitempty"`
	Mode               *Selection `xmlrpc:"mode,omitempty"`
	Model              *String    `xmlrpc:"model,omitempty"`
	ModelDataId        *Many2One  `xmlrpc:"model_data_id,omitempty"`
	ModelIds           *Relation  `xmlrpc:"model_ids,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	Priority           *Int       `xmlrpc:"priority,omitempty"`
	Type               *Selection `xmlrpc:"type,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId              *String    `xmlrpc:"xml_id,omitempty"`
}

// IrUiViews represents array of ir.ui.view model.
type IrUiViews []IrUiView

// IrUiViewModel is the odoo model name.
const IrUiViewModel = "ir.ui.view"

// Many2One convert IrUiView to *Many2One.
func (iuv *IrUiView) Many2One() *Many2One {
	return NewMany2One(iuv.Id.Get(), "")
}

// CreateIrUiView creates a new ir.ui.view model and returns its id.
func (c *Client) CreateIrUiView(iuv *IrUiView) (int64, error) {
	ids, err := c.CreateIrUiViews([]*IrUiView{iuv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrUiViews creates a new ir.ui.view model and returns its id.
func (c *Client) CreateIrUiViews(iuvs []*IrUiView) ([]int64, error) {
	var vv []interface{}
	for _, v := range iuvs {
		vv = append(vv, v)
	}
	return c.Create(IrUiViewModel, vv, nil)
}

// UpdateIrUiView updates an existing ir.ui.view record.
func (c *Client) UpdateIrUiView(iuv *IrUiView) error {
	return c.UpdateIrUiViews([]int64{iuv.Id.Get()}, iuv)
}

// UpdateIrUiViews updates existing ir.ui.view records.
// All records (represented by ids) will be updated by iuv values.
func (c *Client) UpdateIrUiViews(ids []int64, iuv *IrUiView) error {
	return c.Update(IrUiViewModel, ids, iuv, nil)
}

// DeleteIrUiView deletes an existing ir.ui.view record.
func (c *Client) DeleteIrUiView(id int64) error {
	return c.DeleteIrUiViews([]int64{id})
}

// DeleteIrUiViews deletes existing ir.ui.view records.
func (c *Client) DeleteIrUiViews(ids []int64) error {
	return c.Delete(IrUiViewModel, ids)
}

// GetIrUiView gets ir.ui.view existing record.
func (c *Client) GetIrUiView(id int64) (*IrUiView, error) {
	iuvs, err := c.GetIrUiViews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iuvs)[0]), nil
}

// GetIrUiViews gets ir.ui.view existing records.
func (c *Client) GetIrUiViews(ids []int64) (*IrUiViews, error) {
	iuvs := &IrUiViews{}
	if err := c.Read(IrUiViewModel, ids, nil, iuvs); err != nil {
		return nil, err
	}
	return iuvs, nil
}

// FindIrUiView finds ir.ui.view record by querying it with criteria.
func (c *Client) FindIrUiView(criteria *Criteria) (*IrUiView, error) {
	iuvs := &IrUiViews{}
	if err := c.SearchRead(IrUiViewModel, criteria, NewOptions().Limit(1), iuvs); err != nil {
		return nil, err
	}
	return &((*iuvs)[0]), nil
}

// FindIrUiViews finds ir.ui.view records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiViews(criteria *Criteria, options *Options) (*IrUiViews, error) {
	iuvs := &IrUiViews{}
	if err := c.SearchRead(IrUiViewModel, criteria, options, iuvs); err != nil {
		return nil, err
	}
	return iuvs, nil
}

// FindIrUiViewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiViewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrUiViewModel, criteria, options)
}

// FindIrUiViewId finds record id by querying it with criteria.
func (c *Client) FindIrUiViewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrUiViewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
