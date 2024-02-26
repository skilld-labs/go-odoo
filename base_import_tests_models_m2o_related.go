package odoo

// BaseImportTestsModelsM2ORelated represents base_import.tests.models.m2o.related model.
type BaseImportTestsModelsM2ORelated struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Value       *Int      `xmlrpc:"value,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseImportTestsModelsM2ORelateds represents array of base_import.tests.models.m2o.related model.
type BaseImportTestsModelsM2ORelateds []BaseImportTestsModelsM2ORelated

// BaseImportTestsModelsM2ORelatedModel is the odoo model name.
const BaseImportTestsModelsM2ORelatedModel = "base_import.tests.models.m2o.related"

// Many2One convert BaseImportTestsModelsM2ORelated to *Many2One.
func (btmmr *BaseImportTestsModelsM2ORelated) Many2One() *Many2One {
	return NewMany2One(btmmr.Id.Get(), "")
}

// CreateBaseImportTestsModelsM2ORelated creates a new base_import.tests.models.m2o.related model and returns its id.
func (c *Client) CreateBaseImportTestsModelsM2ORelated(btmmr *BaseImportTestsModelsM2ORelated) (int64, error) {
	ids, err := c.CreateBaseImportTestsModelsM2ORelateds([]*BaseImportTestsModelsM2ORelated{btmmr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportTestsModelsM2ORelated creates a new base_import.tests.models.m2o.related model and returns its id.
func (c *Client) CreateBaseImportTestsModelsM2ORelateds(btmmrs []*BaseImportTestsModelsM2ORelated) ([]int64, error) {
	var vv []interface{}
	for _, v := range btmmrs {
		vv = append(vv, v)
	}
	return c.Create(BaseImportTestsModelsM2ORelatedModel, vv, nil)
}

// UpdateBaseImportTestsModelsM2ORelated updates an existing base_import.tests.models.m2o.related record.
func (c *Client) UpdateBaseImportTestsModelsM2ORelated(btmmr *BaseImportTestsModelsM2ORelated) error {
	return c.UpdateBaseImportTestsModelsM2ORelateds([]int64{btmmr.Id.Get()}, btmmr)
}

// UpdateBaseImportTestsModelsM2ORelateds updates existing base_import.tests.models.m2o.related records.
// All records (represented by ids) will be updated by btmmr values.
func (c *Client) UpdateBaseImportTestsModelsM2ORelateds(ids []int64, btmmr *BaseImportTestsModelsM2ORelated) error {
	return c.Update(BaseImportTestsModelsM2ORelatedModel, ids, btmmr, nil)
}

// DeleteBaseImportTestsModelsM2ORelated deletes an existing base_import.tests.models.m2o.related record.
func (c *Client) DeleteBaseImportTestsModelsM2ORelated(id int64) error {
	return c.DeleteBaseImportTestsModelsM2ORelateds([]int64{id})
}

// DeleteBaseImportTestsModelsM2ORelateds deletes existing base_import.tests.models.m2o.related records.
func (c *Client) DeleteBaseImportTestsModelsM2ORelateds(ids []int64) error {
	return c.Delete(BaseImportTestsModelsM2ORelatedModel, ids)
}

// GetBaseImportTestsModelsM2ORelated gets base_import.tests.models.m2o.related existing record.
func (c *Client) GetBaseImportTestsModelsM2ORelated(id int64) (*BaseImportTestsModelsM2ORelated, error) {
	btmmrs, err := c.GetBaseImportTestsModelsM2ORelateds([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btmmrs)[0]), nil
}

// GetBaseImportTestsModelsM2ORelateds gets base_import.tests.models.m2o.related existing records.
func (c *Client) GetBaseImportTestsModelsM2ORelateds(ids []int64) (*BaseImportTestsModelsM2ORelateds, error) {
	btmmrs := &BaseImportTestsModelsM2ORelateds{}
	if err := c.Read(BaseImportTestsModelsM2ORelatedModel, ids, nil, btmmrs); err != nil {
		return nil, err
	}
	return btmmrs, nil
}

// FindBaseImportTestsModelsM2ORelated finds base_import.tests.models.m2o.related record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2ORelated(criteria *Criteria) (*BaseImportTestsModelsM2ORelated, error) {
	btmmrs := &BaseImportTestsModelsM2ORelateds{}
	if err := c.SearchRead(BaseImportTestsModelsM2ORelatedModel, criteria, NewOptions().Limit(1), btmmrs); err != nil {
		return nil, err
	}
	return &((*btmmrs)[0]), nil
}

// FindBaseImportTestsModelsM2ORelateds finds base_import.tests.models.m2o.related records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2ORelateds(criteria *Criteria, options *Options) (*BaseImportTestsModelsM2ORelateds, error) {
	btmmrs := &BaseImportTestsModelsM2ORelateds{}
	if err := c.SearchRead(BaseImportTestsModelsM2ORelatedModel, criteria, options, btmmrs); err != nil {
		return nil, err
	}
	return btmmrs, nil
}

// FindBaseImportTestsModelsM2ORelatedIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsM2ORelatedIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportTestsModelsM2ORelatedModel, criteria, options)
}

// FindBaseImportTestsModelsM2ORelatedId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsM2ORelatedId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsM2ORelatedModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
