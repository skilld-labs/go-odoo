package odoo

// ProjectRole represents project.role model.
type ProjectRole struct {
	Active      *Bool     `xmlrpc:"active,omitempty"`
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectRoles represents array of project.role model.
type ProjectRoles []ProjectRole

// ProjectRoleModel is the odoo model name.
const ProjectRoleModel = "project.role"

// Many2One convert ProjectRole to *Many2One.
func (pr *ProjectRole) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreateProjectRole creates a new project.role model and returns its id.
func (c *Client) CreateProjectRole(pr *ProjectRole) (int64, error) {
	ids, err := c.CreateProjectRoles([]*ProjectRole{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectRole creates a new project.role model and returns its id.
func (c *Client) CreateProjectRoles(prs []*ProjectRole) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(ProjectRoleModel, vv, nil)
}

// UpdateProjectRole updates an existing project.role record.
func (c *Client) UpdateProjectRole(pr *ProjectRole) error {
	return c.UpdateProjectRoles([]int64{pr.Id.Get()}, pr)
}

// UpdateProjectRoles updates existing project.role records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdateProjectRoles(ids []int64, pr *ProjectRole) error {
	return c.Update(ProjectRoleModel, ids, pr, nil)
}

// DeleteProjectRole deletes an existing project.role record.
func (c *Client) DeleteProjectRole(id int64) error {
	return c.DeleteProjectRoles([]int64{id})
}

// DeleteProjectRoles deletes existing project.role records.
func (c *Client) DeleteProjectRoles(ids []int64) error {
	return c.Delete(ProjectRoleModel, ids)
}

// GetProjectRole gets project.role existing record.
func (c *Client) GetProjectRole(id int64) (*ProjectRole, error) {
	prs, err := c.GetProjectRoles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetProjectRoles gets project.role existing records.
func (c *Client) GetProjectRoles(ids []int64) (*ProjectRoles, error) {
	prs := &ProjectRoles{}
	if err := c.Read(ProjectRoleModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProjectRole finds project.role record by querying it with criteria.
func (c *Client) FindProjectRole(criteria *Criteria) (*ProjectRole, error) {
	prs := &ProjectRoles{}
	if err := c.SearchRead(ProjectRoleModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindProjectRoles finds project.role records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectRoles(criteria *Criteria, options *Options) (*ProjectRoles, error) {
	prs := &ProjectRoles{}
	if err := c.SearchRead(ProjectRoleModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProjectRoleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectRoleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectRoleModel, criteria, options)
}

// FindProjectRoleId finds record id by querying it with criteria.
func (c *Client) FindProjectRoleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectRoleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
