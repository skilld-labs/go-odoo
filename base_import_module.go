package odoo

// BaseImportModule represents base.import.module model.
type BaseImportModule struct {
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Force               *Bool      `xmlrpc:"force,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	ImportMessage       *String    `xmlrpc:"import_message,omitempty"`
	ModuleFile          *String    `xmlrpc:"module_file,omitempty"`
	ModulesDependencies *String    `xmlrpc:"modules_dependencies,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	WithDemo            *Bool      `xmlrpc:"with_demo,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// BaseImportModules represents array of base.import.module model.
type BaseImportModules []BaseImportModule

// BaseImportModuleModel is the odoo model name.
const BaseImportModuleModel = "base.import.module"

// Many2One convert BaseImportModule to *Many2One.
func (bim *BaseImportModule) Many2One() *Many2One {
	return NewMany2One(bim.Id.Get(), "")
}

// CreateBaseImportModule creates a new base.import.module model and returns its id.
func (c *Client) CreateBaseImportModule(bim *BaseImportModule) (int64, error) {
	ids, err := c.CreateBaseImportModules([]*BaseImportModule{bim})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseImportModule creates a new base.import.module model and returns its id.
func (c *Client) CreateBaseImportModules(bims []*BaseImportModule) ([]int64, error) {
	var vv []interface{}
	for _, v := range bims {
		vv = append(vv, v)
	}
	return c.Create(BaseImportModuleModel, vv, nil)
}

// UpdateBaseImportModule updates an existing base.import.module record.
func (c *Client) UpdateBaseImportModule(bim *BaseImportModule) error {
	return c.UpdateBaseImportModules([]int64{bim.Id.Get()}, bim)
}

// UpdateBaseImportModules updates existing base.import.module records.
// All records (represented by ids) will be updated by bim values.
func (c *Client) UpdateBaseImportModules(ids []int64, bim *BaseImportModule) error {
	return c.Update(BaseImportModuleModel, ids, bim, nil)
}

// DeleteBaseImportModule deletes an existing base.import.module record.
func (c *Client) DeleteBaseImportModule(id int64) error {
	return c.DeleteBaseImportModules([]int64{id})
}

// DeleteBaseImportModules deletes existing base.import.module records.
func (c *Client) DeleteBaseImportModules(ids []int64) error {
	return c.Delete(BaseImportModuleModel, ids)
}

// GetBaseImportModule gets base.import.module existing record.
func (c *Client) GetBaseImportModule(id int64) (*BaseImportModule, error) {
	bims, err := c.GetBaseImportModules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bims)[0]), nil
}

// GetBaseImportModules gets base.import.module existing records.
func (c *Client) GetBaseImportModules(ids []int64) (*BaseImportModules, error) {
	bims := &BaseImportModules{}
	if err := c.Read(BaseImportModuleModel, ids, nil, bims); err != nil {
		return nil, err
	}
	return bims, nil
}

// FindBaseImportModule finds base.import.module record by querying it with criteria.
func (c *Client) FindBaseImportModule(criteria *Criteria) (*BaseImportModule, error) {
	bims := &BaseImportModules{}
	if err := c.SearchRead(BaseImportModuleModel, criteria, NewOptions().Limit(1), bims); err != nil {
		return nil, err
	}
	return &((*bims)[0]), nil
}

// FindBaseImportModules finds base.import.module records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportModules(criteria *Criteria, options *Options) (*BaseImportModules, error) {
	bims := &BaseImportModules{}
	if err := c.SearchRead(BaseImportModuleModel, criteria, options, bims); err != nil {
		return nil, err
	}
	return bims, nil
}

// FindBaseImportModuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportModuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseImportModuleModel, criteria, options)
}

// FindBaseImportModuleId finds record id by querying it with criteria.
func (c *Client) FindBaseImportModuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportModuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
