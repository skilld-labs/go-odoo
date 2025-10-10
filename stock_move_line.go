package odoo

// StockMoveLine represents stock.move.line model.
type StockMoveLine struct {
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	ConsumeLineIds             *Relation  `xmlrpc:"consume_line_ids,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                       *Time      `xmlrpc:"date,omitempty"`
	DescriptionPicking         *String    `xmlrpc:"description_picking,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	IsInventory                *Bool      `xmlrpc:"is_inventory,omitempty"`
	IsLocked                   *Bool      `xmlrpc:"is_locked,omitempty"`
	LocationDestId             *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationDestUsage          *Selection `xmlrpc:"location_dest_usage,omitempty"`
	LocationId                 *Many2One  `xmlrpc:"location_id,omitempty"`
	LocationUsage              *Selection `xmlrpc:"location_usage,omitempty"`
	LotId                      *Many2One  `xmlrpc:"lot_id,omitempty"`
	LotName                    *String    `xmlrpc:"lot_name,omitempty"`
	LotsVisible                *Bool      `xmlrpc:"lots_visible,omitempty"`
	MoveId                     *Many2One  `xmlrpc:"move_id,omitempty"`
	Origin                     *String    `xmlrpc:"origin,omitempty"`
	OwnerId                    *Many2One  `xmlrpc:"owner_id,omitempty"`
	PackageId                  *Many2One  `xmlrpc:"package_id,omitempty"`
	PackageLevelId             *Many2One  `xmlrpc:"package_level_id,omitempty"`
	Picked                     *Bool      `xmlrpc:"picked,omitempty"`
	PickingCode                *Selection `xmlrpc:"picking_code,omitempty"`
	PickingId                  *Many2One  `xmlrpc:"picking_id,omitempty"`
	PickingLocationDestId      *Many2One  `xmlrpc:"picking_location_dest_id,omitempty"`
	PickingLocationId          *Many2One  `xmlrpc:"picking_location_id,omitempty"`
	PickingPartnerId           *Many2One  `xmlrpc:"picking_partner_id,omitempty"`
	PickingTypeEntirePacks     *Bool      `xmlrpc:"picking_type_entire_packs,omitempty"`
	PickingTypeId              *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	PickingTypeUseCreateLots   *Bool      `xmlrpc:"picking_type_use_create_lots,omitempty"`
	PickingTypeUseExistingLots *Bool      `xmlrpc:"picking_type_use_existing_lots,omitempty"`
	ProduceLineIds             *Relation  `xmlrpc:"produce_line_ids,omitempty"`
	ProductCategoryName        *String    `xmlrpc:"product_category_name,omitempty"`
	ProductId                  *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductPackagingQty        *Float     `xmlrpc:"product_packaging_qty,omitempty"`
	ProductUomCategoryId       *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId               *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	QuantId                    *Many2One  `xmlrpc:"quant_id,omitempty"`
	Quantity                   *Float     `xmlrpc:"quantity,omitempty"`
	QuantityProductUom         *Float     `xmlrpc:"quantity_product_uom,omitempty"`
	Reference                  *String    `xmlrpc:"reference,omitempty"`
	ResultPackageId            *Many2One  `xmlrpc:"result_package_id,omitempty"`
	ScheduledDate              *Time      `xmlrpc:"scheduled_date,omitempty"`
	State                      *Selection `xmlrpc:"state,omitempty"`
	Tracking                   *Selection `xmlrpc:"tracking,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockMoveLines represents array of stock.move.line model.
type StockMoveLines []StockMoveLine

// StockMoveLineModel is the odoo model name.
const StockMoveLineModel = "stock.move.line"

// Many2One convert StockMoveLine to *Many2One.
func (sml *StockMoveLine) Many2One() *Many2One {
	return NewMany2One(sml.Id.Get(), "")
}

// CreateStockMoveLine creates a new stock.move.line model and returns its id.
func (c *Client) CreateStockMoveLine(sml *StockMoveLine) (int64, error) {
	ids, err := c.CreateStockMoveLines([]*StockMoveLine{sml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockMoveLine creates a new stock.move.line model and returns its id.
func (c *Client) CreateStockMoveLines(smls []*StockMoveLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range smls {
		vv = append(vv, v)
	}
	return c.Create(StockMoveLineModel, vv, nil)
}

// UpdateStockMoveLine updates an existing stock.move.line record.
func (c *Client) UpdateStockMoveLine(sml *StockMoveLine) error {
	return c.UpdateStockMoveLines([]int64{sml.Id.Get()}, sml)
}

// UpdateStockMoveLines updates existing stock.move.line records.
// All records (represented by ids) will be updated by sml values.
func (c *Client) UpdateStockMoveLines(ids []int64, sml *StockMoveLine) error {
	return c.Update(StockMoveLineModel, ids, sml, nil)
}

// DeleteStockMoveLine deletes an existing stock.move.line record.
func (c *Client) DeleteStockMoveLine(id int64) error {
	return c.DeleteStockMoveLines([]int64{id})
}

// DeleteStockMoveLines deletes existing stock.move.line records.
func (c *Client) DeleteStockMoveLines(ids []int64) error {
	return c.Delete(StockMoveLineModel, ids)
}

// GetStockMoveLine gets stock.move.line existing record.
func (c *Client) GetStockMoveLine(id int64) (*StockMoveLine, error) {
	smls, err := c.GetStockMoveLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*smls)[0]), nil
}

// GetStockMoveLines gets stock.move.line existing records.
func (c *Client) GetStockMoveLines(ids []int64) (*StockMoveLines, error) {
	smls := &StockMoveLines{}
	if err := c.Read(StockMoveLineModel, ids, nil, smls); err != nil {
		return nil, err
	}
	return smls, nil
}

// FindStockMoveLine finds stock.move.line record by querying it with criteria.
func (c *Client) FindStockMoveLine(criteria *Criteria) (*StockMoveLine, error) {
	smls := &StockMoveLines{}
	if err := c.SearchRead(StockMoveLineModel, criteria, NewOptions().Limit(1), smls); err != nil {
		return nil, err
	}
	return &((*smls)[0]), nil
}

// FindStockMoveLines finds stock.move.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoveLines(criteria *Criteria, options *Options) (*StockMoveLines, error) {
	smls := &StockMoveLines{}
	if err := c.SearchRead(StockMoveLineModel, criteria, options, smls); err != nil {
		return nil, err
	}
	return smls, nil
}

// FindStockMoveLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoveLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockMoveLineModel, criteria, options)
}

// FindStockMoveLineId finds record id by querying it with criteria.
func (c *Client) FindStockMoveLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockMoveLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
