package odoo

import (
	"fmt"
)

// QueueJob represents queue.job model.
type QueueJob struct {
	LastUpdate               *Time       `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline     *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds              *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState            *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary          *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId           *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId           *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	Args                     interface{} `xmlrpc:"args,omitempty"`
	Channel                  *String     `xmlrpc:"channel,omitempty"`
	ChannelMethodName        *String     `xmlrpc:"channel_method_name,omitempty"`
	CompanyId                *Many2One   `xmlrpc:"company_id,omitempty"`
	DateCreated              *Time       `xmlrpc:"date_created,omitempty"`
	DateDone                 *Time       `xmlrpc:"date_done,omitempty"`
	DateEnqueued             *Time       `xmlrpc:"date_enqueued,omitempty"`
	DateStarted              *Time       `xmlrpc:"date_started,omitempty"`
	DisplayName              *String     `xmlrpc:"display_name,omitempty"`
	Eta                      *Time       `xmlrpc:"eta,omitempty"`
	ExcInfo                  *String     `xmlrpc:"exc_info,omitempty"`
	FuncString               *String     `xmlrpc:"func_string,omitempty"`
	Id                       *Int        `xmlrpc:"id,omitempty"`
	JobFunctionId            *Many2One   `xmlrpc:"job_function_id,omitempty"`
	Kwargs                   interface{} `xmlrpc:"kwargs,omitempty"`
	MaxRetries               *Int        `xmlrpc:"max_retries,omitempty"`
	MessageChannelIds        *Relation   `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time       `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool       `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int        `xmlrpc:"message_unread_counter,omitempty"`
	MethodName               *String     `xmlrpc:"method_name,omitempty"`
	ModelName                *String     `xmlrpc:"model_name,omitempty"`
	Name                     *String     `xmlrpc:"name,omitempty"`
	Priority                 *Int        `xmlrpc:"priority,omitempty"`
	RecordIds                interface{} `xmlrpc:"record_ids,omitempty"`
	Result                   *String     `xmlrpc:"result,omitempty"`
	Retry                    *Int        `xmlrpc:"retry,omitempty"`
	State                    *Selection  `xmlrpc:"state,omitempty"`
	UserId                   *Many2One   `xmlrpc:"user_id,omitempty"`
	Uuid                     *String     `xmlrpc:"uuid,omitempty"`
	WebsiteMessageIds        *Relation   `xmlrpc:"website_message_ids,omitempty"`
}

// QueueJobs represents array of queue.job model.
type QueueJobs []QueueJob

// QueueJobModel is the odoo model name.
const QueueJobModel = "queue.job"

// Many2One convert QueueJob to *Many2One.
func (qj *QueueJob) Many2One() *Many2One {
	return NewMany2One(qj.Id.Get(), "")
}

// CreateQueueJob creates a new queue.job model and returns its id.
func (c *Client) CreateQueueJob(qj *QueueJob) (int64, error) {
	return c.Create(QueueJobModel, qj)
}

// UpdateQueueJob updates an existing queue.job record.
func (c *Client) UpdateQueueJob(qj *QueueJob) error {
	return c.UpdateQueueJobs([]int64{qj.Id.Get()}, qj)
}

// UpdateQueueJobs updates existing queue.job records.
// All records (represented by ids) will be updated by qj values.
func (c *Client) UpdateQueueJobs(ids []int64, qj *QueueJob) error {
	return c.Update(QueueJobModel, ids, qj)
}

// DeleteQueueJob deletes an existing queue.job record.
func (c *Client) DeleteQueueJob(id int64) error {
	return c.DeleteQueueJobs([]int64{id})
}

// DeleteQueueJobs deletes existing queue.job records.
func (c *Client) DeleteQueueJobs(ids []int64) error {
	return c.Delete(QueueJobModel, ids)
}

// GetQueueJob gets queue.job existing record.
func (c *Client) GetQueueJob(id int64) (*QueueJob, error) {
	qjs, err := c.GetQueueJobs([]int64{id})
	if err != nil {
		return nil, err
	}
	if qjs != nil && len(*qjs) > 0 {
		return &((*qjs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue.job not found", id)
}

// GetQueueJobs gets queue.job existing records.
func (c *Client) GetQueueJobs(ids []int64) (*QueueJobs, error) {
	qjs := &QueueJobs{}
	if err := c.Read(QueueJobModel, ids, nil, qjs); err != nil {
		return nil, err
	}
	return qjs, nil
}

// FindQueueJob finds queue.job record by querying it with criteria.
func (c *Client) FindQueueJob(criteria *Criteria) (*QueueJob, error) {
	qjs := &QueueJobs{}
	if err := c.SearchRead(QueueJobModel, criteria, NewOptions().Limit(1), qjs); err != nil {
		return nil, err
	}
	if qjs != nil && len(*qjs) > 0 {
		return &((*qjs)[0]), nil
	}
	return nil, fmt.Errorf("no queue.job was found with criteria %v", criteria)
}

// FindQueueJobs finds queue.job records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobs(criteria *Criteria, options *Options) (*QueueJobs, error) {
	qjs := &QueueJobs{}
	if err := c.SearchRead(QueueJobModel, criteria, options, qjs); err != nil {
		return nil, err
	}
	return qjs, nil
}

// FindQueueJobIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueJobModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueJobId finds record id by querying it with criteria.
func (c *Client) FindQueueJobId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueJobModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no queue.job was found with criteria %v and options %v", criteria, options)
}
