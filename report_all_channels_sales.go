package odoo

// ReportAllChannelsSales represents report.all.channels.sales model.
type ReportAllChannelsSales struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omitempty"`
	AnalyticAccountId *Many2One `xmlrpc:"analytic_account_id,omitempty"`
	CategId           *Many2One `xmlrpc:"categ_id,omitempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omitempty"`
	CountryId         *Many2One `xmlrpc:"country_id,omitempty"`
	DateOrder         *Time     `xmlrpc:"date_order,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	Name              *String   `xmlrpc:"name,omitempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omitempty"`
	PriceSubtotal     *Float    `xmlrpc:"price_subtotal,omitempty"`
	PriceTotal        *Float    `xmlrpc:"price_total,omitempty"`
	PricelistId       *Many2One `xmlrpc:"pricelist_id,omitempty"`
	ProductId         *Many2One `xmlrpc:"product_id,omitempty"`
	ProductQty        *Float    `xmlrpc:"product_qty,omitempty"`
	ProductTmplId     *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	TeamId            *Many2One `xmlrpc:"team_id,omitempty"`
	UserId            *Many2One `xmlrpc:"user_id,omitempty"`
}

// ReportAllChannelsSaless represents array of report.all.channels.sales model.
type ReportAllChannelsSaless []ReportAllChannelsSales

// ReportAllChannelsSalesModel is the odoo model name.
const ReportAllChannelsSalesModel = "report.all.channels.sales"

// Many2One convert ReportAllChannelsSales to *Many2One.
func (racs *ReportAllChannelsSales) Many2One() *Many2One {
	return NewMany2One(racs.Id.Get(), "")
}

// CreateReportAllChannelsSales creates a new report.all.channels.sales model and returns its id.
func (c *Client) CreateReportAllChannelsSales(racs *ReportAllChannelsSales) (int64, error) {
	ids, err := c.CreateReportAllChannelsSaless([]*ReportAllChannelsSales{racs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportAllChannelsSaless creates a new report.all.channels.sales model and returns its id.
func (c *Client) CreateReportAllChannelsSaless(racss []*ReportAllChannelsSales) ([]int64, error) {
	var vv []interface{}
	for _, v := range racss {
		vv = append(vv, v)
	}
	return c.Create(ReportAllChannelsSalesModel, vv, nil)
}

// UpdateReportAllChannelsSales updates an existing report.all.channels.sales record.
func (c *Client) UpdateReportAllChannelsSales(racs *ReportAllChannelsSales) error {
	return c.UpdateReportAllChannelsSaless([]int64{racs.Id.Get()}, racs)
}

// UpdateReportAllChannelsSaless updates existing report.all.channels.sales records.
// All records (represented by ids) will be updated by racs values.
func (c *Client) UpdateReportAllChannelsSaless(ids []int64, racs *ReportAllChannelsSales) error {
	return c.Update(ReportAllChannelsSalesModel, ids, racs, nil)
}

// DeleteReportAllChannelsSales deletes an existing report.all.channels.sales record.
func (c *Client) DeleteReportAllChannelsSales(id int64) error {
	return c.DeleteReportAllChannelsSaless([]int64{id})
}

// DeleteReportAllChannelsSaless deletes existing report.all.channels.sales records.
func (c *Client) DeleteReportAllChannelsSaless(ids []int64) error {
	return c.Delete(ReportAllChannelsSalesModel, ids)
}

// GetReportAllChannelsSales gets report.all.channels.sales existing record.
func (c *Client) GetReportAllChannelsSales(id int64) (*ReportAllChannelsSales, error) {
	racss, err := c.GetReportAllChannelsSaless([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*racss)[0]), nil
}

// GetReportAllChannelsSaless gets report.all.channels.sales existing records.
func (c *Client) GetReportAllChannelsSaless(ids []int64) (*ReportAllChannelsSaless, error) {
	racss := &ReportAllChannelsSaless{}
	if err := c.Read(ReportAllChannelsSalesModel, ids, nil, racss); err != nil {
		return nil, err
	}
	return racss, nil
}

// FindReportAllChannelsSales finds report.all.channels.sales record by querying it with criteria.
func (c *Client) FindReportAllChannelsSales(criteria *Criteria) (*ReportAllChannelsSales, error) {
	racss := &ReportAllChannelsSaless{}
	if err := c.SearchRead(ReportAllChannelsSalesModel, criteria, NewOptions().Limit(1), racss); err != nil {
		return nil, err
	}
	return &((*racss)[0]), nil
}

// FindReportAllChannelsSaless finds report.all.channels.sales records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAllChannelsSaless(criteria *Criteria, options *Options) (*ReportAllChannelsSaless, error) {
	racss := &ReportAllChannelsSaless{}
	if err := c.SearchRead(ReportAllChannelsSalesModel, criteria, options, racss); err != nil {
		return nil, err
	}
	return racss, nil
}

// FindReportAllChannelsSalesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportAllChannelsSalesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportAllChannelsSalesModel, criteria, options)
}

// FindReportAllChannelsSalesId finds record id by querying it with criteria.
func (c *Client) FindReportAllChannelsSalesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportAllChannelsSalesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
