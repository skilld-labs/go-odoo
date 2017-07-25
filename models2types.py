#!/usr/bin/env python

import os
import csv

types = {
    'datetime' : 'time.Time',
    'date': 'time.Time',
    'monetary' : 'float64',
    'char': 'string',
    'many2one': 'Many2One',
    'many2many': '[]int64',
    'one2many': '[]int64',
    'integer': 'int64',
    'boolean': 'bool',
    'text': 'string',
    'selection': 'string',
    'float': 'float64',
    'binary': '[]byte',
    'html': 'string',
    'reference': 'string'
}

nil_types = {
    'datetime' : 'interface{}',
    'date': 'interface{}',
    'monetary' : 'interface{}',
    'char': 'interface{}',
    'many2one': 'interface{}',
    'many2many': 'interface{}',
    'one2many': 'interface{}',
    'integer': 'interface{}',
    'boolean': 'bool',
    'text': 'interface{}',
    'selection': 'interface{}',
    'float': 'interface{}',
    'binary': 'interface{}',
    'html': 'interface{}',
    'reference': 'interface{}'
}

def camelcase(string):
    return ''.join(x.capitalize() or '_' for x in string.replace('.','_').split('_'))

def write(model, content):
    p = './python-types/' + '_'.join(model.split('.')) + '.go'
    f = open(p, 'w')
    f.write(content)
    f.close()
    os.system('gofmt -w ' + p)

def plop(model):
    return """\ntype {{model}}s []{{model}}
type {{model}}sNil []{{model}}Nil

func (s *{{model}}) NilableType_() interface{} {
    return &{{model}}Nil{}
}

func (ns *{{model}}Nil) Type_() interface{} {
    s := &{{model}}{}
    return load(ns, s)
}

func (s *{{model}}s) NilableType_() interface{} {
    return &{{model}}sNil{}
}

func (ns *{{model}}sNil) Type_() interface{} {
    s := &{{model}}s{}
    for _, nsi := range *ns {
        *s = append(*s, *nsi.Type_().(*{{model}}))
    }
    return s
}""".replace('{{model}}', camelcase(model))

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
            content += struct + '}\n' + nil_struct + '}\n'
            content += plop(model)
            content = add_imports(content, imports)
            content = 'package types\n' + content
            write(model, content)
        model = row['model']
        content = ''
        imports = []
        struct = 'type ' + camelcase(model) + ' struct {\n'
        nil_struct  = 'type ' + camelcase(model) + 'Nil struct {\n'
    if row['field_id/name'][0:2] != 'x_':
        struct += camelcase(row['field_id/name']) + ' ' + types[row['field_id/ttype']] + ' `xmlrpc:"' + row['field_id/name'] + '"`\n'
        nil_struct += camelcase(row['field_id/name']) + ' ' + nil_types[row['field_id/ttype']] + ' `xmlrpc:"' + row['field_id/name'] + '"`\n'
        if 'time.Time' == types[row['field_id/ttype']]:
            imports.append('time')
content += struct + '}\n' + nil_struct + '}\n'
content += plop(model)
content = add_imports(content, imports)
content = 'package types\n' + content
write(model, content)


