package odoo

import (
	"fmt"
)

// ResLang represents res.lang model.
type ResLang struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omitempty"`
	Active       *Bool      `xmlrpc:"active,omitempty"`
	Code         *String    `xmlrpc:"code,omitempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFormat   *String    `xmlrpc:"date_format,omitempty"`
	DecimalPoint *String    `xmlrpc:"decimal_point,omitempty"`
	Direction    *Selection `xmlrpc:"direction,omitempty"`
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	FlagImage    *String    `xmlrpc:"flag_image,omitempty"`
	FlagImageUrl *String    `xmlrpc:"flag_image_url,omitempty"`
	Grouping     *String    `xmlrpc:"grouping,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	IsoCode      *String    `xmlrpc:"iso_code,omitempty"`
	Name         *String    `xmlrpc:"name,omitempty"`
	ThousandsSep *String    `xmlrpc:"thousands_sep,omitempty"`
	TimeFormat   *String    `xmlrpc:"time_format,omitempty"`
	UrlCode      *String    `xmlrpc:"url_code,omitempty"`
	WeekStart    *Selection `xmlrpc:"week_start,omitempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(ResLangModel, rl)
}

// UpdateResLang updates an existing res.lang record.
func (c *Client) UpdateResLang(rl *ResLang) error {
	return c.UpdateResLangs([]int64{rl.Id.Get()}, rl)
}

// UpdateResLangs updates existing res.lang records.
// All records (represented by ids) will be updated by rl values.
func (c *Client) UpdateResLangs(ids []int64, rl *ResLang) error {
	return c.Update(ResLangModel, ids, rl)
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
	if rls != nil && len(*rls) > 0 {
		return &((*rls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.lang not found", id)
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
	if rls != nil && len(*rls) > 0 {
		return &((*rls)[0]), nil
	}
	return nil, fmt.Errorf("res.lang was not found")
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
	ids, err := c.Search(ResLangModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResLangId finds record id by querying it with criteria.
func (c *Client) FindResLangId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResLangModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.lang was not found")
}
