package odoo

// StockMove represents stock.move model.
type StockMove struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	AccountMoveIds           *Relation  `xmlrpc:"account_move_ids,omitempty"`
	Additional               *Bool      `xmlrpc:"additional,omitempty"`
	Availability             *Float     `xmlrpc:"availability,omitempty"`
	BackorderId              *Many2One  `xmlrpc:"backorder_id,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	CreatedPurchaseLineId    *Many2One  `xmlrpc:"created_purchase_line_id,omitempty"`
	Date                     *Time      `xmlrpc:"date,omitempty"`
	DateExpected             *Time      `xmlrpc:"date_expected,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	GroupId                  *Many2One  `xmlrpc:"group_id,omitempty"`
	HasMoveLines             *Bool      `xmlrpc:"has_move_lines,omitempty"`
	HasTracking              *Selection `xmlrpc:"has_tracking,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	InventoryId              *Many2One  `xmlrpc:"inventory_id,omitempty"`
	IsInitialDemandEditable  *Bool      `xmlrpc:"is_initial_demand_editable,omitempty"`
	IsLocked                 *Bool      `xmlrpc:"is_locked,omitempty"`
	IsQuantityDoneEditable   *Bool      `xmlrpc:"is_quantity_done_editable,omitempty"`
	LocationDestId           *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationId               *Many2One  `xmlrpc:"location_id,omitempty"`
	MoveDestIds              *Relation  `xmlrpc:"move_dest_ids,omitempty"`
	MoveLineIds              *Relation  `xmlrpc:"move_line_ids,omitempty"`
	MoveLineNosuggestIds     *Relation  `xmlrpc:"move_line_nosuggest_ids,omitempty"`
	MoveOrigIds              *Relation  `xmlrpc:"move_orig_ids,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	Note                     *String    `xmlrpc:"note,omitempty"`
	OrderedQty               *Float     `xmlrpc:"ordered_qty,omitempty"`
	Origin                   *String    `xmlrpc:"origin,omitempty"`
	OriginReturnedMoveId     *Many2One  `xmlrpc:"origin_returned_move_id,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickingCode              *Selection `xmlrpc:"picking_code,omitempty"`
	PickingId                *Many2One  `xmlrpc:"picking_id,omitempty"`
	PickingPartnerId         *Many2One  `xmlrpc:"picking_partner_id,omitempty"`
	PickingTypeId            *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	PriceUnit                *Float     `xmlrpc:"price_unit,omitempty"`
	Priority                 *Selection `xmlrpc:"priority,omitempty"`
	ProcureMethod            *Selection `xmlrpc:"procure_method,omitempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductPackaging         *Many2One  `xmlrpc:"product_packaging,omitempty"`
	ProductQty               *Float     `xmlrpc:"product_qty,omitempty"`
	ProductTmplId            *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductType              *Selection `xmlrpc:"product_type,omitempty"`
	ProductUom               *Many2One  `xmlrpc:"product_uom,omitempty"`
	ProductUomQty            *Float     `xmlrpc:"product_uom_qty,omitempty"`
	Propagate                *Bool      `xmlrpc:"propagate,omitempty"`
	PurchaseLineId           *Many2One  `xmlrpc:"purchase_line_id,omitempty"`
	PushRuleId               *Many2One  `xmlrpc:"push_rule_id,omitempty"`
	QuantityDone             *Float     `xmlrpc:"quantity_done,omitempty"`
	Reference                *String    `xmlrpc:"reference,omitempty"`
	RemainingQty             *Float     `xmlrpc:"remaining_qty,omitempty"`
	RemainingValue           *Float     `xmlrpc:"remaining_value,omitempty"`
	ReservedAvailability     *Float     `xmlrpc:"reserved_availability,omitempty"`
	RestrictPartnerId        *Many2One  `xmlrpc:"restrict_partner_id,omitempty"`
	ReturnedMoveIds          *Relation  `xmlrpc:"returned_move_ids,omitempty"`
	RouteIds                 *Relation  `xmlrpc:"route_ids,omitempty"`
	RuleId                   *Many2One  `xmlrpc:"rule_id,omitempty"`
	SaleLineId               *Many2One  `xmlrpc:"sale_line_id,omitempty"`
	ScrapIds                 *Relation  `xmlrpc:"scrap_ids,omitempty"`
	Scrapped                 *Bool      `xmlrpc:"scrapped,omitempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omitempty"`
	ShowDetailsVisible       *Bool      `xmlrpc:"show_details_visible,omitempty"`
	ShowOperations           *Bool      `xmlrpc:"show_operations,omitempty"`
	ShowReservedAvailability *Bool      `xmlrpc:"show_reserved_availability,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	StringAvailabilityInfo   *String    `xmlrpc:"string_availability_info,omitempty"`
	ToRefund                 *Bool      `xmlrpc:"to_refund,omitempty"`
	Value                    *Float     `xmlrpc:"value,omitempty"`
	WarehouseId              *Many2One  `xmlrpc:"warehouse_id,omitempty"`
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

// CreateStockMoves creates a new stock.move model and returns its id.
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
