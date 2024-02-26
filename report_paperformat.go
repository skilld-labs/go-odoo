package odoo

// ReportPaperformat represents report.paperformat model.
type ReportPaperformat struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	Default       *Bool      `xmlrpc:"default,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Dpi           *Int       `xmlrpc:"dpi,omptempty"`
	Format        *Selection `xmlrpc:"format,omptempty"`
	HeaderLine    *Bool      `xmlrpc:"header_line,omptempty"`
	HeaderSpacing *Int       `xmlrpc:"header_spacing,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	MarginBottom  *Float     `xmlrpc:"margin_bottom,omptempty"`
	MarginLeft    *Float     `xmlrpc:"margin_left,omptempty"`
	MarginRight   *Float     `xmlrpc:"margin_right,omptempty"`
	MarginTop     *Float     `xmlrpc:"margin_top,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	Orientation   *Selection `xmlrpc:"orientation,omptempty"`
	PageHeight    *Int       `xmlrpc:"page_height,omptempty"`
	PageWidth     *Int       `xmlrpc:"page_width,omptempty"`
	ReportIds     *Relation  `xmlrpc:"report_ids,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ReportPaperformats represents array of report.paperformat model.
type ReportPaperformats []ReportPaperformat

// ReportPaperformatModel is the odoo model name.
const ReportPaperformatModel = "report.paperformat"

// Many2One convert ReportPaperformat to *Many2One.
func (rp *ReportPaperformat) Many2One() *Many2One {
	return NewMany2One(rp.Id.Get(), "")
}

// CreateReportPaperformat creates a new report.paperformat model and returns its id.
func (c *Client) CreateReportPaperformat(rp *ReportPaperformat) (int64, error) {
	ids, err := c.CreateReportPaperformats([]*ReportPaperformat{rp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateReportPaperformat creates a new report.paperformat model and returns its id.
func (c *Client) CreateReportPaperformats(rps []*ReportPaperformat) ([]int64, error) {
	var vv []interface{}
	for _, v := range rps {
		vv = append(vv, v)
	}
	return c.Create(ReportPaperformatModel, vv, nil)
}

// UpdateReportPaperformat updates an existing report.paperformat record.
func (c *Client) UpdateReportPaperformat(rp *ReportPaperformat) error {
	return c.UpdateReportPaperformats([]int64{rp.Id.Get()}, rp)
}

// UpdateReportPaperformats updates existing report.paperformat records.
// All records (represented by ids) will be updated by rp values.
func (c *Client) UpdateReportPaperformats(ids []int64, rp *ReportPaperformat) error {
	return c.Update(ReportPaperformatModel, ids, rp, nil)
}

// DeleteReportPaperformat deletes an existing report.paperformat record.
func (c *Client) DeleteReportPaperformat(id int64) error {
	return c.DeleteReportPaperformats([]int64{id})
}

// DeleteReportPaperformats deletes existing report.paperformat records.
func (c *Client) DeleteReportPaperformats(ids []int64) error {
	return c.Delete(ReportPaperformatModel, ids)
}

// GetReportPaperformat gets report.paperformat existing record.
func (c *Client) GetReportPaperformat(id int64) (*ReportPaperformat, error) {
	rps, err := c.GetReportPaperformats([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rps)[0]), nil
}

// GetReportPaperformats gets report.paperformat existing records.
func (c *Client) GetReportPaperformats(ids []int64) (*ReportPaperformats, error) {
	rps := &ReportPaperformats{}
	if err := c.Read(ReportPaperformatModel, ids, nil, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindReportPaperformat finds report.paperformat record by querying it with criteria.
func (c *Client) FindReportPaperformat(criteria *Criteria) (*ReportPaperformat, error) {
	rps := &ReportPaperformats{}
	if err := c.SearchRead(ReportPaperformatModel, criteria, NewOptions().Limit(1), rps); err != nil {
		return nil, err
	}
	return &((*rps)[0]), nil
}

// FindReportPaperformats finds report.paperformat records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPaperformats(criteria *Criteria, options *Options) (*ReportPaperformats, error) {
	rps := &ReportPaperformats{}
	if err := c.SearchRead(ReportPaperformatModel, criteria, options, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindReportPaperformatIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPaperformatIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ReportPaperformatModel, criteria, options)
}

// FindReportPaperformatId finds record id by querying it with criteria.
func (c *Client) FindReportPaperformatId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportPaperformatModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
