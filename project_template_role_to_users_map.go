package odoo

// ProjectTemplateRoleToUsersMap represents project.template.role.to.users.map model.
type ProjectTemplateRoleToUsersMap struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	RoleId      *Many2One `xmlrpc:"role_id,omitempty"`
	UserIds     *Relation `xmlrpc:"user_ids,omitempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectTemplateRoleToUsersMaps represents array of project.template.role.to.users.map model.
type ProjectTemplateRoleToUsersMaps []ProjectTemplateRoleToUsersMap

// ProjectTemplateRoleToUsersMapModel is the odoo model name.
const ProjectTemplateRoleToUsersMapModel = "project.template.role.to.users.map"

// Many2One convert ProjectTemplateRoleToUsersMap to *Many2One.
func (ptrtum *ProjectTemplateRoleToUsersMap) Many2One() *Many2One {
	return NewMany2One(ptrtum.Id.Get(), "")
}

// CreateProjectTemplateRoleToUsersMap creates a new project.template.role.to.users.map model and returns its id.
func (c *Client) CreateProjectTemplateRoleToUsersMap(ptrtum *ProjectTemplateRoleToUsersMap) (int64, error) {
	ids, err := c.CreateProjectTemplateRoleToUsersMaps([]*ProjectTemplateRoleToUsersMap{ptrtum})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTemplateRoleToUsersMap creates a new project.template.role.to.users.map model and returns its id.
func (c *Client) CreateProjectTemplateRoleToUsersMaps(ptrtums []*ProjectTemplateRoleToUsersMap) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptrtums {
		vv = append(vv, v)
	}
	return c.Create(ProjectTemplateRoleToUsersMapModel, vv, nil)
}

// UpdateProjectTemplateRoleToUsersMap updates an existing project.template.role.to.users.map record.
func (c *Client) UpdateProjectTemplateRoleToUsersMap(ptrtum *ProjectTemplateRoleToUsersMap) error {
	return c.UpdateProjectTemplateRoleToUsersMaps([]int64{ptrtum.Id.Get()}, ptrtum)
}

// UpdateProjectTemplateRoleToUsersMaps updates existing project.template.role.to.users.map records.
// All records (represented by ids) will be updated by ptrtum values.
func (c *Client) UpdateProjectTemplateRoleToUsersMaps(ids []int64, ptrtum *ProjectTemplateRoleToUsersMap) error {
	return c.Update(ProjectTemplateRoleToUsersMapModel, ids, ptrtum, nil)
}

// DeleteProjectTemplateRoleToUsersMap deletes an existing project.template.role.to.users.map record.
func (c *Client) DeleteProjectTemplateRoleToUsersMap(id int64) error {
	return c.DeleteProjectTemplateRoleToUsersMaps([]int64{id})
}

// DeleteProjectTemplateRoleToUsersMaps deletes existing project.template.role.to.users.map records.
func (c *Client) DeleteProjectTemplateRoleToUsersMaps(ids []int64) error {
	return c.Delete(ProjectTemplateRoleToUsersMapModel, ids)
}

// GetProjectTemplateRoleToUsersMap gets project.template.role.to.users.map existing record.
func (c *Client) GetProjectTemplateRoleToUsersMap(id int64) (*ProjectTemplateRoleToUsersMap, error) {
	ptrtums, err := c.GetProjectTemplateRoleToUsersMaps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ptrtums)[0]), nil
}

// GetProjectTemplateRoleToUsersMaps gets project.template.role.to.users.map existing records.
func (c *Client) GetProjectTemplateRoleToUsersMaps(ids []int64) (*ProjectTemplateRoleToUsersMaps, error) {
	ptrtums := &ProjectTemplateRoleToUsersMaps{}
	if err := c.Read(ProjectTemplateRoleToUsersMapModel, ids, nil, ptrtums); err != nil {
		return nil, err
	}
	return ptrtums, nil
}

// FindProjectTemplateRoleToUsersMap finds project.template.role.to.users.map record by querying it with criteria.
func (c *Client) FindProjectTemplateRoleToUsersMap(criteria *Criteria) (*ProjectTemplateRoleToUsersMap, error) {
	ptrtums := &ProjectTemplateRoleToUsersMaps{}
	if err := c.SearchRead(ProjectTemplateRoleToUsersMapModel, criteria, NewOptions().Limit(1), ptrtums); err != nil {
		return nil, err
	}
	return &((*ptrtums)[0]), nil
}

// FindProjectTemplateRoleToUsersMaps finds project.template.role.to.users.map records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTemplateRoleToUsersMaps(criteria *Criteria, options *Options) (*ProjectTemplateRoleToUsersMaps, error) {
	ptrtums := &ProjectTemplateRoleToUsersMaps{}
	if err := c.SearchRead(ProjectTemplateRoleToUsersMapModel, criteria, options, ptrtums); err != nil {
		return nil, err
	}
	return ptrtums, nil
}

// FindProjectTemplateRoleToUsersMapIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTemplateRoleToUsersMapIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTemplateRoleToUsersMapModel, criteria, options)
}

// FindProjectTemplateRoleToUsersMapId finds record id by querying it with criteria.
func (c *Client) FindProjectTemplateRoleToUsersMapId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTemplateRoleToUsersMapModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
