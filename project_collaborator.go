package odoo

// ProjectCollaborator represents project.collaborator model.
type ProjectCollaborator struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	LimitedAccess *Bool     `xmlrpc:"limited_access,omitempty"`
	PartnerEmail  *String   `xmlrpc:"partner_email,omitempty"`
	PartnerId     *Many2One `xmlrpc:"partner_id,omitempty"`
	ProjectId     *Many2One `xmlrpc:"project_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectCollaborators represents array of project.collaborator model.
type ProjectCollaborators []ProjectCollaborator

// ProjectCollaboratorModel is the odoo model name.
const ProjectCollaboratorModel = "project.collaborator"

// Many2One convert ProjectCollaborator to *Many2One.
func (pc *ProjectCollaborator) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreateProjectCollaborator creates a new project.collaborator model and returns its id.
func (c *Client) CreateProjectCollaborator(pc *ProjectCollaborator) (int64, error) {
	ids, err := c.CreateProjectCollaborators([]*ProjectCollaborator{pc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectCollaborator creates a new project.collaborator model and returns its id.
func (c *Client) CreateProjectCollaborators(pcs []*ProjectCollaborator) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcs {
		vv = append(vv, v)
	}
	return c.Create(ProjectCollaboratorModel, vv, nil)
}

// UpdateProjectCollaborator updates an existing project.collaborator record.
func (c *Client) UpdateProjectCollaborator(pc *ProjectCollaborator) error {
	return c.UpdateProjectCollaborators([]int64{pc.Id.Get()}, pc)
}

// UpdateProjectCollaborators updates existing project.collaborator records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdateProjectCollaborators(ids []int64, pc *ProjectCollaborator) error {
	return c.Update(ProjectCollaboratorModel, ids, pc, nil)
}

// DeleteProjectCollaborator deletes an existing project.collaborator record.
func (c *Client) DeleteProjectCollaborator(id int64) error {
	return c.DeleteProjectCollaborators([]int64{id})
}

// DeleteProjectCollaborators deletes existing project.collaborator records.
func (c *Client) DeleteProjectCollaborators(ids []int64) error {
	return c.Delete(ProjectCollaboratorModel, ids)
}

// GetProjectCollaborator gets project.collaborator existing record.
func (c *Client) GetProjectCollaborator(id int64) (*ProjectCollaborator, error) {
	pcs, err := c.GetProjectCollaborators([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// GetProjectCollaborators gets project.collaborator existing records.
func (c *Client) GetProjectCollaborators(ids []int64) (*ProjectCollaborators, error) {
	pcs := &ProjectCollaborators{}
	if err := c.Read(ProjectCollaboratorModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindProjectCollaborator finds project.collaborator record by querying it with criteria.
func (c *Client) FindProjectCollaborator(criteria *Criteria) (*ProjectCollaborator, error) {
	pcs := &ProjectCollaborators{}
	if err := c.SearchRead(ProjectCollaboratorModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// FindProjectCollaborators finds project.collaborator records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCollaborators(criteria *Criteria, options *Options) (*ProjectCollaborators, error) {
	pcs := &ProjectCollaborators{}
	if err := c.SearchRead(ProjectCollaboratorModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindProjectCollaboratorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCollaboratorIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectCollaboratorModel, criteria, options)
}

// FindProjectCollaboratorId finds record id by querying it with criteria.
func (c *Client) FindProjectCollaboratorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectCollaboratorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
