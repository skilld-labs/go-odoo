package odoo

// WebTourTourStep represents web_tour.tour.step model.
type WebTourTourStep struct {
	Content     *String   `xmlrpc:"content,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Run         *String   `xmlrpc:"run,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	TourId      *Many2One `xmlrpc:"tour_id,omitempty"`
	Trigger     *String   `xmlrpc:"trigger,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebTourTourSteps represents array of web_tour.tour.step model.
type WebTourTourSteps []WebTourTourStep

// WebTourTourStepModel is the odoo model name.
const WebTourTourStepModel = "web_tour.tour.step"

// Many2One convert WebTourTourStep to *Many2One.
func (wts *WebTourTourStep) Many2One() *Many2One {
	return NewMany2One(wts.Id.Get(), "")
}

// CreateWebTourTourStep creates a new web_tour.tour.step model and returns its id.
func (c *Client) CreateWebTourTourStep(wts *WebTourTourStep) (int64, error) {
	ids, err := c.CreateWebTourTourSteps([]*WebTourTourStep{wts})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebTourTourStep creates a new web_tour.tour.step model and returns its id.
func (c *Client) CreateWebTourTourSteps(wtss []*WebTourTourStep) ([]int64, error) {
	var vv []interface{}
	for _, v := range wtss {
		vv = append(vv, v)
	}
	return c.Create(WebTourTourStepModel, vv, nil)
}

// UpdateWebTourTourStep updates an existing web_tour.tour.step record.
func (c *Client) UpdateWebTourTourStep(wts *WebTourTourStep) error {
	return c.UpdateWebTourTourSteps([]int64{wts.Id.Get()}, wts)
}

// UpdateWebTourTourSteps updates existing web_tour.tour.step records.
// All records (represented by ids) will be updated by wts values.
func (c *Client) UpdateWebTourTourSteps(ids []int64, wts *WebTourTourStep) error {
	return c.Update(WebTourTourStepModel, ids, wts, nil)
}

// DeleteWebTourTourStep deletes an existing web_tour.tour.step record.
func (c *Client) DeleteWebTourTourStep(id int64) error {
	return c.DeleteWebTourTourSteps([]int64{id})
}

// DeleteWebTourTourSteps deletes existing web_tour.tour.step records.
func (c *Client) DeleteWebTourTourSteps(ids []int64) error {
	return c.Delete(WebTourTourStepModel, ids)
}

// GetWebTourTourStep gets web_tour.tour.step existing record.
func (c *Client) GetWebTourTourStep(id int64) (*WebTourTourStep, error) {
	wtss, err := c.GetWebTourTourSteps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wtss)[0]), nil
}

// GetWebTourTourSteps gets web_tour.tour.step existing records.
func (c *Client) GetWebTourTourSteps(ids []int64) (*WebTourTourSteps, error) {
	wtss := &WebTourTourSteps{}
	if err := c.Read(WebTourTourStepModel, ids, nil, wtss); err != nil {
		return nil, err
	}
	return wtss, nil
}

// FindWebTourTourStep finds web_tour.tour.step record by querying it with criteria.
func (c *Client) FindWebTourTourStep(criteria *Criteria) (*WebTourTourStep, error) {
	wtss := &WebTourTourSteps{}
	if err := c.SearchRead(WebTourTourStepModel, criteria, NewOptions().Limit(1), wtss); err != nil {
		return nil, err
	}
	return &((*wtss)[0]), nil
}

// FindWebTourTourSteps finds web_tour.tour.step records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebTourTourSteps(criteria *Criteria, options *Options) (*WebTourTourSteps, error) {
	wtss := &WebTourTourSteps{}
	if err := c.SearchRead(WebTourTourStepModel, criteria, options, wtss); err != nil {
		return nil, err
	}
	return wtss, nil
}

// FindWebTourTourStepIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebTourTourStepIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebTourTourStepModel, criteria, options)
}

// FindWebTourTourStepId finds record id by querying it with criteria.
func (c *Client) FindWebTourTourStepId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebTourTourStepModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
