package odoo

// ProjectTemplateCreateWizard represents project.template.create.wizard model.
type ProjectTemplateCreateWizard struct {
	AliasDomainId    *Many2One `xmlrpc:"alias_domain_id,omitempty"`
	AliasName        *String   `xmlrpc:"alias_name,omitempty"`
	AllowBillable    *Bool     `xmlrpc:"allow_billable,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	Date             *Time     `xmlrpc:"date,omitempty"`
	DateStart        *Time     `xmlrpc:"date_start,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	PartnerId        *Many2One `xmlrpc:"partner_id,omitempty"`
	RoleToUsersIds   *Relation `xmlrpc:"role_to_users_ids,omitempty"`
	TemplateHasDates *Bool     `xmlrpc:"template_has_dates,omitempty"`
	TemplateId       *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectTemplateCreateWizards represents array of project.template.create.wizard model.
type ProjectTemplateCreateWizards []ProjectTemplateCreateWizard

// ProjectTemplateCreateWizardModel is the odoo model name.
const ProjectTemplateCreateWizardModel = "project.template.create.wizard"

// Many2One convert ProjectTemplateCreateWizard to *Many2One.
func (ptcw *ProjectTemplateCreateWizard) Many2One() *Many2One {
	return NewMany2One(ptcw.Id.Get(), "")
}

// CreateProjectTemplateCreateWizard creates a new project.template.create.wizard model and returns its id.
func (c *Client) CreateProjectTemplateCreateWizard(ptcw *ProjectTemplateCreateWizard) (int64, error) {
	ids, err := c.CreateProjectTemplateCreateWizards([]*ProjectTemplateCreateWizard{ptcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTemplateCreateWizard creates a new project.template.create.wizard model and returns its id.
func (c *Client) CreateProjectTemplateCreateWizards(ptcws []*ProjectTemplateCreateWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptcws {
		vv = append(vv, v)
	}
	return c.Create(ProjectTemplateCreateWizardModel, vv, nil)
}

// UpdateProjectTemplateCreateWizard updates an existing project.template.create.wizard record.
func (c *Client) UpdateProjectTemplateCreateWizard(ptcw *ProjectTemplateCreateWizard) error {
	return c.UpdateProjectTemplateCreateWizards([]int64{ptcw.Id.Get()}, ptcw)
}

// UpdateProjectTemplateCreateWizards updates existing project.template.create.wizard records.
// All records (represented by ids) will be updated by ptcw values.
func (c *Client) UpdateProjectTemplateCreateWizards(ids []int64, ptcw *ProjectTemplateCreateWizard) error {
	return c.Update(ProjectTemplateCreateWizardModel, ids, ptcw, nil)
}

// DeleteProjectTemplateCreateWizard deletes an existing project.template.create.wizard record.
func (c *Client) DeleteProjectTemplateCreateWizard(id int64) error {
	return c.DeleteProjectTemplateCreateWizards([]int64{id})
}

// DeleteProjectTemplateCreateWizards deletes existing project.template.create.wizard records.
func (c *Client) DeleteProjectTemplateCreateWizards(ids []int64) error {
	return c.Delete(ProjectTemplateCreateWizardModel, ids)
}

// GetProjectTemplateCreateWizard gets project.template.create.wizard existing record.
func (c *Client) GetProjectTemplateCreateWizard(id int64) (*ProjectTemplateCreateWizard, error) {
	ptcws, err := c.GetProjectTemplateCreateWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ptcws)[0]), nil
}

// GetProjectTemplateCreateWizards gets project.template.create.wizard existing records.
func (c *Client) GetProjectTemplateCreateWizards(ids []int64) (*ProjectTemplateCreateWizards, error) {
	ptcws := &ProjectTemplateCreateWizards{}
	if err := c.Read(ProjectTemplateCreateWizardModel, ids, nil, ptcws); err != nil {
		return nil, err
	}
	return ptcws, nil
}

// FindProjectTemplateCreateWizard finds project.template.create.wizard record by querying it with criteria.
func (c *Client) FindProjectTemplateCreateWizard(criteria *Criteria) (*ProjectTemplateCreateWizard, error) {
	ptcws := &ProjectTemplateCreateWizards{}
	if err := c.SearchRead(ProjectTemplateCreateWizardModel, criteria, NewOptions().Limit(1), ptcws); err != nil {
		return nil, err
	}
	return &((*ptcws)[0]), nil
}

// FindProjectTemplateCreateWizards finds project.template.create.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTemplateCreateWizards(criteria *Criteria, options *Options) (*ProjectTemplateCreateWizards, error) {
	ptcws := &ProjectTemplateCreateWizards{}
	if err := c.SearchRead(ProjectTemplateCreateWizardModel, criteria, options, ptcws); err != nil {
		return nil, err
	}
	return ptcws, nil
}

// FindProjectTemplateCreateWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTemplateCreateWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTemplateCreateWizardModel, criteria, options)
}

// FindProjectTemplateCreateWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectTemplateCreateWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTemplateCreateWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
