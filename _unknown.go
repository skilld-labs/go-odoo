package odoo

// Unknown represents _unknown model.
type Unknown struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// Unknowns represents array of _unknown model.
type Unknowns []Unknown

// UnknownModel is the odoo model name.
const UnknownModel = "_unknown"

// Many2One convert Unknown to *Many2One.
func (_ *Unknown) Many2One() *Many2One {
	return NewMany2One(_.Id.Get(), "")
}

// CreateUnknown creates a new _unknown model and returns its id.
func (c *Client) CreateUnknown(_ *Unknown) (int64, error) {
	ids, err := c.CreateUnknowns([]*Unknown{_})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateUnknown creates a new _unknown model and returns its id.
func (c *Client) CreateUnknowns(_s []*Unknown) ([]int64, error) {
	var vv []interface{}
	for _, v := range _s {
		vv = append(vv, v)
	}
	return c.Create(UnknownModel, vv, nil)
}

// UpdateUnknown updates an existing _unknown record.
func (c *Client) UpdateUnknown(_ *Unknown) error {
	return c.UpdateUnknowns([]int64{_.Id.Get()}, _)
}

// UpdateUnknowns updates existing _unknown records.
// All records (represented by ids) will be updated by _ values.
func (c *Client) UpdateUnknowns(ids []int64, _ *Unknown) error {
	return c.Update(UnknownModel, ids, _, nil)
}

// DeleteUnknown deletes an existing _unknown record.
func (c *Client) DeleteUnknown(id int64) error {
	return c.DeleteUnknowns([]int64{id})
}

// DeleteUnknowns deletes existing _unknown records.
func (c *Client) DeleteUnknowns(ids []int64) error {
	return c.Delete(UnknownModel, ids)
}

// GetUnknown gets _unknown existing record.
func (c *Client) GetUnknown(id int64) (*Unknown, error) {
	_s, err := c.GetUnknowns([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*_s)[0]), nil
}

// GetUnknowns gets _unknown existing records.
func (c *Client) GetUnknowns(ids []int64) (*Unknowns, error) {
	_s := &Unknowns{}
	if err := c.Read(UnknownModel, ids, nil, _s); err != nil {
		return nil, err
	}
	return _s, nil
}

// FindUnknown finds _unknown record by querying it with criteria.
func (c *Client) FindUnknown(criteria *Criteria) (*Unknown, error) {
	_s := &Unknowns{}
	if err := c.SearchRead(UnknownModel, criteria, NewOptions().Limit(1), _s); err != nil {
		return nil, err
	}
	return &((*_s)[0]), nil
}

// FindUnknowns finds _unknown records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUnknowns(criteria *Criteria, options *Options) (*Unknowns, error) {
	_s := &Unknowns{}
	if err := c.SearchRead(UnknownModel, criteria, options, _s); err != nil {
		return nil, err
	}
	return _s, nil
}

// FindUnknownIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUnknownIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(UnknownModel, criteria, options)
}

// FindUnknownId finds record id by querying it with criteria.
func (c *Client) FindUnknownId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UnknownModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
