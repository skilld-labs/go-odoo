package odoo

// BaseImportTestsModelsPreview represents base_import.tests.models.preview model.
type BaseImportTestsModelsPreview struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Othervalue  *Int      `xmlrpc:"othervalue,omitempty"`
	Somevalue   *Int      `xmlrpc:"somevalue,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseImportTestsModelsPreviews represents array of base_import.tests.models.preview model.
type BaseImportTestsModelsPreviews []BaseImportTestsModelsPreview

// BaseImportTestsModelsPreviewModel is the odoo model name.
const BaseImportTestsModelsPreviewModel = "base_import.tests.models.preview"

// Many2One convert BaseImportTestsModelsPreview to *Many2One.
func (btmp *BaseImportTestsModelsPreview) Many2One() *Many2One {
	return NewMany2One(btmp.Id.Get(), "")
}

// CreateBaseImportTestsModelsPreview creates a new base_import.tests.models.preview model and returns its id.
func (c *Client) CreateBaseImportTestsModelsPreview(btmp *BaseImportTestsModelsPreview) (int64, error) {
	ids, err := c.CreateBaseImportTestsModelsPreviews([]*BaseImportTestsModelsPreview{btmp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportTestsModelsPreviews creates a new base_import.tests.models.preview model and returns its id.
func (c *Client) CreateBaseImportTestsModelsPreviews(btmps []*BaseImportTestsModelsPreview) ([]int64, error) {
	var vv []interface{}
	for _, v := range btmps {
		vv = append(vv, v)
	}
	return c.Create(BaseImportTestsModelsPreviewModel, vv, nil)
}

// UpdateBaseImportTestsModelsPreview updates an existing base_import.tests.models.preview record.
func (c *Client) UpdateBaseImportTestsModelsPreview(btmp *BaseImportTestsModelsPreview) error {
	return c.UpdateBaseImportTestsModelsPreviews([]int64{btmp.Id.Get()}, btmp)
}

// UpdateBaseImportTestsModelsPreviews updates existing base_import.tests.models.preview records.
// All records (represented by ids) will be updated by btmp values.
func (c *Client) UpdateBaseImportTestsModelsPreviews(ids []int64, btmp *BaseImportTestsModelsPreview) error {
	return c.Update(BaseImportTestsModelsPreviewModel, ids, btmp, nil)
}

// DeleteBaseImportTestsModelsPreview deletes an existing base_import.tests.models.preview record.
func (c *Client) DeleteBaseImportTestsModelsPreview(id int64) error {
	return c.DeleteBaseImportTestsModelsPreviews([]int64{id})
}

// DeleteBaseImportTestsModelsPreviews deletes existing base_import.tests.models.preview records.
func (c *Client) DeleteBaseImportTestsModelsPreviews(ids []int64) error {
	return c.Delete(BaseImportTestsModelsPreviewModel, ids)
}

// GetBaseImportTestsModelsPreview gets base_import.tests.models.preview existing record.
func (c *Client) GetBaseImportTestsModelsPreview(id int64) (*BaseImportTestsModelsPreview, error) {
	btmps, err := c.GetBaseImportTestsModelsPreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btmps)[0]), nil
}

// GetBaseImportTestsModelsPreviews gets base_import.tests.models.preview existing records.
func (c *Client) GetBaseImportTestsModelsPreviews(ids []int64) (*BaseImportTestsModelsPreviews, error) {
	btmps := &BaseImportTestsModelsPreviews{}
	if err := c.Read(BaseImportTestsModelsPreviewModel, ids, nil, btmps); err != nil {
		return nil, err
	}
	return btmps, nil
}

// FindBaseImportTestsModelsPreview finds base_import.tests.models.preview record by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsPreview(criteria *Criteria) (*BaseImportTestsModelsPreview, error) {
	btmps := &BaseImportTestsModelsPreviews{}
	if err := c.SearchRead(BaseImportTestsModelsPreviewModel, criteria, NewOptions().Limit(1), btmps); err != nil {
		return nil, err
	}
	return &((*btmps)[0]), nil
}

// FindBaseImportTestsModelsPreviews finds base_import.tests.models.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsPreviews(criteria *Criteria, options *Options) (*BaseImportTestsModelsPreviews, error) {
	btmps := &BaseImportTestsModelsPreviews{}
	if err := c.SearchRead(BaseImportTestsModelsPreviewModel, criteria, options, btmps); err != nil {
		return nil, err
	}
	return btmps, nil
}

// FindBaseImportTestsModelsPreviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportTestsModelsPreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportTestsModelsPreviewModel, criteria, options)
}

// FindBaseImportTestsModelsPreviewId finds record id by querying it with criteria.
func (c *Client) FindBaseImportTestsModelsPreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportTestsModelsPreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
