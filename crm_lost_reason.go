package odoo

import (
	"fmt"
)

// CrmLostReason represents crm.lost.reason model.
type CrmLostReason struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmLostReasons represents array of crm.lost.reason model.
type CrmLostReasons []CrmLostReason

// CrmLostReasonModel is the odoo model name.
const CrmLostReasonModel = "crm.lost.reason"

// Many2One convert CrmLostReason to *Many2One.
func (clr *CrmLostReason) Many2One() *Many2One {
	return NewMany2One(clr.Id.Get(), "")
}

// CreateCrmLostReason creates a new crm.lost.reason model and returns its id.
func (c *Client) CreateCrmLostReason(clr *CrmLostReason) (int64, error) {
	return c.Create(CrmLostReasonModel, clr)
}

// UpdateCrmLostReason updates an existing crm.lost.reason record.
func (c *Client) UpdateCrmLostReason(clr *CrmLostReason) error {
	return c.UpdateCrmLostReasons([]int64{clr.Id.Get()}, clr)
}

// UpdateCrmLostReasons updates existing crm.lost.reason records.
// All records (represented by ids) will be updated by clr values.
func (c *Client) UpdateCrmLostReasons(ids []int64, clr *CrmLostReason) error {
	return c.Update(CrmLostReasonModel, ids, clr)
}

// DeleteCrmLostReason deletes an existing crm.lost.reason record.
func (c *Client) DeleteCrmLostReason(id int64) error {
	return c.DeleteCrmLostReasons([]int64{id})
}

// DeleteCrmLostReasons deletes existing crm.lost.reason records.
func (c *Client) DeleteCrmLostReasons(ids []int64) error {
	return c.Delete(CrmLostReasonModel, ids)
}

// GetCrmLostReason gets crm.lost.reason existing record.
func (c *Client) GetCrmLostReason(id int64) (*CrmLostReason, error) {
	clrs, err := c.GetCrmLostReasons([]int64{id})
	if err != nil {
		return nil, err
	}
	if clrs != nil && len(*clrs) > 0 {
		return &((*clrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.lost.reason not found", id)
}

// GetCrmLostReasons gets crm.lost.reason existing records.
func (c *Client) GetCrmLostReasons(ids []int64) (*CrmLostReasons, error) {
	clrs := &CrmLostReasons{}
	if err := c.Read(CrmLostReasonModel, ids, nil, clrs); err != nil {
		return nil, err
	}
	return clrs, nil
}

// FindCrmLostReason finds crm.lost.reason record by querying it with criteria.
func (c *Client) FindCrmLostReason(criteria *Criteria) (*CrmLostReason, error) {
	clrs := &CrmLostReasons{}
	if err := c.SearchRead(CrmLostReasonModel, criteria, NewOptions().Limit(1), clrs); err != nil {
		return nil, err
	}
	if clrs != nil && len(*clrs) > 0 {
		return &((*clrs)[0]), nil
	}
	return nil, fmt.Errorf("no crm.lost.reason was found with criteria %v", criteria)
}

// FindCrmLostReasons finds crm.lost.reason records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLostReasons(criteria *Criteria, options *Options) (*CrmLostReasons, error) {
	clrs := &CrmLostReasons{}
	if err := c.SearchRead(CrmLostReasonModel, criteria, options, clrs); err != nil {
		return nil, err
	}
	return clrs, nil
}

// FindCrmLostReasonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLostReasonIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmLostReasonModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmLostReasonId finds record id by querying it with criteria.
func (c *Client) FindCrmLostReasonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLostReasonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no crm.lost.reason was found with criteria %v and options %v", criteria, options)
}
