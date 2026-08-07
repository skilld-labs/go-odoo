package odoo

// TaskShareWizard represents task.share.wizard model.
type TaskShareWizard struct {
	AccessWarning            *String    `xmlrpc:"access_warning,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	Note                     *String    `xmlrpc:"note,omitempty"`
	PartnerIds               *Relation  `xmlrpc:"partner_ids,omitempty"`
	ProjectPrivacyVisibility *Selection `xmlrpc:"project_privacy_visibility,omitempty"`
	ResId                    *Int       `xmlrpc:"res_id,omitempty"`
	ResModel                 *String    `xmlrpc:"res_model,omitempty"`
	ResourceRef              *String    `xmlrpc:"resource_ref,omitempty"`
	ShareLink                *String    `xmlrpc:"share_link,omitempty"`
	TaskId                   *Many2One  `xmlrpc:"task_id,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// TaskShareWizards represents array of task.share.wizard model.
type TaskShareWizards []TaskShareWizard

// TaskShareWizardModel is the odoo model name.
const TaskShareWizardModel = "task.share.wizard"

// Many2One convert TaskShareWizard to *Many2One.
func (tsw *TaskShareWizard) Many2One() *Many2One {
	return NewMany2One(tsw.Id.Get(), "")
}

// CreateTaskShareWizard creates a new task.share.wizard model and returns its id.
func (c *Client) CreateTaskShareWizard(tsw *TaskShareWizard) (int64, error) {
	ids, err := c.CreateTaskShareWizards([]*TaskShareWizard{tsw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTaskShareWizard creates a new task.share.wizard model and returns its id.
func (c *Client) CreateTaskShareWizards(tsws []*TaskShareWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range tsws {
		vv = append(vv, v)
	}
	return c.Create(TaskShareWizardModel, vv, nil)
}

// UpdateTaskShareWizard updates an existing task.share.wizard record.
func (c *Client) UpdateTaskShareWizard(tsw *TaskShareWizard) error {
	return c.UpdateTaskShareWizards([]int64{tsw.Id.Get()}, tsw)
}

// UpdateTaskShareWizards updates existing task.share.wizard records.
// All records (represented by ids) will be updated by tsw values.
func (c *Client) UpdateTaskShareWizards(ids []int64, tsw *TaskShareWizard) error {
	return c.Update(TaskShareWizardModel, ids, tsw, nil)
}

// DeleteTaskShareWizard deletes an existing task.share.wizard record.
func (c *Client) DeleteTaskShareWizard(id int64) error {
	return c.DeleteTaskShareWizards([]int64{id})
}

// DeleteTaskShareWizards deletes existing task.share.wizard records.
func (c *Client) DeleteTaskShareWizards(ids []int64) error {
	return c.Delete(TaskShareWizardModel, ids)
}

// GetTaskShareWizard gets task.share.wizard existing record.
func (c *Client) GetTaskShareWizard(id int64) (*TaskShareWizard, error) {
	tsws, err := c.GetTaskShareWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*tsws)[0]), nil
}

// GetTaskShareWizards gets task.share.wizard existing records.
func (c *Client) GetTaskShareWizards(ids []int64) (*TaskShareWizards, error) {
	tsws := &TaskShareWizards{}
	if err := c.Read(TaskShareWizardModel, ids, nil, tsws); err != nil {
		return nil, err
	}
	return tsws, nil
}

// FindTaskShareWizard finds task.share.wizard record by querying it with criteria.
func (c *Client) FindTaskShareWizard(criteria *Criteria) (*TaskShareWizard, error) {
	tsws := &TaskShareWizards{}
	if err := c.SearchRead(TaskShareWizardModel, criteria, NewOptions().Limit(1), tsws); err != nil {
		return nil, err
	}
	return &((*tsws)[0]), nil
}

// FindTaskShareWizards finds task.share.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTaskShareWizards(criteria *Criteria, options *Options) (*TaskShareWizards, error) {
	tsws := &TaskShareWizards{}
	if err := c.SearchRead(TaskShareWizardModel, criteria, options, tsws); err != nil {
		return nil, err
	}
	return tsws, nil
}

// FindTaskShareWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTaskShareWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(TaskShareWizardModel, criteria, options)
}

// FindTaskShareWizardId finds record id by querying it with criteria.
func (c *Client) FindTaskShareWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TaskShareWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
