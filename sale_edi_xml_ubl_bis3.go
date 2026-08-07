package odoo

// SaleEdiXmlUblBis3 represents sale.edi.xml.ubl_bis3 model.
type SaleEdiXmlUblBis3 struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// SaleEdiXmlUblBis3s represents array of sale.edi.xml.ubl_bis3 model.
type SaleEdiXmlUblBis3s []SaleEdiXmlUblBis3

// SaleEdiXmlUblBis3Model is the odoo model name.
const SaleEdiXmlUblBis3Model = "sale.edi.xml.ubl_bis3"

// Many2One convert SaleEdiXmlUblBis3 to *Many2One.
func (sexu *SaleEdiXmlUblBis3) Many2One() *Many2One {
	return NewMany2One(sexu.Id.Get(), "")
}

// CreateSaleEdiXmlUblBis3 creates a new sale.edi.xml.ubl_bis3 model and returns its id.
func (c *Client) CreateSaleEdiXmlUblBis3(sexu *SaleEdiXmlUblBis3) (int64, error) {
	ids, err := c.CreateSaleEdiXmlUblBis3s([]*SaleEdiXmlUblBis3{sexu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleEdiXmlUblBis3 creates a new sale.edi.xml.ubl_bis3 model and returns its id.
func (c *Client) CreateSaleEdiXmlUblBis3s(sexus []*SaleEdiXmlUblBis3) ([]int64, error) {
	var vv []interface{}
	for _, v := range sexus {
		vv = append(vv, v)
	}
	return c.Create(SaleEdiXmlUblBis3Model, vv, nil)
}

// UpdateSaleEdiXmlUblBis3 updates an existing sale.edi.xml.ubl_bis3 record.
func (c *Client) UpdateSaleEdiXmlUblBis3(sexu *SaleEdiXmlUblBis3) error {
	return c.UpdateSaleEdiXmlUblBis3s([]int64{sexu.Id.Get()}, sexu)
}

// UpdateSaleEdiXmlUblBis3s updates existing sale.edi.xml.ubl_bis3 records.
// All records (represented by ids) will be updated by sexu values.
func (c *Client) UpdateSaleEdiXmlUblBis3s(ids []int64, sexu *SaleEdiXmlUblBis3) error {
	return c.Update(SaleEdiXmlUblBis3Model, ids, sexu, nil)
}

// DeleteSaleEdiXmlUblBis3 deletes an existing sale.edi.xml.ubl_bis3 record.
func (c *Client) DeleteSaleEdiXmlUblBis3(id int64) error {
	return c.DeleteSaleEdiXmlUblBis3s([]int64{id})
}

// DeleteSaleEdiXmlUblBis3s deletes existing sale.edi.xml.ubl_bis3 records.
func (c *Client) DeleteSaleEdiXmlUblBis3s(ids []int64) error {
	return c.Delete(SaleEdiXmlUblBis3Model, ids)
}

// GetSaleEdiXmlUblBis3 gets sale.edi.xml.ubl_bis3 existing record.
func (c *Client) GetSaleEdiXmlUblBis3(id int64) (*SaleEdiXmlUblBis3, error) {
	sexus, err := c.GetSaleEdiXmlUblBis3s([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sexus)[0]), nil
}

// GetSaleEdiXmlUblBis3s gets sale.edi.xml.ubl_bis3 existing records.
func (c *Client) GetSaleEdiXmlUblBis3s(ids []int64) (*SaleEdiXmlUblBis3s, error) {
	sexus := &SaleEdiXmlUblBis3s{}
	if err := c.Read(SaleEdiXmlUblBis3Model, ids, nil, sexus); err != nil {
		return nil, err
	}
	return sexus, nil
}

// FindSaleEdiXmlUblBis3 finds sale.edi.xml.ubl_bis3 record by querying it with criteria.
func (c *Client) FindSaleEdiXmlUblBis3(criteria *Criteria) (*SaleEdiXmlUblBis3, error) {
	sexus := &SaleEdiXmlUblBis3s{}
	if err := c.SearchRead(SaleEdiXmlUblBis3Model, criteria, NewOptions().Limit(1), sexus); err != nil {
		return nil, err
	}
	return &((*sexus)[0]), nil
}

// FindSaleEdiXmlUblBis3s finds sale.edi.xml.ubl_bis3 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleEdiXmlUblBis3s(criteria *Criteria, options *Options) (*SaleEdiXmlUblBis3s, error) {
	sexus := &SaleEdiXmlUblBis3s{}
	if err := c.SearchRead(SaleEdiXmlUblBis3Model, criteria, options, sexus); err != nil {
		return nil, err
	}
	return sexus, nil
}

// FindSaleEdiXmlUblBis3Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleEdiXmlUblBis3Ids(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleEdiXmlUblBis3Model, criteria, options)
}

// FindSaleEdiXmlUblBis3Id finds record id by querying it with criteria.
func (c *Client) FindSaleEdiXmlUblBis3Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleEdiXmlUblBis3Model, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
