# go-odoo

An Odoo API client enabling Go programs to interact with Odoo in a simple and uniform way.

[![GitHub license](https://img.shields.io/github/license/ahuret/go-odoo.svg)](https://github.com/ahuret/go-odoo/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/ahuret/go-odoo?status.svg)](https://pkg.go.dev/github.com/ahuret/go-odoo?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/ahuret/go-odoo)](https://goreportcard.com/report/github.com/ahuret/go-odoo)
[![GitHub issues](https://img.shields.io/github/issues/ahuret/go-odoo.svg)](https://github.com/ahuret/go-odoo/issues)

## Usage

### Generate your models


Define the environment variables to be able to connect to your odoo instance :
If you don't set `ODOO_MODELS` all models will be generated.

```
export ODOO_ADMIN=admin
export ODOO_PASSWORD=password
export ODOO_DATABASE=odoo
export ODOO_URL=http://localhost:8069
export ODOO_MODELS="crm.lead"
export ODOO_REPO_PATH=$GOPATH/src/github.com/ahuret/go-odoo
```

Download library and generate models :
```
go get github.com/ahuret/go-odoo
cd $ODOO_REPO_PATH
go generate
```

That's it ! Your models have been generated !

### Enjoy coding!

(All exemples on this README are based on model `crm.lead`)

```go
package main
	
import (
	odoo "github.com/ahuret/go-odoo"
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

Generated models contains high level functions for interact with models in an easy and golang way.
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
Find is powerful and allow you to query a model and filtering results by given [Criteria and Options](#criteria-and-options)

```go
func (c *Client) FindCrmLeads(criteria *Criteria, options *Options) (*CrmLeads, error) {}
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
one2many and many2many are represented by the `Relation` type and allow you to execute special actions as define [here](https://www.odoo.com/documentation/13.0/reference/orm.html#odoo.models.Model.write).

### Criteria and Options

`Criteria` is a set of criterion and allow you to query models. [More informations](https://www.odoo.com/documentation/13.0/reference/orm.html#search-domains)

`Options` allow you to filter results.

```go
cls, err := c.FindCrmLeads(odoo.NewCriteria().Add("user_id.name", "=", "John Doe"), odoo.NewOptions().Limit(2))
```

## Low level functions

All high level functions are based on basic odoo webservices functions. 

These functions give you more flexibility but less usability. We recommand you to use models functions (high level).

Here are disponible low level functions :

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

- If you have an issue, please report it on the [issue tracker](https://github.com/ahuret/go-odoo/issues)

## Contributors

Antoine Huret (<ahuret@skilld.cloud>)

Jean-Baptiste Guerraz (<jbguerraz@skilld.cloud>)
