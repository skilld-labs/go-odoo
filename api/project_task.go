package api

import (
	"../types"
)

var ProjectTaskModel string = "project.task"

type ProjectTaskService struct {
	client *Client
}

func (c *Client) NewProjectTaskService() {
	c.ProjectTask = &ProjectTaskService{client: c}
	return
}

func (s *ProjectTaskService) GetIdByName(name string) ([]int, error) {
	var args []interface{}
	var ids []int
	args = append(args, []interface{}{"name", "=", name})
	err := s.client.Search(ProjectTaskModel, args, nil, &ids)
	return ids, err
}

func (s *ProjectTaskService) GetByIds(ids []int) (*types.ProjectTasks, error) {
	var args []interface{}
	args = append(args, ids)
	so := &types.ProjectTasks{}
	err := s.client.Read(ProjectTaskModel, args, nil, so)
	return so, err
}

func (s *ProjectTaskService) GetByName(name string) (*types.ProjectTasks, error) {
	var args []interface{}
	var args2 []interface{}
	args2 = append(args2, []string{"name", "=", name})
	args = append(args, args2)
	so := &types.ProjectTasks{}
	err := s.client.SearchRead(ProjectTaskModel, args, nil, so)
	return so, err
}

func (s *ProjectTaskService) GetAll() (*types.ProjectTasks, error) {
	var args []interface{}
	var args2 []interface{}
	args = append(args, args2)
	so := &types.ProjectTasks{}
	err := s.client.SearchRead(ProjectTaskModel, args, nil, so)
	return so, err
}

func (s *ProjectTaskService) Create(required *types.RequiredProjectTaskFields, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) (int, error) {
	var args []interface{}
	var id int
	fields["name"] = required.Name
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Create(ProjectTaskModel, args, &id)
	return id, err
}

func (s *ProjectTaskService) Update(ids []int, fields map[string]interface{}, fields2Many *types.Handle2ManysStruct) error {
	var args []interface{}
	args = append(args, ids)
	if fields2Many != nil {
		fields = types.Handle2Manys(fields, fields2Many)
	}
	args = append(args, fields)
	err := s.client.Update(ProjectTaskModel, args)
	return err
}

func (s *ProjectTaskService) Delete(ids []int) error {
	var args []interface{}
	args = append(args, ids)
	return s.client.Delete("sale.order", args)
}
