package odoo

// HrResumeLineType represents hr.resume.line.type model.
type HrResumeLineType struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrResumeLineTypes represents array of hr.resume.line.type model.
type HrResumeLineTypes []HrResumeLineType

// HrResumeLineTypeModel is the odoo model name.
const HrResumeLineTypeModel = "hr.resume.line.type"

// Many2One convert HrResumeLineType to *Many2One.
func (hrlt *HrResumeLineType) Many2One() *Many2One {
	return NewMany2One(hrlt.Id.Get(), "")
}

// CreateHrResumeLineType creates a new hr.resume.line.type model and returns its id.
func (c *Client) CreateHrResumeLineType(hrlt *HrResumeLineType) (int64, error) {
	ids, err := c.CreateHrResumeLineTypes([]*HrResumeLineType{hrlt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrResumeLineType creates a new hr.resume.line.type model and returns its id.
func (c *Client) CreateHrResumeLineTypes(hrlts []*HrResumeLineType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrlts {
		vv = append(vv, v)
	}
	return c.Create(HrResumeLineTypeModel, vv, nil)
}

// UpdateHrResumeLineType updates an existing hr.resume.line.type record.
func (c *Client) UpdateHrResumeLineType(hrlt *HrResumeLineType) error {
	return c.UpdateHrResumeLineTypes([]int64{hrlt.Id.Get()}, hrlt)
}

// UpdateHrResumeLineTypes updates existing hr.resume.line.type records.
// All records (represented by ids) will be updated by hrlt values.
func (c *Client) UpdateHrResumeLineTypes(ids []int64, hrlt *HrResumeLineType) error {
	return c.Update(HrResumeLineTypeModel, ids, hrlt, nil)
}

// DeleteHrResumeLineType deletes an existing hr.resume.line.type record.
func (c *Client) DeleteHrResumeLineType(id int64) error {
	return c.DeleteHrResumeLineTypes([]int64{id})
}

// DeleteHrResumeLineTypes deletes existing hr.resume.line.type records.
func (c *Client) DeleteHrResumeLineTypes(ids []int64) error {
	return c.Delete(HrResumeLineTypeModel, ids)
}

// GetHrResumeLineType gets hr.resume.line.type existing record.
func (c *Client) GetHrResumeLineType(id int64) (*HrResumeLineType, error) {
	hrlts, err := c.GetHrResumeLineTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hrlts)[0]), nil
}

// GetHrResumeLineTypes gets hr.resume.line.type existing records.
func (c *Client) GetHrResumeLineTypes(ids []int64) (*HrResumeLineTypes, error) {
	hrlts := &HrResumeLineTypes{}
	if err := c.Read(HrResumeLineTypeModel, ids, nil, hrlts); err != nil {
		return nil, err
	}
	return hrlts, nil
}

// FindHrResumeLineType finds hr.resume.line.type record by querying it with criteria.
func (c *Client) FindHrResumeLineType(criteria *Criteria) (*HrResumeLineType, error) {
	hrlts := &HrResumeLineTypes{}
	if err := c.SearchRead(HrResumeLineTypeModel, criteria, NewOptions().Limit(1), hrlts); err != nil {
		return nil, err
	}
	return &((*hrlts)[0]), nil
}

// FindHrResumeLineTypes finds hr.resume.line.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrResumeLineTypes(criteria *Criteria, options *Options) (*HrResumeLineTypes, error) {
	hrlts := &HrResumeLineTypes{}
	if err := c.SearchRead(HrResumeLineTypeModel, criteria, options, hrlts); err != nil {
		return nil, err
	}
	return hrlts, nil
}

// FindHrResumeLineTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrResumeLineTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrResumeLineTypeModel, criteria, options)
}

// FindHrResumeLineTypeId finds record id by querying it with criteria.
func (c *Client) FindHrResumeLineTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrResumeLineTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
