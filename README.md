# go-odoo

A Odoo API client enabling Go programs to interact with Odoo in a simple and uniform way.

## NOTE

Release v0.0.1 (released on 26-07-2017).

## Coverage

This API client package covers all basic functions from the odoo API.
This include all calls to the following services :

- [x] Login
- [x] Create
- [x] Update
- [x] Delete
- [x] Search
- [x] Read
- [x] SearchRead
- [x] SearchCount
- [x] DoRequest

Services listed above are basic low-level functions from Odoo API, there accessible by any client.

There also are high-level functions based on these low-level functions. Each model has its own functions.
Actually we got:

- [x] GetIdsByName
- [x] GetByIds
- [x] GetByName
- [x] GetByField
- [x] GetAll
- [x] Create
- [x] Update
- [x] Delete

## Usage

```go
import "github.com/skilld-labs/go-odoo/api"
```

### First, generate the models you need to use

Generate the binary:
```
cd go-odoo-cli
go build
```

Using go-odoo-cli binary, you can add and update your models
```
./go-odoo-cli add account.invoice --uri http://localhost:8069 -d database -u admin -p password
./go-odoo-cli update base account.analytic.account --uri http://localhost:8069 -d database -u admin -p password
./go-odoo-cli add all --uri http://localhost:8069 -d database -u admin -p password
```

All the needed files are generated and you can start using the api

### Then you have to construct a new client.

```go
	c, err := api.NewClient("http://localhost:8069", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Login("dbName", "Admin", "password")
	if err != nil {
		fmt.Println(err.Error())
	}
```

## Examples

This is an example of how create a new sale order :

```go
package main

import (
	"fmt"

	"github.com/skilld-labs/go-odoo/api"
)

func main() {
	c, err := api.NewClient("http://localhost:8069", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Login("dbName", "admin", "password")
	if err != nil {
		fmt.Println(err.Error())
	}
	//get the sale order service
	s := api.NewSaleOrderService(c)
	//call the function GetAll() linked to the sale order service
	so, err := s.GetAll()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(so)
}
```

## ToDo

- Tests
- New Odoo API functions (ex: report printing)

## Issues

- If you have an issue, please report it on the [issue tracker](https://github.com/skilld-labs/go-odoo/issues)

## Contributors

Antoine Huret (<huret.antoine@yahoo.fr>)

Jean-Baptiste Guerraz (<jbguerraz@gmail.com>)
