package odoo

// SpreadsheetDashboardShare represents spreadsheet.dashboard.share model.
type SpreadsheetDashboardShare struct {
	AccessToken           *String   `xmlrpc:"access_token,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DashboardId           *Many2One `xmlrpc:"dashboard_id,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	ExcelExport           *String   `xmlrpc:"excel_export,omitempty"`
	FullUrl               *String   `xmlrpc:"full_url,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	Name                  *String   `xmlrpc:"name,omitempty"`
	SpreadsheetBinaryData *String   `xmlrpc:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData       *String   `xmlrpc:"spreadsheet_data,omitempty"`
	SpreadsheetFileName   *String   `xmlrpc:"spreadsheet_file_name,omitempty"`
	Thumbnail             *String   `xmlrpc:"thumbnail,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SpreadsheetDashboardShares represents array of spreadsheet.dashboard.share model.
type SpreadsheetDashboardShares []SpreadsheetDashboardShare

// SpreadsheetDashboardShareModel is the odoo model name.
const SpreadsheetDashboardShareModel = "spreadsheet.dashboard.share"

// Many2One convert SpreadsheetDashboardShare to *Many2One.
func (sds *SpreadsheetDashboardShare) Many2One() *Many2One {
	return NewMany2One(sds.Id.Get(), "")
}

// CreateSpreadsheetDashboardShare creates a new spreadsheet.dashboard.share model and returns its id.
func (c *Client) CreateSpreadsheetDashboardShare(sds *SpreadsheetDashboardShare) (int64, error) {
	ids, err := c.CreateSpreadsheetDashboardShares([]*SpreadsheetDashboardShare{sds})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetDashboardShare creates a new spreadsheet.dashboard.share model and returns its id.
func (c *Client) CreateSpreadsheetDashboardShares(sdss []*SpreadsheetDashboardShare) ([]int64, error) {
	var vv []interface{}
	for _, v := range sdss {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetDashboardShareModel, vv, nil)
}

// UpdateSpreadsheetDashboardShare updates an existing spreadsheet.dashboard.share record.
func (c *Client) UpdateSpreadsheetDashboardShare(sds *SpreadsheetDashboardShare) error {
	return c.UpdateSpreadsheetDashboardShares([]int64{sds.Id.Get()}, sds)
}

// UpdateSpreadsheetDashboardShares updates existing spreadsheet.dashboard.share records.
// All records (represented by ids) will be updated by sds values.
func (c *Client) UpdateSpreadsheetDashboardShares(ids []int64, sds *SpreadsheetDashboardShare) error {
	return c.Update(SpreadsheetDashboardShareModel, ids, sds, nil)
}

// DeleteSpreadsheetDashboardShare deletes an existing spreadsheet.dashboard.share record.
func (c *Client) DeleteSpreadsheetDashboardShare(id int64) error {
	return c.DeleteSpreadsheetDashboardShares([]int64{id})
}

// DeleteSpreadsheetDashboardShares deletes existing spreadsheet.dashboard.share records.
func (c *Client) DeleteSpreadsheetDashboardShares(ids []int64) error {
	return c.Delete(SpreadsheetDashboardShareModel, ids)
}

// GetSpreadsheetDashboardShare gets spreadsheet.dashboard.share existing record.
func (c *Client) GetSpreadsheetDashboardShare(id int64) (*SpreadsheetDashboardShare, error) {
	sdss, err := c.GetSpreadsheetDashboardShares([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sdss)[0]), nil
}

// GetSpreadsheetDashboardShares gets spreadsheet.dashboard.share existing records.
func (c *Client) GetSpreadsheetDashboardShares(ids []int64) (*SpreadsheetDashboardShares, error) {
	sdss := &SpreadsheetDashboardShares{}
	if err := c.Read(SpreadsheetDashboardShareModel, ids, nil, sdss); err != nil {
		return nil, err
	}
	return sdss, nil
}

// FindSpreadsheetDashboardShare finds spreadsheet.dashboard.share record by querying it with criteria.
func (c *Client) FindSpreadsheetDashboardShare(criteria *Criteria) (*SpreadsheetDashboardShare, error) {
	sdss := &SpreadsheetDashboardShares{}
	if err := c.SearchRead(SpreadsheetDashboardShareModel, criteria, NewOptions().Limit(1), sdss); err != nil {
		return nil, err
	}
	return &((*sdss)[0]), nil
}

// FindSpreadsheetDashboardShares finds spreadsheet.dashboard.share records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboardShares(criteria *Criteria, options *Options) (*SpreadsheetDashboardShares, error) {
	sdss := &SpreadsheetDashboardShares{}
	if err := c.SearchRead(SpreadsheetDashboardShareModel, criteria, options, sdss); err != nil {
		return nil, err
	}
	return sdss, nil
}

// FindSpreadsheetDashboardShareIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboardShareIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SpreadsheetDashboardShareModel, criteria, options)
}

// FindSpreadsheetDashboardShareId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetDashboardShareId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetDashboardShareModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
