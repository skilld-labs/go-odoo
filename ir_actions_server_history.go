package odoo

// IrActionsServerHistory represents ir.actions.server.history model.
type IrActionsServerHistory struct {
	ActionId    *Many2One `xmlrpc:"action_id,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrActionsServerHistorys represents array of ir.actions.server.history model.
type IrActionsServerHistorys []IrActionsServerHistory

// IrActionsServerHistoryModel is the odoo model name.
const IrActionsServerHistoryModel = "ir.actions.server.history"

// Many2One convert IrActionsServerHistory to *Many2One.
func (iash *IrActionsServerHistory) Many2One() *Many2One {
	return NewMany2One(iash.Id.Get(), "")
}

// CreateIrActionsServerHistory creates a new ir.actions.server.history model and returns its id.
func (c *Client) CreateIrActionsServerHistory(iash *IrActionsServerHistory) (int64, error) {
	ids, err := c.CreateIrActionsServerHistorys([]*IrActionsServerHistory{iash})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrActionsServerHistory creates a new ir.actions.server.history model and returns its id.
func (c *Client) CreateIrActionsServerHistorys(iashs []*IrActionsServerHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range iashs {
		vv = append(vv, v)
	}
	return c.Create(IrActionsServerHistoryModel, vv, nil)
}

// UpdateIrActionsServerHistory updates an existing ir.actions.server.history record.
func (c *Client) UpdateIrActionsServerHistory(iash *IrActionsServerHistory) error {
	return c.UpdateIrActionsServerHistorys([]int64{iash.Id.Get()}, iash)
}

// UpdateIrActionsServerHistorys updates existing ir.actions.server.history records.
// All records (represented by ids) will be updated by iash values.
func (c *Client) UpdateIrActionsServerHistorys(ids []int64, iash *IrActionsServerHistory) error {
	return c.Update(IrActionsServerHistoryModel, ids, iash, nil)
}

// DeleteIrActionsServerHistory deletes an existing ir.actions.server.history record.
func (c *Client) DeleteIrActionsServerHistory(id int64) error {
	return c.DeleteIrActionsServerHistorys([]int64{id})
}

// DeleteIrActionsServerHistorys deletes existing ir.actions.server.history records.
func (c *Client) DeleteIrActionsServerHistorys(ids []int64) error {
	return c.Delete(IrActionsServerHistoryModel, ids)
}

// GetIrActionsServerHistory gets ir.actions.server.history existing record.
func (c *Client) GetIrActionsServerHistory(id int64) (*IrActionsServerHistory, error) {
	iashs, err := c.GetIrActionsServerHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iashs)[0]), nil
}

// GetIrActionsServerHistorys gets ir.actions.server.history existing records.
func (c *Client) GetIrActionsServerHistorys(ids []int64) (*IrActionsServerHistorys, error) {
	iashs := &IrActionsServerHistorys{}
	if err := c.Read(IrActionsServerHistoryModel, ids, nil, iashs); err != nil {
		return nil, err
	}
	return iashs, nil
}

// FindIrActionsServerHistory finds ir.actions.server.history record by querying it with criteria.
func (c *Client) FindIrActionsServerHistory(criteria *Criteria) (*IrActionsServerHistory, error) {
	iashs := &IrActionsServerHistorys{}
	if err := c.SearchRead(IrActionsServerHistoryModel, criteria, NewOptions().Limit(1), iashs); err != nil {
		return nil, err
	}
	return &((*iashs)[0]), nil
}

// FindIrActionsServerHistorys finds ir.actions.server.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsServerHistorys(criteria *Criteria, options *Options) (*IrActionsServerHistorys, error) {
	iashs := &IrActionsServerHistorys{}
	if err := c.SearchRead(IrActionsServerHistoryModel, criteria, options, iashs); err != nil {
		return nil, err
	}
	return iashs, nil
}

// FindIrActionsServerHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsServerHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrActionsServerHistoryModel, criteria, options)
}

// FindIrActionsServerHistoryId finds record id by querying it with criteria.
func (c *Client) FindIrActionsServerHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsServerHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
