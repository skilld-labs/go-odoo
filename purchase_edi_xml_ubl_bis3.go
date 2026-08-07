package odoo

// PurchaseEdiXmlUblBis3 represents purchase.edi.xml.ubl_bis3 model.
type PurchaseEdiXmlUblBis3 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// PurchaseEdiXmlUblBis3s represents array of purchase.edi.xml.ubl_bis3 model.
type PurchaseEdiXmlUblBis3s []PurchaseEdiXmlUblBis3

// PurchaseEdiXmlUblBis3Model is the odoo model name.
const PurchaseEdiXmlUblBis3Model = "purchase.edi.xml.ubl_bis3"

// Many2One convert PurchaseEdiXmlUblBis3 to *Many2One.
func (pexu *PurchaseEdiXmlUblBis3) Many2One() *Many2One {
	return NewMany2One(pexu.Id.Get(), "")
}

// CreatePurchaseEdiXmlUblBis3 creates a new purchase.edi.xml.ubl_bis3 model and returns its id.
func (c *Client) CreatePurchaseEdiXmlUblBis3(pexu *PurchaseEdiXmlUblBis3) (int64, error) {
	ids, err := c.CreatePurchaseEdiXmlUblBis3s([]*PurchaseEdiXmlUblBis3{pexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePurchaseEdiXmlUblBis3 creates a new purchase.edi.xml.ubl_bis3 model and returns its id.
func (c *Client) CreatePurchaseEdiXmlUblBis3s(pexus []*PurchaseEdiXmlUblBis3) ([]int64, error) {
	var vv []interface{}
	for _, v := range pexus {
		vv = append(vv, v)
	}
	return c.Create(PurchaseEdiXmlUblBis3Model, vv, nil)
}

// UpdatePurchaseEdiXmlUblBis3 updates an existing purchase.edi.xml.ubl_bis3 record.
func (c *Client) UpdatePurchaseEdiXmlUblBis3(pexu *PurchaseEdiXmlUblBis3) error {
	return c.UpdatePurchaseEdiXmlUblBis3s([]int64{pexu.Id.Get()}, pexu)
}

// UpdatePurchaseEdiXmlUblBis3s updates existing purchase.edi.xml.ubl_bis3 records.
// All records (represented by ids) will be updated by pexu values.
func (c *Client) UpdatePurchaseEdiXmlUblBis3s(ids []int64, pexu *PurchaseEdiXmlUblBis3) error {
	return c.Update(PurchaseEdiXmlUblBis3Model, ids, pexu, nil)
}

// DeletePurchaseEdiXmlUblBis3 deletes an existing purchase.edi.xml.ubl_bis3 record.
func (c *Client) DeletePurchaseEdiXmlUblBis3(id int64) error {
	return c.DeletePurchaseEdiXmlUblBis3s([]int64{id})
}

// DeletePurchaseEdiXmlUblBis3s deletes existing purchase.edi.xml.ubl_bis3 records.
func (c *Client) DeletePurchaseEdiXmlUblBis3s(ids []int64) error {
	return c.Delete(PurchaseEdiXmlUblBis3Model, ids)
}

// GetPurchaseEdiXmlUblBis3 gets purchase.edi.xml.ubl_bis3 existing record.
func (c *Client) GetPurchaseEdiXmlUblBis3(id int64) (*PurchaseEdiXmlUblBis3, error) {
	pexus, err := c.GetPurchaseEdiXmlUblBis3s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pexus)[0]), nil
}

// GetPurchaseEdiXmlUblBis3s gets purchase.edi.xml.ubl_bis3 existing records.
func (c *Client) GetPurchaseEdiXmlUblBis3s(ids []int64) (*PurchaseEdiXmlUblBis3s, error) {
	pexus := &PurchaseEdiXmlUblBis3s{}
	if err := c.Read(PurchaseEdiXmlUblBis3Model, ids, nil, pexus); err != nil {
		return nil, err
	}
	return pexus, nil
}

// FindPurchaseEdiXmlUblBis3 finds purchase.edi.xml.ubl_bis3 record by querying it with criteria.
func (c *Client) FindPurchaseEdiXmlUblBis3(criteria *Criteria) (*PurchaseEdiXmlUblBis3, error) {
	pexus := &PurchaseEdiXmlUblBis3s{}
	if err := c.SearchRead(PurchaseEdiXmlUblBis3Model, criteria, NewOptions().Limit(1), pexus); err != nil {
		return nil, err
	}
	return &((*pexus)[0]), nil
}

// FindPurchaseEdiXmlUblBis3s finds purchase.edi.xml.ubl_bis3 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseEdiXmlUblBis3s(criteria *Criteria, options *Options) (*PurchaseEdiXmlUblBis3s, error) {
	pexus := &PurchaseEdiXmlUblBis3s{}
	if err := c.SearchRead(PurchaseEdiXmlUblBis3Model, criteria, options, pexus); err != nil {
		return nil, err
	}
	return pexus, nil
}

// FindPurchaseEdiXmlUblBis3Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseEdiXmlUblBis3Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PurchaseEdiXmlUblBis3Model, criteria, options)
}

// FindPurchaseEdiXmlUblBis3Id finds record id by querying it with criteria.
func (c *Client) FindPurchaseEdiXmlUblBis3Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseEdiXmlUblBis3Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
