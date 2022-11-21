package odoo

import (
	"fmt"
)

// SaleReport represents sale.report model.
type SaleReport struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	AmtInvoiced         *Float     `xmlrpc:"amt_invoiced,omptempty"`
	AmtToInvoice        *Float     `xmlrpc:"amt_to_invoice,omptempty"`
	AnalyticAccountId   *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	CategId             *Many2One  `xmlrpc:"categ_id,omptempty"`
	CommercialPartnerId *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	ConfirmationDate    *Time      `xmlrpc:"confirmation_date,omptempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omptempty"`
	Date                *Time      `xmlrpc:"date,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	Name                *String    `xmlrpc:"name,omptempty"`
	Nbr                 *Int       `xmlrpc:"nbr,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	PriceSubtotal       *Float     `xmlrpc:"price_subtotal,omptempty"`
	PriceTotal          *Float     `xmlrpc:"price_total,omptempty"`
	PricelistId         *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductTmplId       *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductUom          *Many2One  `xmlrpc:"product_uom,omptempty"`
	ProductUomQty       *Float     `xmlrpc:"product_uom_qty,omptempty"`
	QtyDelivered        *Float     `xmlrpc:"qty_delivered,omptempty"`
	QtyInvoiced         *Float     `xmlrpc:"qty_invoiced,omptempty"`
	QtyToInvoice        *Float     `xmlrpc:"qty_to_invoice,omptempty"`
	State               *Selection `xmlrpc:"state,omptempty"`
	TeamId              *Many2One  `xmlrpc:"team_id,omptempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omptempty"`
	Volume              *Float     `xmlrpc:"volume,omptempty"`
	WarehouseId         *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	Weight              *Float     `xmlrpc:"weight,omptempty"`
}

// SaleReports represents array of sale.report model.
type SaleReports []SaleReport

// SaleReportModel is the odoo model name.
const SaleReportModel = "sale.report"

// Many2One convert SaleReport to *Many2One.
func (sr *SaleReport) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateSaleReport creates a new sale.report model and returns its id.
func (c *Client) CreateSaleReport(sr *SaleReport) (int64, error) {
	return c.Create(SaleReportModel, sr)
}

// UpdateSaleReport updates an existing sale.report record.
func (c *Client) UpdateSaleReport(sr *SaleReport) error {
	return c.UpdateSaleReports([]int64{sr.Id.Get()}, sr)
}

// UpdateSaleReports updates existing sale.report records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateSaleReports(ids []int64, sr *SaleReport) error {
	return c.Update(SaleReportModel, ids, sr)
}

// DeleteSaleReport deletes an existing sale.report record.
func (c *Client) DeleteSaleReport(id int64) error {
	return c.DeleteSaleReports([]int64{id})
}

// DeleteSaleReports deletes existing sale.report records.
func (c *Client) DeleteSaleReports(ids []int64) error {
	return c.Delete(SaleReportModel, ids)
}

// GetSaleReport gets sale.report existing record.
func (c *Client) GetSaleReport(id int64) (*SaleReport, error) {
	srs, err := c.GetSaleReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.report not found", id)
}

// GetSaleReports gets sale.report existing records.
func (c *Client) GetSaleReports(ids []int64) (*SaleReports, error) {
	srs := &SaleReports{}
	if err := c.Read(SaleReportModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSaleReport finds sale.report record by querying it with criteria.
func (c *Client) FindSaleReport(criteria *Criteria) (*SaleReport, error) {
	srs := &SaleReports{}
	if err := c.SearchRead(SaleReportModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("no sale.report was found with criteria %v", criteria)
}

// FindSaleReports finds sale.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleReports(criteria *Criteria, options *Options) (*SaleReports, error) {
	srs := &SaleReports{}
	if err := c.SearchRead(SaleReportModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSaleReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleReportId finds record id by querying it with criteria.
func (c *Client) FindSaleReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no sale.report was found with criteria %v and options %v", criteria, options)
}
