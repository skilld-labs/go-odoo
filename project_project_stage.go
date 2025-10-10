package odoo

// ProjectProjectStage represents project.project.stage model.
type ProjectProjectStage struct {
	Active         *Bool     `xmlrpc:"active,omitempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Fold           *Bool     `xmlrpc:"fold,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	MailTemplateId *Many2One `xmlrpc:"mail_template_id,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	SmsTemplateId  *Many2One `xmlrpc:"sms_template_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectProjectStages represents array of project.project.stage model.
type ProjectProjectStages []ProjectProjectStage

// ProjectProjectStageModel is the odoo model name.
const ProjectProjectStageModel = "project.project.stage"

// Many2One convert ProjectProjectStage to *Many2One.
func (pps *ProjectProjectStage) Many2One() *Many2One {
	return NewMany2One(pps.Id.Get(), "")
}

// CreateProjectProjectStage creates a new project.project.stage model and returns its id.
func (c *Client) CreateProjectProjectStage(pps *ProjectProjectStage) (int64, error) {
	ids, err := c.CreateProjectProjectStages([]*ProjectProjectStage{pps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectProjectStage creates a new project.project.stage model and returns its id.
func (c *Client) CreateProjectProjectStages(ppss []*ProjectProjectStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range ppss {
		vv = append(vv, v)
	}
	return c.Create(ProjectProjectStageModel, vv, nil)
}

// UpdateProjectProjectStage updates an existing project.project.stage record.
func (c *Client) UpdateProjectProjectStage(pps *ProjectProjectStage) error {
	return c.UpdateProjectProjectStages([]int64{pps.Id.Get()}, pps)
}

// UpdateProjectProjectStages updates existing project.project.stage records.
// All records (represented by ids) will be updated by pps values.
func (c *Client) UpdateProjectProjectStages(ids []int64, pps *ProjectProjectStage) error {
	return c.Update(ProjectProjectStageModel, ids, pps, nil)
}

// DeleteProjectProjectStage deletes an existing project.project.stage record.
func (c *Client) DeleteProjectProjectStage(id int64) error {
	return c.DeleteProjectProjectStages([]int64{id})
}

// DeleteProjectProjectStages deletes existing project.project.stage records.
func (c *Client) DeleteProjectProjectStages(ids []int64) error {
	return c.Delete(ProjectProjectStageModel, ids)
}

// GetProjectProjectStage gets project.project.stage existing record.
func (c *Client) GetProjectProjectStage(id int64) (*ProjectProjectStage, error) {
	ppss, err := c.GetProjectProjectStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ppss)[0]), nil
}

// GetProjectProjectStages gets project.project.stage existing records.
func (c *Client) GetProjectProjectStages(ids []int64) (*ProjectProjectStages, error) {
	ppss := &ProjectProjectStages{}
	if err := c.Read(ProjectProjectStageModel, ids, nil, ppss); err != nil {
		return nil, err
	}
	return ppss, nil
}

// FindProjectProjectStage finds project.project.stage record by querying it with criteria.
func (c *Client) FindProjectProjectStage(criteria *Criteria) (*ProjectProjectStage, error) {
	ppss := &ProjectProjectStages{}
	if err := c.SearchRead(ProjectProjectStageModel, criteria, NewOptions().Limit(1), ppss); err != nil {
		return nil, err
	}
	return &((*ppss)[0]), nil
}

// FindProjectProjectStages finds project.project.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectStages(criteria *Criteria, options *Options) (*ProjectProjectStages, error) {
	ppss := &ProjectProjectStages{}
	if err := c.SearchRead(ProjectProjectStageModel, criteria, options, ppss); err != nil {
		return nil, err
	}
	return ppss, nil
}

// FindProjectProjectStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectProjectStageModel, criteria, options)
}

// FindProjectProjectStageId finds record id by querying it with criteria.
func (c *Client) FindProjectProjectStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProjectStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
