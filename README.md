# go-odoo

An Odoo API client enabling Go programs to interact with Odoo in a simple and uniform way.

[![GitHub license](https://img.shields.io/github/license/skilld-labs/go-odoo.svg)](https://github.com/skilld-labs/go-odoo/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/skilld-labs/go-odoo?status.svg)](https://pkg.go.dev/github.com/skilld-labs/go-odoo?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/skilld-labs/go-odoo)](https://goreportcard.com/report/github.com/skilld-labs/go-odoo)
[![GitHub issues](https://img.shields.io/github/issues/skilld-labs/go-odoo.svg)](https://github.com/skilld-labs/go-odoo/issues)

## Usage

### Generate your models

**Note: Generating models require to follow instructions in GOPATH mode. Refactoring for go modules will come soon.**

Define the environment variables to be able to connect to your odoo instance :

(Don't set `ODOO_MODELS` if you want all your models to be generated)

```
export ODOO_ADMIN=admin // ensure the user has sufficient permissions to generate models
export ODOO_PASSWORD=password
export ODOO_DATABASE=odoo
export ODOO_URL=http://localhost:8069
export ODOO_MODELS="crm.lead"
```

`ODOO_REPO_PATH` is the path where the repository will be downloaded (by default its GOPATH):
```
export ODOO_REPO_PATH=$(echo $GOPATH | awk -F ':' '{ print $1 }')/src/github.com/skilld-labs/go-odoo
```

Download library and generate models :
```
go get github.com/skilld-labs/go-odoo
cd $ODOO_REPO_PATH
ls | grep -v "conversion.go\|generator\|go.mod\|go-odoo-generator\|go.sum\|ir_model_fields.go\|ir_model.go\|LICENSE\|odoo.go\|README.md\|types.go\|version.go" // keep only go-odoo core files
go generate
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
		Admin:    "admin",
		Password: "password",
		Database: "odoo",
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
func (m *Many2One) Get() int64 {}

func NewRelation() *Relation {}
func (r *Relation) Get() []int64 {}
```
one2many and many2many are represented by the `Relation` type and allow you to execute special actions as defined [here](https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write).

### Criteria and Options

`Criteria` is a set of criterion and allow you to query models. [More informations](https://www.odoo.com/documentation/13.0/reference/orm.html#search-domains)

`Options` allow you to filter results.

```go
cls, err := c.FindCrmLeads(odoo.NewCriteria().Add("user_id.name", "=", "John Doe"), odoo.NewOptions().Limit(2))
```

## Low level functions

All high level functions are based on basic odoo webservices functions.

These functions give you more flexibility but less usability. We recommand you to use models functions (high level).

Here are available low level functions :

```go
func (c *Client) Create(model string, values interface{}) (int64, error) {}
func (c *Client) Update(model string, ids []int64, values interface{}) error {}
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

## Contributors

Antoine Huret (https://github.com/ahuret)

Jean-Baptiste Guerraz (https://github.com/jbguerraz)
