package odoo

// BaseImportTestsModelsCharNoreadonly represents base_import.tests.models.char.noreadonly model.
type BaseImportTestsModelsCharNoreadonly struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsCharNoreadonlys represents array of base_import.tests.models.char.noreadonly model.
type BaseImportTestsModelsCharNoreadonlys []BaseImportTestsModelsCharNoreadonly

// BaseImportTestsModelsCharNoreadonlyModel is the odoo model name.
const BaseImportTestsModelsCharNoreadonlyModel = "base_import.tests.models.char.noreadonly"

// Many2One convert BaseImportTestsModelsCharNoreadonly to *Many2One.
func (btmcn *BaseImportTestsModelsCharNoreadonly) Many2One() *Many2One {
	return NewMany2One(btmcn.Id.Get(), "")
}

// CreateBaseImportTestsModelsCharNoreadonly creates a new base_import.tests.models.char.noreadonly model and returns its id.
func (c *Client) CreateBaseImportTestsModelsCharNoreadonly(btmcn *BaseImportTestsModelsCharNoreadonly) (int64, error) {
	ids, err := c.CreateBaseImportTestsModelsCharNoreadonlys([]*BaseImportTestsModelsCharNoreadonly{btmcn})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportTestsModelsCharNoreadonlys creates a new base_import.tests.models.char.noreadonly model and returns its id.
func (c *Client) CreateBaseImportTestsModelsCharNoreadonlys(btmcns []*BaseImportTestsModelsCharNoreadonly) ([]int64, error) {
	var vv []interface{}
	for _, v := range btmcns {
		vv = append(vv, v)
	}
	return c.Create(BaseImportTestsModelsCharNoreadonlyModel, vv, nil)
}

// UpdateBaseImportTestsModelsCharNoreadonly updates an existing base_import.tests.models.char.noreadonly record.
func (c *Client) UpdateBaseImportTestsModelsCharNoreadonly(btmcn *BaseImportTestsModelsCharNoreadonly) error {
	return c.UpdateBaseImportTestsModelsCharNoreadonlys([]int64{btmcn.Id.Get()}, btmcn)
}

// UpdateBaseImportTestsModelsCharNoreadonlys updates existing base_import.tests.models.char.noreadonly records.
// All records (represented by ids) will be updated by btmcn values.
func (c *Client) UpdateBaseImportTestsModelsCharNoreadonlys(ids []int64, btmcn *BaseImportTestsModelsCharNoreadonly) error {
	return c.Update(BaseImportTestsModelsCharNoreadonlyModel, ids, btmcn, nil)
}

// DeleteBaseImportTestsModelsCharNoreadonly deletes an existing base_import.tests.models.char.noreadonly record.
func (c *Client) DeleteBaseImportTestsModelsCharNoreadonly(id int64) error {
	return c.DeleteBaseImportTestsModelsCharNoreadonlys([]int64{id})
}

// DeleteBaseImportTestsModelsCharNoreadonlys deletes existing base_import.tests.models.char.noreadonly records.
func (c *Client) DeleteBaseImportTestsModelsCharNoreadonlys(ids []int64) error {
	return c.Delete(BaseImportTestsModelsCharNoreadonlyModel, ids)
}

// GetBaseImportTestsModelsCharNoreadonly gets base_import.tests.models.char.noreadonly existing record.
func (c *Client) GetBaseImportTestsModelsCharNoreadonly(id int64) (*BaseImportTestsModelsCharNoreadonly, error) {
	btmcns, err := c.GetBaseImportTestsModelsCharNoreadonlys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btmcns)[0]), nil
}

// GetBaseImportTestsModelsCharNoreadonlys gets base_import.tests.models.char.noreadonly existing records.
func (c *Client) GetBaseImportTestsModelsCharNoreadonlys(ids []int64) (*BaseImportTestsModelsCharNoreadonlys, error) {
	btmcns := &BaseImportTestsModelsCharNoreadonlys{}
	if err := c.Read(BaseImportTestsModelsCharNoreadonlyModel, ids, nil, btmcns); err != nil {
		return nil, err
	}
	return btmcns, nil
}

// FindBaseImportTestsModelsCharNoreadonly finds base_import.tests.models.char.noreadonly record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharNoreadonly(criteria *Criteria) (*BaseImportTestsModelsCharNoreadonly, error) {
	btmcns := &BaseImportTestsModelsCharNoreadonlys{}
	if err := c.SearchRead(BaseImportTestsModelsCharNoreadonlyModel, criteria, NewOptions().Limit(1), btmcns); err != nil {
		return nil, err
	}
	return &((*btmcns)[0]), nil
}

// FindBaseImportTestsModelsCharNoreadonlys finds base_import.tests.models.char.noreadonly records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharNoreadonlys(criteria *Criteria, options *Options) (*BaseImportTestsModelsCharNoreadonlys, error) {
	btmcns := &BaseImportTestsModelsCharNoreadonlys{}
	if err := c.SearchRead(BaseImportTestsModelsCharNoreadonlyModel, criteria, options, btmcns); err != nil {
		return nil, err
	}
	return btmcns, nil
}

// FindBaseImportTestsModelsCharNoreadonlyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharNoreadonlyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportTestsModelsCharNoreadonlyModel, criteria, options)
}

// FindBaseImportTestsModelsCharNoreadonlyId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharNoreadonlyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharNoreadonlyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
