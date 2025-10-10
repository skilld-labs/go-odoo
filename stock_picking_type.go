package odoo

// StockPickingType represents stock.picking.type model.
type StockPickingType struct {
	Active                         *Bool       `xmlrpc:"active,omitempty"`
	AnalyticCosts                  *Bool       `xmlrpc:"analytic_costs,omitempty"`
	AutoPrintDeliverySlip          *Bool       `xmlrpc:"auto_print_delivery_slip,omitempty"`
	AutoPrintLotLabels             *Bool       `xmlrpc:"auto_print_lot_labels,omitempty"`
	AutoPrintPackageLabel          *Bool       `xmlrpc:"auto_print_package_label,omitempty"`
	AutoPrintPackages              *Bool       `xmlrpc:"auto_print_packages,omitempty"`
	AutoPrintProductLabels         *Bool       `xmlrpc:"auto_print_product_labels,omitempty"`
	AutoPrintReceptionReport       *Bool       `xmlrpc:"auto_print_reception_report,omitempty"`
	AutoPrintReceptionReportLabels *Bool       `xmlrpc:"auto_print_reception_report_labels,omitempty"`
	AutoPrintReturnSlip            *Bool       `xmlrpc:"auto_print_return_slip,omitempty"`
	AutoShowReceptionReport        *Bool       `xmlrpc:"auto_show_reception_report,omitempty"`
	Barcode                        *String     `xmlrpc:"barcode,omitempty"`
	Code                           *Selection  `xmlrpc:"code,omitempty"`
	Color                          *Int        `xmlrpc:"color,omitempty"`
	CompanyId                      *Many2One   `xmlrpc:"company_id,omitempty"`
	CountMoveReady                 *Int        `xmlrpc:"count_move_ready,omitempty"`
	CountPicking                   *Int        `xmlrpc:"count_picking,omitempty"`
	CountPickingBackorders         *Int        `xmlrpc:"count_picking_backorders,omitempty"`
	CountPickingDraft              *Int        `xmlrpc:"count_picking_draft,omitempty"`
	CountPickingLate               *Int        `xmlrpc:"count_picking_late,omitempty"`
	CountPickingReady              *Int        `xmlrpc:"count_picking_ready,omitempty"`
	CountPickingWaiting            *Int        `xmlrpc:"count_picking_waiting,omitempty"`
	CreateBackorder                *Selection  `xmlrpc:"create_backorder,omitempty"`
	CreateDate                     *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One   `xmlrpc:"create_uid,omitempty"`
	DefaultLocationDestId          *Many2One   `xmlrpc:"default_location_dest_id,omitempty"`
	DefaultLocationSrcId           *Many2One   `xmlrpc:"default_location_src_id,omitempty"`
	DisplayName                    *String     `xmlrpc:"display_name,omitempty"`
	FavoriteUserIds                *Relation   `xmlrpc:"favorite_user_ids,omitempty"`
	HideReservationMethod          *Bool       `xmlrpc:"hide_reservation_method,omitempty"`
	Id                             *Int        `xmlrpc:"id,omitempty"`
	IsFavorite                     *Bool       `xmlrpc:"is_favorite,omitempty"`
	KanbanDashboardGraph           *String     `xmlrpc:"kanban_dashboard_graph,omitempty"`
	LotLabelFormat                 *Selection  `xmlrpc:"lot_label_format,omitempty"`
	MoveType                       *Selection  `xmlrpc:"move_type,omitempty"`
	Name                           *String     `xmlrpc:"name,omitempty"`
	PackageLabelToPrint            *Selection  `xmlrpc:"package_label_to_print,omitempty"`
	PickingPropertiesDefinition    interface{} `xmlrpc:"picking_properties_definition,omitempty"`
	PrintLabel                     *Bool       `xmlrpc:"print_label,omitempty"`
	ProductLabelFormat             *Selection  `xmlrpc:"product_label_format,omitempty"`
	ReservationDaysBefore          *Int        `xmlrpc:"reservation_days_before,omitempty"`
	ReservationDaysBeforePriority  *Int        `xmlrpc:"reservation_days_before_priority,omitempty"`
	ReservationMethod              *Selection  `xmlrpc:"reservation_method,omitempty"`
	ReturnPickingTypeId            *Many2One   `xmlrpc:"return_picking_type_id,omitempty"`
	Sequence                       *Int        `xmlrpc:"sequence,omitempty"`
	SequenceCode                   *String     `xmlrpc:"sequence_code,omitempty"`
	SequenceId                     *Many2One   `xmlrpc:"sequence_id,omitempty"`
	ShowEntirePacks                *Bool       `xmlrpc:"show_entire_packs,omitempty"`
	ShowOperations                 *Bool       `xmlrpc:"show_operations,omitempty"`
	ShowPickingType                *Bool       `xmlrpc:"show_picking_type,omitempty"`
	UseCreateLots                  *Bool       `xmlrpc:"use_create_lots,omitempty"`
	UseExistingLots                *Bool       `xmlrpc:"use_existing_lots,omitempty"`
	WarehouseId                    *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WriteDate                      *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// StockPickingTypes represents array of stock.picking.type model.
type StockPickingTypes []StockPickingType

// StockPickingTypeModel is the odoo model name.
const StockPickingTypeModel = "stock.picking.type"

// Many2One convert StockPickingType to *Many2One.
func (spt *StockPickingType) Many2One() *Many2One {
	return NewMany2One(spt.Id.Get(), "")
}

// CreateStockPickingType creates a new stock.picking.type model and returns its id.
func (c *Client) CreateStockPickingType(spt *StockPickingType) (int64, error) {
	ids, err := c.CreateStockPickingTypes([]*StockPickingType{spt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPickingType creates a new stock.picking.type model and returns its id.
func (c *Client) CreateStockPickingTypes(spts []*StockPickingType) ([]int64, error) {
	var vv []interface{}
	for _, v := range spts {
		vv = append(vv, v)
	}
	return c.Create(StockPickingTypeModel, vv, nil)
}

// UpdateStockPickingType updates an existing stock.picking.type record.
func (c *Client) UpdateStockPickingType(spt *StockPickingType) error {
	return c.UpdateStockPickingTypes([]int64{spt.Id.Get()}, spt)
}

// UpdateStockPickingTypes updates existing stock.picking.type records.
// All records (represented by ids) will be updated by spt values.
func (c *Client) UpdateStockPickingTypes(ids []int64, spt *StockPickingType) error {
	return c.Update(StockPickingTypeModel, ids, spt, nil)
}

// DeleteStockPickingType deletes an existing stock.picking.type record.
func (c *Client) DeleteStockPickingType(id int64) error {
	return c.DeleteStockPickingTypes([]int64{id})
}

// DeleteStockPickingTypes deletes existing stock.picking.type records.
func (c *Client) DeleteStockPickingTypes(ids []int64) error {
	return c.Delete(StockPickingTypeModel, ids)
}

// GetStockPickingType gets stock.picking.type existing record.
func (c *Client) GetStockPickingType(id int64) (*StockPickingType, error) {
	spts, err := c.GetStockPickingTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*spts)[0]), nil
}

// GetStockPickingTypes gets stock.picking.type existing records.
func (c *Client) GetStockPickingTypes(ids []int64) (*StockPickingTypes, error) {
	spts := &StockPickingTypes{}
	if err := c.Read(StockPickingTypeModel, ids, nil, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindStockPickingType finds stock.picking.type record by querying it with criteria.
func (c *Client) FindStockPickingType(criteria *Criteria) (*StockPickingType, error) {
	spts := &StockPickingTypes{}
	if err := c.SearchRead(StockPickingTypeModel, criteria, NewOptions().Limit(1), spts); err != nil {
		return nil, err
	}
	return &((*spts)[0]), nil
}

// FindStockPickingTypes finds stock.picking.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingTypes(criteria *Criteria, options *Options) (*StockPickingTypes, error) {
	spts := &StockPickingTypes{}
	if err := c.SearchRead(StockPickingTypeModel, criteria, options, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindStockPickingTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPickingTypeModel, criteria, options)
}

// FindStockPickingTypeId finds record id by querying it with criteria.
func (c *Client) FindStockPickingTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPickingTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
