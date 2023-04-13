package odoo

import (
	"fmt"
)

// ProjectTags represents project.tags model.
type ProjectTags struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Color       *Int      `xmlrpc:"color,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTagss represents array of project.tags model.
type ProjectTagss []ProjectTags

// ProjectTagsModel is the odoo model name.
const ProjectTagsModel = "project.tags"

// Many2One convert ProjectTags to *Many2One.
func (pt *ProjectTags) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProjectTags creates a new project.tags model and returns its id.
func (c *Client) CreateProjectTags(pt *ProjectTags) (int64, error) {
	ids, err := c.CreateProjectTagss([]*ProjectTags{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTags creates a new project.tags model and returns its id.
func (c *Client) CreateProjectTagss(pts []*ProjectTags) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProjectTagsModel, vv)
}

// UpdateProjectTags updates an existing project.tags record.
func (c *Client) UpdateProjectTags(pt *ProjectTags) error {
	return c.UpdateProjectTagss([]int64{pt.Id.Get()}, pt)
}

// UpdateProjectTagss updates existing project.tags records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProjectTagss(ids []int64, pt *ProjectTags) error {
	return c.Update(ProjectTagsModel, ids, pt)
}

// DeleteProjectTags deletes an existing project.tags record.
func (c *Client) DeleteProjectTags(id int64) error {
	return c.DeleteProjectTagss([]int64{id})
}

// DeleteProjectTagss deletes existing project.tags records.
func (c *Client) DeleteProjectTagss(ids []int64) error {
	return c.Delete(ProjectTagsModel, ids)
}

// GetProjectTags gets project.tags existing record.
func (c *Client) GetProjectTags(id int64) (*ProjectTags, error) {
	pts, err := c.GetProjectTagss([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.tags not found", id)
}

// GetProjectTagss gets project.tags existing records.
func (c *Client) GetProjectTagss(ids []int64) (*ProjectTagss, error) {
	pts := &ProjectTagss{}
	if err := c.Read(ProjectTagsModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTags finds project.tags record by querying it with criteria.
func (c *Client) FindProjectTags(criteria *Criteria) (*ProjectTags, error) {
	pts := &ProjectTagss{}
	if err := c.SearchRead(ProjectTagsModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("project.tags was not found with criteria %v", criteria)
}

// FindProjectTagss finds project.tags records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTagss(criteria *Criteria, options *Options) (*ProjectTagss, error) {
	pts := &ProjectTagss{}
	if err := c.SearchRead(ProjectTagsModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTagsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTagsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTagsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTagsId finds record id by querying it with criteria.
func (c *Client) FindProjectTagsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTagsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.tags was not found with criteria %v and options %v", criteria, options)
}
