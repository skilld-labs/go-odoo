package odoo

import (
	"fmt"
)

// HrDepartment represents hr.department model.
type HrDepartment struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	AbsenceOfToday           *Int      `xmlrpc:"absence_of_today,omptempty"`
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	AllocationToApproveCount *Int      `xmlrpc:"allocation_to_approve_count,omptempty"`
	ChildIds                 *Relation `xmlrpc:"child_ids,omptempty"`
	Color                    *Int      `xmlrpc:"color,omptempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omptempty"`
	CompleteName             *String   `xmlrpc:"complete_name,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	JobsIds                  *Relation `xmlrpc:"jobs_ids,omptempty"`
	LeaveToApproveCount      *Int      `xmlrpc:"leave_to_approve_count,omptempty"`
	ManagerId                *Many2One `xmlrpc:"manager_id,omptempty"`
	MemberIds                *Relation `xmlrpc:"member_ids,omptempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time     `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	Note                     *String   `xmlrpc:"note,omptempty"`
	ParentId                 *Many2One `xmlrpc:"parent_id,omptempty"`
	TotalEmployee            *Int      `xmlrpc:"total_employee,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrDepartments represents array of hr.department model.
type HrDepartments []HrDepartment

// HrDepartmentModel is the odoo model name.
const HrDepartmentModel = "hr.department"

// Many2One convert HrDepartment to *Many2One.
func (hd *HrDepartment) Many2One() *Many2One {
	return NewMany2One(hd.Id.Get(), "")
}

// CreateHrDepartment creates a new hr.department model and returns its id.
func (c *Client) CreateHrDepartment(hd *HrDepartment) (int64, error) {
	return c.Create(HrDepartmentModel, hd)
}

// UpdateHrDepartment updates an existing hr.department record.
func (c *Client) UpdateHrDepartment(hd *HrDepartment) error {
	return c.UpdateHrDepartments([]int64{hd.Id.Get()}, hd)
}

// UpdateHrDepartments updates existing hr.department records.
// All records (represented by ids) will be updated by hd values.
func (c *Client) UpdateHrDepartments(ids []int64, hd *HrDepartment) error {
	return c.Update(HrDepartmentModel, ids, hd)
}

// DeleteHrDepartment deletes an existing hr.department record.
func (c *Client) DeleteHrDepartment(id int64) error {
	return c.DeleteHrDepartments([]int64{id})
}

// DeleteHrDepartments deletes existing hr.department records.
func (c *Client) DeleteHrDepartments(ids []int64) error {
	return c.Delete(HrDepartmentModel, ids)
}

// GetHrDepartment gets hr.department existing record.
func (c *Client) GetHrDepartment(id int64) (*HrDepartment, error) {
	hds, err := c.GetHrDepartments([]int64{id})
	if err != nil {
		return nil, err
	}
	if hds != nil && len(*hds) > 0 {
		return &((*hds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.department not found", id)
}

// GetHrDepartments gets hr.department existing records.
func (c *Client) GetHrDepartments(ids []int64) (*HrDepartments, error) {
	hds := &HrDepartments{}
	if err := c.Read(HrDepartmentModel, ids, nil, hds); err != nil {
		return nil, err
	}
	return hds, nil
}

// FindHrDepartment finds hr.department record by querying it with criteria.
func (c *Client) FindHrDepartment(criteria *Criteria) (*HrDepartment, error) {
	hds := &HrDepartments{}
	if err := c.SearchRead(HrDepartmentModel, criteria, NewOptions().Limit(1), hds); err != nil {
		return nil, err
	}
	if hds != nil && len(*hds) > 0 {
		return &((*hds)[0]), nil
	}
	return nil, fmt.Errorf("hr.department was not found with criteria %v", criteria)
}

// FindHrDepartments finds hr.department records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrDepartments(criteria *Criteria, options *Options) (*HrDepartments, error) {
	hds := &HrDepartments{}
	if err := c.SearchRead(HrDepartmentModel, criteria, options, hds); err != nil {
		return nil, err
	}
	return hds, nil
}

// FindHrDepartmentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrDepartmentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrDepartmentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrDepartmentId finds record id by querying it with criteria.
func (c *Client) FindHrDepartmentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrDepartmentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.department was not found with criteria %v and options %v", criteria, options)
}
