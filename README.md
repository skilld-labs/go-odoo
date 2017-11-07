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
- [x] GetAll
- [x] Create
- [x] Update
- [x] Delete

All models are automatically generated by the model2types.py and models.csv files who take datas and retranscribe into .go files.
All api functions are automatically generated by the model2api.py and models.csv files.

## Usage

```go
import "github.com/skilld-labs/go-odoo/api"
```

You have to construct a new client.

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

If you want to generate your own model, you have to check if it is into the 'models.csv' file or add it.
Then, you generate all models with the python command
```python
python model2types.py
```
Recover your .go file in the 'go-types' folder and add it into the 'types' folder.

# Examples

This is an example of how to create a new sale order :

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
