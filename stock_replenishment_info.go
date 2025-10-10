package odoo

// StockReplenishmentInfo represents stock.replenishment.info model.
type StockReplenishmentInfo struct {
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	JsonLeadDays             *String   `xmlrpc:"json_lead_days,omitempty"`
	JsonReplenishmentHistory *String   `xmlrpc:"json_replenishment_history,omitempty"`
	OrderpointId             *Many2One `xmlrpc:"orderpoint_id,omitempty"`
	ProductId                *Many2One `xmlrpc:"product_id,omitempty"`
	QtyToOrder               *Float    `xmlrpc:"qty_to_order,omitempty"`
	SupplierinfoId           *Many2One `xmlrpc:"supplierinfo_id,omitempty"`
	SupplierinfoIds          *Relation `xmlrpc:"supplierinfo_ids,omitempty"`
	WarehouseinfoIds         *Relation `xmlrpc:"warehouseinfo_ids,omitempty"`
	WhReplenishmentOptionIds *Relation `xmlrpc:"wh_replenishment_option_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockReplenishmentInfos represents array of stock.replenishment.info model.
type StockReplenishmentInfos []StockReplenishmentInfo

// StockReplenishmentInfoModel is the odoo model name.
const StockReplenishmentInfoModel = "stock.replenishment.info"

// Many2One convert StockReplenishmentInfo to *Many2One.
func (sri *StockReplenishmentInfo) Many2One() *Many2One {
	return NewMany2One(sri.Id.Get(), "")
}

// CreateStockReplenishmentInfo creates a new stock.replenishment.info model and returns its id.
func (c *Client) CreateStockReplenishmentInfo(sri *StockReplenishmentInfo) (int64, error) {
	ids, err := c.CreateStockReplenishmentInfos([]*StockReplenishmentInfo{sri})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockReplenishmentInfo creates a new stock.replenishment.info model and returns its id.
func (c *Client) CreateStockReplenishmentInfos(sris []*StockReplenishmentInfo) ([]int64, error) {
	var vv []interface{}
	for _, v := range sris {
		vv = append(vv, v)
	}
	return c.Create(StockReplenishmentInfoModel, vv, nil)
}

// UpdateStockReplenishmentInfo updates an existing stock.replenishment.info record.
func (c *Client) UpdateStockReplenishmentInfo(sri *StockReplenishmentInfo) error {
	return c.UpdateStockReplenishmentInfos([]int64{sri.Id.Get()}, sri)
}

// UpdateStockReplenishmentInfos updates existing stock.replenishment.info records.
// All records (represented by ids) will be updated by sri values.
func (c *Client) UpdateStockReplenishmentInfos(ids []int64, sri *StockReplenishmentInfo) error {
	return c.Update(StockReplenishmentInfoModel, ids, sri, nil)
}

// DeleteStockReplenishmentInfo deletes an existing stock.replenishment.info record.
func (c *Client) DeleteStockReplenishmentInfo(id int64) error {
	return c.DeleteStockReplenishmentInfos([]int64{id})
}

// DeleteStockReplenishmentInfos deletes existing stock.replenishment.info records.
func (c *Client) DeleteStockReplenishmentInfos(ids []int64) error {
	return c.Delete(StockReplenishmentInfoModel, ids)
}

// GetStockReplenishmentInfo gets stock.replenishment.info existing record.
func (c *Client) GetStockReplenishmentInfo(id int64) (*StockReplenishmentInfo, error) {
	sris, err := c.GetStockReplenishmentInfos([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sris)[0]), nil
}

// GetStockReplenishmentInfos gets stock.replenishment.info existing records.
func (c *Client) GetStockReplenishmentInfos(ids []int64) (*StockReplenishmentInfos, error) {
	sris := &StockReplenishmentInfos{}
	if err := c.Read(StockReplenishmentInfoModel, ids, nil, sris); err != nil {
		return nil, err
	}
	return sris, nil
}

// FindStockReplenishmentInfo finds stock.replenishment.info record by querying it with criteria.
func (c *Client) FindStockReplenishmentInfo(criteria *Criteria) (*StockReplenishmentInfo, error) {
	sris := &StockReplenishmentInfos{}
	if err := c.SearchRead(StockReplenishmentInfoModel, criteria, NewOptions().Limit(1), sris); err != nil {
		return nil, err
	}
	return &((*sris)[0]), nil
}

// FindStockReplenishmentInfos finds stock.replenishment.info records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishmentInfos(criteria *Criteria, options *Options) (*StockReplenishmentInfos, error) {
	sris := &StockReplenishmentInfos{}
	if err := c.SearchRead(StockReplenishmentInfoModel, criteria, options, sris); err != nil {
		return nil, err
	}
	return sris, nil
}

// FindStockReplenishmentInfoIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishmentInfoIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockReplenishmentInfoModel, criteria, options)
}

// FindStockReplenishmentInfoId finds record id by querying it with criteria.
func (c *Client) FindStockReplenishmentInfoId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReplenishmentInfoModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
