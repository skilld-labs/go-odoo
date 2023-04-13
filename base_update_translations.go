package odoo

import (
	"fmt"
)

// BaseUpdateTranslations represents base.update.translations model.
type BaseUpdateTranslations struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Lang        *Selection `xmlrpc:"lang,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// BaseUpdateTranslationss represents array of base.update.translations model.
type BaseUpdateTranslationss []BaseUpdateTranslations

// BaseUpdateTranslationsModel is the odoo model name.
const BaseUpdateTranslationsModel = "base.update.translations"

// Many2One convert BaseUpdateTranslations to *Many2One.
func (but *BaseUpdateTranslations) Many2One() *Many2One {
	return NewMany2One(but.Id.Get(), "")
}

// CreateBaseUpdateTranslations creates a new base.update.translations model and returns its id.
func (c *Client) CreateBaseUpdateTranslations(but *BaseUpdateTranslations) (int64, error) {
	ids, err := c.Create(BaseUpdateTranslationsModel, []interface{}{but})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseUpdateTranslations creates a new base.update.translations model and returns its id.
func (c *Client) CreateBaseUpdateTranslationss(buts []*BaseUpdateTranslations) ([]int64, error) {
	var vv []interface{}
	for _, v := range buts {
		vv = append(vv, v)
	}
	return c.Create(BaseUpdateTranslationsModel, vv)
}

// UpdateBaseUpdateTranslations updates an existing base.update.translations record.
func (c *Client) UpdateBaseUpdateTranslations(but *BaseUpdateTranslations) error {
	return c.UpdateBaseUpdateTranslationss([]int64{but.Id.Get()}, but)
}

// UpdateBaseUpdateTranslationss updates existing base.update.translations records.
// All records (represented by ids) will be updated by but values.
func (c *Client) UpdateBaseUpdateTranslationss(ids []int64, but *BaseUpdateTranslations) error {
	return c.Update(BaseUpdateTranslationsModel, ids, but)
}

// DeleteBaseUpdateTranslations deletes an existing base.update.translations record.
func (c *Client) DeleteBaseUpdateTranslations(id int64) error {
	return c.DeleteBaseUpdateTranslationss([]int64{id})
}

// DeleteBaseUpdateTranslationss deletes existing base.update.translations records.
func (c *Client) DeleteBaseUpdateTranslationss(ids []int64) error {
	return c.Delete(BaseUpdateTranslationsModel, ids)
}

// GetBaseUpdateTranslations gets base.update.translations existing record.
func (c *Client) GetBaseUpdateTranslations(id int64) (*BaseUpdateTranslations, error) {
	buts, err := c.GetBaseUpdateTranslationss([]int64{id})
	if err != nil {
		return nil, err
	}
	if buts != nil && len(*buts) > 0 {
		return &((*buts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.update.translations not found", id)
}

// GetBaseUpdateTranslationss gets base.update.translations existing records.
func (c *Client) GetBaseUpdateTranslationss(ids []int64) (*BaseUpdateTranslationss, error) {
	buts := &BaseUpdateTranslationss{}
	if err := c.Read(BaseUpdateTranslationsModel, ids, nil, buts); err != nil {
		return nil, err
	}
	return buts, nil
}

// FindBaseUpdateTranslations finds base.update.translations record by querying it with criteria.
func (c *Client) FindBaseUpdateTranslations(criteria *Criteria) (*BaseUpdateTranslations, error) {
	buts := &BaseUpdateTranslationss{}
	if err := c.SearchRead(BaseUpdateTranslationsModel, criteria, NewOptions().Limit(1), buts); err != nil {
		return nil, err
	}
	if buts != nil && len(*buts) > 0 {
		return &((*buts)[0]), nil
	}
	return nil, fmt.Errorf("base.update.translations was not found with criteria %v", criteria)
}

// FindBaseUpdateTranslationss finds base.update.translations records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseUpdateTranslationss(criteria *Criteria, options *Options) (*BaseUpdateTranslationss, error) {
	buts := &BaseUpdateTranslationss{}
	if err := c.SearchRead(BaseUpdateTranslationsModel, criteria, options, buts); err != nil {
		return nil, err
	}
	return buts, nil
}

// FindBaseUpdateTranslationsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseUpdateTranslationsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseUpdateTranslationsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseUpdateTranslationsId finds record id by querying it with criteria.
func (c *Client) FindBaseUpdateTranslationsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseUpdateTranslationsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.update.translations was not found with criteria %v and options %v", criteria, options)
}
