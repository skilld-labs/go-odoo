package odoo

// BaseImportTestsModelsM2O represents base_import.tests.models.m2o model.
type BaseImportTestsModelsM2O struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *Many2One `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsM2Os represents array of base_import.tests.models.m2o model.
type BaseImportTestsModelsM2Os []BaseImportTestsModelsM2O

// BaseImportTestsModelsM2OModel is the odoo model name.
const BaseImportTestsModelsM2OModel = "base_import.tests.models.m2o"

// Many2One convert BaseImportTestsModelsM2O to *Many2One.
func (btmm *BaseImportTestsModelsM2O) Many2One() *Many2One {
	return NewMany2One(btmm.Id.Get(), "")
}

// CreateBaseImportTestsModelsM2O creates a new base_import.tests.models.m2o model and returns its id.
func (c *Client) CreateBaseImportTestsModelsM2O(btmm *BaseImportTestsModelsM2O) (int64, error) {
	ids, err := c.CreateBaseImportTestsModelsM2Os([]*BaseImportTestsModelsM2O{btmm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportTestsModelsM2O creates a new base_import.tests.models.m2o model and returns its id.
func (c *Client) CreateBaseImportTestsModelsM2Os(btmms []*BaseImportTestsModelsM2O) ([]int64, error) {
	var vv []interface{}
	for _, v := range btmms {
		vv = append(vv, v)
	}
	return c.Create(BaseImportTestsModelsM2OModel, vv, nil)
}

// UpdateBaseImportTestsModelsM2O updates an existing base_import.tests.models.m2o record.
func (c *Client) UpdateBaseImportTestsModelsM2O(btmm *BaseImportTestsModelsM2O) error {
	return c.UpdateBaseImportTestsModelsM2Os([]int64{btmm.Id.Get()}, btmm)
}

// UpdateBaseImportTestsModelsM2Os updates existing base_import.tests.models.m2o records.
// All records (represented by ids) will be updated by btmm values.
func (c *Client) UpdateBaseImportTestsModelsM2Os(ids []int64, btmm *BaseImportTestsModelsM2O) error {
	return c.Update(BaseImportTestsModelsM2OModel, ids, btmm, nil)
}

// DeleteBaseImportTestsModelsM2O deletes an existing base_import.tests.models.m2o record.
func (c *Client) DeleteBaseImportTestsModelsM2O(id int64) error {
	return c.DeleteBaseImportTestsModelsM2Os([]int64{id})
}

// DeleteBaseImportTestsModelsM2Os deletes existing base_import.tests.models.m2o records.
func (c *Client) DeleteBaseImportTestsModelsM2Os(ids []int64) error {
	return c.Delete(BaseImportTestsModelsM2OModel, ids)
}

// GetBaseImportTestsModelsM2O gets base_import.tests.models.m2o existing record.
func (c *Client) GetBaseImportTestsModelsM2O(id int64) (*BaseImportTestsModelsM2O, error) {
	btmms, err := c.GetBaseImportTestsModelsM2Os([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btmms)[0]), nil
}

// GetBaseImportTestsModelsM2Os gets base_import.tests.models.m2o existing records.
func (c *Client) GetBaseImportTestsModelsM2Os(ids []int64) (*BaseImportTestsModelsM2Os, error) {
	btmms := &BaseImportTestsModelsM2Os{}
	if err := c.Read(BaseImportTestsModelsM2OModel, ids, nil, btmms); err != nil {
		return nil, err
	}
	return btmms, nil
}

// FindBaseImportTestsModelsM2O finds base_import.tests.models.m2o record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2O(criteria *Criteria) (*BaseImportTestsModelsM2O, error) {
	btmms := &BaseImportTestsModelsM2Os{}
	if err := c.SearchRead(BaseImportTestsModelsM2OModel, criteria, NewOptions().Limit(1), btmms); err != nil {
		return nil, err
	}
	return &((*btmms)[0]), nil
}

// FindBaseImportTestsModelsM2Os finds base_import.tests.models.m2o records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2Os(criteria *Criteria, options *Options) (*BaseImportTestsModelsM2Os, error) {
	btmms := &BaseImportTestsModelsM2Os{}
	if err := c.SearchRead(BaseImportTestsModelsM2OModel, criteria, options, btmms); err != nil {
		return nil, err
	}
	return btmms, nil
}

// FindBaseImportTestsModelsM2OIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2OIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportTestsModelsM2OModel, criteria, options)
}

// FindBaseImportTestsModelsM2OId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2OId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsM2OModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
