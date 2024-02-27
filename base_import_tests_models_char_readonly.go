package odoo

// BaseImportTestsModelsCharReadonly represents base_import.tests.models.char.readonly model.
type BaseImportTestsModelsCharReadonly struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsCharReadonlys represents array of base_import.tests.models.char.readonly model.
type BaseImportTestsModelsCharReadonlys []BaseImportTestsModelsCharReadonly

// BaseImportTestsModelsCharReadonlyModel is the odoo model name.
const BaseImportTestsModelsCharReadonlyModel = "base_import.tests.models.char.readonly"

// Many2One convert BaseImportTestsModelsCharReadonly to *Many2One.
func (btmcr *BaseImportTestsModelsCharReadonly) Many2One() *Many2One {
	return NewMany2One(btmcr.Id.Get(), "")
}

// CreateBaseImportTestsModelsCharReadonly creates a new base_import.tests.models.char.readonly model and returns its id.
func (c *Client) CreateBaseImportTestsModelsCharReadonly(btmcr *BaseImportTestsModelsCharReadonly) (int64, error) {
	ids, err := c.CreateBaseImportTestsModelsCharReadonlys([]*BaseImportTestsModelsCharReadonly{btmcr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportTestsModelsCharReadonly creates a new base_import.tests.models.char.readonly model and returns its id.
func (c *Client) CreateBaseImportTestsModelsCharReadonlys(btmcrs []*BaseImportTestsModelsCharReadonly) ([]int64, error) {
	var vv []interface{}
	for _, v := range btmcrs {
		vv = append(vv, v)
	}
	return c.Create(BaseImportTestsModelsCharReadonlyModel, vv, nil)
}

// UpdateBaseImportTestsModelsCharReadonly updates an existing base_import.tests.models.char.readonly record.
func (c *Client) UpdateBaseImportTestsModelsCharReadonly(btmcr *BaseImportTestsModelsCharReadonly) error {
	return c.UpdateBaseImportTestsModelsCharReadonlys([]int64{btmcr.Id.Get()}, btmcr)
}

// UpdateBaseImportTestsModelsCharReadonlys updates existing base_import.tests.models.char.readonly records.
// All records (represented by ids) will be updated by btmcr values.
func (c *Client) UpdateBaseImportTestsModelsCharReadonlys(ids []int64, btmcr *BaseImportTestsModelsCharReadonly) error {
	return c.Update(BaseImportTestsModelsCharReadonlyModel, ids, btmcr, nil)
}

// DeleteBaseImportTestsModelsCharReadonly deletes an existing base_import.tests.models.char.readonly record.
func (c *Client) DeleteBaseImportTestsModelsCharReadonly(id int64) error {
	return c.DeleteBaseImportTestsModelsCharReadonlys([]int64{id})
}

// DeleteBaseImportTestsModelsCharReadonlys deletes existing base_import.tests.models.char.readonly records.
func (c *Client) DeleteBaseImportTestsModelsCharReadonlys(ids []int64) error {
	return c.Delete(BaseImportTestsModelsCharReadonlyModel, ids)
}

// GetBaseImportTestsModelsCharReadonly gets base_import.tests.models.char.readonly existing record.
func (c *Client) GetBaseImportTestsModelsCharReadonly(id int64) (*BaseImportTestsModelsCharReadonly, error) {
	btmcrs, err := c.GetBaseImportTestsModelsCharReadonlys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btmcrs)[0]), nil
}

// GetBaseImportTestsModelsCharReadonlys gets base_import.tests.models.char.readonly existing records.
func (c *Client) GetBaseImportTestsModelsCharReadonlys(ids []int64) (*BaseImportTestsModelsCharReadonlys, error) {
	btmcrs := &BaseImportTestsModelsCharReadonlys{}
	if err := c.Read(BaseImportTestsModelsCharReadonlyModel, ids, nil, btmcrs); err != nil {
		return nil, err
	}
	return btmcrs, nil
}

// FindBaseImportTestsModelsCharReadonly finds base_import.tests.models.char.readonly record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharReadonly(criteria *Criteria) (*BaseImportTestsModelsCharReadonly, error) {
	btmcrs := &BaseImportTestsModelsCharReadonlys{}
	if err := c.SearchRead(BaseImportTestsModelsCharReadonlyModel, criteria, NewOptions().Limit(1), btmcrs); err != nil {
		return nil, err
	}
	return &((*btmcrs)[0]), nil
}

// FindBaseImportTestsModelsCharReadonlys finds base_import.tests.models.char.readonly records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharReadonlys(criteria *Criteria, options *Options) (*BaseImportTestsModelsCharReadonlys, error) {
	btmcrs := &BaseImportTestsModelsCharReadonlys{}
	if err := c.SearchRead(BaseImportTestsModelsCharReadonlyModel, criteria, options, btmcrs); err != nil {
		return nil, err
	}
	return btmcrs, nil
}

// FindBaseImportTestsModelsCharReadonlyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharReadonlyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportTestsModelsCharReadonlyModel, criteria, options)
}

// FindBaseImportTestsModelsCharReadonlyId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharReadonlyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharReadonlyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
