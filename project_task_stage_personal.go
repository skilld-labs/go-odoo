package odoo

// ProjectTaskStagePersonal represents project.task.stage.personal model.
type ProjectTaskStagePersonal struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	StageId     *Many2One `xmlrpc:"stage_id,omitempty"`
	TaskId      *Many2One `xmlrpc:"task_id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectTaskStagePersonals represents array of project.task.stage.personal model.
type ProjectTaskStagePersonals []ProjectTaskStagePersonal

// ProjectTaskStagePersonalModel is the odoo model name.
const ProjectTaskStagePersonalModel = "project.task.stage.personal"

// Many2One convert ProjectTaskStagePersonal to *Many2One.
func (ptsp *ProjectTaskStagePersonal) Many2One() *Many2One {
	return NewMany2One(ptsp.Id.Get(), "")
}

// CreateProjectTaskStagePersonal creates a new project.task.stage.personal model and returns its id.
func (c *Client) CreateProjectTaskStagePersonal(ptsp *ProjectTaskStagePersonal) (int64, error) {
	ids, err := c.CreateProjectTaskStagePersonals([]*ProjectTaskStagePersonal{ptsp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskStagePersonal creates a new project.task.stage.personal model and returns its id.
func (c *Client) CreateProjectTaskStagePersonals(ptsps []*ProjectTaskStagePersonal) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptsps {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskStagePersonalModel, vv, nil)
}

// UpdateProjectTaskStagePersonal updates an existing project.task.stage.personal record.
func (c *Client) UpdateProjectTaskStagePersonal(ptsp *ProjectTaskStagePersonal) error {
	return c.UpdateProjectTaskStagePersonals([]int64{ptsp.Id.Get()}, ptsp)
}

// UpdateProjectTaskStagePersonals updates existing project.task.stage.personal records.
// All records (represented by ids) will be updated by ptsp values.
func (c *Client) UpdateProjectTaskStagePersonals(ids []int64, ptsp *ProjectTaskStagePersonal) error {
	return c.Update(ProjectTaskStagePersonalModel, ids, ptsp, nil)
}

// DeleteProjectTaskStagePersonal deletes an existing project.task.stage.personal record.
func (c *Client) DeleteProjectTaskStagePersonal(id int64) error {
	return c.DeleteProjectTaskStagePersonals([]int64{id})
}

// DeleteProjectTaskStagePersonals deletes existing project.task.stage.personal records.
func (c *Client) DeleteProjectTaskStagePersonals(ids []int64) error {
	return c.Delete(ProjectTaskStagePersonalModel, ids)
}

// GetProjectTaskStagePersonal gets project.task.stage.personal existing record.
func (c *Client) GetProjectTaskStagePersonal(id int64) (*ProjectTaskStagePersonal, error) {
	ptsps, err := c.GetProjectTaskStagePersonals([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ptsps)[0]), nil
}

// GetProjectTaskStagePersonals gets project.task.stage.personal existing records.
func (c *Client) GetProjectTaskStagePersonals(ids []int64) (*ProjectTaskStagePersonals, error) {
	ptsps := &ProjectTaskStagePersonals{}
	if err := c.Read(ProjectTaskStagePersonalModel, ids, nil, ptsps); err != nil {
		return nil, err
	}
	return ptsps, nil
}

// FindProjectTaskStagePersonal finds project.task.stage.personal record by querying it with criteria.
func (c *Client) FindProjectTaskStagePersonal(criteria *Criteria) (*ProjectTaskStagePersonal, error) {
	ptsps := &ProjectTaskStagePersonals{}
	if err := c.SearchRead(ProjectTaskStagePersonalModel, criteria, NewOptions().Limit(1), ptsps); err != nil {
		return nil, err
	}
	return &((*ptsps)[0]), nil
}

// FindProjectTaskStagePersonals finds project.task.stage.personal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskStagePersonals(criteria *Criteria, options *Options) (*ProjectTaskStagePersonals, error) {
	ptsps := &ProjectTaskStagePersonals{}
	if err := c.SearchRead(ProjectTaskStagePersonalModel, criteria, options, ptsps); err != nil {
		return nil, err
	}
	return ptsps, nil
}

// FindProjectTaskStagePersonalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskStagePersonalIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTaskStagePersonalModel, criteria, options)
}

// FindProjectTaskStagePersonalId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskStagePersonalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskStagePersonalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
