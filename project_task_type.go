package odoo

import (
	"fmt"
)

// ProjectTaskType represents project.task.type model.
type ProjectTaskType struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	Description    *String   `xmlrpc:"description,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Fold           *Bool     `xmlrpc:"fold,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LegendBlocked  *String   `xmlrpc:"legend_blocked,omptempty"`
	LegendDone     *String   `xmlrpc:"legend_done,omptempty"`
	LegendNormal   *String   `xmlrpc:"legend_normal,omptempty"`
	LegendPriority *String   `xmlrpc:"legend_priority,omptempty"`
	MailTemplateId *Many2One `xmlrpc:"mail_template_id,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	ProjectIds     *Relation `xmlrpc:"project_ids,omptempty"`
	Sequence       *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTaskTypes represents array of project.task.type model.
type ProjectTaskTypes []ProjectTaskType

// ProjectTaskTypeModel is the odoo model name.
const ProjectTaskTypeModel = "project.task.type"

// Many2One convert ProjectTaskType to *Many2One.
func (ptt *ProjectTaskType) Many2One() *Many2One {
	return NewMany2One(ptt.Id.Get(), "")
}

// CreateProjectTaskType creates a new project.task.type model and returns its id.
func (c *Client) CreateProjectTaskType(ptt *ProjectTaskType) (int64, error) {
	ids, err := c.Create(ProjectTaskTypeModel, []interface{}{ptt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskType creates a new project.task.type model and returns its id.
func (c *Client) CreateProjectTaskTypes(ptts []*ProjectTaskType) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptts {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskTypeModel, vv)
}

// UpdateProjectTaskType updates an existing project.task.type record.
func (c *Client) UpdateProjectTaskType(ptt *ProjectTaskType) error {
	return c.UpdateProjectTaskTypes([]int64{ptt.Id.Get()}, ptt)
}

// UpdateProjectTaskTypes updates existing project.task.type records.
// All records (represented by ids) will be updated by ptt values.
func (c *Client) UpdateProjectTaskTypes(ids []int64, ptt *ProjectTaskType) error {
	return c.Update(ProjectTaskTypeModel, ids, ptt)
}

// DeleteProjectTaskType deletes an existing project.task.type record.
func (c *Client) DeleteProjectTaskType(id int64) error {
	return c.DeleteProjectTaskTypes([]int64{id})
}

// DeleteProjectTaskTypes deletes existing project.task.type records.
func (c *Client) DeleteProjectTaskTypes(ids []int64) error {
	return c.Delete(ProjectTaskTypeModel, ids)
}

// GetProjectTaskType gets project.task.type existing record.
func (c *Client) GetProjectTaskType(id int64) (*ProjectTaskType, error) {
	ptts, err := c.GetProjectTaskTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptts != nil && len(*ptts) > 0 {
		return &((*ptts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task.type not found", id)
}

// GetProjectTaskTypes gets project.task.type existing records.
func (c *Client) GetProjectTaskTypes(ids []int64) (*ProjectTaskTypes, error) {
	ptts := &ProjectTaskTypes{}
	if err := c.Read(ProjectTaskTypeModel, ids, nil, ptts); err != nil {
		return nil, err
	}
	return ptts, nil
}

// FindProjectTaskType finds project.task.type record by querying it with criteria.
func (c *Client) FindProjectTaskType(criteria *Criteria) (*ProjectTaskType, error) {
	ptts := &ProjectTaskTypes{}
	if err := c.SearchRead(ProjectTaskTypeModel, criteria, NewOptions().Limit(1), ptts); err != nil {
		return nil, err
	}
	if ptts != nil && len(*ptts) > 0 {
		return &((*ptts)[0]), nil
	}
	return nil, fmt.Errorf("project.task.type was not found with criteria %v", criteria)
}

// FindProjectTaskTypes finds project.task.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskTypes(criteria *Criteria, options *Options) (*ProjectTaskTypes, error) {
	ptts := &ProjectTaskTypes{}
	if err := c.SearchRead(ProjectTaskTypeModel, criteria, options, ptts); err != nil {
		return nil, err
	}
	return ptts, nil
}

// FindProjectTaskTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskTypeId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task.type was not found with criteria %v and options %v", criteria, options)
}
