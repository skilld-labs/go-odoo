package odoo

// ResCountryState represents res.country.state model.
type ResCountryState struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CountryId   *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResCountryStates represents array of res.country.state model.
type ResCountryStates []ResCountryState

// ResCountryStateModel is the odoo model name.
const ResCountryStateModel = "res.country.state"

// Many2One convert ResCountryState to *Many2One.
func (rcs *ResCountryState) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResCountryState creates a new res.country.state model and returns its id.
func (c *Client) CreateResCountryState(rcs *ResCountryState) (int64, error) {
	ids, err := c.CreateResCountryStates([]*ResCountryState{rcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCountryState creates a new res.country.state model and returns its id.
func (c *Client) CreateResCountryStates(rcss []*ResCountryState) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcss {
		vv = append(vv, v)
	}
	return c.Create(ResCountryStateModel, vv, nil)
}

// UpdateResCountryState updates an existing res.country.state record.
func (c *Client) UpdateResCountryState(rcs *ResCountryState) error {
	return c.UpdateResCountryStates([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResCountryStates updates existing res.country.state records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResCountryStates(ids []int64, rcs *ResCountryState) error {
	return c.Update(ResCountryStateModel, ids, rcs, nil)
}

// DeleteResCountryState deletes an existing res.country.state record.
func (c *Client) DeleteResCountryState(id int64) error {
	return c.DeleteResCountryStates([]int64{id})
}

// DeleteResCountryStates deletes existing res.country.state records.
func (c *Client) DeleteResCountryStates(ids []int64) error {
	return c.Delete(ResCountryStateModel, ids)
}

// GetResCountryState gets res.country.state existing record.
func (c *Client) GetResCountryState(id int64) (*ResCountryState, error) {
	rcss, err := c.GetResCountryStates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// GetResCountryStates gets res.country.state existing records.
func (c *Client) GetResCountryStates(ids []int64) (*ResCountryStates, error) {
	rcss := &ResCountryStates{}
	if err := c.Read(ResCountryStateModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResCountryState finds res.country.state record by querying it with criteria.
func (c *Client) FindResCountryState(criteria *Criteria) (*ResCountryState, error) {
	rcss := &ResCountryStates{}
	if err := c.SearchRead(ResCountryStateModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// FindResCountryStates finds res.country.state records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountryStates(criteria *Criteria, options *Options) (*ResCountryStates, error) {
	rcss := &ResCountryStates{}
	if err := c.SearchRead(ResCountryStateModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResCountryStateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountryStateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResCountryStateModel, criteria, options)
}

// FindResCountryStateId finds record id by querying it with criteria.
func (c *Client) FindResCountryStateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCountryStateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
