package odoo

// StockMoveLine represents stock.move.line model.
type StockMoveLine struct {
	CarrierId                    *Many2One  `xmlrpc:"carrier_id,omitempty"`
	CheckIds                     *Relation  `xmlrpc:"check_ids,omitempty"`
	CheckState                   *Selection `xmlrpc:"check_state,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	ConsumeLineIds               *Relation  `xmlrpc:"consume_line_ids,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                         *Time      `xmlrpc:"date,omitempty"`
	DescriptionBomLine           *String    `xmlrpc:"description_bom_line,omitempty"`
	DescriptionPicking           *String    `xmlrpc:"description_picking,omitempty"`
	DestinationCountryCode       *String    `xmlrpc:"destination_country_code,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	DummyId                      *String    `xmlrpc:"dummy_id,omitempty"`
	ExpirationDate               *Time      `xmlrpc:"expiration_date,omitempty"`
	HideLot                      *Bool      `xmlrpc:"hide_lot,omitempty"`
	HideLotName                  *Bool      `xmlrpc:"hide_lot_name,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	Image1920                    *String    `xmlrpc:"image_1920,omitempty"`
	IsExpired                    *Bool      `xmlrpc:"is_expired,omitempty"`
	IsInventory                  *Bool      `xmlrpc:"is_inventory,omitempty"`
	IsLocked                     *Bool      `xmlrpc:"is_locked,omitempty"`
	LocationDestId               *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationDestUsage            *Selection `xmlrpc:"location_dest_usage,omitempty"`
	LocationId                   *Many2One  `xmlrpc:"location_id,omitempty"`
	LocationProcessed            *Bool      `xmlrpc:"location_processed,omitempty"`
	LocationUsage                *Selection `xmlrpc:"location_usage,omitempty"`
	LotId                        *Many2One  `xmlrpc:"lot_id,omitempty"`
	LotName                      *String    `xmlrpc:"lot_name,omitempty"`
	LotsVisible                  *Bool      `xmlrpc:"lots_visible,omitempty"`
	ManualConsumption            *Bool      `xmlrpc:"manual_consumption,omitempty"`
	MoveId                       *Many2One  `xmlrpc:"move_id,omitempty"`
	Origin                       *String    `xmlrpc:"origin,omitempty"`
	OwnerId                      *Many2One  `xmlrpc:"owner_id,omitempty"`
	PackageId                    *Many2One  `xmlrpc:"package_id,omitempty"`
	PackageLevelId               *Many2One  `xmlrpc:"package_level_id,omitempty"`
	ParentLocationDestId         *Many2One  `xmlrpc:"parent_location_dest_id,omitempty"`
	ParentLocationId             *Many2One  `xmlrpc:"parent_location_id,omitempty"`
	PickTypeCreateComponentsLots *Bool      `xmlrpc:"pick_type_create_components_lots,omitempty"`
	Picked                       *Bool      `xmlrpc:"picked,omitempty"`
	PickingCode                  *Selection `xmlrpc:"picking_code,omitempty"`
	PickingId                    *Many2One  `xmlrpc:"picking_id,omitempty"`
	PickingLocationDestId        *Many2One  `xmlrpc:"picking_location_dest_id,omitempty"`
	PickingLocationId            *Many2One  `xmlrpc:"picking_location_id,omitempty"`
	PickingPartnerId             *Many2One  `xmlrpc:"picking_partner_id,omitempty"`
	PickingTypeEntirePacks       *Bool      `xmlrpc:"picking_type_entire_packs,omitempty"`
	PickingTypeId                *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	PickingTypeUseCreateLots     *Bool      `xmlrpc:"picking_type_use_create_lots,omitempty"`
	PickingTypeUseExistingLots   *Bool      `xmlrpc:"picking_type_use_existing_lots,omitempty"`
	ProduceLineIds               *Relation  `xmlrpc:"produce_line_ids,omitempty"`
	ProductBarcode               *String    `xmlrpc:"product_barcode,omitempty"`
	ProductCategoryName          *String    `xmlrpc:"product_category_name,omitempty"`
	ProductId                    *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductPackagingId           *Many2One  `xmlrpc:"product_packaging_id,omitempty"`
	ProductPackagingQty          *Float     `xmlrpc:"product_packaging_qty,omitempty"`
	ProductPackagingUomQty       *Float     `xmlrpc:"product_packaging_uom_qty,omitempty"`
	ProductReferenceCode         *String    `xmlrpc:"product_reference_code,omitempty"`
	ProductStockQuantIds         *Relation  `xmlrpc:"product_stock_quant_ids,omitempty"`
	ProductUomCategoryId         *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId                 *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	ProductionId                 *Many2One  `xmlrpc:"production_id,omitempty"`
	QtyDone                      *Float     `xmlrpc:"qty_done,omitempty"`
	QualityCheckIds              *Relation  `xmlrpc:"quality_check_ids,omitempty"`
	QuantId                      *Many2One  `xmlrpc:"quant_id,omitempty"`
	Quantity                     *Float     `xmlrpc:"quantity,omitempty"`
	QuantityProductUom           *Float     `xmlrpc:"quantity_product_uom,omitempty"`
	Reference                    *String    `xmlrpc:"reference,omitempty"`
	ResultPackageId              *Many2One  `xmlrpc:"result_package_id,omitempty"`
	SalePrice                    *Float     `xmlrpc:"sale_price,omitempty"`
	State                        *Selection `xmlrpc:"state,omitempty"`
	Tracking                     *Selection `xmlrpc:"tracking,omitempty"`
	UseExpirationDate            *Bool      `xmlrpc:"use_expiration_date,omitempty"`
	WorkorderId                  *Many2One  `xmlrpc:"workorder_id,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
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
