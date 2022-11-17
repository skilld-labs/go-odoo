package odoo

import (
	"fmt"
)

// ProcurementGroup represents procurement.group model.
type ProcurementGroup struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	MoveType    *Selection `xmlrpc:"move_type,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	PartnerId   *Many2One  `xmlrpc:"partner_id,omitempty"`
	SaleId      *Many2One  `xmlrpc:"sale_id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProcurementGroups represents array of procurement.group model.
type ProcurementGroups []ProcurementGroup

// ProcurementGroupModel is the odoo model name.
const ProcurementGroupModel = "procurement.group"

// Many2One convert ProcurementGroup to *Many2One.
func (pg *ProcurementGroup) Many2One() *Many2One {
	return NewMany2One(pg.Id.Get(), "")
}

// CreateProcurementGroup creates a new procurement.group model and returns its id.
func (c *Client) CreateProcurementGroup(pg *ProcurementGroup) (int64, error) {
	return c.Create(ProcurementGroupModel, pg)
}

// UpdateProcurementGroup updates an existing procurement.group record.
func (c *Client) UpdateProcurementGroup(pg *ProcurementGroup) error {
	return c.UpdateProcurementGroups([]int64{pg.Id.Get()}, pg)
}

// UpdateProcurementGroups updates existing procurement.group records.
// All records (represented by ids) will be updated by pg values.
func (c *Client) UpdateProcurementGroups(ids []int64, pg *ProcurementGroup) error {
	return c.Update(ProcurementGroupModel, ids, pg)
}

// DeleteProcurementGroup deletes an existing procurement.group record.
func (c *Client) DeleteProcurementGroup(id int64) error {
	return c.DeleteProcurementGroups([]int64{id})
}

// DeleteProcurementGroups deletes existing procurement.group records.
func (c *Client) DeleteProcurementGroups(ids []int64) error {
	return c.Delete(ProcurementGroupModel, ids)
}

// GetProcurementGroup gets procurement.group existing record.
func (c *Client) GetProcurementGroup(id int64) (*ProcurementGroup, error) {
	pgs, err := c.GetProcurementGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if pgs != nil && len(*pgs) > 0 {
		return &((*pgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of procurement.group not found", id)
}

// GetProcurementGroups gets procurement.group existing records.
func (c *Client) GetProcurementGroups(ids []int64) (*ProcurementGroups, error) {
	pgs := &ProcurementGroups{}
	if err := c.Read(ProcurementGroupModel, ids, nil, pgs); err != nil {
		return nil, err
	}
	return pgs, nil
}

// FindProcurementGroup finds procurement.group record by querying it with criteria.
func (c *Client) FindProcurementGroup(criteria *Criteria) (*ProcurementGroup, error) {
	pgs := &ProcurementGroups{}
	if err := c.SearchRead(ProcurementGroupModel, criteria, NewOptions().Limit(1), pgs); err != nil {
		return nil, err
	}
	if pgs != nil && len(*pgs) > 0 {
		return &((*pgs)[0]), nil
	}
	return nil, fmt.Errorf("no procurement.group was found with criteria %v", criteria)
}

// FindProcurementGroups finds procurement.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProcurementGroups(criteria *Criteria, options *Options) (*ProcurementGroups, error) {
	pgs := &ProcurementGroups{}
	if err := c.SearchRead(ProcurementGroupModel, criteria, options, pgs); err != nil {
		return nil, err
	}
	return pgs, nil
}

// FindProcurementGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProcurementGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProcurementGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProcurementGroupId finds record id by querying it with criteria.
func (c *Client) FindProcurementGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProcurementGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no procurement.group was found with criteria %v and options %v", criteria, options)
}
