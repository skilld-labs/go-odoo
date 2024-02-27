package odoo

// StockInventoryLine represents stock.inventory.line model.
type StockInventoryLine struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	InventoryId         *Many2One  `xmlrpc:"inventory_id,omitempty"`
	InventoryLocationId *Many2One  `xmlrpc:"inventory_location_id,omitempty"`
	LocationId          *Many2One  `xmlrpc:"location_id,omitempty"`
	LocationName        *String    `xmlrpc:"location_name,omitempty"`
	PackageId           *Many2One  `xmlrpc:"package_id,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	ProdLotId           *Many2One  `xmlrpc:"prod_lot_id,omitempty"`
	ProdlotName         *String    `xmlrpc:"prodlot_name,omitempty"`
	ProductCode         *String    `xmlrpc:"product_code,omitempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductName         *String    `xmlrpc:"product_name,omitempty"`
	ProductQty          *Float     `xmlrpc:"product_qty,omitempty"`
	ProductUomId        *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	TheoreticalQty      *Float     `xmlrpc:"theoretical_qty,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockInventoryLines represents array of stock.inventory.line model.
type StockInventoryLines []StockInventoryLine

// StockInventoryLineModel is the odoo model name.
const StockInventoryLineModel = "stock.inventory.line"

// Many2One convert StockInventoryLine to *Many2One.
func (sil *StockInventoryLine) Many2One() *Many2One {
	return NewMany2One(sil.Id.Get(), "")
}

// CreateStockInventoryLine creates a new stock.inventory.line model and returns its id.
func (c *Client) CreateStockInventoryLine(sil *StockInventoryLine) (int64, error) {
	ids, err := c.CreateStockInventoryLines([]*StockInventoryLine{sil})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockInventoryLine creates a new stock.inventory.line model and returns its id.
func (c *Client) CreateStockInventoryLines(sils []*StockInventoryLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range sils {
		vv = append(vv, v)
	}
	return c.Create(StockInventoryLineModel, vv, nil)
}

// UpdateStockInventoryLine updates an existing stock.inventory.line record.
func (c *Client) UpdateStockInventoryLine(sil *StockInventoryLine) error {
	return c.UpdateStockInventoryLines([]int64{sil.Id.Get()}, sil)
}

// UpdateStockInventoryLines updates existing stock.inventory.line records.
// All records (represented by ids) will be updated by sil values.
func (c *Client) UpdateStockInventoryLines(ids []int64, sil *StockInventoryLine) error {
	return c.Update(StockInventoryLineModel, ids, sil, nil)
}

// DeleteStockInventoryLine deletes an existing stock.inventory.line record.
func (c *Client) DeleteStockInventoryLine(id int64) error {
	return c.DeleteStockInventoryLines([]int64{id})
}

// DeleteStockInventoryLines deletes existing stock.inventory.line records.
func (c *Client) DeleteStockInventoryLines(ids []int64) error {
	return c.Delete(StockInventoryLineModel, ids)
}

// GetStockInventoryLine gets stock.inventory.line existing record.
func (c *Client) GetStockInventoryLine(id int64) (*StockInventoryLine, error) {
	sils, err := c.GetStockInventoryLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sils)[0]), nil
}

// GetStockInventoryLines gets stock.inventory.line existing records.
func (c *Client) GetStockInventoryLines(ids []int64) (*StockInventoryLines, error) {
	sils := &StockInventoryLines{}
	if err := c.Read(StockInventoryLineModel, ids, nil, sils); err != nil {
		return nil, err
	}
	return sils, nil
}

// FindStockInventoryLine finds stock.inventory.line record by querying it with criteria.
func (c *Client) FindStockInventoryLine(criteria *Criteria) (*StockInventoryLine, error) {
	sils := &StockInventoryLines{}
	if err := c.SearchRead(StockInventoryLineModel, criteria, NewOptions().Limit(1), sils); err != nil {
		return nil, err
	}
	return &((*sils)[0]), nil
}

// FindStockInventoryLines finds stock.inventory.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryLines(criteria *Criteria, options *Options) (*StockInventoryLines, error) {
	sils := &StockInventoryLines{}
	if err := c.SearchRead(StockInventoryLineModel, criteria, options, sils); err != nil {
		return nil, err
	}
	return sils, nil
}

// FindStockInventoryLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockInventoryLineModel, criteria, options)
}

// FindStockInventoryLineId finds record id by querying it with criteria.
func (c *Client) FindStockInventoryLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockInventoryLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
