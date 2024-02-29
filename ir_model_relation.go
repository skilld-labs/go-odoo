package odoo

// IrModelRelation represents ir.model.relation model.
type IrModelRelation struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DateInit    *Time     `xmlrpc:"date_init,omitempty"`
	DateUpdate  *Time     `xmlrpc:"date_update,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Model       *Many2One `xmlrpc:"model,omitempty"`
	Module      *Many2One `xmlrpc:"module,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrModelRelations represents array of ir.model.relation model.
type IrModelRelations []IrModelRelation

// IrModelRelationModel is the odoo model name.
const IrModelRelationModel = "ir.model.relation"

// Many2One convert IrModelRelation to *Many2One.
func (imr *IrModelRelation) Many2One() *Many2One {
	return NewMany2One(imr.Id.Get(), "")
}

// CreateIrModelRelation creates a new ir.model.relation model and returns its id.
func (c *Client) CreateIrModelRelation(imr *IrModelRelation) (int64, error) {
	ids, err := c.CreateIrModelRelations([]*IrModelRelation{imr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModelRelations creates a new ir.model.relation model and returns its id.
func (c *Client) CreateIrModelRelations(imrs []*IrModelRelation) ([]int64, error) {
	var vv []interface{}
	for _, v := range imrs {
		vv = append(vv, v)
	}
	return c.Create(IrModelRelationModel, vv, nil)
}

// UpdateIrModelRelation updates an existing ir.model.relation record.
func (c *Client) UpdateIrModelRelation(imr *IrModelRelation) error {
	return c.UpdateIrModelRelations([]int64{imr.Id.Get()}, imr)
}

// UpdateIrModelRelations updates existing ir.model.relation records.
// All records (represented by ids) will be updated by imr values.
func (c *Client) UpdateIrModelRelations(ids []int64, imr *IrModelRelation) error {
	return c.Update(IrModelRelationModel, ids, imr, nil)
}

// DeleteIrModelRelation deletes an existing ir.model.relation record.
func (c *Client) DeleteIrModelRelation(id int64) error {
	return c.DeleteIrModelRelations([]int64{id})
}

// DeleteIrModelRelations deletes existing ir.model.relation records.
func (c *Client) DeleteIrModelRelations(ids []int64) error {
	return c.Delete(IrModelRelationModel, ids)
}

// GetIrModelRelation gets ir.model.relation existing record.
func (c *Client) GetIrModelRelation(id int64) (*IrModelRelation, error) {
	imrs, err := c.GetIrModelRelations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imrs)[0]), nil
}

// GetIrModelRelations gets ir.model.relation existing records.
func (c *Client) GetIrModelRelations(ids []int64) (*IrModelRelations, error) {
	imrs := &IrModelRelations{}
	if err := c.Read(IrModelRelationModel, ids, nil, imrs); err != nil {
		return nil, err
	}
	return imrs, nil
}

// FindIrModelRelation finds ir.model.relation record by querying it with criteria.
func (c *Client) FindIrModelRelation(criteria *Criteria) (*IrModelRelation, error) {
	imrs := &IrModelRelations{}
	if err := c.SearchRead(IrModelRelationModel, criteria, NewOptions().Limit(1), imrs); err != nil {
		return nil, err
	}
	return &((*imrs)[0]), nil
}

// FindIrModelRelations finds ir.model.relation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelRelations(criteria *Criteria, options *Options) (*IrModelRelations, error) {
	imrs := &IrModelRelations{}
	if err := c.SearchRead(IrModelRelationModel, criteria, options, imrs); err != nil {
		return nil, err
	}
	return imrs, nil
}

// FindIrModelRelationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelRelationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrModelRelationModel, criteria, options)
}

// FindIrModelRelationId finds record id by querying it with criteria.
func (c *Client) FindIrModelRelationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelRelationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
