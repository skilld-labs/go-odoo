package api

import (
	"../types"
)

var ProjectProjectModel string = "project.project"

type ProjectProjectService struct {
	client *Client
}

func (c *Client) NewProjectProjectService() {
	c.ProjectProject = &ProjectProjectService{client: c}
	return
}

func (s *ProjectProjectService) GetIdByName(name string) ([]int, error) {
	var args []interface{}
	var ids []int
	args = append(args, []interface{}{"name", "=", name})
	err := s.client.Search(ProjectProjectModel, args, nil, &ids)
	return ids, err
}

func (s *ProjectProjectService) GetByIds(ids []int) (*types.ProjectProjects, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.ProjectProjects{}
	err := s.client.Read(ProjectProjectModel, args, nil, so)
	return so, err
}

func (s *ProjectProjectService) GetByName(name string) (*types.ProjectProjects, error) {
	var args []interface{}
	var args2 []interface{}
	args2 = append(args2, []string{"name", "=", name})
	args = append(args, args2)
	so := &types.ProjectProjects{}
	err := s.client.SearchRead(ProjectProjectModel, args, nil, so)
	return so, err
}

func (s *ProjectProjectService) GetAll() (*types.ProjectProjects, error) {
	var args []interface{}
	var args2 []interface{}
	args = append(args, args2)
	so := &types.ProjectProjects{}
	err := s.client.SearchRead(ProjectProjectModel, args, nil, so)
	return so, err
}

func (s *ProjectProjectService) Create(required *types.RequiredProjectProjectFields, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) (int, error) {
	var args []interface{}
	var id int
	fields["name"] = required.Name
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Create(ProjectProjectModel, args, &id)
	return id, err
}

func (s *ProjectProjectService) Update(ids []int, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) error {
	var args []interface{}
	args = append(args, ids)
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Update(ProjectProjectModel, args)
	return err
}

func (s *ProjectProjectService) Delete(ids []int) error {
	var args []interface{}
	args = append(args, ids)
	return s.client.Delete("sale.order", args)
}
