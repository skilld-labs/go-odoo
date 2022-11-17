package odoo

import (
	"fmt"
)

// HrJob represents hr.job model.
type HrJob struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	EmployeeIds              *Relation  `xmlrpc:"employee_ids,omitempty"`
	ExpectedEmployees        *Int       `xmlrpc:"expected_employees,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	NoOfEmployee             *Int       `xmlrpc:"no_of_employee,omitempty"`
	NoOfHiredEmployee        *Int       `xmlrpc:"no_of_hired_employee,omitempty"`
	NoOfRecruitment          *Int       `xmlrpc:"no_of_recruitment,omitempty"`
	Requirements             *String    `xmlrpc:"requirements,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrJobs represents array of hr.job model.
type HrJobs []HrJob

// HrJobModel is the odoo model name.
const HrJobModel = "hr.job"

// Many2One convert HrJob to *Many2One.
func (hj *HrJob) Many2One() *Many2One {
	return NewMany2One(hj.Id.Get(), "")
}

// CreateHrJob creates a new hr.job model and returns its id.
func (c *Client) CreateHrJob(hj *HrJob) (int64, error) {
	return c.Create(HrJobModel, hj)
}

// UpdateHrJob updates an existing hr.job record.
func (c *Client) UpdateHrJob(hj *HrJob) error {
	return c.UpdateHrJobs([]int64{hj.Id.Get()}, hj)
}

// UpdateHrJobs updates existing hr.job records.
// All records (represented by ids) will be updated by hj values.
func (c *Client) UpdateHrJobs(ids []int64, hj *HrJob) error {
	return c.Update(HrJobModel, ids, hj)
}

// DeleteHrJob deletes an existing hr.job record.
func (c *Client) DeleteHrJob(id int64) error {
	return c.DeleteHrJobs([]int64{id})
}

// DeleteHrJobs deletes existing hr.job records.
func (c *Client) DeleteHrJobs(ids []int64) error {
	return c.Delete(HrJobModel, ids)
}

// GetHrJob gets hr.job existing record.
func (c *Client) GetHrJob(id int64) (*HrJob, error) {
	hjs, err := c.GetHrJobs([]int64{id})
	if err != nil {
		return nil, err
	}
	if hjs != nil && len(*hjs) > 0 {
		return &((*hjs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.job not found", id)
}

// GetHrJobs gets hr.job existing records.
func (c *Client) GetHrJobs(ids []int64) (*HrJobs, error) {
	hjs := &HrJobs{}
	if err := c.Read(HrJobModel, ids, nil, hjs); err != nil {
		return nil, err
	}
	return hjs, nil
}

// FindHrJob finds hr.job record by querying it with criteria.
func (c *Client) FindHrJob(criteria *Criteria) (*HrJob, error) {
	hjs := &HrJobs{}
	if err := c.SearchRead(HrJobModel, criteria, NewOptions().Limit(1), hjs); err != nil {
		return nil, err
	}
	if hjs != nil && len(*hjs) > 0 {
		return &((*hjs)[0]), nil
	}
	return nil, fmt.Errorf("no hr.job was found with criteria %v", criteria)
}

// FindHrJobs finds hr.job records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobs(criteria *Criteria, options *Options) (*HrJobs, error) {
	hjs := &HrJobs{}
	if err := c.SearchRead(HrJobModel, criteria, options, hjs); err != nil {
		return nil, err
	}
	return hjs, nil
}

// FindHrJobIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrJobIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrJobModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrJobId finds record id by querying it with criteria.
func (c *Client) FindHrJobId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrJobModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no hr.job was found with criteria %v and options %v", criteria, options)
}
