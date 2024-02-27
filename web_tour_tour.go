package odoo

// WebTourTour represents web_tour.tour model.
type WebTourTour struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
}

// WebTourTours represents array of web_tour.tour model.
type WebTourTours []WebTourTour

// WebTourTourModel is the odoo model name.
const WebTourTourModel = "web_tour.tour"

// Many2One convert WebTourTour to *Many2One.
func (wt *WebTourTour) Many2One() *Many2One {
	return NewMany2One(wt.Id.Get(), "")
}

// CreateWebTourTour creates a new web_tour.tour model and returns its id.
func (c *Client) CreateWebTourTour(wt *WebTourTour) (int64, error) {
	ids, err := c.CreateWebTourTours([]*WebTourTour{wt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebTourTours creates a new web_tour.tour model and returns its id.
func (c *Client) CreateWebTourTours(wts []*WebTourTour) ([]int64, error) {
	var vv []interface{}
	for _, v := range wts {
		vv = append(vv, v)
	}
	return c.Create(WebTourTourModel, vv, nil)
}

// UpdateWebTourTour updates an existing web_tour.tour record.
func (c *Client) UpdateWebTourTour(wt *WebTourTour) error {
	return c.UpdateWebTourTours([]int64{wt.Id.Get()}, wt)
}

// UpdateWebTourTours updates existing web_tour.tour records.
// All records (represented by ids) will be updated by wt values.
func (c *Client) UpdateWebTourTours(ids []int64, wt *WebTourTour) error {
	return c.Update(WebTourTourModel, ids, wt, nil)
}

// DeleteWebTourTour deletes an existing web_tour.tour record.
func (c *Client) DeleteWebTourTour(id int64) error {
	return c.DeleteWebTourTours([]int64{id})
}

// DeleteWebTourTours deletes existing web_tour.tour records.
func (c *Client) DeleteWebTourTours(ids []int64) error {
	return c.Delete(WebTourTourModel, ids)
}

// GetWebTourTour gets web_tour.tour existing record.
func (c *Client) GetWebTourTour(id int64) (*WebTourTour, error) {
	wts, err := c.GetWebTourTours([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wts)[0]), nil
}

// GetWebTourTours gets web_tour.tour existing records.
func (c *Client) GetWebTourTours(ids []int64) (*WebTourTours, error) {
	wts := &WebTourTours{}
	if err := c.Read(WebTourTourModel, ids, nil, wts); err != nil {
		return nil, err
	}
	return wts, nil
}

// FindWebTourTour finds web_tour.tour record by querying it with criteria.
func (c *Client) FindWebTourTour(criteria *Criteria) (*WebTourTour, error) {
	wts := &WebTourTours{}
	if err := c.SearchRead(WebTourTourModel, criteria, NewOptions().Limit(1), wts); err != nil {
		return nil, err
	}
	return &((*wts)[0]), nil
}

// FindWebTourTours finds web_tour.tour records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebTourTours(criteria *Criteria, options *Options) (*WebTourTours, error) {
	wts := &WebTourTours{}
	if err := c.SearchRead(WebTourTourModel, criteria, options, wts); err != nil {
		return nil, err
	}
	return wts, nil
}

// FindWebTourTourIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebTourTourIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebTourTourModel, criteria, options)
}

// FindWebTourTourId finds record id by querying it with criteria.
func (c *Client) FindWebTourTourId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebTourTourModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
