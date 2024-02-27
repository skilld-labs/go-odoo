package odoo

// ResCountry represents res.country model.
type ResCountry struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	AddressFormat   *String    `xmlrpc:"address_format,omitempty"`
	AddressViewId   *Many2One  `xmlrpc:"address_view_id,omitempty"`
	Code            *String    `xmlrpc:"code,omitempty"`
	CountryGroupIds *Relation  `xmlrpc:"country_group_ids,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId      *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	Image           *String    `xmlrpc:"image,omitempty"`
	Name            *String    `xmlrpc:"name,omitempty"`
	NamePosition    *Selection `xmlrpc:"name_position,omitempty"`
	PhoneCode       *Int       `xmlrpc:"phone_code,omitempty"`
	StateIds        *Relation  `xmlrpc:"state_ids,omitempty"`
	VatLabel        *String    `xmlrpc:"vat_label,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResCountrys represents array of res.country model.
type ResCountrys []ResCountry

// ResCountryModel is the odoo model name.
const ResCountryModel = "res.country"

// Many2One convert ResCountry to *Many2One.
func (rc *ResCountry) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCountry creates a new res.country model and returns its id.
func (c *Client) CreateResCountry(rc *ResCountry) (int64, error) {
	ids, err := c.CreateResCountrys([]*ResCountry{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCountry creates a new res.country model and returns its id.
func (c *Client) CreateResCountrys(rcs []*ResCountry) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResCountryModel, vv, nil)
}

// UpdateResCountry updates an existing res.country record.
func (c *Client) UpdateResCountry(rc *ResCountry) error {
	return c.UpdateResCountrys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCountrys updates existing res.country records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCountrys(ids []int64, rc *ResCountry) error {
	return c.Update(ResCountryModel, ids, rc, nil)
}

// DeleteResCountry deletes an existing res.country record.
func (c *Client) DeleteResCountry(id int64) error {
	return c.DeleteResCountrys([]int64{id})
}

// DeleteResCountrys deletes existing res.country records.
func (c *Client) DeleteResCountrys(ids []int64) error {
	return c.Delete(ResCountryModel, ids)
}

// GetResCountry gets res.country existing record.
func (c *Client) GetResCountry(id int64) (*ResCountry, error) {
	rcs, err := c.GetResCountrys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// GetResCountrys gets res.country existing records.
func (c *Client) GetResCountrys(ids []int64) (*ResCountrys, error) {
	rcs := &ResCountrys{}
	if err := c.Read(ResCountryModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCountry finds res.country record by querying it with criteria.
func (c *Client) FindResCountry(criteria *Criteria) (*ResCountry, error) {
	rcs := &ResCountrys{}
	if err := c.SearchRead(ResCountryModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// FindResCountrys finds res.country records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountrys(criteria *Criteria, options *Options) (*ResCountrys, error) {
	rcs := &ResCountrys{}
	if err := c.SearchRead(ResCountryModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCountryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCountryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResCountryModel, criteria, options)
}

// FindResCountryId finds record id by querying it with criteria.
func (c *Client) FindResCountryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCountryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
