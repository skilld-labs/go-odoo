package odoo

// StockPackage represents stock.package model.
type StockPackage struct {
	AllChildrenPackageIds *Relation `xmlrpc:"all_children_package_ids,omitempty"`
	ChildPackageDestIds   *Relation `xmlrpc:"child_package_dest_ids,omitempty"`
	ChildPackageIds       *Relation `xmlrpc:"child_package_ids,omitempty"`
	CompanyId             *Many2One `xmlrpc:"company_id,omitempty"`
	CompleteName          *String   `xmlrpc:"complete_name,omitempty"`
	ContainedQuantIds     *Relation `xmlrpc:"contained_quant_ids,omitempty"`
	ContentDescription    *String   `xmlrpc:"content_description,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DestCompleteName      *String   `xmlrpc:"dest_complete_name,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	JsonPopover           *String   `xmlrpc:"json_popover,omitempty"`
	LocationDestId        *Many2One `xmlrpc:"location_dest_id,omitempty"`
	LocationId            *Many2One `xmlrpc:"location_id,omitempty"`
	MoveLineIds           *Relation `xmlrpc:"move_line_ids,omitempty"`
	Name                  *String   `xmlrpc:"name,omitempty"`
	OutermostPackageId    *Many2One `xmlrpc:"outermost_package_id,omitempty"`
	OwnerId               *Many2One `xmlrpc:"owner_id,omitempty"`
	PackDate              *Time     `xmlrpc:"pack_date,omitempty"`
	PackageDestId         *Many2One `xmlrpc:"package_dest_id,omitempty"`
	PackageTypeId         *Many2One `xmlrpc:"package_type_id,omitempty"`
	ParentPackageId       *Many2One `xmlrpc:"parent_package_id,omitempty"`
	ParentPath            *String   `xmlrpc:"parent_path,omitempty"`
	PickingIds            *Relation `xmlrpc:"picking_ids,omitempty"`
	QuantIds              *Relation `xmlrpc:"quant_ids,omitempty"`
	ShippingWeight        *Float    `xmlrpc:"shipping_weight,omitempty"`
	ValidSscc             *Bool     `xmlrpc:"valid_sscc,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockPackages represents array of stock.package model.
type StockPackages []StockPackage

// StockPackageModel is the odoo model name.
const StockPackageModel = "stock.package"

// Many2One convert StockPackage to *Many2One.
func (sp *StockPackage) Many2One() *Many2One {
	return NewMany2One(sp.Id.Get(), "")
}

// CreateStockPackage creates a new stock.package model and returns its id.
func (c *Client) CreateStockPackage(sp *StockPackage) (int64, error) {
	ids, err := c.CreateStockPackages([]*StockPackage{sp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockPackage creates a new stock.package model and returns its id.
func (c *Client) CreateStockPackages(sps []*StockPackage) ([]int64, error) {
	var vv []interface{}
	for _, v := range sps {
		vv = append(vv, v)
	}
	return c.Create(StockPackageModel, vv, nil)
}

// UpdateStockPackage updates an existing stock.package record.
func (c *Client) UpdateStockPackage(sp *StockPackage) error {
	return c.UpdateStockPackages([]int64{sp.Id.Get()}, sp)
}

// UpdateStockPackages updates existing stock.package records.
// All records (represented by ids) will be updated by sp values.
func (c *Client) UpdateStockPackages(ids []int64, sp *StockPackage) error {
	return c.Update(StockPackageModel, ids, sp, nil)
}

// DeleteStockPackage deletes an existing stock.package record.
func (c *Client) DeleteStockPackage(id int64) error {
	return c.DeleteStockPackages([]int64{id})
}

// DeleteStockPackages deletes existing stock.package records.
func (c *Client) DeleteStockPackages(ids []int64) error {
	return c.Delete(StockPackageModel, ids)
}

// GetStockPackage gets stock.package existing record.
func (c *Client) GetStockPackage(id int64) (*StockPackage, error) {
	sps, err := c.GetStockPackages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sps)[0]), nil
}

// GetStockPackages gets stock.package existing records.
func (c *Client) GetStockPackages(ids []int64) (*StockPackages, error) {
	sps := &StockPackages{}
	if err := c.Read(StockPackageModel, ids, nil, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPackage finds stock.package record by querying it with criteria.
func (c *Client) FindStockPackage(criteria *Criteria) (*StockPackage, error) {
	sps := &StockPackages{}
	if err := c.SearchRead(StockPackageModel, criteria, NewOptions().Limit(1), sps); err != nil {
		return nil, err
	}
	return &((*sps)[0]), nil
}

// FindStockPackages finds stock.package records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackages(criteria *Criteria, options *Options) (*StockPackages, error) {
	sps := &StockPackages{}
	if err := c.SearchRead(StockPackageModel, criteria, options, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPackageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockPackageModel, criteria, options)
}

// FindStockPackageId finds record id by querying it with criteria.
func (c *Client) FindStockPackageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPackageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
