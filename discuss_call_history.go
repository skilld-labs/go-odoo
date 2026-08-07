package odoo

// DiscussCallHistory represents discuss.call.history model.
type DiscussCallHistory struct {
	ChannelId                     *Many2One `xmlrpc:"channel_id,omitempty"`
	CreateDate                    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName                   *String   `xmlrpc:"display_name,omitempty"`
	DurationHour                  *Float    `xmlrpc:"duration_hour,omitempty"`
	EndDt                         *Time     `xmlrpc:"end_dt,omitempty"`
	Id                            *Int      `xmlrpc:"id,omitempty"`
	LivechatParticipantHistoryIds *Relation `xmlrpc:"livechat_participant_history_ids,omitempty"`
	StartCallMessageId            *Many2One `xmlrpc:"start_call_message_id,omitempty"`
	StartDt                       *Time     `xmlrpc:"start_dt,omitempty"`
	WriteDate                     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DiscussCallHistorys represents array of discuss.call.history model.
type DiscussCallHistorys []DiscussCallHistory

// DiscussCallHistoryModel is the odoo model name.
const DiscussCallHistoryModel = "discuss.call.history"

// Many2One convert DiscussCallHistory to *Many2One.
func (dch *DiscussCallHistory) Many2One() *Many2One {
	return NewMany2One(dch.Id.Get(), "")
}

// CreateDiscussCallHistory creates a new discuss.call.history model and returns its id.
func (c *Client) CreateDiscussCallHistory(dch *DiscussCallHistory) (int64, error) {
	ids, err := c.CreateDiscussCallHistorys([]*DiscussCallHistory{dch})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDiscussCallHistory creates a new discuss.call.history model and returns its id.
func (c *Client) CreateDiscussCallHistorys(dchs []*DiscussCallHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range dchs {
		vv = append(vv, v)
	}
	return c.Create(DiscussCallHistoryModel, vv, nil)
}

// UpdateDiscussCallHistory updates an existing discuss.call.history record.
func (c *Client) UpdateDiscussCallHistory(dch *DiscussCallHistory) error {
	return c.UpdateDiscussCallHistorys([]int64{dch.Id.Get()}, dch)
}

// UpdateDiscussCallHistorys updates existing discuss.call.history records.
// All records (represented by ids) will be updated by dch values.
func (c *Client) UpdateDiscussCallHistorys(ids []int64, dch *DiscussCallHistory) error {
	return c.Update(DiscussCallHistoryModel, ids, dch, nil)
}

// DeleteDiscussCallHistory deletes an existing discuss.call.history record.
func (c *Client) DeleteDiscussCallHistory(id int64) error {
	return c.DeleteDiscussCallHistorys([]int64{id})
}

// DeleteDiscussCallHistorys deletes existing discuss.call.history records.
func (c *Client) DeleteDiscussCallHistorys(ids []int64) error {
	return c.Delete(DiscussCallHistoryModel, ids)
}

// GetDiscussCallHistory gets discuss.call.history existing record.
func (c *Client) GetDiscussCallHistory(id int64) (*DiscussCallHistory, error) {
	dchs, err := c.GetDiscussCallHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dchs)[0]), nil
}

// GetDiscussCallHistorys gets discuss.call.history existing records.
func (c *Client) GetDiscussCallHistorys(ids []int64) (*DiscussCallHistorys, error) {
	dchs := &DiscussCallHistorys{}
	if err := c.Read(DiscussCallHistoryModel, ids, nil, dchs); err != nil {
		return nil, err
	}
	return dchs, nil
}

// FindDiscussCallHistory finds discuss.call.history record by querying it with criteria.
func (c *Client) FindDiscussCallHistory(criteria *Criteria) (*DiscussCallHistory, error) {
	dchs := &DiscussCallHistorys{}
	if err := c.SearchRead(DiscussCallHistoryModel, criteria, NewOptions().Limit(1), dchs); err != nil {
		return nil, err
	}
	return &((*dchs)[0]), nil
}

// FindDiscussCallHistorys finds discuss.call.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussCallHistorys(criteria *Criteria, options *Options) (*DiscussCallHistorys, error) {
	dchs := &DiscussCallHistorys{}
	if err := c.SearchRead(DiscussCallHistoryModel, criteria, options, dchs); err != nil {
		return nil, err
	}
	return dchs, nil
}

// FindDiscussCallHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussCallHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DiscussCallHistoryModel, criteria, options)
}

// FindDiscussCallHistoryId finds record id by querying it with criteria.
func (c *Client) FindDiscussCallHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DiscussCallHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
