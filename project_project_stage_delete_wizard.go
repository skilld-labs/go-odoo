package odoo

// ProjectProjectStageDeleteWizard represents project.project.stage.delete.wizard model.
type ProjectProjectStageDeleteWizard struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	ProjectsCount *Int      `xmlrpc:"projects_count,omitempty"`
	StageIds      *Relation `xmlrpc:"stage_ids,omitempty"`
	StagesActive  *Bool     `xmlrpc:"stages_active,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectProjectStageDeleteWizards represents array of project.project.stage.delete.wizard model.
type ProjectProjectStageDeleteWizards []ProjectProjectStageDeleteWizard

// ProjectProjectStageDeleteWizardModel is the odoo model name.
const ProjectProjectStageDeleteWizardModel = "project.project.stage.delete.wizard"

// Many2One convert ProjectProjectStageDeleteWizard to *Many2One.
func (ppsdw *ProjectProjectStageDeleteWizard) Many2One() *Many2One {
	return NewMany2One(ppsdw.Id.Get(), "")
}

// CreateProjectProjectStageDeleteWizard creates a new project.project.stage.delete.wizard model and returns its id.
func (c *Client) CreateProjectProjectStageDeleteWizard(ppsdw *ProjectProjectStageDeleteWizard) (int64, error) {
	ids, err := c.CreateProjectProjectStageDeleteWizards([]*ProjectProjectStageDeleteWizard{ppsdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectProjectStageDeleteWizard creates a new project.project.stage.delete.wizard model and returns its id.
func (c *Client) CreateProjectProjectStageDeleteWizards(ppsdws []*ProjectProjectStageDeleteWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ppsdws {
		vv = append(vv, v)
	}
	return c.Create(ProjectProjectStageDeleteWizardModel, vv, nil)
}

// UpdateProjectProjectStageDeleteWizard updates an existing project.project.stage.delete.wizard record.
func (c *Client) UpdateProjectProjectStageDeleteWizard(ppsdw *ProjectProjectStageDeleteWizard) error {
	return c.UpdateProjectProjectStageDeleteWizards([]int64{ppsdw.Id.Get()}, ppsdw)
}

// UpdateProjectProjectStageDeleteWizards updates existing project.project.stage.delete.wizard records.
// All records (represented by ids) will be updated by ppsdw values.
func (c *Client) UpdateProjectProjectStageDeleteWizards(ids []int64, ppsdw *ProjectProjectStageDeleteWizard) error {
	return c.Update(ProjectProjectStageDeleteWizardModel, ids, ppsdw, nil)
}

// DeleteProjectProjectStageDeleteWizard deletes an existing project.project.stage.delete.wizard record.
func (c *Client) DeleteProjectProjectStageDeleteWizard(id int64) error {
	return c.DeleteProjectProjectStageDeleteWizards([]int64{id})
}

// DeleteProjectProjectStageDeleteWizards deletes existing project.project.stage.delete.wizard records.
func (c *Client) DeleteProjectProjectStageDeleteWizards(ids []int64) error {
	return c.Delete(ProjectProjectStageDeleteWizardModel, ids)
}

// GetProjectProjectStageDeleteWizard gets project.project.stage.delete.wizard existing record.
func (c *Client) GetProjectProjectStageDeleteWizard(id int64) (*ProjectProjectStageDeleteWizard, error) {
	ppsdws, err := c.GetProjectProjectStageDeleteWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ppsdws)[0]), nil
}

// GetProjectProjectStageDeleteWizards gets project.project.stage.delete.wizard existing records.
func (c *Client) GetProjectProjectStageDeleteWizards(ids []int64) (*ProjectProjectStageDeleteWizards, error) {
	ppsdws := &ProjectProjectStageDeleteWizards{}
	if err := c.Read(ProjectProjectStageDeleteWizardModel, ids, nil, ppsdws); err != nil {
		return nil, err
	}
	return ppsdws, nil
}

// FindProjectProjectStageDeleteWizard finds project.project.stage.delete.wizard record by querying it with criteria.
func (c *Client) FindProjectProjectStageDeleteWizard(criteria *Criteria) (*ProjectProjectStageDeleteWizard, error) {
	ppsdws := &ProjectProjectStageDeleteWizards{}
	if err := c.SearchRead(ProjectProjectStageDeleteWizardModel, criteria, NewOptions().Limit(1), ppsdws); err != nil {
		return nil, err
	}
	return &((*ppsdws)[0]), nil
}

// FindProjectProjectStageDeleteWizards finds project.project.stage.delete.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectStageDeleteWizards(criteria *Criteria, options *Options) (*ProjectProjectStageDeleteWizards, error) {
	ppsdws := &ProjectProjectStageDeleteWizards{}
	if err := c.SearchRead(ProjectProjectStageDeleteWizardModel, criteria, options, ppsdws); err != nil {
		return nil, err
	}
	return ppsdws, nil
}

// FindProjectProjectStageDeleteWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectStageDeleteWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectProjectStageDeleteWizardModel, criteria, options)
}

// FindProjectProjectStageDeleteWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectProjectStageDeleteWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProjectStageDeleteWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
