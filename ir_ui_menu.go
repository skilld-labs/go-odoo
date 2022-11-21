package odoo

import (
	"fmt"
)

// IrUiMenu represents ir.ui.menu model.
type IrUiMenu struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	Action       *String   `xmlrpc:"action,omptempty"`
	Active       *Bool     `xmlrpc:"active,omptempty"`
	ChildId      *Relation `xmlrpc:"child_id,omptempty"`
	CompleteName *String   `xmlrpc:"complete_name,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	GroupsId     *Relation `xmlrpc:"groups_id,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	ParentId     *Many2One `xmlrpc:"parent_id,omptempty"`
	ParentLeft   *Int      `xmlrpc:"parent_left,omptempty"`
	ParentRight  *Int      `xmlrpc:"parent_right,omptempty"`
	Sequence     *Int      `xmlrpc:"sequence,omptempty"`
	WebIcon      *String   `xmlrpc:"web_icon,omptempty"`
	WebIconData  *String   `xmlrpc:"web_icon_data,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// IrUiMenus represents array of ir.ui.menu model.
type IrUiMenus []IrUiMenu

// IrUiMenuModel is the odoo model name.
const IrUiMenuModel = "ir.ui.menu"

// Many2One convert IrUiMenu to *Many2One.
func (ium *IrUiMenu) Many2One() *Many2One {
	return NewMany2One(ium.Id.Get(), "")
}

// CreateIrUiMenu creates a new ir.ui.menu model and returns its id.
func (c *Client) CreateIrUiMenu(ium *IrUiMenu) (int64, error) {
	return c.Create(IrUiMenuModel, ium)
}

// UpdateIrUiMenu updates an existing ir.ui.menu record.
func (c *Client) UpdateIrUiMenu(ium *IrUiMenu) error {
	return c.UpdateIrUiMenus([]int64{ium.Id.Get()}, ium)
}

// UpdateIrUiMenus updates existing ir.ui.menu records.
// All records (represented by ids) will be updated by ium values.
func (c *Client) UpdateIrUiMenus(ids []int64, ium *IrUiMenu) error {
	return c.Update(IrUiMenuModel, ids, ium)
}

// DeleteIrUiMenu deletes an existing ir.ui.menu record.
func (c *Client) DeleteIrUiMenu(id int64) error {
	return c.DeleteIrUiMenus([]int64{id})
}

// DeleteIrUiMenus deletes existing ir.ui.menu records.
func (c *Client) DeleteIrUiMenus(ids []int64) error {
	return c.Delete(IrUiMenuModel, ids)
}

// GetIrUiMenu gets ir.ui.menu existing record.
func (c *Client) GetIrUiMenu(id int64) (*IrUiMenu, error) {
	iums, err := c.GetIrUiMenus([]int64{id})
	if err != nil {
		return nil, err
	}
	if iums != nil && len(*iums) > 0 {
		return &((*iums)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.ui.menu not found", id)
}

// GetIrUiMenus gets ir.ui.menu existing records.
func (c *Client) GetIrUiMenus(ids []int64) (*IrUiMenus, error) {
	iums := &IrUiMenus{}
	if err := c.Read(IrUiMenuModel, ids, nil, iums); err != nil {
		return nil, err
	}
	return iums, nil
}

// FindIrUiMenu finds ir.ui.menu record by querying it with criteria.
func (c *Client) FindIrUiMenu(criteria *Criteria) (*IrUiMenu, error) {
	iums := &IrUiMenus{}
	if err := c.SearchRead(IrUiMenuModel, criteria, NewOptions().Limit(1), iums); err != nil {
		return nil, err
	}
	if iums != nil && len(*iums) > 0 {
		return &((*iums)[0]), nil
	}
	return nil, fmt.Errorf("no ir.ui.menu was found with criteria %v", criteria)
}

// FindIrUiMenus finds ir.ui.menu records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiMenus(criteria *Criteria, options *Options) (*IrUiMenus, error) {
	iums := &IrUiMenus{}
	if err := c.SearchRead(IrUiMenuModel, criteria, options, iums); err != nil {
		return nil, err
	}
	return iums, nil
}

// FindIrUiMenuIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrUiMenuIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrUiMenuModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrUiMenuId finds record id by querying it with criteria.
func (c *Client) FindIrUiMenuId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrUiMenuModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no ir.ui.menu was found with criteria %v and options %v", criteria, options)
}
