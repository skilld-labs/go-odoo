package odoo

// HrJob represents hr.job model.
type HrJob struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	EmployeeIds              *Relation  `xmlrpc:"employee_ids,omptempty"`
	ExpectedEmployees        *Int       `xmlrpc:"expected_employees,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	NoOfEmployee             *Int       `xmlrpc:"no_of_employee,omptempty"`
	NoOfHiredEmployee        *Int       `xmlrpc:"no_of_hired_employee,omptempty"`
	NoOfRecruitment          *Int       `xmlrpc:"no_of_recruitment,omptempty"`
	Requirements             *String    `xmlrpc:"requirements,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.CreateHrJobs([]*HrJob{hj})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrJob creates a new hr.job model and returns its id.
func (c *Client) CreateHrJobs(hjs []*HrJob) ([]int64, error) {
	var vv []interface{}
	for _, v := range hjs {
		vv = append(vv, v)
	}
	return c.Create(HrJobModel, vv, nil)
}

// UpdateHrJob updates an existing hr.job record.
func (c *Client) UpdateHrJob(hj *HrJob) error {
	return c.UpdateHrJobs([]int64{hj.Id.Get()}, hj)
}

// UpdateHrJobs updates existing hr.job records.
// All records (represented by ids) will be updated by hj values.
func (c *Client) UpdateHrJobs(ids []int64, hj *HrJob) error {
	return c.Update(HrJobModel, ids, hj, nil)
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
	return &((*hjs)[0]), nil
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
	return &((*hjs)[0]), nil
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
	return c.Search(HrJobModel, criteria, options)
}

// FindHrJobId finds record id by querying it with criteria.
func (c *Client) FindHrJobId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrJobModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
