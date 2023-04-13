package odoo

import (
	"fmt"
)

// BasePartnerMergeLine represents base.partner.merge.line model.
type BasePartnerMergeLine struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	AggrIds     *String   `xmlrpc:"aggr_ids,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	MinId       *Int      `xmlrpc:"min_id,omptempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BasePartnerMergeLines represents array of base.partner.merge.line model.
type BasePartnerMergeLines []BasePartnerMergeLine

// BasePartnerMergeLineModel is the odoo model name.
const BasePartnerMergeLineModel = "base.partner.merge.line"

// Many2One convert BasePartnerMergeLine to *Many2One.
func (bpml *BasePartnerMergeLine) Many2One() *Many2One {
	return NewMany2One(bpml.Id.Get(), "")
}

// CreateBasePartnerMergeLine creates a new base.partner.merge.line model and returns its id.
func (c *Client) CreateBasePartnerMergeLine(bpml *BasePartnerMergeLine) (int64, error) {
	ids, err := c.Create(BasePartnerMergeLineModel, []interface{}{bpml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBasePartnerMergeLine creates a new base.partner.merge.line model and returns its id.
func (c *Client) CreateBasePartnerMergeLines(bpmls []*BasePartnerMergeLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range bpmls {
		vv = append(vv, v)
	}
	return c.Create(BasePartnerMergeLineModel, vv)
}

// UpdateBasePartnerMergeLine updates an existing base.partner.merge.line record.
func (c *Client) UpdateBasePartnerMergeLine(bpml *BasePartnerMergeLine) error {
	return c.UpdateBasePartnerMergeLines([]int64{bpml.Id.Get()}, bpml)
}

// UpdateBasePartnerMergeLines updates existing base.partner.merge.line records.
// All records (represented by ids) will be updated by bpml values.
func (c *Client) UpdateBasePartnerMergeLines(ids []int64, bpml *BasePartnerMergeLine) error {
	return c.Update(BasePartnerMergeLineModel, ids, bpml)
}

// DeleteBasePartnerMergeLine deletes an existing base.partner.merge.line record.
func (c *Client) DeleteBasePartnerMergeLine(id int64) error {
	return c.DeleteBasePartnerMergeLines([]int64{id})
}

// DeleteBasePartnerMergeLines deletes existing base.partner.merge.line records.
func (c *Client) DeleteBasePartnerMergeLines(ids []int64) error {
	return c.Delete(BasePartnerMergeLineModel, ids)
}

// GetBasePartnerMergeLine gets base.partner.merge.line existing record.
func (c *Client) GetBasePartnerMergeLine(id int64) (*BasePartnerMergeLine, error) {
	bpmls, err := c.GetBasePartnerMergeLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if bpmls != nil && len(*bpmls) > 0 {
		return &((*bpmls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.partner.merge.line not found", id)
}

// GetBasePartnerMergeLines gets base.partner.merge.line existing records.
func (c *Client) GetBasePartnerMergeLines(ids []int64) (*BasePartnerMergeLines, error) {
	bpmls := &BasePartnerMergeLines{}
	if err := c.Read(BasePartnerMergeLineModel, ids, nil, bpmls); err != nil {
		return nil, err
	}
	return bpmls, nil
}

// FindBasePartnerMergeLine finds base.partner.merge.line record by querying it with criteria.
func (c *Client) FindBasePartnerMergeLine(criteria *Criteria) (*BasePartnerMergeLine, error) {
	bpmls := &BasePartnerMergeLines{}
	if err := c.SearchRead(BasePartnerMergeLineModel, criteria, NewOptions().Limit(1), bpmls); err != nil {
		return nil, err
	}
	if bpmls != nil && len(*bpmls) > 0 {
		return &((*bpmls)[0]), nil
	}
	return nil, fmt.Errorf("base.partner.merge.line was not found with criteria %v", criteria)
}

// FindBasePartnerMergeLines finds base.partner.merge.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBasePartnerMergeLines(criteria *Criteria, options *Options) (*BasePartnerMergeLines, error) {
	bpmls := &BasePartnerMergeLines{}
	if err := c.SearchRead(BasePartnerMergeLineModel, criteria, options, bpmls); err != nil {
		return nil, err
	}
	return bpmls, nil
}

// FindBasePartnerMergeLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBasePartnerMergeLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BasePartnerMergeLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBasePartnerMergeLineId finds record id by querying it with criteria.
func (c *Client) FindBasePartnerMergeLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BasePartnerMergeLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.partner.merge.line was not found with criteria %v and options %v", criteria, options)
}
