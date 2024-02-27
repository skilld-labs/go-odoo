package odoo

// BaseUpdateTranslations represents base.update.translations model.
type BaseUpdateTranslations struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Lang        *Selection `xmlrpc:"lang,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateBaseUpdateTranslationss([]*BaseUpdateTranslations{but})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseUpdateTranslationss creates a new base.update.translations model and returns its id.
func (c *Client) CreateBaseUpdateTranslationss(buts []*BaseUpdateTranslations) ([]int64, error) {
	var vv []interface{}
	for _, v := range buts {
		vv = append(vv, v)
	}
	return c.Create(BaseUpdateTranslationsModel, vv, nil)
}

// UpdateBaseUpdateTranslations updates an existing base.update.translations record.
func (c *Client) UpdateBaseUpdateTranslations(but *BaseUpdateTranslations) error {
	return c.UpdateBaseUpdateTranslationss([]int64{but.Id.Get()}, but)
}

// UpdateBaseUpdateTranslationss updates existing base.update.translations records.
// All records (represented by ids) will be updated by but values.
func (c *Client) UpdateBaseUpdateTranslationss(ids []int64, but *BaseUpdateTranslations) error {
	return c.Update(BaseUpdateTranslationsModel, ids, but, nil)
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
	return &((*buts)[0]), nil
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
	return &((*buts)[0]), nil
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
	return c.Search(BaseUpdateTranslationsModel, criteria, options)
}

// FindBaseUpdateTranslationsId finds record id by querying it with criteria.
func (c *Client) FindBaseUpdateTranslationsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseUpdateTranslationsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
