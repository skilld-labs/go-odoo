package odoo

import (
	"fmt"
)

// HrDepartment represents hr.department model.
type HrDepartment struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omitempty"`
	AbsenceOfToday           *Int      `xmlrpc:"absence_of_today,omitempty"`
	Active                   *Bool     `xmlrpc:"active,omitempty"`
	AllocationToApproveCount *Int      `xmlrpc:"allocation_to_approve_count,omitempty"`
	ChildIds                 *Relation `xmlrpc:"child_ids,omitempty"`
	Color                    *Int      `xmlrpc:"color,omitempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omitempty"`
	CompleteName             *String   `xmlrpc:"complete_name,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	JobsIds                  *Relation `xmlrpc:"jobs_ids,omitempty"`
	LeaveToApproveCount      *Int      `xmlrpc:"leave_to_approve_count,omitempty"`
	ManagerId                *Many2One `xmlrpc:"manager_id,omitempty"`
	MemberIds                *Relation `xmlrpc:"member_ids,omitempty"`
	MessageChannelIds        *Relation `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time     `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool     `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int      `xmlrpc:"message_unread_counter,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	Note                     *String   `xmlrpc:"note,omitempty"`
	ParentId                 *Many2One `xmlrpc:"parent_id,omitempty"`
	TotalEmployee            *Int      `xmlrpc:"total_employee,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return nil, fmt.Errorf("no hr.department was found with criteria %v", criteria)
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
	return -1, fmt.Errorf("no hr.department was found with criteria %v and options %v", criteria, options)
}
