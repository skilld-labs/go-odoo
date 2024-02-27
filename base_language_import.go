package odoo

// BaseLanguageImport represents base.language.import model.
type BaseLanguageImport struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Data        *String   `xmlrpc:"data,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Filename    *String   `xmlrpc:"filename,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Overwrite   *Bool     `xmlrpc:"overwrite,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseLanguageImports represents array of base.language.import model.
type BaseLanguageImports []BaseLanguageImport

// BaseLanguageImportModel is the odoo model name.
const BaseLanguageImportModel = "base.language.import"

// Many2One convert BaseLanguageImport to *Many2One.
func (bli *BaseLanguageImport) Many2One() *Many2One {
	return NewMany2One(bli.Id.Get(), "")
}

// CreateBaseLanguageImport creates a new base.language.import model and returns its id.
func (c *Client) CreateBaseLanguageImport(bli *BaseLanguageImport) (int64, error) {
	ids, err := c.CreateBaseLanguageImports([]*BaseLanguageImport{bli})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseLanguageImport creates a new base.language.import model and returns its id.
func (c *Client) CreateBaseLanguageImports(blis []*BaseLanguageImport) ([]int64, error) {
	var vv []interface{}
	for _, v := range blis {
		vv = append(vv, v)
	}
	return c.Create(BaseLanguageImportModel, vv, nil)
}

// UpdateBaseLanguageImport updates an existing base.language.import record.
func (c *Client) UpdateBaseLanguageImport(bli *BaseLanguageImport) error {
	return c.UpdateBaseLanguageImports([]int64{bli.Id.Get()}, bli)
}

// UpdateBaseLanguageImports updates existing base.language.import records.
// All records (represented by ids) will be updated by bli values.
func (c *Client) UpdateBaseLanguageImports(ids []int64, bli *BaseLanguageImport) error {
	return c.Update(BaseLanguageImportModel, ids, bli, nil)
}

// DeleteBaseLanguageImport deletes an existing base.language.import record.
func (c *Client) DeleteBaseLanguageImport(id int64) error {
	return c.DeleteBaseLanguageImports([]int64{id})
}

// DeleteBaseLanguageImports deletes existing base.language.import records.
func (c *Client) DeleteBaseLanguageImports(ids []int64) error {
	return c.Delete(BaseLanguageImportModel, ids)
}

// GetBaseLanguageImport gets base.language.import existing record.
func (c *Client) GetBaseLanguageImport(id int64) (*BaseLanguageImport, error) {
	blis, err := c.GetBaseLanguageImports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*blis)[0]), nil
}

// GetBaseLanguageImports gets base.language.import existing records.
func (c *Client) GetBaseLanguageImports(ids []int64) (*BaseLanguageImports, error) {
	blis := &BaseLanguageImports{}
	if err := c.Read(BaseLanguageImportModel, ids, nil, blis); err != nil {
		return nil, err
	}
	return blis, nil
}

// FindBaseLanguageImport finds base.language.import record by querying it with criteria.
func (c *Client) FindBaseLanguageImport(criteria *Criteria) (*BaseLanguageImport, error) {
	blis := &BaseLanguageImports{}
	if err := c.SearchRead(BaseLanguageImportModel, criteria, NewOptions().Limit(1), blis); err != nil {
		return nil, err
	}
	return &((*blis)[0]), nil
}

// FindBaseLanguageImports finds base.language.import records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseLanguageImports(criteria *Criteria, options *Options) (*BaseLanguageImports, error) {
	blis := &BaseLanguageImports{}
	if err := c.SearchRead(BaseLanguageImportModel, criteria, options, blis); err != nil {
		return nil, err
	}
	return blis, nil
}

// FindBaseLanguageImportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseLanguageImportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseLanguageImportModel, criteria, options)
}

// FindBaseLanguageImportId finds record id by querying it with criteria.
func (c *Client) FindBaseLanguageImportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseLanguageImportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
