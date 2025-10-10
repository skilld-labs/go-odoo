package odoo

// HrResumeLine represents hr.resume.line model.
type HrResumeLine struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateEnd     *Time      `xmlrpc:"date_end,omitempty"`
	DateStart   *Time      `xmlrpc:"date_start,omitempty"`
	Description *String    `xmlrpc:"description,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	DisplayType *Selection `xmlrpc:"display_type,omitempty"`
	EmployeeId  *Many2One  `xmlrpc:"employee_id,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	LineTypeId  *Many2One  `xmlrpc:"line_type_id,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrResumeLines represents array of hr.resume.line model.
type HrResumeLines []HrResumeLine

// HrResumeLineModel is the odoo model name.
const HrResumeLineModel = "hr.resume.line"

// Many2One convert HrResumeLine to *Many2One.
func (hrl *HrResumeLine) Many2One() *Many2One {
	return NewMany2One(hrl.Id.Get(), "")
}

// CreateHrResumeLine creates a new hr.resume.line model and returns its id.
func (c *Client) CreateHrResumeLine(hrl *HrResumeLine) (int64, error) {
	ids, err := c.CreateHrResumeLines([]*HrResumeLine{hrl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrResumeLine creates a new hr.resume.line model and returns its id.
func (c *Client) CreateHrResumeLines(hrls []*HrResumeLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrls {
		vv = append(vv, v)
	}
	return c.Create(HrResumeLineModel, vv, nil)
}

// UpdateHrResumeLine updates an existing hr.resume.line record.
func (c *Client) UpdateHrResumeLine(hrl *HrResumeLine) error {
	return c.UpdateHrResumeLines([]int64{hrl.Id.Get()}, hrl)
}

// UpdateHrResumeLines updates existing hr.resume.line records.
// All records (represented by ids) will be updated by hrl values.
func (c *Client) UpdateHrResumeLines(ids []int64, hrl *HrResumeLine) error {
	return c.Update(HrResumeLineModel, ids, hrl, nil)
}

// DeleteHrResumeLine deletes an existing hr.resume.line record.
func (c *Client) DeleteHrResumeLine(id int64) error {
	return c.DeleteHrResumeLines([]int64{id})
}

// DeleteHrResumeLines deletes existing hr.resume.line records.
func (c *Client) DeleteHrResumeLines(ids []int64) error {
	return c.Delete(HrResumeLineModel, ids)
}

// GetHrResumeLine gets hr.resume.line existing record.
func (c *Client) GetHrResumeLine(id int64) (*HrResumeLine, error) {
	hrls, err := c.GetHrResumeLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hrls)[0]), nil
}

// GetHrResumeLines gets hr.resume.line existing records.
func (c *Client) GetHrResumeLines(ids []int64) (*HrResumeLines, error) {
	hrls := &HrResumeLines{}
	if err := c.Read(HrResumeLineModel, ids, nil, hrls); err != nil {
		return nil, err
	}
	return hrls, nil
}

// FindHrResumeLine finds hr.resume.line record by querying it with criteria.
func (c *Client) FindHrResumeLine(criteria *Criteria) (*HrResumeLine, error) {
	hrls := &HrResumeLines{}
	if err := c.SearchRead(HrResumeLineModel, criteria, NewOptions().Limit(1), hrls); err != nil {
		return nil, err
	}
	return &((*hrls)[0]), nil
}

// FindHrResumeLines finds hr.resume.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrResumeLines(criteria *Criteria, options *Options) (*HrResumeLines, error) {
	hrls := &HrResumeLines{}
	if err := c.SearchRead(HrResumeLineModel, criteria, options, hrls); err != nil {
		return nil, err
	}
	return hrls, nil
}

// FindHrResumeLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrResumeLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrResumeLineModel, criteria, options)
}

// FindHrResumeLineId finds record id by querying it with criteria.
func (c *Client) FindHrResumeLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrResumeLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
