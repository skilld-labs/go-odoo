package odoo

import (
	"fmt"
)

// FormatAddressMixin represents format.address.mixin model.
type FormatAddressMixin struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// FormatAddressMixins represents array of format.address.mixin model.
type FormatAddressMixins []FormatAddressMixin

// FormatAddressMixinModel is the odoo model name.
const FormatAddressMixinModel = "format.address.mixin"

// Many2One convert FormatAddressMixin to *Many2One.
func (fam *FormatAddressMixin) Many2One() *Many2One {
	return NewMany2One(fam.Id.Get(), "")
}

// CreateFormatAddressMixin creates a new format.address.mixin model and returns its id.
func (c *Client) CreateFormatAddressMixin(fam *FormatAddressMixin) (int64, error) {
	return c.Create(FormatAddressMixinModel, fam)
}

// UpdateFormatAddressMixin updates an existing format.address.mixin record.
func (c *Client) UpdateFormatAddressMixin(fam *FormatAddressMixin) error {
	return c.UpdateFormatAddressMixins([]int64{fam.Id.Get()}, fam)
}

// UpdateFormatAddressMixins updates existing format.address.mixin records.
// All records (represented by IDs) will be updated by fam values.
func (c *Client) UpdateFormatAddressMixins(ids []int64, fam *FormatAddressMixin) error {
	return c.Update(FormatAddressMixinModel, ids, fam)
}

// DeleteFormatAddressMixin deletes an existing format.address.mixin record.
func (c *Client) DeleteFormatAddressMixin(id int64) error {
	return c.DeleteFormatAddressMixins([]int64{id})
}

// DeleteFormatAddressMixins deletes existing format.address.mixin records.
func (c *Client) DeleteFormatAddressMixins(ids []int64) error {
	return c.Delete(FormatAddressMixinModel, ids)
}

// GetFormatAddressMixin gets format.address.mixin existing record.
func (c *Client) GetFormatAddressMixin(id int64) (*FormatAddressMixin, error) {
	fams, err := c.GetFormatAddressMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if fams != nil && len(*fams) > 0 {
		return &((*fams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of format.address.mixin not found", id)
}

// GetFormatAddressMixins gets format.address.mixin existing records.
func (c *Client) GetFormatAddressMixins(ids []int64) (*FormatAddressMixins, error) {
	fams := &FormatAddressMixins{}
	if err := c.Read(FormatAddressMixinModel, ids, nil, fams); err != nil {
		return nil, err
	}
	return fams, nil
}

// FindFormatAddressMixin finds format.address.mixin record by querying it with criteria.
func (c *Client) FindFormatAddressMixin(criteria *Criteria) (*FormatAddressMixin, error) {
	fams := &FormatAddressMixins{}
	if err := c.SearchRead(FormatAddressMixinModel, criteria, NewOptions().Limit(1), fams); err != nil {
		return nil, err
	}
	if fams != nil && len(*fams) > 0 {
		return &((*fams)[0]), nil
	}
	return nil, fmt.Errorf("format.address.mixin was not found")
}

// FindFormatAddressMixins finds format.address.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindFormatAddressMixins(criteria *Criteria, options *Options) (*FormatAddressMixins, error) {
	fams := &FormatAddressMixins{}
	if err := c.SearchRead(FormatAddressMixinModel, criteria, options, fams); err != nil {
		return nil, err
	}
	return fams, nil
}

// FindFormatAddressMixinIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindFormatAddressMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(FormatAddressMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindFormatAddressMixinId finds record id by querying it with criteria.
func (c *Client) FindFormatAddressMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(FormatAddressMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("format.address.mixin was not found")
}
