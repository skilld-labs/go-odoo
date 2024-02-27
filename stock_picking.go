package odoo

// StockPicking represents stock.picking model.
type StockPicking struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	ActivityDateDeadline     *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds              *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState            *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary          *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId           *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId           *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	BackorderId              *Many2One  `xmlrpc:"backorder_id,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date                     *Time      `xmlrpc:"date,omitempty"`
	DateDone                 *Time      `xmlrpc:"date_done,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	EntirePackageDetailIds   *Relation  `xmlrpc:"entire_package_detail_ids,omitempty"`
	EntirePackageIds         *Relation  `xmlrpc:"entire_package_ids,omitempty"`
	GroupId                  *Many2One  `xmlrpc:"group_id,omitempty"`
	HasPackages              *Bool      `xmlrpc:"has_packages,omitempty"`
	HasScrapMove             *Bool      `xmlrpc:"has_scrap_move,omitempty"`
	HasTracking              *Bool      `xmlrpc:"has_tracking,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IsLocked                 *Bool      `xmlrpc:"is_locked,omitempty"`
	LocationDestId           *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationId               *Many2One  `xmlrpc:"location_id,omitempty"`
	MessageChannelIds        *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost          *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread            *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter     *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MoveLineExist            *Bool      `xmlrpc:"move_line_exist,omitempty"`
	MoveLineIds              *Relation  `xmlrpc:"move_line_ids,omitempty"`
	MoveLines                *Relation  `xmlrpc:"move_lines,omitempty"`
	MoveType                 *Selection `xmlrpc:"move_type,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	Note                     *String    `xmlrpc:"note,omitempty"`
	Origin                   *String    `xmlrpc:"origin,omitempty"`
	OwnerId                  *Many2One  `xmlrpc:"owner_id,omitempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickingTypeCode          *Selection `xmlrpc:"picking_type_code,omitempty"`
	PickingTypeEntirePacks   *Bool      `xmlrpc:"picking_type_entire_packs,omitempty"`
	PickingTypeId            *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	Printed                  *Bool      `xmlrpc:"printed,omitempty"`
	Priority                 *Selection `xmlrpc:"priority,omitempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omitempty"`
	PurchaseId               *Many2One  `xmlrpc:"purchase_id,omitempty"`
	SaleId                   *Many2One  `xmlrpc:"sale_id,omitempty"`
	ScheduledDate            *Time      `xmlrpc:"scheduled_date,omitempty"`
	ShowCheckAvailability    *Bool      `xmlrpc:"show_check_availability,omitempty"`
	ShowLotsText             *Bool      `xmlrpc:"show_lots_text,omitempty"`
	ShowMarkAsTodo           *Bool      `xmlrpc:"show_mark_as_todo,omitempty"`
	ShowOperations           *Bool      `xmlrpc:"show_operations,omitempty"`
	ShowValidate             *Bool      `xmlrpc:"show_validate,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockPickings represents array of stock.picking model.
type StockPickings []StockPicking

// StockPickingModel is the odoo model name.
const StockPickingModel = "stock.picking"

// Many2One convert StockPicking to *Many2One.
func (sp *StockPicking) Many2One() *Many2One {
	return NewMany2One(sp.Id.Get(), "")
}

// CreateStockPicking creates a new stock.picking model and returns its id.
func (c *Client) CreateStockPicking(sp *StockPicking) (int64, error) {
	ids, err := c.CreateStockPickings([]*StockPicking{sp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPickings creates a new stock.picking model and returns its id.
func (c *Client) CreateStockPickings(sps []*StockPicking) ([]int64, error) {
	var vv []interface{}
	for _, v := range sps {
		vv = append(vv, v)
	}
	return c.Create(StockPickingModel, vv, nil)
}

// UpdateStockPicking updates an existing stock.picking record.
func (c *Client) UpdateStockPicking(sp *StockPicking) error {
	return c.UpdateStockPickings([]int64{sp.Id.Get()}, sp)
}

// UpdateStockPickings updates existing stock.picking records.
// All records (represented by ids) will be updated by sp values.
func (c *Client) UpdateStockPickings(ids []int64, sp *StockPicking) error {
	return c.Update(StockPickingModel, ids, sp, nil)
}

// DeleteStockPicking deletes an existing stock.picking record.
func (c *Client) DeleteStockPicking(id int64) error {
	return c.DeleteStockPickings([]int64{id})
}

// DeleteStockPickings deletes existing stock.picking records.
func (c *Client) DeleteStockPickings(ids []int64) error {
	return c.Delete(StockPickingModel, ids)
}

// GetStockPicking gets stock.picking existing record.
func (c *Client) GetStockPicking(id int64) (*StockPicking, error) {
	sps, err := c.GetStockPickings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sps)[0]), nil
}

// GetStockPickings gets stock.picking existing records.
func (c *Client) GetStockPickings(ids []int64) (*StockPickings, error) {
	sps := &StockPickings{}
	if err := c.Read(StockPickingModel, ids, nil, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPicking finds stock.picking record by querying it with criteria.
func (c *Client) FindStockPicking(criteria *Criteria) (*StockPicking, error) {
	sps := &StockPickings{}
	if err := c.SearchRead(StockPickingModel, criteria, NewOptions().Limit(1), sps); err != nil {
		return nil, err
	}
	return &((*sps)[0]), nil
}

// FindStockPickings finds stock.picking records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickings(criteria *Criteria, options *Options) (*StockPickings, error) {
	sps := &StockPickings{}
	if err := c.SearchRead(StockPickingModel, criteria, options, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPickingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPickingModel, criteria, options)
}

// FindStockPickingId finds record id by querying it with criteria.
func (c *Client) FindStockPickingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPickingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
