package odoo

// ResLang represents res.lang model.
type ResLang struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	Active       *Bool      `xmlrpc:"active,omptempty"`
	Code         *String    `xmlrpc:"code,omptempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFormat   *String    `xmlrpc:"date_format,omptempty"`
	DecimalPoint *String    `xmlrpc:"decimal_point,omptempty"`
	Direction    *Selection `xmlrpc:"direction,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	Grouping     *String    `xmlrpc:"grouping,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	IsoCode      *String    `xmlrpc:"iso_code,omptempty"`
	Name         *String    `xmlrpc:"name,omptempty"`
	ThousandsSep *String    `xmlrpc:"thousands_sep,omptempty"`
	TimeFormat   *String    `xmlrpc:"time_format,omptempty"`
	Translatable *Bool      `xmlrpc:"translatable,omptempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResLangs represents array of res.lang model.
type ResLangs []ResLang

// ResLangModel is the odoo model name.
const ResLangModel = "res.lang"

// Many2One convert ResLang to *Many2One.
func (rl *ResLang) Many2One() *Many2One {
	return NewMany2One(rl.Id.Get(), "")
}

// CreateResLang creates a new res.lang model and returns its id.
func (c *Client) CreateResLang(rl *ResLang) (int64, error) {
	ids, err := c.CreateResLangs([]*ResLang{rl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResLang creates a new res.lang model and returns its id.
func (c *Client) CreateResLangs(rls []*ResLang) ([]int64, error) {
	var vv []interface{}
	for _, v := range rls {
		vv = append(vv, v)
	}
	return c.Create(ResLangModel, vv, nil)
}

// UpdateResLang updates an existing res.lang record.
func (c *Client) UpdateResLang(rl *ResLang) error {
	return c.UpdateResLangs([]int64{rl.Id.Get()}, rl)
}

// UpdateResLangs updates existing res.lang records.
// All records (represented by ids) will be updated by rl values.
func (c *Client) UpdateResLangs(ids []int64, rl *ResLang) error {
	return c.Update(ResLangModel, ids, rl, nil)
}

// DeleteResLang deletes an existing res.lang record.
func (c *Client) DeleteResLang(id int64) error {
	return c.DeleteResLangs([]int64{id})
}

// DeleteResLangs deletes existing res.lang records.
func (c *Client) DeleteResLangs(ids []int64) error {
	return c.Delete(ResLangModel, ids)
}

// GetResLang gets res.lang existing record.
func (c *Client) GetResLang(id int64) (*ResLang, error) {
	rls, err := c.GetResLangs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rls)[0]), nil
}

// GetResLangs gets res.lang existing records.
func (c *Client) GetResLangs(ids []int64) (*ResLangs, error) {
	rls := &ResLangs{}
	if err := c.Read(ResLangModel, ids, nil, rls); err != nil {
		return nil, err
	}
	return rls, nil
}

// FindResLang finds res.lang record by querying it with criteria.
func (c *Client) FindResLang(criteria *Criteria) (*ResLang, error) {
	rls := &ResLangs{}
	if err := c.SearchRead(ResLangModel, criteria, NewOptions().Limit(1), rls); err != nil {
		return nil, err
	}
	return &((*rls)[0]), nil
}

// FindResLangs finds res.lang records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResLangs(criteria *Criteria, options *Options) (*ResLangs, error) {
	rls := &ResLangs{}
	if err := c.SearchRead(ResLangModel, criteria, options, rls); err != nil {
		return nil, err
	}
	return rls, nil
}

// FindResLangIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResLangIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResLangModel, criteria, options)
}

// FindResLangId finds record id by querying it with criteria.
func (c *Client) FindResLangId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResLangModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
