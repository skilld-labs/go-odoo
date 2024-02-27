package odoo

// HrHolidaysRemainingLeavesUser represents hr.holidays.remaining.leaves.user model.
type HrHolidaysRemainingLeavesUser struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	LeaveType   *String   `xmlrpc:"leave_type,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	NoOfLeaves  *Int      `xmlrpc:"no_of_leaves,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
}

// HrHolidaysRemainingLeavesUsers represents array of hr.holidays.remaining.leaves.user model.
type HrHolidaysRemainingLeavesUsers []HrHolidaysRemainingLeavesUser

// HrHolidaysRemainingLeavesUserModel is the odoo model name.
const HrHolidaysRemainingLeavesUserModel = "hr.holidays.remaining.leaves.user"

// Many2One convert HrHolidaysRemainingLeavesUser to *Many2One.
func (hhrlu *HrHolidaysRemainingLeavesUser) Many2One() *Many2One {
	return NewMany2One(hhrlu.Id.Get(), "")
}

// CreateHrHolidaysRemainingLeavesUser creates a new hr.holidays.remaining.leaves.user model and returns its id.
func (c *Client) CreateHrHolidaysRemainingLeavesUser(hhrlu *HrHolidaysRemainingLeavesUser) (int64, error) {
	ids, err := c.CreateHrHolidaysRemainingLeavesUsers([]*HrHolidaysRemainingLeavesUser{hhrlu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrHolidaysRemainingLeavesUser creates a new hr.holidays.remaining.leaves.user model and returns its id.
func (c *Client) CreateHrHolidaysRemainingLeavesUsers(hhrlus []*HrHolidaysRemainingLeavesUser) ([]int64, error) {
	var vv []interface{}
	for _, v := range hhrlus {
		vv = append(vv, v)
	}
	return c.Create(HrHolidaysRemainingLeavesUserModel, vv, nil)
}

// UpdateHrHolidaysRemainingLeavesUser updates an existing hr.holidays.remaining.leaves.user record.
func (c *Client) UpdateHrHolidaysRemainingLeavesUser(hhrlu *HrHolidaysRemainingLeavesUser) error {
	return c.UpdateHrHolidaysRemainingLeavesUsers([]int64{hhrlu.Id.Get()}, hhrlu)
}

// UpdateHrHolidaysRemainingLeavesUsers updates existing hr.holidays.remaining.leaves.user records.
// All records (represented by ids) will be updated by hhrlu values.
func (c *Client) UpdateHrHolidaysRemainingLeavesUsers(ids []int64, hhrlu *HrHolidaysRemainingLeavesUser) error {
	return c.Update(HrHolidaysRemainingLeavesUserModel, ids, hhrlu, nil)
}

// DeleteHrHolidaysRemainingLeavesUser deletes an existing hr.holidays.remaining.leaves.user record.
func (c *Client) DeleteHrHolidaysRemainingLeavesUser(id int64) error {
	return c.DeleteHrHolidaysRemainingLeavesUsers([]int64{id})
}

// DeleteHrHolidaysRemainingLeavesUsers deletes existing hr.holidays.remaining.leaves.user records.
func (c *Client) DeleteHrHolidaysRemainingLeavesUsers(ids []int64) error {
	return c.Delete(HrHolidaysRemainingLeavesUserModel, ids)
}

// GetHrHolidaysRemainingLeavesUser gets hr.holidays.remaining.leaves.user existing record.
func (c *Client) GetHrHolidaysRemainingLeavesUser(id int64) (*HrHolidaysRemainingLeavesUser, error) {
	hhrlus, err := c.GetHrHolidaysRemainingLeavesUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hhrlus)[0]), nil
}

// GetHrHolidaysRemainingLeavesUsers gets hr.holidays.remaining.leaves.user existing records.
func (c *Client) GetHrHolidaysRemainingLeavesUsers(ids []int64) (*HrHolidaysRemainingLeavesUsers, error) {
	hhrlus := &HrHolidaysRemainingLeavesUsers{}
	if err := c.Read(HrHolidaysRemainingLeavesUserModel, ids, nil, hhrlus); err != nil {
		return nil, err
	}
	return hhrlus, nil
}

// FindHrHolidaysRemainingLeavesUser finds hr.holidays.remaining.leaves.user record by querying it with criteria.
func (c *Client) FindHrHolidaysRemainingLeavesUser(criteria *Criteria) (*HrHolidaysRemainingLeavesUser, error) {
	hhrlus := &HrHolidaysRemainingLeavesUsers{}
	if err := c.SearchRead(HrHolidaysRemainingLeavesUserModel, criteria, NewOptions().Limit(1), hhrlus); err != nil {
		return nil, err
	}
	return &((*hhrlus)[0]), nil
}

// FindHrHolidaysRemainingLeavesUsers finds hr.holidays.remaining.leaves.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysRemainingLeavesUsers(criteria *Criteria, options *Options) (*HrHolidaysRemainingLeavesUsers, error) {
	hhrlus := &HrHolidaysRemainingLeavesUsers{}
	if err := c.SearchRead(HrHolidaysRemainingLeavesUserModel, criteria, options, hhrlus); err != nil {
		return nil, err
	}
	return hhrlus, nil
}

// FindHrHolidaysRemainingLeavesUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysRemainingLeavesUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrHolidaysRemainingLeavesUserModel, criteria, options)
}

// FindHrHolidaysRemainingLeavesUserId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysRemainingLeavesUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysRemainingLeavesUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
