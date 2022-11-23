package odoo

import (
	"fmt"
)

// QueueJobChannel represents queue.job.channel model.
type QueueJobChannel struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CompleteName   *String   `xmlrpc:"complete_name,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	JobFunctionIds *Relation `xmlrpc:"job_function_ids,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	ParentId       *Many2One `xmlrpc:"parent_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// QueueJobChannels represents array of queue.job.channel model.
type QueueJobChannels []QueueJobChannel

// QueueJobChannelModel is the odoo model name.
const QueueJobChannelModel = "queue.job.channel"

// Many2One convert QueueJobChannel to *Many2One.
func (qjc *QueueJobChannel) Many2One() *Many2One {
	return NewMany2One(qjc.Id.Get(), "")
}

// CreateQueueJobChannel creates a new queue.job.channel model and returns its id.
func (c *Client) CreateQueueJobChannel(qjc *QueueJobChannel) (int64, error) {
	return c.Create(QueueJobChannelModel, qjc)
}

// UpdateQueueJobChannel updates an existing queue.job.channel record.
func (c *Client) UpdateQueueJobChannel(qjc *QueueJobChannel) error {
	return c.UpdateQueueJobChannels([]int64{qjc.Id.Get()}, qjc)
}

// UpdateQueueJobChannels updates existing queue.job.channel records.
// All records (represented by ids) will be updated by qjc values.
func (c *Client) UpdateQueueJobChannels(ids []int64, qjc *QueueJobChannel) error {
	return c.Update(QueueJobChannelModel, ids, qjc)
}

// DeleteQueueJobChannel deletes an existing queue.job.channel record.
func (c *Client) DeleteQueueJobChannel(id int64) error {
	return c.DeleteQueueJobChannels([]int64{id})
}

// DeleteQueueJobChannels deletes existing queue.job.channel records.
func (c *Client) DeleteQueueJobChannels(ids []int64) error {
	return c.Delete(QueueJobChannelModel, ids)
}

// GetQueueJobChannel gets queue.job.channel existing record.
func (c *Client) GetQueueJobChannel(id int64) (*QueueJobChannel, error) {
	qjcs, err := c.GetQueueJobChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if qjcs != nil && len(*qjcs) > 0 {
		return &((*qjcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue.job.channel not found", id)
}

// GetQueueJobChannels gets queue.job.channel existing records.
func (c *Client) GetQueueJobChannels(ids []int64) (*QueueJobChannels, error) {
	qjcs := &QueueJobChannels{}
	if err := c.Read(QueueJobChannelModel, ids, nil, qjcs); err != nil {
		return nil, err
	}
	return qjcs, nil
}

// FindQueueJobChannel finds queue.job.channel record by querying it with criteria.
func (c *Client) FindQueueJobChannel(criteria *Criteria) (*QueueJobChannel, error) {
	qjcs := &QueueJobChannels{}
	if err := c.SearchRead(QueueJobChannelModel, criteria, NewOptions().Limit(1), qjcs); err != nil {
		return nil, err
	}
	if qjcs != nil && len(*qjcs) > 0 {
		return &((*qjcs)[0]), nil
	}
	return nil, fmt.Errorf("queue.job.channel was not found with criteria %v", criteria)
}

// FindQueueJobChannels finds queue.job.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobChannels(criteria *Criteria, options *Options) (*QueueJobChannels, error) {
	qjcs := &QueueJobChannels{}
	if err := c.SearchRead(QueueJobChannelModel, criteria, options, qjcs); err != nil {
		return nil, err
	}
	return qjcs, nil
}

// FindQueueJobChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueJobChannelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueJobChannelId finds record id by querying it with criteria.
func (c *Client) FindQueueJobChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueJobChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("queue.job.channel was not found with criteria %v and options %v", criteria, options)
}
