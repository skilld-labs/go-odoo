package odoo

// IrModuleCategory represents ir.module.category model.
type IrModuleCategory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	ChildIds    *Relation `xmlrpc:"child_ids,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Description *String   `xmlrpc:"description,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Exclusive   *Bool     `xmlrpc:"exclusive,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ModuleIds   *Relation `xmlrpc:"module_ids,omitempty"`
	ModuleNr    *Int      `xmlrpc:"module_nr,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	Visible     *Bool     `xmlrpc:"visible,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
	XmlId       *String   `xmlrpc:"xml_id,omitempty"`
}

// IrModuleCategorys represents array of ir.module.category model.
type IrModuleCategorys []IrModuleCategory

// IrModuleCategoryModel is the odoo model name.
const IrModuleCategoryModel = "ir.module.category"

// Many2One convert IrModuleCategory to *Many2One.
func (imc *IrModuleCategory) Many2One() *Many2One {
	return NewMany2One(imc.Id.Get(), "")
}

// CreateIrModuleCategory creates a new ir.module.category model and returns its id.
func (c *Client) CreateIrModuleCategory(imc *IrModuleCategory) (int64, error) {
	ids, err := c.CreateIrModuleCategorys([]*IrModuleCategory{imc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModuleCategorys creates a new ir.module.category model and returns its id.
func (c *Client) CreateIrModuleCategorys(imcs []*IrModuleCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range imcs {
		vv = append(vv, v)
	}
	return c.Create(IrModuleCategoryModel, vv, nil)
}

// UpdateIrModuleCategory updates an existing ir.module.category record.
func (c *Client) UpdateIrModuleCategory(imc *IrModuleCategory) error {
	return c.UpdateIrModuleCategorys([]int64{imc.Id.Get()}, imc)
}

// UpdateIrModuleCategorys updates existing ir.module.category records.
// All records (represented by ids) will be updated by imc values.
func (c *Client) UpdateIrModuleCategorys(ids []int64, imc *IrModuleCategory) error {
	return c.Update(IrModuleCategoryModel, ids, imc, nil)
}

// DeleteIrModuleCategory deletes an existing ir.module.category record.
func (c *Client) DeleteIrModuleCategory(id int64) error {
	return c.DeleteIrModuleCategorys([]int64{id})
}

// DeleteIrModuleCategorys deletes existing ir.module.category records.
func (c *Client) DeleteIrModuleCategorys(ids []int64) error {
	return c.Delete(IrModuleCategoryModel, ids)
}

// GetIrModuleCategory gets ir.module.category existing record.
func (c *Client) GetIrModuleCategory(id int64) (*IrModuleCategory, error) {
	imcs, err := c.GetIrModuleCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imcs)[0]), nil
}

// GetIrModuleCategorys gets ir.module.category existing records.
func (c *Client) GetIrModuleCategorys(ids []int64) (*IrModuleCategorys, error) {
	imcs := &IrModuleCategorys{}
	if err := c.Read(IrModuleCategoryModel, ids, nil, imcs); err != nil {
		return nil, err
	}
	return imcs, nil
}

// FindIrModuleCategory finds ir.module.category record by querying it with criteria.
func (c *Client) FindIrModuleCategory(criteria *Criteria) (*IrModuleCategory, error) {
	imcs := &IrModuleCategorys{}
	if err := c.SearchRead(IrModuleCategoryModel, criteria, NewOptions().Limit(1), imcs); err != nil {
		return nil, err
	}
	return &((*imcs)[0]), nil
}

// FindIrModuleCategorys finds ir.module.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleCategorys(criteria *Criteria, options *Options) (*IrModuleCategorys, error) {
	imcs := &IrModuleCategorys{}
	if err := c.SearchRead(IrModuleCategoryModel, criteria, options, imcs); err != nil {
		return nil, err
	}
	return imcs, nil
}

// FindIrModuleCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModuleCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrModuleCategoryModel, criteria, options)
}

// FindIrModuleCategoryId finds record id by querying it with criteria.
func (c *Client) FindIrModuleCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModuleCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
