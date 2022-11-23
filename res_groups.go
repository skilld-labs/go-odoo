package odoo

import (
	"fmt"
)

// ResGroups represents res.groups model.
type ResGroups struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CategoryId      *Many2One `xmlrpc:"category_id,omptempty"`
	Color           *Int      `xmlrpc:"color,omptempty"`
	Comment         *String   `xmlrpc:"comment,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	FullName        *String   `xmlrpc:"full_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	ImpliedIds      *Relation `xmlrpc:"implied_ids,omptempty"`
	IsPortal        *Bool     `xmlrpc:"is_portal,omptempty"`
	MenuAccess      *Relation `xmlrpc:"menu_access,omptempty"`
	ModelAccess     *Relation `xmlrpc:"model_access,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	RuleGroups      *Relation `xmlrpc:"rule_groups,omptempty"`
	Share           *Bool     `xmlrpc:"share,omptempty"`
	TransImpliedIds *Relation `xmlrpc:"trans_implied_ids,omptempty"`
	Users           *Relation `xmlrpc:"users,omptempty"`
	ViewAccess      *Relation `xmlrpc:"view_access,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResGroupss represents array of res.groups model.
type ResGroupss []ResGroups

// ResGroupsModel is the odoo model name.
const ResGroupsModel = "res.groups"

// Many2One convert ResGroups to *Many2One.
func (rg *ResGroups) Many2One() *Many2One {
	return NewMany2One(rg.Id.Get(), "")
}

// CreateResGroups creates a new res.groups model and returns its id.
func (c *Client) CreateResGroups(rg *ResGroups) (int64, error) {
	return c.Create(ResGroupsModel, rg)
}

// UpdateResGroups updates an existing res.groups record.
func (c *Client) UpdateResGroups(rg *ResGroups) error {
	return c.UpdateResGroupss([]int64{rg.Id.Get()}, rg)
}

// UpdateResGroupss updates existing res.groups records.
// All records (represented by ids) will be updated by rg values.
func (c *Client) UpdateResGroupss(ids []int64, rg *ResGroups) error {
	return c.Update(ResGroupsModel, ids, rg)
}

// DeleteResGroups deletes an existing res.groups record.
func (c *Client) DeleteResGroups(id int64) error {
	return c.DeleteResGroupss([]int64{id})
}

// DeleteResGroupss deletes existing res.groups records.
func (c *Client) DeleteResGroupss(ids []int64) error {
	return c.Delete(ResGroupsModel, ids)
}

// GetResGroups gets res.groups existing record.
func (c *Client) GetResGroups(id int64) (*ResGroups, error) {
	rgs, err := c.GetResGroupss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rgs != nil && len(*rgs) > 0 {
		return &((*rgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.groups not found", id)
}

// GetResGroupss gets res.groups existing records.
func (c *Client) GetResGroupss(ids []int64) (*ResGroupss, error) {
	rgs := &ResGroupss{}
	if err := c.Read(ResGroupsModel, ids, nil, rgs); err != nil {
		return nil, err
	}
	return rgs, nil
}

// FindResGroups finds res.groups record by querying it with criteria.
func (c *Client) FindResGroups(criteria *Criteria) (*ResGroups, error) {
	rgs := &ResGroupss{}
	if err := c.SearchRead(ResGroupsModel, criteria, NewOptions().Limit(1), rgs); err != nil {
		return nil, err
	}
	if rgs != nil && len(*rgs) > 0 {
		return &((*rgs)[0]), nil
	}
	return nil, fmt.Errorf("res.groups was not found with criteria %v", criteria)
}

// FindResGroupss finds res.groups records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResGroupss(criteria *Criteria, options *Options) (*ResGroupss, error) {
	rgs := &ResGroupss{}
	if err := c.SearchRead(ResGroupsModel, criteria, options, rgs); err != nil {
		return nil, err
	}
	return rgs, nil
}

// FindResGroupsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResGroupsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResGroupsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResGroupsId finds record id by querying it with criteria.
func (c *Client) FindResGroupsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResGroupsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.groups was not found with criteria %v and options %v", criteria, options)
}
