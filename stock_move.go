package odoo

// StockMove represents stock.move model.
type StockMove struct {
	AccountMoveId                         *Many2One   `xmlrpc:"account_move_id,omitempty"`
	Additional                            *Bool       `xmlrpc:"additional,omitempty"`
	AllowedUomIds                         *Relation   `xmlrpc:"allowed_uom_ids,omitempty"`
	AnalyticAccountLineIds                *Relation   `xmlrpc:"analytic_account_line_ids,omitempty"`
	Availability                          *Float      `xmlrpc:"availability,omitempty"`
	CompanyCurrencyId                     *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                            *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One   `xmlrpc:"create_uid,omitempty"`
	CreatedPurchaseLineIds                *Relation   `xmlrpc:"created_purchase_line_ids,omitempty"`
	Date                                  *Time       `xmlrpc:"date,omitempty"`
	DateDeadline                          *Time       `xmlrpc:"date_deadline,omitempty"`
	DelayAlertDate                        *Time       `xmlrpc:"delay_alert_date,omitempty"`
	DescriptionPicking                    *String     `xmlrpc:"description_picking,omitempty"`
	DescriptionPickingManual              *String     `xmlrpc:"description_picking_manual,omitempty"`
	DisplayAssignSerial                   *Bool       `xmlrpc:"display_assign_serial,omitempty"`
	DisplayImportLot                      *Bool       `xmlrpc:"display_import_lot,omitempty"`
	DisplayName                           *String     `xmlrpc:"display_name,omitempty"`
	ForecastAvailability                  *Float      `xmlrpc:"forecast_availability,omitempty"`
	ForecastExpectedDate                  *Time       `xmlrpc:"forecast_expected_date,omitempty"`
	HasLinesWithoutResultPackage          *Bool       `xmlrpc:"has_lines_without_result_package,omitempty"`
	HasTracking                           *Selection  `xmlrpc:"has_tracking,omitempty"`
	Id                                    *Int        `xmlrpc:"id,omitempty"`
	InventoryName                         *String     `xmlrpc:"inventory_name,omitempty"`
	IsDateEditable                        *Bool       `xmlrpc:"is_date_editable,omitempty"`
	IsDropship                            *Bool       `xmlrpc:"is_dropship,omitempty"`
	IsIn                                  *Bool       `xmlrpc:"is_in,omitempty"`
	IsInitialDemandEditable               *Bool       `xmlrpc:"is_initial_demand_editable,omitempty"`
	IsInventory                           *Bool       `xmlrpc:"is_inventory,omitempty"`
	IsLocked                              *Bool       `xmlrpc:"is_locked,omitempty"`
	IsOut                                 *Bool       `xmlrpc:"is_out,omitempty"`
	IsQuantityDoneEditable                *Bool       `xmlrpc:"is_quantity_done_editable,omitempty"`
	IsStorable                            *Bool       `xmlrpc:"is_storable,omitempty"`
	IsValued                              *Bool       `xmlrpc:"is_valued,omitempty"`
	LocationDestId                        *Many2One   `xmlrpc:"location_dest_id,omitempty"`
	LocationDestUsage                     *Selection  `xmlrpc:"location_dest_usage,omitempty"`
	LocationFinalId                       *Many2One   `xmlrpc:"location_final_id,omitempty"`
	LocationId                            *Many2One   `xmlrpc:"location_id,omitempty"`
	LocationUsage                         *Selection  `xmlrpc:"location_usage,omitempty"`
	LotIds                                *Relation   `xmlrpc:"lot_ids,omitempty"`
	MoveDestIds                           *Relation   `xmlrpc:"move_dest_ids,omitempty"`
	MoveLineIds                           *Relation   `xmlrpc:"move_line_ids,omitempty"`
	MoveLinesCount                        *Int        `xmlrpc:"move_lines_count,omitempty"`
	MoveOrigIds                           *Relation   `xmlrpc:"move_orig_ids,omitempty"`
	NeverProductTemplateAttributeValueIds *Relation   `xmlrpc:"never_product_template_attribute_value_ids,omitempty"`
	NextSerial                            *String     `xmlrpc:"next_serial,omitempty"`
	NextSerialCount                       *Int        `xmlrpc:"next_serial_count,omitempty"`
	OrderpointId                          *Many2One   `xmlrpc:"orderpoint_id,omitempty"`
	Origin                                *String     `xmlrpc:"origin,omitempty"`
	OriginReturnedMoveId                  *Many2One   `xmlrpc:"origin_returned_move_id,omitempty"`
	PackageIds                            *Relation   `xmlrpc:"package_ids,omitempty"`
	PackagingUomId                        *Many2One   `xmlrpc:"packaging_uom_id,omitempty"`
	PackagingUomQty                       *Float      `xmlrpc:"packaging_uom_qty,omitempty"`
	PartnerId                             *Many2One   `xmlrpc:"partner_id,omitempty"`
	Picked                                *Bool       `xmlrpc:"picked,omitempty"`
	PickingCode                           *Selection  `xmlrpc:"picking_code,omitempty"`
	PickingId                             *Many2One   `xmlrpc:"picking_id,omitempty"`
	PickingTypeId                         *Many2One   `xmlrpc:"picking_type_id,omitempty"`
	PriceUnit                             *Float      `xmlrpc:"price_unit,omitempty"`
	Priority                              *Selection  `xmlrpc:"priority,omitempty"`
	ProcureMethod                         *Selection  `xmlrpc:"procure_method,omitempty"`
	ProcurementValues                     interface{} `xmlrpc:"procurement_values,omitempty"`
	ProductCategoryId                     *Many2One   `xmlrpc:"product_category_id,omitempty"`
	ProductId                             *Many2One   `xmlrpc:"product_id,omitempty"`
	ProductQty                            *Float      `xmlrpc:"product_qty,omitempty"`
	ProductTmplId                         *Many2One   `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUom                            *Many2One   `xmlrpc:"product_uom,omitempty"`
	ProductUomQty                         *Float      `xmlrpc:"product_uom_qty,omitempty"`
	PropagateCancel                       *Bool       `xmlrpc:"propagate_cancel,omitempty"`
	PurchaseLineId                        *Many2One   `xmlrpc:"purchase_line_id,omitempty"`
	Quantity                              *Float      `xmlrpc:"quantity,omitempty"`
	Reference                             *String     `xmlrpc:"reference,omitempty"`
	ReferenceIds                          *Relation   `xmlrpc:"reference_ids,omitempty"`
	RemainingQty                          *Float      `xmlrpc:"remaining_qty,omitempty"`
	RemainingValue                        *Float      `xmlrpc:"remaining_value,omitempty"`
	ReservationDate                       *Time       `xmlrpc:"reservation_date,omitempty"`
	RestrictPartnerId                     *Many2One   `xmlrpc:"restrict_partner_id,omitempty"`
	ReturnedMoveIds                       *Relation   `xmlrpc:"returned_move_ids,omitempty"`
	RouteIds                              *Relation   `xmlrpc:"route_ids,omitempty"`
	RuleId                                *Many2One   `xmlrpc:"rule_id,omitempty"`
	SaleLineId                            *Many2One   `xmlrpc:"sale_line_id,omitempty"`
	ScrapId                               *Many2One   `xmlrpc:"scrap_id,omitempty"`
	Sequence                              *Int        `xmlrpc:"sequence,omitempty"`
	ShowDetailsVisible                    *Bool       `xmlrpc:"show_details_visible,omitempty"`
	ShowLotsM2O                           *Bool       `xmlrpc:"show_lots_m2o,omitempty"`
	ShowLotsText                          *Bool       `xmlrpc:"show_lots_text,omitempty"`
	ShowOperations                        *Bool       `xmlrpc:"show_operations,omitempty"`
	ShowQuant                             *Bool       `xmlrpc:"show_quant,omitempty"`
	StandardPrice                         *Float      `xmlrpc:"standard_price,omitempty"`
	State                                 *Selection  `xmlrpc:"state,omitempty"`
	ToRefund                              *Bool       `xmlrpc:"to_refund,omitempty"`
	Value                                 *Float      `xmlrpc:"value,omitempty"`
	ValueComputedJustification            *String     `xmlrpc:"value_computed_justification,omitempty"`
	ValueJustification                    *String     `xmlrpc:"value_justification,omitempty"`
	ValueManual                           *Float      `xmlrpc:"value_manual,omitempty"`
	WarehouseId                           *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WriteDate                             *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One   `xmlrpc:"write_uid,omitempty"`
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
