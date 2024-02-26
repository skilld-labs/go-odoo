package odoo

// BaseImportTestsModelsCharRequired represents base_import.tests.models.char.required model.
type BaseImportTestsModelsCharRequired struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Value       *String   `xmlrpc:"value,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseImportTestsModelsCharRequireds represents array of base_import.tests.models.char.required model.
type BaseImportTestsModelsCharRequireds []BaseImportTestsModelsCharRequired

// BaseImportTestsModelsCharRequiredModel is the odoo model name.
const BaseImportTestsModelsCharRequiredModel = "base_import.tests.models.char.required"

// Many2One convert BaseImportTestsModelsCharRequired to *Many2One.
func (btmcr *BaseImportTestsModelsCharRequired) Many2One() *Many2One {
	return NewMany2One(btmcr.Id.Get(), "")
}

// CreateBaseImportTestsModelsCharRequired creates a new base_import.tests.models.char.required model and returns its id.
func (c *Client) CreateBaseImportTestsModelsCharRequired(btmcr *BaseImportTestsModelsCharRequired) (int64, error) {
	ids, err := c.CreateBaseImportTestsModelsCharRequireds([]*BaseImportTestsModelsCharRequired{btmcr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportTestsModelsCharRequired creates a new base_import.tests.models.char.required model and returns its id.
func (c *Client) CreateBaseImportTestsModelsCharRequireds(btmcrs []*BaseImportTestsModelsCharRequired) ([]int64, error) {
	var vv []interface{}
	for _, v := range btmcrs {
		vv = append(vv, v)
	}
	return c.Create(BaseImportTestsModelsCharRequiredModel, vv, nil)
}

// UpdateBaseImportTestsModelsCharRequired updates an existing base_import.tests.models.char.required record.
func (c *Client) UpdateBaseImportTestsModelsCharRequired(btmcr *BaseImportTestsModelsCharRequired) error {
	return c.UpdateBaseImportTestsModelsCharRequireds([]int64{btmcr.Id.Get()}, btmcr)
}

// UpdateBaseImportTestsModelsCharRequireds updates existing base_import.tests.models.char.required records.
// All records (represented by ids) will be updated by btmcr values.
func (c *Client) UpdateBaseImportTestsModelsCharRequireds(ids []int64, btmcr *BaseImportTestsModelsCharRequired) error {
	return c.Update(BaseImportTestsModelsCharRequiredModel, ids, btmcr, nil)
}

// DeleteBaseImportTestsModelsCharRequired deletes an existing base_import.tests.models.char.required record.
func (c *Client) DeleteBaseImportTestsModelsCharRequired(id int64) error {
	return c.DeleteBaseImportTestsModelsCharRequireds([]int64{id})
}

// DeleteBaseImportTestsModelsCharRequireds deletes existing base_import.tests.models.char.required records.
func (c *Client) DeleteBaseImportTestsModelsCharRequireds(ids []int64) error {
	return c.Delete(BaseImportTestsModelsCharRequiredModel, ids)
}

// GetBaseImportTestsModelsCharRequired gets base_import.tests.models.char.required existing record.
func (c *Client) GetBaseImportTestsModelsCharRequired(id int64) (*BaseImportTestsModelsCharRequired, error) {
	btmcrs, err := c.GetBaseImportTestsModelsCharRequireds([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btmcrs)[0]), nil
}

// GetBaseImportTestsModelsCharRequireds gets base_import.tests.models.char.required existing records.
func (c *Client) GetBaseImportTestsModelsCharRequireds(ids []int64) (*BaseImportTestsModelsCharRequireds, error) {
	btmcrs := &BaseImportTestsModelsCharRequireds{}
	if err := c.Read(BaseImportTestsModelsCharRequiredModel, ids, nil, btmcrs); err != nil {
		return nil, err
	}
	return btmcrs, nil
}

// FindBaseImportTestsModelsCharRequired finds base_import.tests.models.char.required record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharRequired(criteria *Criteria) (*BaseImportTestsModelsCharRequired, error) {
	btmcrs := &BaseImportTestsModelsCharRequireds{}
	if err := c.SearchRead(BaseImportTestsModelsCharRequiredModel, criteria, NewOptions().Limit(1), btmcrs); err != nil {
		return nil, err
	}
	return &((*btmcrs)[0]), nil
}

// FindBaseImportTestsModelsCharRequireds finds base_import.tests.models.char.required records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharRequireds(criteria *Criteria, options *Options) (*BaseImportTestsModelsCharRequireds, error) {
	btmcrs := &BaseImportTestsModelsCharRequireds{}
	if err := c.SearchRead(BaseImportTestsModelsCharRequiredModel, criteria, options, btmcrs); err != nil {
		return nil, err
	}
	return btmcrs, nil
}

// FindBaseImportTestsModelsCharRequiredIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsCharRequiredIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportTestsModelsCharRequiredModel, criteria, options)
}

// FindBaseImportTestsModelsCharRequiredId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsCharRequiredId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsCharRequiredModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
