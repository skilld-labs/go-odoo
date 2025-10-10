package odoo

// ProjectTaskTypeDeleteWizard represents project.task.type.delete.wizard model.
type ProjectTaskTypeDeleteWizard struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	ProjectIds   *Relation `xmlrpc:"project_ids,omitempty"`
	StageIds     *Relation `xmlrpc:"stage_ids,omitempty"`
	StagesActive *Bool     `xmlrpc:"stages_active,omitempty"`
	TasksCount   *Int      `xmlrpc:"tasks_count,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectTaskTypeDeleteWizards represents array of project.task.type.delete.wizard model.
type ProjectTaskTypeDeleteWizards []ProjectTaskTypeDeleteWizard

// ProjectTaskTypeDeleteWizardModel is the odoo model name.
const ProjectTaskTypeDeleteWizardModel = "project.task.type.delete.wizard"

// Many2One convert ProjectTaskTypeDeleteWizard to *Many2One.
func (pttdw *ProjectTaskTypeDeleteWizard) Many2One() *Many2One {
	return NewMany2One(pttdw.Id.Get(), "")
}

// CreateProjectTaskTypeDeleteWizard creates a new project.task.type.delete.wizard model and returns its id.
func (c *Client) CreateProjectTaskTypeDeleteWizard(pttdw *ProjectTaskTypeDeleteWizard) (int64, error) {
	ids, err := c.CreateProjectTaskTypeDeleteWizards([]*ProjectTaskTypeDeleteWizard{pttdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskTypeDeleteWizard creates a new project.task.type.delete.wizard model and returns its id.
func (c *Client) CreateProjectTaskTypeDeleteWizards(pttdws []*ProjectTaskTypeDeleteWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pttdws {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskTypeDeleteWizardModel, vv, nil)
}

// UpdateProjectTaskTypeDeleteWizard updates an existing project.task.type.delete.wizard record.
func (c *Client) UpdateProjectTaskTypeDeleteWizard(pttdw *ProjectTaskTypeDeleteWizard) error {
	return c.UpdateProjectTaskTypeDeleteWizards([]int64{pttdw.Id.Get()}, pttdw)
}

// UpdateProjectTaskTypeDeleteWizards updates existing project.task.type.delete.wizard records.
// All records (represented by ids) will be updated by pttdw values.
func (c *Client) UpdateProjectTaskTypeDeleteWizards(ids []int64, pttdw *ProjectTaskTypeDeleteWizard) error {
	return c.Update(ProjectTaskTypeDeleteWizardModel, ids, pttdw, nil)
}

// DeleteProjectTaskTypeDeleteWizard deletes an existing project.task.type.delete.wizard record.
func (c *Client) DeleteProjectTaskTypeDeleteWizard(id int64) error {
	return c.DeleteProjectTaskTypeDeleteWizards([]int64{id})
}

// DeleteProjectTaskTypeDeleteWizards deletes existing project.task.type.delete.wizard records.
func (c *Client) DeleteProjectTaskTypeDeleteWizards(ids []int64) error {
	return c.Delete(ProjectTaskTypeDeleteWizardModel, ids)
}

// GetProjectTaskTypeDeleteWizard gets project.task.type.delete.wizard existing record.
func (c *Client) GetProjectTaskTypeDeleteWizard(id int64) (*ProjectTaskTypeDeleteWizard, error) {
	pttdws, err := c.GetProjectTaskTypeDeleteWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pttdws)[0]), nil
}

// GetProjectTaskTypeDeleteWizards gets project.task.type.delete.wizard existing records.
func (c *Client) GetProjectTaskTypeDeleteWizards(ids []int64) (*ProjectTaskTypeDeleteWizards, error) {
	pttdws := &ProjectTaskTypeDeleteWizards{}
	if err := c.Read(ProjectTaskTypeDeleteWizardModel, ids, nil, pttdws); err != nil {
		return nil, err
	}
	return pttdws, nil
}

// FindProjectTaskTypeDeleteWizard finds project.task.type.delete.wizard record by querying it with criteria.
func (c *Client) FindProjectTaskTypeDeleteWizard(criteria *Criteria) (*ProjectTaskTypeDeleteWizard, error) {
	pttdws := &ProjectTaskTypeDeleteWizards{}
	if err := c.SearchRead(ProjectTaskTypeDeleteWizardModel, criteria, NewOptions().Limit(1), pttdws); err != nil {
		return nil, err
	}
	return &((*pttdws)[0]), nil
}

// FindProjectTaskTypeDeleteWizards finds project.task.type.delete.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskTypeDeleteWizards(criteria *Criteria, options *Options) (*ProjectTaskTypeDeleteWizards, error) {
	pttdws := &ProjectTaskTypeDeleteWizards{}
	if err := c.SearchRead(ProjectTaskTypeDeleteWizardModel, criteria, options, pttdws); err != nil {
		return nil, err
	}
	return pttdws, nil
}

// FindProjectTaskTypeDeleteWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskTypeDeleteWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTaskTypeDeleteWizardModel, criteria, options)
}

// FindProjectTaskTypeDeleteWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskTypeDeleteWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskTypeDeleteWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
