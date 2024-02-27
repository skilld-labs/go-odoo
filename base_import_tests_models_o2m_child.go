package odoo

// BaseImportTestsModelsO2MChild represents base_import.tests.models.o2m.child model.
type BaseImportTestsModelsO2MChild struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omitempty"`
	Value       *Int      `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsO2MChilds represents array of base_import.tests.models.o2m.child model.
type BaseImportTestsModelsO2MChilds []BaseImportTestsModelsO2MChild

// BaseImportTestsModelsO2MChildModel is the odoo model name.
const BaseImportTestsModelsO2MChildModel = "base_import.tests.models.o2m.child"

// Many2One convert BaseImportTestsModelsO2MChild to *Many2One.
func (btmoc *BaseImportTestsModelsO2MChild) Many2One() *Many2One {
	return NewMany2One(btmoc.Id.Get(), "")
}

// CreateBaseImportTestsModelsO2MChild creates a new base_import.tests.models.o2m.child model and returns its id.
func (c *Client) CreateBaseImportTestsModelsO2MChild(btmoc *BaseImportTestsModelsO2MChild) (int64, error) {
	ids, err := c.CreateBaseImportTestsModelsO2MChilds([]*BaseImportTestsModelsO2MChild{btmoc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportTestsModelsO2MChild creates a new base_import.tests.models.o2m.child model and returns its id.
func (c *Client) CreateBaseImportTestsModelsO2MChilds(btmocs []*BaseImportTestsModelsO2MChild) ([]int64, error) {
	var vv []interface{}
	for _, v := range btmocs {
		vv = append(vv, v)
	}
	return c.Create(BaseImportTestsModelsO2MChildModel, vv, nil)
}

// UpdateBaseImportTestsModelsO2MChild updates an existing base_import.tests.models.o2m.child record.
func (c *Client) UpdateBaseImportTestsModelsO2MChild(btmoc *BaseImportTestsModelsO2MChild) error {
	return c.UpdateBaseImportTestsModelsO2MChilds([]int64{btmoc.Id.Get()}, btmoc)
}

// UpdateBaseImportTestsModelsO2MChilds updates existing base_import.tests.models.o2m.child records.
// All records (represented by ids) will be updated by btmoc values.
func (c *Client) UpdateBaseImportTestsModelsO2MChilds(ids []int64, btmoc *BaseImportTestsModelsO2MChild) error {
	return c.Update(BaseImportTestsModelsO2MChildModel, ids, btmoc, nil)
}

// DeleteBaseImportTestsModelsO2MChild deletes an existing base_import.tests.models.o2m.child record.
func (c *Client) DeleteBaseImportTestsModelsO2MChild(id int64) error {
	return c.DeleteBaseImportTestsModelsO2MChilds([]int64{id})
}

// DeleteBaseImportTestsModelsO2MChilds deletes existing base_import.tests.models.o2m.child records.
func (c *Client) DeleteBaseImportTestsModelsO2MChilds(ids []int64) error {
	return c.Delete(BaseImportTestsModelsO2MChildModel, ids)
}

// GetBaseImportTestsModelsO2MChild gets base_import.tests.models.o2m.child existing record.
func (c *Client) GetBaseImportTestsModelsO2MChild(id int64) (*BaseImportTestsModelsO2MChild, error) {
	btmocs, err := c.GetBaseImportTestsModelsO2MChilds([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btmocs)[0]), nil
}

// GetBaseImportTestsModelsO2MChilds gets base_import.tests.models.o2m.child existing records.
func (c *Client) GetBaseImportTestsModelsO2MChilds(ids []int64) (*BaseImportTestsModelsO2MChilds, error) {
	btmocs := &BaseImportTestsModelsO2MChilds{}
	if err := c.Read(BaseImportTestsModelsO2MChildModel, ids, nil, btmocs); err != nil {
		return nil, err
	}
	return btmocs, nil
}

// FindBaseImportTestsModelsO2MChild finds base_import.tests.models.o2m.child record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsO2MChild(criteria *Criteria) (*BaseImportTestsModelsO2MChild, error) {
	btmocs := &BaseImportTestsModelsO2MChilds{}
	if err := c.SearchRead(BaseImportTestsModelsO2MChildModel, criteria, NewOptions().Limit(1), btmocs); err != nil {
		return nil, err
	}
	return &((*btmocs)[0]), nil
}

// FindBaseImportTestsModelsO2MChilds finds base_import.tests.models.o2m.child records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsO2MChilds(criteria *Criteria, options *Options) (*BaseImportTestsModelsO2MChilds, error) {
	btmocs := &BaseImportTestsModelsO2MChilds{}
	if err := c.SearchRead(BaseImportTestsModelsO2MChildModel, criteria, options, btmocs); err != nil {
		return nil, err
	}
	return btmocs, nil
}

// FindBaseImportTestsModelsO2MChildIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsO2MChildIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportTestsModelsO2MChildModel, criteria, options)
}

// FindBaseImportTestsModelsO2MChildId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsO2MChildId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsO2MChildModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
