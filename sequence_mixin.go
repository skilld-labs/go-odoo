package odoo

import (
	"fmt"
)

// SequenceMixin represents sequence.mixin model.
type SequenceMixin struct {
	LastUpdate     *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName    *String `xmlrpc:"display_name,omitempty"`
	Id             *Int    `xmlrpc:"id,omitempty"`
	SequenceNumber *Int    `xmlrpc:"sequence_number,omitempty"`
	SequencePrefix *String `xmlrpc:"sequence_prefix,omitempty"`
}

// SequenceMixins represents array of sequence.mixin model.
type SequenceMixins []SequenceMixin

// SequenceMixinModel is the odoo model name.
const SequenceMixinModel = "sequence.mixin"

// Many2One convert SequenceMixin to *Many2One.
func (sm *SequenceMixin) Many2One() *Many2One {
	return NewMany2One(sm.Id.Get(), "")
}

// CreateSequenceMixin creates a new sequence.mixin model and returns its id.
func (c *Client) CreateSequenceMixin(sm *SequenceMixin) (int64, error) {
	return c.Create(SequenceMixinModel, sm)
}

// UpdateSequenceMixin updates an existing sequence.mixin record.
func (c *Client) UpdateSequenceMixin(sm *SequenceMixin) error {
	return c.UpdateSequenceMixins([]int64{sm.Id.Get()}, sm)
}

// UpdateSequenceMixins updates existing sequence.mixin records.
// All records (represented by IDs) will be updated by sm values.
func (c *Client) UpdateSequenceMixins(ids []int64, sm *SequenceMixin) error {
	return c.Update(SequenceMixinModel, ids, sm)
}

// DeleteSequenceMixin deletes an existing sequence.mixin record.
func (c *Client) DeleteSequenceMixin(id int64) error {
	return c.DeleteSequenceMixins([]int64{id})
}

// DeleteSequenceMixins deletes existing sequence.mixin records.
func (c *Client) DeleteSequenceMixins(ids []int64) error {
	return c.Delete(SequenceMixinModel, ids)
}

// GetSequenceMixin gets sequence.mixin existing record.
func (c *Client) GetSequenceMixin(id int64) (*SequenceMixin, error) {
	sms, err := c.GetSequenceMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if sms != nil && len(*sms) > 0 {
		return &((*sms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sequence.mixin not found", id)
}

// GetSequenceMixins gets sequence.mixin existing records.
func (c *Client) GetSequenceMixins(ids []int64) (*SequenceMixins, error) {
	sms := &SequenceMixins{}
	if err := c.Read(SequenceMixinModel, ids, nil, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSequenceMixin finds sequence.mixin record by querying it with criteria.
func (c *Client) FindSequenceMixin(criteria *Criteria) (*SequenceMixin, error) {
	sms := &SequenceMixins{}
	if err := c.SearchRead(SequenceMixinModel, criteria, NewOptions().Limit(1), sms); err != nil {
		return nil, err
	}
	if sms != nil && len(*sms) > 0 {
		return &((*sms)[0]), nil
	}
	return nil, fmt.Errorf("sequence.mixin was not found")
}

// FindSequenceMixins finds sequence.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSequenceMixins(criteria *Criteria, options *Options) (*SequenceMixins, error) {
	sms := &SequenceMixins{}
	if err := c.SearchRead(SequenceMixinModel, criteria, options, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSequenceMixinIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSequenceMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SequenceMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSequenceMixinId finds record id by querying it with criteria.
func (c *Client) FindSequenceMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SequenceMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sequence.mixin was not found")
}
