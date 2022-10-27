# go-odoo

An Odoo API client enabling Go programs to interact with Odoo in a simple and uniform way.

## Usage

### Generate your models

Make sure the generator binary is compiled

`cd generator && go build`

Now you can generate your models:

`./generator -u nick@xiatech.co.uk -p test -d big_test --url http://localhost:8069`

That's it ! Your models have been generated !

### Current generated models

#### Core models

Core models are `ir_model.go` and `ir_model_fields.go` since there are used to generate models.

It is **highly recommanded** to not remove them, since you would not be able to generate models again.


### Enjoy coding!

(All examples on this README are based on model `crm.lead`)

```go
package main

import (
	odoo "github.com/xiatechs/go-odoo"
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
