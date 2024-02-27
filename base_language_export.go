package odoo

// BaseLanguageExport represents base.language.export model.
type BaseLanguageExport struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Data        *String    `xmlrpc:"data,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Format      *Selection `xmlrpc:"format,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Lang        *Selection `xmlrpc:"lang,omitempty"`
	Modules     *Relation  `xmlrpc:"modules,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	State       *Selection `xmlrpc:"state,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// BaseLanguageExports represents array of base.language.export model.
type BaseLanguageExports []BaseLanguageExport

// BaseLanguageExportModel is the odoo model name.
const BaseLanguageExportModel = "base.language.export"

// Many2One convert BaseLanguageExport to *Many2One.
func (ble *BaseLanguageExport) Many2One() *Many2One {
	return NewMany2One(ble.Id.Get(), "")
}

// CreateBaseLanguageExport creates a new base.language.export model and returns its id.
func (c *Client) CreateBaseLanguageExport(ble *BaseLanguageExport) (int64, error) {
	ids, err := c.CreateBaseLanguageExports([]*BaseLanguageExport{ble})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseLanguageExports creates a new base.language.export model and returns its id.
func (c *Client) CreateBaseLanguageExports(bles []*BaseLanguageExport) ([]int64, error) {
	var vv []interface{}
	for _, v := range bles {
		vv = append(vv, v)
	}
	return c.Create(BaseLanguageExportModel, vv, nil)
}

// UpdateBaseLanguageExport updates an existing base.language.export record.
func (c *Client) UpdateBaseLanguageExport(ble *BaseLanguageExport) error {
	return c.UpdateBaseLanguageExports([]int64{ble.Id.Get()}, ble)
}

// UpdateBaseLanguageExports updates existing base.language.export records.
// All records (represented by ids) will be updated by ble values.
func (c *Client) UpdateBaseLanguageExports(ids []int64, ble *BaseLanguageExport) error {
	return c.Update(BaseLanguageExportModel, ids, ble, nil)
}

// DeleteBaseLanguageExport deletes an existing base.language.export record.
func (c *Client) DeleteBaseLanguageExport(id int64) error {
	return c.DeleteBaseLanguageExports([]int64{id})
}

// DeleteBaseLanguageExports deletes existing base.language.export records.
func (c *Client) DeleteBaseLanguageExports(ids []int64) error {
	return c.Delete(BaseLanguageExportModel, ids)
}

// GetBaseLanguageExport gets base.language.export existing record.
func (c *Client) GetBaseLanguageExport(id int64) (*BaseLanguageExport, error) {
	bles, err := c.GetBaseLanguageExports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bles)[0]), nil
}

// GetBaseLanguageExports gets base.language.export existing records.
func (c *Client) GetBaseLanguageExports(ids []int64) (*BaseLanguageExports, error) {
	bles := &BaseLanguageExports{}
	if err := c.Read(BaseLanguageExportModel, ids, nil, bles); err != nil {
		return nil, err
	}
	return bles, nil
}

// FindBaseLanguageExport finds base.language.export record by querying it with criteria.
func (c *Client) FindBaseLanguageExport(criteria *Criteria) (*BaseLanguageExport, error) {
	bles := &BaseLanguageExports{}
	if err := c.SearchRead(BaseLanguageExportModel, criteria, NewOptions().Limit(1), bles); err != nil {
		return nil, err
	}
	return &((*bles)[0]), nil
}

// FindBaseLanguageExports finds base.language.export records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseLanguageExports(criteria *Criteria, options *Options) (*BaseLanguageExports, error) {
	bles := &BaseLanguageExports{}
	if err := c.SearchRead(BaseLanguageExportModel, criteria, options, bles); err != nil {
		return nil, err
	}
	return bles, nil
}

// FindBaseLanguageExportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseLanguageExportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseLanguageExportModel, criteria, options)
}

// FindBaseLanguageExportId finds record id by querying it with criteria.
func (c *Client) FindBaseLanguageExportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseLanguageExportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
