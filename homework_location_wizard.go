package odoo

// HomeworkLocationWizard represents homework.location.wizard model.
type HomeworkLocationWizard struct {
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date             *Time      `xmlrpc:"date,omitempty"`
	DayWeekString    *String    `xmlrpc:"day_week_string,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId       *Many2One  `xmlrpc:"employee_id,omitempty"`
	EmployeeName     *String    `xmlrpc:"employee_name,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Weekly           *Bool      `xmlrpc:"weekly,omitempty"`
	WorkLocationId   *Many2One  `xmlrpc:"work_location_id,omitempty"`
	WorkLocationName *String    `xmlrpc:"work_location_name,omitempty"`
	WorkLocationType *Selection `xmlrpc:"work_location_type,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HomeworkLocationWizards represents array of homework.location.wizard model.
type HomeworkLocationWizards []HomeworkLocationWizard

// HomeworkLocationWizardModel is the odoo model name.
const HomeworkLocationWizardModel = "homework.location.wizard"

// Many2One convert HomeworkLocationWizard to *Many2One.
func (hlw *HomeworkLocationWizard) Many2One() *Many2One {
	return NewMany2One(hlw.Id.Get(), "")
}

// CreateHomeworkLocationWizard creates a new homework.location.wizard model and returns its id.
func (c *Client) CreateHomeworkLocationWizard(hlw *HomeworkLocationWizard) (int64, error) {
	ids, err := c.CreateHomeworkLocationWizards([]*HomeworkLocationWizard{hlw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHomeworkLocationWizard creates a new homework.location.wizard model and returns its id.
func (c *Client) CreateHomeworkLocationWizards(hlws []*HomeworkLocationWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlws {
		vv = append(vv, v)
	}
	return c.Create(HomeworkLocationWizardModel, vv, nil)
}

// UpdateHomeworkLocationWizard updates an existing homework.location.wizard record.
func (c *Client) UpdateHomeworkLocationWizard(hlw *HomeworkLocationWizard) error {
	return c.UpdateHomeworkLocationWizards([]int64{hlw.Id.Get()}, hlw)
}

// UpdateHomeworkLocationWizards updates existing homework.location.wizard records.
// All records (represented by ids) will be updated by hlw values.
func (c *Client) UpdateHomeworkLocationWizards(ids []int64, hlw *HomeworkLocationWizard) error {
	return c.Update(HomeworkLocationWizardModel, ids, hlw, nil)
}

// DeleteHomeworkLocationWizard deletes an existing homework.location.wizard record.
func (c *Client) DeleteHomeworkLocationWizard(id int64) error {
	return c.DeleteHomeworkLocationWizards([]int64{id})
}

// DeleteHomeworkLocationWizards deletes existing homework.location.wizard records.
func (c *Client) DeleteHomeworkLocationWizards(ids []int64) error {
	return c.Delete(HomeworkLocationWizardModel, ids)
}

// GetHomeworkLocationWizard gets homework.location.wizard existing record.
func (c *Client) GetHomeworkLocationWizard(id int64) (*HomeworkLocationWizard, error) {
	hlws, err := c.GetHomeworkLocationWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlws)[0]), nil
}

// GetHomeworkLocationWizards gets homework.location.wizard existing records.
func (c *Client) GetHomeworkLocationWizards(ids []int64) (*HomeworkLocationWizards, error) {
	hlws := &HomeworkLocationWizards{}
	if err := c.Read(HomeworkLocationWizardModel, ids, nil, hlws); err != nil {
		return nil, err
	}
	return hlws, nil
}

// FindHomeworkLocationWizard finds homework.location.wizard record by querying it with criteria.
func (c *Client) FindHomeworkLocationWizard(criteria *Criteria) (*HomeworkLocationWizard, error) {
	hlws := &HomeworkLocationWizards{}
	if err := c.SearchRead(HomeworkLocationWizardModel, criteria, NewOptions().Limit(1), hlws); err != nil {
		return nil, err
	}
	return &((*hlws)[0]), nil
}

// FindHomeworkLocationWizards finds homework.location.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHomeworkLocationWizards(criteria *Criteria, options *Options) (*HomeworkLocationWizards, error) {
	hlws := &HomeworkLocationWizards{}
	if err := c.SearchRead(HomeworkLocationWizardModel, criteria, options, hlws); err != nil {
		return nil, err
	}
	return hlws, nil
}

// FindHomeworkLocationWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHomeworkLocationWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HomeworkLocationWizardModel, criteria, options)
}

// FindHomeworkLocationWizardId finds record id by querying it with criteria.
func (c *Client) FindHomeworkLocationWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HomeworkLocationWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
