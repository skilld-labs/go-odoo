#!/usr/bin/env python

import os
import csv

def camelcase(string):
    return ''.join(x.capitalize() or '_' for x in string.replace('.','_').split('_')).replace('_', '')

def write(model, content):
    p = './api/' + '_'.join(model.split('.')) + '.go'
    f = open(p, 'w')
    f.write(content)
    f.close()
    os.system('gofmt -w ' + p)


def type_model(model):
    return """\ntype {{camelModel}}Service struct {
	client *Client
}

func New{{camelModel}}Service(c *Client) *{{camelModel}}Service {
	return &{{camelModel}}Service{client: c}
}

func (svc *{{camelModel}}Service) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.{{camelModel}}Model, name)
}

func (svc *{{camelModel}}Service) GetByIds(ids []int64) (*types.{{camelModel}}s, error) {
	{{variable_name}} := &types.{{camelModel}}s{}
	return {{variable_name}}, svc.client.getByIds(types.{{camelModel}}Model, ids, {{variable_name}})
}

func (svc *{{camelModel}}Service) GetByName(name string) (*types.{{camelModel}}s, error) {
	{{variable_name}} := &types.{{camelModel}}s{}
	return {{variable_name}}, svc.client.getByName(types.{{camelModel}}Model, name, {{variable_name}})
}

func (svc *{{camelModel}}Service) GetByField(field string, value string) (*types.{{camelModel}}s, error) {
	{{variable_name}} := &types.{{camelModel}}s{}
	return {{variable_name}}, svc.client.getByField(types.{{camelModel}}Model, field, value, {{variable_name}})
}

func (svc *{{camelModel}}Service) GetAll() (*types.{{camelModel}}s, error) {
	{{variable_name}} := &types.{{camelModel}}s{}
	return {{variable_name}}, svc.client.getAll(types.{{camelModel}}Model, {{variable_name}})
}

func (svc *{{camelModel}}Service) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.{{camelModel}}Model, fields, relations)
}

func (svc *{{camelModel}}Service) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.{{camelModel}}Model, ids, fields, relations)
}

func (svc *{{camelModel}}Service) Delete(ids []int64) error {
	return svc.client.delete(types.{{camelModel}}Model, ids)
}
""".replace('{{camelModel}}', camelcase(model)).replace('{{variable_name}}', camelcase(model)[0].lower())

def add_imports(content, imports):
    if len(imports) > 0:
        s = 'import (\n'
        for i in imports:
            s += '"' + i + '"\n'
        s += ')\n\n'
        content = s + content
    return content

input_file = csv.DictReader(open('./models.csv'))
model = ''
content = ''
imports = []

for row in input_file:
    if row['model'] != '':
        if model != '':
            content += type_model(model)
            content = add_imports(content, imports)
            content = 'package api\n' + content
            write(model, content)
        model = row['model']
        content = ''
        imports = []
        imports.append("github.com/morezig/go-odoo/types")
        struct = 'type ' + camelcase(model) + ' struct {\n'
        nil_struct  = 'type ' + camelcase(model) + 'Nil struct {\n'
content += type_model(model)
content = add_imports(content, imports)
content = 'package api\n' + content
write(model, content)
