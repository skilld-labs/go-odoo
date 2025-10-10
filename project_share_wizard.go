package odoo

// ProjectShareWizard represents project.share.wizard model.
type ProjectShareWizard struct {
	AccessWarning      *String   `xmlrpc:"access_warning,omitempty"`
	CollaboratorIds    *Relation `xmlrpc:"collaborator_ids,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	ExistingPartnerIds *Relation `xmlrpc:"existing_partner_ids,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	Note               *String   `xmlrpc:"note,omitempty"`
	PartnerIds         *Relation `xmlrpc:"partner_ids,omitempty"`
	ResId              *Int      `xmlrpc:"res_id,omitempty"`
	ResModel           *String   `xmlrpc:"res_model,omitempty"`
	ResourceRef        *String   `xmlrpc:"resource_ref,omitempty"`
	ShareLink          *String   `xmlrpc:"share_link,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectShareWizards represents array of project.share.wizard model.
type ProjectShareWizards []ProjectShareWizard

// ProjectShareWizardModel is the odoo model name.
const ProjectShareWizardModel = "project.share.wizard"

// Many2One convert ProjectShareWizard to *Many2One.
func (psw *ProjectShareWizard) Many2One() *Many2One {
	return NewMany2One(psw.Id.Get(), "")
}

// CreateProjectShareWizard creates a new project.share.wizard model and returns its id.
func (c *Client) CreateProjectShareWizard(psw *ProjectShareWizard) (int64, error) {
	ids, err := c.CreateProjectShareWizards([]*ProjectShareWizard{psw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectShareWizard creates a new project.share.wizard model and returns its id.
func (c *Client) CreateProjectShareWizards(psws []*ProjectShareWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range psws {
		vv = append(vv, v)
	}
	return c.Create(ProjectShareWizardModel, vv, nil)
}

// UpdateProjectShareWizard updates an existing project.share.wizard record.
func (c *Client) UpdateProjectShareWizard(psw *ProjectShareWizard) error {
	return c.UpdateProjectShareWizards([]int64{psw.Id.Get()}, psw)
}

// UpdateProjectShareWizards updates existing project.share.wizard records.
// All records (represented by ids) will be updated by psw values.
func (c *Client) UpdateProjectShareWizards(ids []int64, psw *ProjectShareWizard) error {
	return c.Update(ProjectShareWizardModel, ids, psw, nil)
}

// DeleteProjectShareWizard deletes an existing project.share.wizard record.
func (c *Client) DeleteProjectShareWizard(id int64) error {
	return c.DeleteProjectShareWizards([]int64{id})
}

// DeleteProjectShareWizards deletes existing project.share.wizard records.
func (c *Client) DeleteProjectShareWizards(ids []int64) error {
	return c.Delete(ProjectShareWizardModel, ids)
}

// GetProjectShareWizard gets project.share.wizard existing record.
func (c *Client) GetProjectShareWizard(id int64) (*ProjectShareWizard, error) {
	psws, err := c.GetProjectShareWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*psws)[0]), nil
}

// GetProjectShareWizards gets project.share.wizard existing records.
func (c *Client) GetProjectShareWizards(ids []int64) (*ProjectShareWizards, error) {
	psws := &ProjectShareWizards{}
	if err := c.Read(ProjectShareWizardModel, ids, nil, psws); err != nil {
		return nil, err
	}
	return psws, nil
}

// FindProjectShareWizard finds project.share.wizard record by querying it with criteria.
func (c *Client) FindProjectShareWizard(criteria *Criteria) (*ProjectShareWizard, error) {
	psws := &ProjectShareWizards{}
	if err := c.SearchRead(ProjectShareWizardModel, criteria, NewOptions().Limit(1), psws); err != nil {
		return nil, err
	}
	return &((*psws)[0]), nil
}

// FindProjectShareWizards finds project.share.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectShareWizards(criteria *Criteria, options *Options) (*ProjectShareWizards, error) {
	psws := &ProjectShareWizards{}
	if err := c.SearchRead(ProjectShareWizardModel, criteria, options, psws); err != nil {
		return nil, err
	}
	return psws, nil
}

// FindProjectShareWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectShareWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectShareWizardModel, criteria, options)
}

// FindProjectShareWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectShareWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectShareWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
