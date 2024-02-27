package odoo

// ProjectTaskMergeWizard represents project.task.merge.wizard model.
type ProjectTaskMergeWizard struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateNewTask   *Bool     `xmlrpc:"create_new_task,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	TargetProjectId *Many2One `xmlrpc:"target_project_id,omitempty"`
	TargetTaskId    *Many2One `xmlrpc:"target_task_id,omitempty"`
	TargetTaskName  *String   `xmlrpc:"target_task_name,omitempty"`
	TaskIds         *Relation `xmlrpc:"task_ids,omitempty"`
	UserId          *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectTaskMergeWizards represents array of project.task.merge.wizard model.
type ProjectTaskMergeWizards []ProjectTaskMergeWizard

// ProjectTaskMergeWizardModel is the odoo model name.
const ProjectTaskMergeWizardModel = "project.task.merge.wizard"

// Many2One convert ProjectTaskMergeWizard to *Many2One.
func (ptmw *ProjectTaskMergeWizard) Many2One() *Many2One {
	return NewMany2One(ptmw.Id.Get(), "")
}

// CreateProjectTaskMergeWizard creates a new project.task.merge.wizard model and returns its id.
func (c *Client) CreateProjectTaskMergeWizard(ptmw *ProjectTaskMergeWizard) (int64, error) {
	ids, err := c.CreateProjectTaskMergeWizards([]*ProjectTaskMergeWizard{ptmw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskMergeWizard creates a new project.task.merge.wizard model and returns its id.
func (c *Client) CreateProjectTaskMergeWizards(ptmws []*ProjectTaskMergeWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptmws {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskMergeWizardModel, vv, nil)
}

// UpdateProjectTaskMergeWizard updates an existing project.task.merge.wizard record.
func (c *Client) UpdateProjectTaskMergeWizard(ptmw *ProjectTaskMergeWizard) error {
	return c.UpdateProjectTaskMergeWizards([]int64{ptmw.Id.Get()}, ptmw)
}

// UpdateProjectTaskMergeWizards updates existing project.task.merge.wizard records.
// All records (represented by ids) will be updated by ptmw values.
func (c *Client) UpdateProjectTaskMergeWizards(ids []int64, ptmw *ProjectTaskMergeWizard) error {
	return c.Update(ProjectTaskMergeWizardModel, ids, ptmw, nil)
}

// DeleteProjectTaskMergeWizard deletes an existing project.task.merge.wizard record.
func (c *Client) DeleteProjectTaskMergeWizard(id int64) error {
	return c.DeleteProjectTaskMergeWizards([]int64{id})
}

// DeleteProjectTaskMergeWizards deletes existing project.task.merge.wizard records.
func (c *Client) DeleteProjectTaskMergeWizards(ids []int64) error {
	return c.Delete(ProjectTaskMergeWizardModel, ids)
}

// GetProjectTaskMergeWizard gets project.task.merge.wizard existing record.
func (c *Client) GetProjectTaskMergeWizard(id int64) (*ProjectTaskMergeWizard, error) {
	ptmws, err := c.GetProjectTaskMergeWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ptmws)[0]), nil
}

// GetProjectTaskMergeWizards gets project.task.merge.wizard existing records.
func (c *Client) GetProjectTaskMergeWizards(ids []int64) (*ProjectTaskMergeWizards, error) {
	ptmws := &ProjectTaskMergeWizards{}
	if err := c.Read(ProjectTaskMergeWizardModel, ids, nil, ptmws); err != nil {
		return nil, err
	}
	return ptmws, nil
}

// FindProjectTaskMergeWizard finds project.task.merge.wizard record by querying it with criteria.
func (c *Client) FindProjectTaskMergeWizard(criteria *Criteria) (*ProjectTaskMergeWizard, error) {
	ptmws := &ProjectTaskMergeWizards{}
	if err := c.SearchRead(ProjectTaskMergeWizardModel, criteria, NewOptions().Limit(1), ptmws); err != nil {
		return nil, err
	}
	return &((*ptmws)[0]), nil
}

// FindProjectTaskMergeWizards finds project.task.merge.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskMergeWizards(criteria *Criteria, options *Options) (*ProjectTaskMergeWizards, error) {
	ptmws := &ProjectTaskMergeWizards{}
	if err := c.SearchRead(ProjectTaskMergeWizardModel, criteria, options, ptmws); err != nil {
		return nil, err
	}
	return ptmws, nil
}

// FindProjectTaskMergeWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskMergeWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTaskMergeWizardModel, criteria, options)
}

// FindProjectTaskMergeWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskMergeWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskMergeWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
