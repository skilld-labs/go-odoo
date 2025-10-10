package odoo

// ProjectShareCollaboratorWizard represents project.share.collaborator.wizard model.
type ProjectShareCollaboratorWizard struct {
	AccessMode     *Selection `xmlrpc:"access_mode,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	ParentWizardId *Many2One  `xmlrpc:"parent_wizard_id,omitempty"`
	PartnerId      *Many2One  `xmlrpc:"partner_id,omitempty"`
	SendInvitation *Bool      `xmlrpc:"send_invitation,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProjectShareCollaboratorWizards represents array of project.share.collaborator.wizard model.
type ProjectShareCollaboratorWizards []ProjectShareCollaboratorWizard

// ProjectShareCollaboratorWizardModel is the odoo model name.
const ProjectShareCollaboratorWizardModel = "project.share.collaborator.wizard"

// Many2One convert ProjectShareCollaboratorWizard to *Many2One.
func (pscw *ProjectShareCollaboratorWizard) Many2One() *Many2One {
	return NewMany2One(pscw.Id.Get(), "")
}

// CreateProjectShareCollaboratorWizard creates a new project.share.collaborator.wizard model and returns its id.
func (c *Client) CreateProjectShareCollaboratorWizard(pscw *ProjectShareCollaboratorWizard) (int64, error) {
	ids, err := c.CreateProjectShareCollaboratorWizards([]*ProjectShareCollaboratorWizard{pscw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectShareCollaboratorWizard creates a new project.share.collaborator.wizard model and returns its id.
func (c *Client) CreateProjectShareCollaboratorWizards(pscws []*ProjectShareCollaboratorWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pscws {
		vv = append(vv, v)
	}
	return c.Create(ProjectShareCollaboratorWizardModel, vv, nil)
}

// UpdateProjectShareCollaboratorWizard updates an existing project.share.collaborator.wizard record.
func (c *Client) UpdateProjectShareCollaboratorWizard(pscw *ProjectShareCollaboratorWizard) error {
	return c.UpdateProjectShareCollaboratorWizards([]int64{pscw.Id.Get()}, pscw)
}

// UpdateProjectShareCollaboratorWizards updates existing project.share.collaborator.wizard records.
// All records (represented by ids) will be updated by pscw values.
func (c *Client) UpdateProjectShareCollaboratorWizards(ids []int64, pscw *ProjectShareCollaboratorWizard) error {
	return c.Update(ProjectShareCollaboratorWizardModel, ids, pscw, nil)
}

// DeleteProjectShareCollaboratorWizard deletes an existing project.share.collaborator.wizard record.
func (c *Client) DeleteProjectShareCollaboratorWizard(id int64) error {
	return c.DeleteProjectShareCollaboratorWizards([]int64{id})
}

// DeleteProjectShareCollaboratorWizards deletes existing project.share.collaborator.wizard records.
func (c *Client) DeleteProjectShareCollaboratorWizards(ids []int64) error {
	return c.Delete(ProjectShareCollaboratorWizardModel, ids)
}

// GetProjectShareCollaboratorWizard gets project.share.collaborator.wizard existing record.
func (c *Client) GetProjectShareCollaboratorWizard(id int64) (*ProjectShareCollaboratorWizard, error) {
	pscws, err := c.GetProjectShareCollaboratorWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pscws)[0]), nil
}

// GetProjectShareCollaboratorWizards gets project.share.collaborator.wizard existing records.
func (c *Client) GetProjectShareCollaboratorWizards(ids []int64) (*ProjectShareCollaboratorWizards, error) {
	pscws := &ProjectShareCollaboratorWizards{}
	if err := c.Read(ProjectShareCollaboratorWizardModel, ids, nil, pscws); err != nil {
		return nil, err
	}
	return pscws, nil
}

// FindProjectShareCollaboratorWizard finds project.share.collaborator.wizard record by querying it with criteria.
func (c *Client) FindProjectShareCollaboratorWizard(criteria *Criteria) (*ProjectShareCollaboratorWizard, error) {
	pscws := &ProjectShareCollaboratorWizards{}
	if err := c.SearchRead(ProjectShareCollaboratorWizardModel, criteria, NewOptions().Limit(1), pscws); err != nil {
		return nil, err
	}
	return &((*pscws)[0]), nil
}

// FindProjectShareCollaboratorWizards finds project.share.collaborator.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectShareCollaboratorWizards(criteria *Criteria, options *Options) (*ProjectShareCollaboratorWizards, error) {
	pscws := &ProjectShareCollaboratorWizards{}
	if err := c.SearchRead(ProjectShareCollaboratorWizardModel, criteria, options, pscws); err != nil {
		return nil, err
	}
	return pscws, nil
}

// FindProjectShareCollaboratorWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectShareCollaboratorWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectShareCollaboratorWizardModel, criteria, options)
}

// FindProjectShareCollaboratorWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectShareCollaboratorWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectShareCollaboratorWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
