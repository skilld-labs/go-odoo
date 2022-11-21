package odoo

import (
	"fmt"
)

// QueueJobFunction represents queue.job.function model.
type QueueJobFunction struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Channel     *String   `xmlrpc:"channel,omptempty"`
	ChannelId   *Many2One `xmlrpc:"channel_id,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
}

// QueueJobFunctions represents array of queue.job.function model.
type QueueJobFunctions []QueueJobFunction

// QueueJobFunctionModel is the odoo model name.
const QueueJobFunctionModel = "queue.job.function"

// Many2One convert QueueJobFunction to *Many2One.
func (qjf *QueueJobFunction) Many2One() *Many2One {
	return NewMany2One(qjf.Id.Get(), "")
}

// CreateQueueJobFunction creates a new queue.job.function model and returns its id.
func (c *Client) CreateQueueJobFunction(qjf *QueueJobFunction) (int64, error) {
	return c.Create(QueueJobFunctionModel, qjf)
}

// UpdateQueueJobFunction updates an existing queue.job.function record.
func (c *Client) UpdateQueueJobFunction(qjf *QueueJobFunction) error {
	return c.UpdateQueueJobFunctions([]int64{qjf.Id.Get()}, qjf)
}

// UpdateQueueJobFunctions updates existing queue.job.function records.
// All records (represented by ids) will be updated by qjf values.
func (c *Client) UpdateQueueJobFunctions(ids []int64, qjf *QueueJobFunction) error {
	return c.Update(QueueJobFunctionModel, ids, qjf)
}

// DeleteQueueJobFunction deletes an existing queue.job.function record.
func (c *Client) DeleteQueueJobFunction(id int64) error {
	return c.DeleteQueueJobFunctions([]int64{id})
}

// DeleteQueueJobFunctions deletes existing queue.job.function records.
func (c *Client) DeleteQueueJobFunctions(ids []int64) error {
	return c.Delete(QueueJobFunctionModel, ids)
}

// GetQueueJobFunction gets queue.job.function existing record.
func (c *Client) GetQueueJobFunction(id int64) (*QueueJobFunction, error) {
	qjfs, err := c.GetQueueJobFunctions([]int64{id})
	if err != nil {
		return nil, err
	}
	if qjfs != nil && len(*qjfs) > 0 {
		return &((*qjfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue.job.function not found", id)
}

// GetQueueJobFunctions gets queue.job.function existing records.
func (c *Client) GetQueueJobFunctions(ids []int64) (*QueueJobFunctions, error) {
	qjfs := &QueueJobFunctions{}
	if err := c.Read(QueueJobFunctionModel, ids, nil, qjfs); err != nil {
		return nil, err
	}
	return qjfs, nil
}

// FindQueueJobFunction finds queue.job.function record by querying it with criteria.
func (c *Client) FindQueueJobFunction(criteria *Criteria) (*QueueJobFunction, error) {
	qjfs := &QueueJobFunctions{}
	if err := c.SearchRead(QueueJobFunctionModel, criteria, NewOptions().Limit(1), qjfs); err != nil {
		return nil, err
	}
	if qjfs != nil && len(*qjfs) > 0 {
		return &((*qjfs)[0]), nil
	}
	return nil, fmt.Errorf("no queue.job.function was found with criteria %v", criteria)
}

// FindQueueJobFunctions finds queue.job.function records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobFunctions(criteria *Criteria, options *Options) (*QueueJobFunctions, error) {
	qjfs := &QueueJobFunctions{}
	if err := c.SearchRead(QueueJobFunctionModel, criteria, options, qjfs); err != nil {
		return nil, err
	}
	return qjfs, nil
}

// FindQueueJobFunctionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobFunctionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueJobFunctionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueJobFunctionId finds record id by querying it with criteria.
func (c *Client) FindQueueJobFunctionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueJobFunctionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no queue.job.function was found with criteria %v and options %v", criteria, options)
}
