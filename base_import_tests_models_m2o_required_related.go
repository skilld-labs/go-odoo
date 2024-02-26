package odoo

// BaseImportTestsModelsM2ORequiredRelated represents base_import.tests.models.m2o.required.related model.
type BaseImportTestsModelsM2ORequiredRelated struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Value       *Int      `xmlrpc:"value,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseImportTestsModelsM2ORequiredRelateds represents array of base_import.tests.models.m2o.required.related model.
type BaseImportTestsModelsM2ORequiredRelateds []BaseImportTestsModelsM2ORequiredRelated

// BaseImportTestsModelsM2ORequiredRelatedModel is the odoo model name.
const BaseImportTestsModelsM2ORequiredRelatedModel = "base_import.tests.models.m2o.required.related"

// Many2One convert BaseImportTestsModelsM2ORequiredRelated to *Many2One.
func (btmmrr *BaseImportTestsModelsM2ORequiredRelated) Many2One() *Many2One {
	return NewMany2One(btmmrr.Id.Get(), "")
}

// CreateBaseImportTestsModelsM2ORequiredRelated creates a new base_import.tests.models.m2o.required.related model and returns its id.
func (c *Client) CreateBaseImportTestsModelsM2ORequiredRelated(btmmrr *BaseImportTestsModelsM2ORequiredRelated) (int64, error) {
	ids, err := c.CreateBaseImportTestsModelsM2ORequiredRelateds([]*BaseImportTestsModelsM2ORequiredRelated{btmmrr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportTestsModelsM2ORequiredRelated creates a new base_import.tests.models.m2o.required.related model and returns its id.
func (c *Client) CreateBaseImportTestsModelsM2ORequiredRelateds(btmmrrs []*BaseImportTestsModelsM2ORequiredRelated) ([]int64, error) {
	var vv []interface{}
	for _, v := range btmmrrs {
		vv = append(vv, v)
	}
	return c.Create(BaseImportTestsModelsM2ORequiredRelatedModel, vv, nil)
}

// UpdateBaseImportTestsModelsM2ORequiredRelated updates an existing base_import.tests.models.m2o.required.related record.
func (c *Client) UpdateBaseImportTestsModelsM2ORequiredRelated(btmmrr *BaseImportTestsModelsM2ORequiredRelated) error {
	return c.UpdateBaseImportTestsModelsM2ORequiredRelateds([]int64{btmmrr.Id.Get()}, btmmrr)
}

// UpdateBaseImportTestsModelsM2ORequiredRelateds updates existing base_import.tests.models.m2o.required.related records.
// All records (represented by ids) will be updated by btmmrr values.
func (c *Client) UpdateBaseImportTestsModelsM2ORequiredRelateds(ids []int64, btmmrr *BaseImportTestsModelsM2ORequiredRelated) error {
	return c.Update(BaseImportTestsModelsM2ORequiredRelatedModel, ids, btmmrr, nil)
}

// DeleteBaseImportTestsModelsM2ORequiredRelated deletes an existing base_import.tests.models.m2o.required.related record.
func (c *Client) DeleteBaseImportTestsModelsM2ORequiredRelated(id int64) error {
	return c.DeleteBaseImportTestsModelsM2ORequiredRelateds([]int64{id})
}

// DeleteBaseImportTestsModelsM2ORequiredRelateds deletes existing base_import.tests.models.m2o.required.related records.
func (c *Client) DeleteBaseImportTestsModelsM2ORequiredRelateds(ids []int64) error {
	return c.Delete(BaseImportTestsModelsM2ORequiredRelatedModel, ids)
}

// GetBaseImportTestsModelsM2ORequiredRelated gets base_import.tests.models.m2o.required.related existing record.
func (c *Client) GetBaseImportTestsModelsM2ORequiredRelated(id int64) (*BaseImportTestsModelsM2ORequiredRelated, error) {
	btmmrrs, err := c.GetBaseImportTestsModelsM2ORequiredRelateds([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btmmrrs)[0]), nil
}

// GetBaseImportTestsModelsM2ORequiredRelateds gets base_import.tests.models.m2o.required.related existing records.
func (c *Client) GetBaseImportTestsModelsM2ORequiredRelateds(ids []int64) (*BaseImportTestsModelsM2ORequiredRelateds, error) {
	btmmrrs := &BaseImportTestsModelsM2ORequiredRelateds{}
	if err := c.Read(BaseImportTestsModelsM2ORequiredRelatedModel, ids, nil, btmmrrs); err != nil {
		return nil, err
	}
	return btmmrrs, nil
}

// FindBaseImportTestsModelsM2ORequiredRelated finds base_import.tests.models.m2o.required.related record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2ORequiredRelated(criteria *Criteria) (*BaseImportTestsModelsM2ORequiredRelated, error) {
	btmmrrs := &BaseImportTestsModelsM2ORequiredRelateds{}
	if err := c.SearchRead(BaseImportTestsModelsM2ORequiredRelatedModel, criteria, NewOptions().Limit(1), btmmrrs); err != nil {
		return nil, err
	}
	return &((*btmmrrs)[0]), nil
}

// FindBaseImportTestsModelsM2ORequiredRelateds finds base_import.tests.models.m2o.required.related records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2ORequiredRelateds(criteria *Criteria, options *Options) (*BaseImportTestsModelsM2ORequiredRelateds, error) {
	btmmrrs := &BaseImportTestsModelsM2ORequiredRelateds{}
	if err := c.SearchRead(BaseImportTestsModelsM2ORequiredRelatedModel, criteria, options, btmmrrs); err != nil {
		return nil, err
	}
	return btmmrrs, nil
}

// FindBaseImportTestsModelsM2ORequiredRelatedIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2ORequiredRelatedIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportTestsModelsM2ORequiredRelatedModel, criteria, options)
}

// FindBaseImportTestsModelsM2ORequiredRelatedId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2ORequiredRelatedId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsM2ORequiredRelatedModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
