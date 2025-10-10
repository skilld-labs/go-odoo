package odoo

// VendorDelayReport represents vendor.delay.report model.
type VendorDelayReport struct {
	CategoryId  *Many2One `xmlrpc:"category_id,omitempty"`
	Date        *Time     `xmlrpc:"date,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	OnTimeRate  *Float    `xmlrpc:"on_time_rate,omitempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omitempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omitempty"`
	QtyOnTime   *Float    `xmlrpc:"qty_on_time,omitempty"`
	QtyTotal    *Float    `xmlrpc:"qty_total,omitempty"`
}

// VendorDelayReports represents array of vendor.delay.report model.
type VendorDelayReports []VendorDelayReport

// VendorDelayReportModel is the odoo model name.
const VendorDelayReportModel = "vendor.delay.report"

// Many2One convert VendorDelayReport to *Many2One.
func (vdr *VendorDelayReport) Many2One() *Many2One {
	return NewMany2One(vdr.Id.Get(), "")
}

// CreateVendorDelayReport creates a new vendor.delay.report model and returns its id.
func (c *Client) CreateVendorDelayReport(vdr *VendorDelayReport) (int64, error) {
	ids, err := c.CreateVendorDelayReports([]*VendorDelayReport{vdr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateVendorDelayReport creates a new vendor.delay.report model and returns its id.
func (c *Client) CreateVendorDelayReports(vdrs []*VendorDelayReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range vdrs {
		vv = append(vv, v)
	}
	return c.Create(VendorDelayReportModel, vv, nil)
}

// UpdateVendorDelayReport updates an existing vendor.delay.report record.
func (c *Client) UpdateVendorDelayReport(vdr *VendorDelayReport) error {
	return c.UpdateVendorDelayReports([]int64{vdr.Id.Get()}, vdr)
}

// UpdateVendorDelayReports updates existing vendor.delay.report records.
// All records (represented by ids) will be updated by vdr values.
func (c *Client) UpdateVendorDelayReports(ids []int64, vdr *VendorDelayReport) error {
	return c.Update(VendorDelayReportModel, ids, vdr, nil)
}

// DeleteVendorDelayReport deletes an existing vendor.delay.report record.
func (c *Client) DeleteVendorDelayReport(id int64) error {
	return c.DeleteVendorDelayReports([]int64{id})
}

// DeleteVendorDelayReports deletes existing vendor.delay.report records.
func (c *Client) DeleteVendorDelayReports(ids []int64) error {
	return c.Delete(VendorDelayReportModel, ids)
}

// GetVendorDelayReport gets vendor.delay.report existing record.
func (c *Client) GetVendorDelayReport(id int64) (*VendorDelayReport, error) {
	vdrs, err := c.GetVendorDelayReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*vdrs)[0]), nil
}

// GetVendorDelayReports gets vendor.delay.report existing records.
func (c *Client) GetVendorDelayReports(ids []int64) (*VendorDelayReports, error) {
	vdrs := &VendorDelayReports{}
	if err := c.Read(VendorDelayReportModel, ids, nil, vdrs); err != nil {
		return nil, err
	}
	return vdrs, nil
}

// FindVendorDelayReport finds vendor.delay.report record by querying it with criteria.
func (c *Client) FindVendorDelayReport(criteria *Criteria) (*VendorDelayReport, error) {
	vdrs := &VendorDelayReports{}
	if err := c.SearchRead(VendorDelayReportModel, criteria, NewOptions().Limit(1), vdrs); err != nil {
		return nil, err
	}
	return &((*vdrs)[0]), nil
}

// FindVendorDelayReports finds vendor.delay.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindVendorDelayReports(criteria *Criteria, options *Options) (*VendorDelayReports, error) {
	vdrs := &VendorDelayReports{}
	if err := c.SearchRead(VendorDelayReportModel, criteria, options, vdrs); err != nil {
		return nil, err
	}
	return vdrs, nil
}

// FindVendorDelayReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindVendorDelayReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(VendorDelayReportModel, criteria, options)
}

// FindVendorDelayReportId finds record id by querying it with criteria.
func (c *Client) FindVendorDelayReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(VendorDelayReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
