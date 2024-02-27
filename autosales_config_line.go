package odoo

// AutosalesConfigLine represents autosales.config.line model.
type AutosalesConfigLine struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omitempty"`
	AutosalesConfigId  *Many2One `xmlrpc:"autosales_config_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	PercentProductBase *Float    `xmlrpc:"percent_product_base,omitempty"`
	ProductAuto        *Many2One `xmlrpc:"product_auto,omitempty"`
	ProductBase        *Many2One `xmlrpc:"product_base,omitempty"`
	Sequence           *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AutosalesConfigLines represents array of autosales.config.line model.
type AutosalesConfigLines []AutosalesConfigLine

// AutosalesConfigLineModel is the odoo model name.
const AutosalesConfigLineModel = "autosales.config.line"

// Many2One convert AutosalesConfigLine to *Many2One.
func (acl *AutosalesConfigLine) Many2One() *Many2One {
	return NewMany2One(acl.Id.Get(), "")
}

// CreateAutosalesConfigLine creates a new autosales.config.line model and returns its id.
func (c *Client) CreateAutosalesConfigLine(acl *AutosalesConfigLine) (int64, error) {
	ids, err := c.CreateAutosalesConfigLines([]*AutosalesConfigLine{acl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAutosalesConfigLine creates a new autosales.config.line model and returns its id.
func (c *Client) CreateAutosalesConfigLines(acls []*AutosalesConfigLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range acls {
		vv = append(vv, v)
	}
	return c.Create(AutosalesConfigLineModel, vv, nil)
}

// UpdateAutosalesConfigLine updates an existing autosales.config.line record.
func (c *Client) UpdateAutosalesConfigLine(acl *AutosalesConfigLine) error {
	return c.UpdateAutosalesConfigLines([]int64{acl.Id.Get()}, acl)
}

// UpdateAutosalesConfigLines updates existing autosales.config.line records.
// All records (represented by ids) will be updated by acl values.
func (c *Client) UpdateAutosalesConfigLines(ids []int64, acl *AutosalesConfigLine) error {
	return c.Update(AutosalesConfigLineModel, ids, acl, nil)
}

// DeleteAutosalesConfigLine deletes an existing autosales.config.line record.
func (c *Client) DeleteAutosalesConfigLine(id int64) error {
	return c.DeleteAutosalesConfigLines([]int64{id})
}

// DeleteAutosalesConfigLines deletes existing autosales.config.line records.
func (c *Client) DeleteAutosalesConfigLines(ids []int64) error {
	return c.Delete(AutosalesConfigLineModel, ids)
}

// GetAutosalesConfigLine gets autosales.config.line existing record.
func (c *Client) GetAutosalesConfigLine(id int64) (*AutosalesConfigLine, error) {
	acls, err := c.GetAutosalesConfigLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*acls)[0]), nil
}

// GetAutosalesConfigLines gets autosales.config.line existing records.
func (c *Client) GetAutosalesConfigLines(ids []int64) (*AutosalesConfigLines, error) {
	acls := &AutosalesConfigLines{}
	if err := c.Read(AutosalesConfigLineModel, ids, nil, acls); err != nil {
		return nil, err
	}
	return acls, nil
}

// FindAutosalesConfigLine finds autosales.config.line record by querying it with criteria.
func (c *Client) FindAutosalesConfigLine(criteria *Criteria) (*AutosalesConfigLine, error) {
	acls := &AutosalesConfigLines{}
	if err := c.SearchRead(AutosalesConfigLineModel, criteria, NewOptions().Limit(1), acls); err != nil {
		return nil, err
	}
	return &((*acls)[0]), nil
}

// FindAutosalesConfigLines finds autosales.config.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAutosalesConfigLines(criteria *Criteria, options *Options) (*AutosalesConfigLines, error) {
	acls := &AutosalesConfigLines{}
	if err := c.SearchRead(AutosalesConfigLineModel, criteria, options, acls); err != nil {
		return nil, err
	}
	return acls, nil
}

// FindAutosalesConfigLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAutosalesConfigLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AutosalesConfigLineModel, criteria, options)
}

// FindAutosalesConfigLineId finds record id by querying it with criteria.
func (c *Client) FindAutosalesConfigLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AutosalesConfigLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
