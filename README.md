# go-odoo

An Odoo API client enabling Go programs to interact with Odoo in a simple and uniform way.

[![GitHub license](https://img.shields.io/github/license/skilld-labs/go-odoo.svg)](https://github.com/skilld-labs/go-odoo/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/skilld-labs/go-odoo?status.svg)](https://pkg.go.dev/github.com/skilld-labs/go-odoo?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/skilld-labs/go-odoo)](https://goreportcard.com/report/github.com/skilld-labs/go-odoo)
[![GitHub issues](https://img.shields.io/github/issues/skilld-labs/go-odoo.svg)](https://github.com/skilld-labs/go-odoo/issues)

## Usage

### Generate your models

```bash
./generator/generator -u admin_name -p admin_password -d database_name -o /the/directory/you/want/the/files/to/be/generated/in --url http://localhost:8069 -t ./generator/cmd/tmpl/model.tmpl -m crm.lead,res.users
```

That's it ! Your models have been generated !

### Current generated models

#### Core models

Core models are `ir_model.go` and `ir_model_fields.go` since there are used to generate models.

It is **highly recommanded** to not remove them, since you would not be able to generate models again.

#### Custom skilld-labs models

All others models (not core one) are specific to skilld-labs usage. They use our own odoo instance which is **version 11**. (note that models structure changed between odoo major versions).

If you're ok to work with those models, you can use this library instance, if not you should fork the repository and generate you own models by following steps above.

### Enjoy coding!

(All exemples on this README are based on model `crm.lead`)

```go
package main

import (
	odoo "github.com/skilld-labs/go-odoo"
)

func main() {
	c, err := odoo.NewClient(&odoo.ClientConfig{
		Admin:    "admin_name",
		Password: "admin_password",
		Database: "database_name",
		URL:      "http://localhost:8069",
	})
	if err != nil {
		log.Fatal(err)
	}
	crm := &odoo.CrmLead{
		Name: odoo.NewString("my first opportunity"),
	}
	if id, err := c.CreateCrmLead(crm); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("the id of the new crm.lead is %d", id)
	}
}
```

## Models

Generated models contains high level functions to interact with models in an easy and golang way.
It covers the most common usage and contains for each models those functions :

### Create
```go
func (c *Client) CreateCrmLead(cl *CrmLead) (int64, error) {}
func (c *Client) CreateCrmLeads(cls []*CrmLead) ([]int64, error) {} // !! Only for odoo 12+ versions !!
```

### Update
```go
func (c *Client) UpdateCrmLead(cl *CrmLead) error {}
func (c *Client) UpdateCrmLeads(ids []int64, cl *CrmLead) error {}
```

### Delete
```go
func (c *Client) DeleteCrmLead(id int64) error {}
func (c *Client) DeleteCrmLeads(ids []int64) error {}
```

### Get
```go
func (c *Client) GetCrmLead(id int64) (*CrmLead, error) {}
func (c *Client) GetCrmLeads(ids []int64) (*CrmLeads, error) {}
```

### Find
Find is powerful and allow you to query a model and filter results. [Criteria and Options](#criteria-and-options)

```go
func (c *Client) FindCrmLead(criteria *Criteria) (*CrmLead, error) {}
func (c *Client) FindCrmLeads(criteria *Criteria, options *Options) (*CrmLeads, error) {}
```

### Conversion
Generated models can be converted to `Many2One` easily.
```go
func (cl *CrmLead) Many2One() *Many2One {}
```

## Types

The library contains custom types to improve the usability :

### Basic types

```go
func NewString(v string) *String {}
func (s *String) Get() string {}

func NewInt(v int64) *Int {}
func (i *Int) Get() int64 {}

func NewBool(v bool) *Bool {}
func (b *Bool) Get() bool {}

func NewSelection(v interface{}) *Selection {}
func (s *Selection) Get() (interface{}) {}

func NewTime(v time.Time) *Time {}
func (t *Time) Get() time.Time {}

func NewFloat(v float64) *Float {}
func (f *Float) Get() float64 {}
```

### Relational types

```go
func NewMany2One(id int64, name string) *Many2One {}
func NewUnassignedMany2One() *Many2One {}
func (m *Many2One) Get() int64 {}

func NewRelation() *Relation {}
func (r *Relation) Get() []int64 {}
```
one2many and many2many are represented by the `Relation` type and allow you to execute special actions as defined [here](https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write).

### Criteria and Options

`Criteria` is a set of Criterion and allow you to query models. [More informations](https://www.odoo.com/documentation/13.0/reference/orm.html#search-domains)

#### Combined Criterions

Criterions can be combined using AND (arity 2), OR (arity 2) and NOT (arity 1) operators.
Criteria have And, Or and Not methods to be able to do such query eg:

```go
c := odoo.NewCriteria().Or(
	odoo.NewCriterion("user_id.name", "=", "Jane Doe"),
	odoo.NewCriterion("user_id.name", "=", "John Doe"),
)
```

`Options` allow you to filter results.

```go
cls, err := c.FindCrmLeads(odoo.NewCriteria().Add("user_id.name", "=", "John Doe"), odoo.NewOptions().Limit(2))
```

## Low level functions

All high level functions are based on basic odoo webservices functions.

These functions give you more flexibility but less usability. We recommand you to use models functions (high level).

Here are available low level functions :

```go
func (c *Client) Create(model string, values []interface{}, options *Options) ([]int64, error) {} !! Creating multiple instances is only for odoo 12+ versions !!
func (c *Client) Update(model string, ids []int64, values interface{}, options *Options) error {}
func (c *Client) Delete(model string, ids []int64) error {}
func (c *Client) SearchRead(model string, criteria *Criteria, options *Options, elem interface{}) error {}
func (c *Client) Read(model string, ids []int64, options *Options, elem interface{}) error {}
func (c *Client) Count(model string, criteria *Criteria, options *Options) (int64, error) {}
func (c *Client) Search(model string, criteria *Criteria, options *Options) ([]int64, error) {}
func (c *Client) FieldsGet(model string, options *Options) (map[string]interface{}, error) {}
func (c *Client) ExecuteKw(method, model string, args []interface{}, options *Options) (interface{}, error) {}
```

## Todo

- Tests
- Modular template

## Issues

- If you have an issue, please report it on the [issue tracker](https://github.com/skilld-labs/go-odoo/issues)
