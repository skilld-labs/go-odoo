package odoo

import (
	"fmt"
)

// PhoneValidationMixin represents phone.validation.mixin model.
type PhoneValidationMixin struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// PhoneValidationMixins represents array of phone.validation.mixin model.
type PhoneValidationMixins []PhoneValidationMixin

// PhoneValidationMixinModel is the odoo model name.
const PhoneValidationMixinModel = "phone.validation.mixin"

// Many2One convert PhoneValidationMixin to *Many2One.
func (pvm *PhoneValidationMixin) Many2One() *Many2One {
	return NewMany2One(pvm.Id.Get(), "")
}

// CreatePhoneValidationMixin creates a new phone.validation.mixin model and returns its id.
func (c *Client) CreatePhoneValidationMixin(pvm *PhoneValidationMixin) (int64, error) {
	return c.Create(PhoneValidationMixinModel, pvm)
}

// UpdatePhoneValidationMixin updates an existing phone.validation.mixin record.
func (c *Client) UpdatePhoneValidationMixin(pvm *PhoneValidationMixin) error {
	return c.UpdatePhoneValidationMixins([]int64{pvm.Id.Get()}, pvm)
}

// UpdatePhoneValidationMixins updates existing phone.validation.mixin records.
// All records (represented by IDs) will be updated by pvm values.
func (c *Client) UpdatePhoneValidationMixins(ids []int64, pvm *PhoneValidationMixin) error {
	return c.Update(PhoneValidationMixinModel, ids, pvm)
}

// DeletePhoneValidationMixin deletes an existing phone.validation.mixin record.
func (c *Client) DeletePhoneValidationMixin(id int64) error {
	return c.DeletePhoneValidationMixins([]int64{id})
}

// DeletePhoneValidationMixins deletes existing phone.validation.mixin records.
func (c *Client) DeletePhoneValidationMixins(ids []int64) error {
	return c.Delete(PhoneValidationMixinModel, ids)
}

// GetPhoneValidationMixin gets phone.validation.mixin existing record.
func (c *Client) GetPhoneValidationMixin(id int64) (*PhoneValidationMixin, error) {
	pvms, err := c.GetPhoneValidationMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if pvms != nil && len(*pvms) > 0 {
		return &((*pvms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of phone.validation.mixin not found", id)
}

// GetPhoneValidationMixins gets phone.validation.mixin existing records.
func (c *Client) GetPhoneValidationMixins(ids []int64) (*PhoneValidationMixins, error) {
	pvms := &PhoneValidationMixins{}
	if err := c.Read(PhoneValidationMixinModel, ids, nil, pvms); err != nil {
		return nil, err
	}
	return pvms, nil
}

// FindPhoneValidationMixin finds phone.validation.mixin record by querying it with criteria.
func (c *Client) FindPhoneValidationMixin(criteria *Criteria) (*PhoneValidationMixin, error) {
	pvms := &PhoneValidationMixins{}
	if err := c.SearchRead(PhoneValidationMixinModel, criteria, NewOptions().Limit(1), pvms); err != nil {
		return nil, err
	}
	if pvms != nil && len(*pvms) > 0 {
		return &((*pvms)[0]), nil
	}
	return nil, fmt.Errorf("phone.validation.mixin was not found")
}

// FindPhoneValidationMixins finds phone.validation.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPhoneValidationMixins(criteria *Criteria, options *Options) (*PhoneValidationMixins, error) {
	pvms := &PhoneValidationMixins{}
	if err := c.SearchRead(PhoneValidationMixinModel, criteria, options, pvms); err != nil {
		return nil, err
	}
	return pvms, nil
}

// FindPhoneValidationMixinIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindPhoneValidationMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PhoneValidationMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPhoneValidationMixinId finds record id by querying it with criteria.
func (c *Client) FindPhoneValidationMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PhoneValidationMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("phone.validation.mixin was not found")
}
