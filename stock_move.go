package odoo

// StockMove represents stock.move model.
type StockMove struct {
	AccountMoveIds           *Relation  `xmlrpc:"account_move_ids,omitempty"`
	Additional               *Bool      `xmlrpc:"additional,omitempty"`
	AllowedOperationIds      *Relation  `xmlrpc:"allowed_operation_ids,omitempty"`
	AnalyticAccountLineIds   *Relation  `xmlrpc:"analytic_account_line_ids,omitempty"`
	Availability             *Float     `xmlrpc:"availability,omitempty"`
	BomLineId                *Many2One  `xmlrpc:"bom_line_id,omitempty"`
	ByproductId              *Many2One  `xmlrpc:"byproduct_id,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	ConsumeUnbuildId         *Many2One  `xmlrpc:"consume_unbuild_id,omitempty"`
	CostShare                *Float     `xmlrpc:"cost_share,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CreatedProductionId      *Many2One  `xmlrpc:"created_production_id,omitempty"`
	CreatedPurchaseLineIds   *Relation  `xmlrpc:"created_purchase_line_ids,omitempty"`
	Date                     *Time      `xmlrpc:"date,omitempty"`
	DateDeadline             *Time      `xmlrpc:"date_deadline,omitempty"`
	DelayAlertDate           *Time      `xmlrpc:"delay_alert_date,omitempty"`
	DescriptionBomLine       *String    `xmlrpc:"description_bom_line,omitempty"`
	DescriptionPicking       *String    `xmlrpc:"description_picking,omitempty"`
	DisplayAssignSerial      *Bool      `xmlrpc:"display_assign_serial,omitempty"`
	DisplayImportLot         *Bool      `xmlrpc:"display_import_lot,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	ForecastAvailability     *Float     `xmlrpc:"forecast_availability,omitempty"`
	ForecastExpectedDate     *Time      `xmlrpc:"forecast_expected_date,omitempty"`
	GroupId                  *Many2One  `xmlrpc:"group_id,omitempty"`
	HasTracking              *Selection `xmlrpc:"has_tracking,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IsDone                   *Bool      `xmlrpc:"is_done,omitempty"`
	IsInitialDemandEditable  *Bool      `xmlrpc:"is_initial_demand_editable,omitempty"`
	IsInventory              *Bool      `xmlrpc:"is_inventory,omitempty"`
	IsLocked                 *Bool      `xmlrpc:"is_locked,omitempty"`
	IsQuantityDoneEditable   *Bool      `xmlrpc:"is_quantity_done_editable,omitempty"`
	LocationDestId           *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationDestUsage        *Selection `xmlrpc:"location_dest_usage,omitempty"`
	LocationId               *Many2One  `xmlrpc:"location_id,omitempty"`
	LocationUsage            *Selection `xmlrpc:"location_usage,omitempty"`
	LotIds                   *Relation  `xmlrpc:"lot_ids,omitempty"`
	ManualConsumption        *Bool      `xmlrpc:"manual_consumption,omitempty"`
	MoveDestIds              *Relation  `xmlrpc:"move_dest_ids,omitempty"`
	MoveLineIds              *Relation  `xmlrpc:"move_line_ids,omitempty"`
	MoveLinesCount           *Int       `xmlrpc:"move_lines_count,omitempty"`
	MoveOrigIds              *Relation  `xmlrpc:"move_orig_ids,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	NextSerial               *String    `xmlrpc:"next_serial,omitempty"`
	NextSerialCount          *Int       `xmlrpc:"next_serial_count,omitempty"`
	OperationId              *Many2One  `xmlrpc:"operation_id,omitempty"`
	OrderFinishedLotId       *Many2One  `xmlrpc:"order_finished_lot_id,omitempty"`
	OrderpointId             *Many2One  `xmlrpc:"orderpoint_id,omitempty"`
	Origin                   *String    `xmlrpc:"origin,omitempty"`
	OriginReturnedMoveId     *Many2One  `xmlrpc:"origin_returned_move_id,omitempty"`
	PackageLevelId           *Many2One  `xmlrpc:"package_level_id,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	Picked                   *Bool      `xmlrpc:"picked,omitempty"`
	PickingCode              *Selection `xmlrpc:"picking_code,omitempty"`
	PickingId                *Many2One  `xmlrpc:"picking_id,omitempty"`
	PickingTypeEntirePacks   *Bool      `xmlrpc:"picking_type_entire_packs,omitempty"`
	PickingTypeId            *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	PriceUnit                *Float     `xmlrpc:"price_unit,omitempty"`
	Priority                 *Selection `xmlrpc:"priority,omitempty"`
	ProcureMethod            *Selection `xmlrpc:"procure_method,omitempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductPackagingId       *Many2One  `xmlrpc:"product_packaging_id,omitempty"`
	ProductPackagingQty      *Float     `xmlrpc:"product_packaging_qty,omitempty"`
	ProductPackagingQuantity *Float     `xmlrpc:"product_packaging_quantity,omitempty"`
	ProductQty               *Float     `xmlrpc:"product_qty,omitempty"`
	ProductQtyAvailable      *Float     `xmlrpc:"product_qty_available,omitempty"`
	ProductTmplId            *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductType              *Selection `xmlrpc:"product_type,omitempty"`
	ProductUom               *Many2One  `xmlrpc:"product_uom,omitempty"`
	ProductUomCategoryId     *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomQty            *Float     `xmlrpc:"product_uom_qty,omitempty"`
	ProductVirtualAvailable  *Float     `xmlrpc:"product_virtual_available,omitempty"`
	ProductionId             *Many2One  `xmlrpc:"production_id,omitempty"`
	PropagateCancel          *Bool      `xmlrpc:"propagate_cancel,omitempty"`
	PurchaseLineId           *Many2One  `xmlrpc:"purchase_line_id,omitempty"`
	Quantity                 *Float     `xmlrpc:"quantity,omitempty"`
	RawMaterialProductionId  *Many2One  `xmlrpc:"raw_material_production_id,omitempty"`
	Reference                *String    `xmlrpc:"reference,omitempty"`
	RequisitionLineIds       *Relation  `xmlrpc:"requisition_line_ids,omitempty"`
	ReservationDate          *Time      `xmlrpc:"reservation_date,omitempty"`
	RestrictPartnerId        *Many2One  `xmlrpc:"restrict_partner_id,omitempty"`
	ReturnedMoveIds          *Relation  `xmlrpc:"returned_move_ids,omitempty"`
	RouteIds                 *Relation  `xmlrpc:"route_ids,omitempty"`
	RuleId                   *Many2One  `xmlrpc:"rule_id,omitempty"`
	SaleLineId               *Many2One  `xmlrpc:"sale_line_id,omitempty"`
	ScrapId                  *Many2One  `xmlrpc:"scrap_id,omitempty"`
	Scrapped                 *Bool      `xmlrpc:"scrapped,omitempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omitempty"`
	ShouldConsumeQty         *Float     `xmlrpc:"should_consume_qty,omitempty"`
	ShowDetailsVisible       *Bool      `xmlrpc:"show_details_visible,omitempty"`
	ShowLotsM2O              *Bool      `xmlrpc:"show_lots_m2o,omitempty"`
	ShowLotsText             *Bool      `xmlrpc:"show_lots_text,omitempty"`
	ShowOperations           *Bool      `xmlrpc:"show_operations,omitempty"`
	ShowQuant                *Bool      `xmlrpc:"show_quant,omitempty"`
	ShowReserved             *Bool      `xmlrpc:"show_reserved,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	StockValuationLayerIds   *Relation  `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	ToRefund                 *Bool      `xmlrpc:"to_refund,omitempty"`
	UnbuildId                *Many2One  `xmlrpc:"unbuild_id,omitempty"`
	UnitFactor               *Float     `xmlrpc:"unit_factor,omitempty"`
	UseExpirationDate        *Bool      `xmlrpc:"use_expiration_date,omitempty"`
	WarehouseId              *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	Weight                   *Float     `xmlrpc:"weight,omitempty"`
	WorkorderId              *Many2One  `xmlrpc:"workorder_id,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockMoves represents array of stock.move model.
type StockMoves []StockMove

// StockMoveModel is the odoo model name.
const StockMoveModel = "stock.move"

// Many2One convert StockMove to *Many2One.
func (sm *StockMove) Many2One() *Many2One {
	return NewMany2One(sm.Id.Get(), "")
}

// CreateStockMove creates a new stock.move model and returns its id.
func (c *Client) CreateStockMove(sm *StockMove) (int64, error) {
	ids, err := c.CreateStockMoves([]*StockMove{sm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockMove creates a new stock.move model and returns its id.
func (c *Client) CreateStockMoves(sms []*StockMove) ([]int64, error) {
	var vv []interface{}
	for _, v := range sms {
		vv = append(vv, v)
	}
	return c.Create(StockMoveModel, vv, nil)
}

// UpdateStockMove updates an existing stock.move record.
func (c *Client) UpdateStockMove(sm *StockMove) error {
	return c.UpdateStockMoves([]int64{sm.Id.Get()}, sm)
}

// UpdateStockMoves updates existing stock.move records.
// All records (represented by ids) will be updated by sm values.
func (c *Client) UpdateStockMoves(ids []int64, sm *StockMove) error {
	return c.Update(StockMoveModel, ids, sm, nil)
}

// DeleteStockMove deletes an existing stock.move record.
func (c *Client) DeleteStockMove(id int64) error {
	return c.DeleteStockMoves([]int64{id})
}

// DeleteStockMoves deletes existing stock.move records.
func (c *Client) DeleteStockMoves(ids []int64) error {
	return c.Delete(StockMoveModel, ids)
}

// GetStockMove gets stock.move existing record.
func (c *Client) GetStockMove(id int64) (*StockMove, error) {
	sms, err := c.GetStockMoves([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sms)[0]), nil
}

// GetStockMoves gets stock.move existing records.
func (c *Client) GetStockMoves(ids []int64) (*StockMoves, error) {
	sms := &StockMoves{}
	if err := c.Read(StockMoveModel, ids, nil, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindStockMove finds stock.move record by querying it with criteria.
func (c *Client) FindStockMove(criteria *Criteria) (*StockMove, error) {
	sms := &StockMoves{}
	if err := c.SearchRead(StockMoveModel, criteria, NewOptions().Limit(1), sms); err != nil {
		return nil, err
	}
	return &((*sms)[0]), nil
}

// FindStockMoves finds stock.move records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoves(criteria *Criteria, options *Options) (*StockMoves, error) {
	sms := &StockMoves{}
	if err := c.SearchRead(StockMoveModel, criteria, options, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindStockMoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockMoveModel, criteria, options)
}

// FindStockMoveId finds record id by querying it with criteria.
func (c *Client) FindStockMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
