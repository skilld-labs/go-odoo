package odoo

// PdpFlow10XmlBuilder represents pdp.flow.10.xml.builder model.
type PdpFlow10XmlBuilder struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// PdpFlow10XmlBuilders represents array of pdp.flow.10.xml.builder model.
type PdpFlow10XmlBuilders []PdpFlow10XmlBuilder

// PdpFlow10XmlBuilderModel is the odoo model name.
const PdpFlow10XmlBuilderModel = "pdp.flow.10.xml.builder"

// Many2One convert PdpFlow10XmlBuilder to *Many2One.
func (pf1xb *PdpFlow10XmlBuilder) Many2One() *Many2One {
	return NewMany2One(pf1xb.Id.Get(), "")
}

// CreatePdpFlow10XmlBuilder creates a new pdp.flow.10.xml.builder model and returns its id.
func (c *Client) CreatePdpFlow10XmlBuilder(pf1xb *PdpFlow10XmlBuilder) (int64, error) {
	ids, err := c.CreatePdpFlow10XmlBuilders([]*PdpFlow10XmlBuilder{pf1xb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePdpFlow10XmlBuilder creates a new pdp.flow.10.xml.builder model and returns its id.
func (c *Client) CreatePdpFlow10XmlBuilders(pf1xbs []*PdpFlow10XmlBuilder) ([]int64, error) {
	var vv []interface{}
	for _, v := range pf1xbs {
		vv = append(vv, v)
	}
	return c.Create(PdpFlow10XmlBuilderModel, vv, nil)
}

// UpdatePdpFlow10XmlBuilder updates an existing pdp.flow.10.xml.builder record.
func (c *Client) UpdatePdpFlow10XmlBuilder(pf1xb *PdpFlow10XmlBuilder) error {
	return c.UpdatePdpFlow10XmlBuilders([]int64{pf1xb.Id.Get()}, pf1xb)
}

// UpdatePdpFlow10XmlBuilders updates existing pdp.flow.10.xml.builder records.
// All records (represented by ids) will be updated by pf1xb values.
func (c *Client) UpdatePdpFlow10XmlBuilders(ids []int64, pf1xb *PdpFlow10XmlBuilder) error {
	return c.Update(PdpFlow10XmlBuilderModel, ids, pf1xb, nil)
}

// DeletePdpFlow10XmlBuilder deletes an existing pdp.flow.10.xml.builder record.
func (c *Client) DeletePdpFlow10XmlBuilder(id int64) error {
	return c.DeletePdpFlow10XmlBuilders([]int64{id})
}

// DeletePdpFlow10XmlBuilders deletes existing pdp.flow.10.xml.builder records.
func (c *Client) DeletePdpFlow10XmlBuilders(ids []int64) error {
	return c.Delete(PdpFlow10XmlBuilderModel, ids)
}

// GetPdpFlow10XmlBuilder gets pdp.flow.10.xml.builder existing record.
func (c *Client) GetPdpFlow10XmlBuilder(id int64) (*PdpFlow10XmlBuilder, error) {
	pf1xbs, err := c.GetPdpFlow10XmlBuilders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pf1xbs)[0]), nil
}

// GetPdpFlow10XmlBuilders gets pdp.flow.10.xml.builder existing records.
func (c *Client) GetPdpFlow10XmlBuilders(ids []int64) (*PdpFlow10XmlBuilders, error) {
	pf1xbs := &PdpFlow10XmlBuilders{}
	if err := c.Read(PdpFlow10XmlBuilderModel, ids, nil, pf1xbs); err != nil {
		return nil, err
	}
	return pf1xbs, nil
}

// FindPdpFlow10XmlBuilder finds pdp.flow.10.xml.builder record by querying it with criteria.
func (c *Client) FindPdpFlow10XmlBuilder(criteria *Criteria) (*PdpFlow10XmlBuilder, error) {
	pf1xbs := &PdpFlow10XmlBuilders{}
	if err := c.SearchRead(PdpFlow10XmlBuilderModel, criteria, NewOptions().Limit(1), pf1xbs); err != nil {
		return nil, err
	}
	return &((*pf1xbs)[0]), nil
}

// FindPdpFlow10XmlBuilders finds pdp.flow.10.xml.builder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPdpFlow10XmlBuilders(criteria *Criteria, options *Options) (*PdpFlow10XmlBuilders, error) {
	pf1xbs := &PdpFlow10XmlBuilders{}
	if err := c.SearchRead(PdpFlow10XmlBuilderModel, criteria, options, pf1xbs); err != nil {
		return nil, err
	}
	return pf1xbs, nil
}

// FindPdpFlow10XmlBuilderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPdpFlow10XmlBuilderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PdpFlow10XmlBuilderModel, criteria, options)
}

// FindPdpFlow10XmlBuilderId finds record id by querying it with criteria.
func (c *Client) FindPdpFlow10XmlBuilderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PdpFlow10XmlBuilderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
