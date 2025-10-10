package odoo

// UomUom represents uom.uom model.
type UomUom struct {
	Active             *Bool      `xmlrpc:"active,omitempty"`
	CategoryId         *Many2One  `xmlrpc:"category_id,omitempty"`
	Color              *Int       `xmlrpc:"color,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Factor             *Float     `xmlrpc:"factor,omitempty"`
	FactorInv          *Float     `xmlrpc:"factor_inv,omitempty"`
	FiscalCountryCodes *String    `xmlrpc:"fiscal_country_codes,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	Ratio              *Float     `xmlrpc:"ratio,omitempty"`
	Rounding           *Float     `xmlrpc:"rounding,omitempty"`
	TimesheetWidget    *String    `xmlrpc:"timesheet_widget,omitempty"`
	UomType            *Selection `xmlrpc:"uom_type,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// UomUoms represents array of uom.uom model.
type UomUoms []UomUom

// UomUomModel is the odoo model name.
const UomUomModel = "uom.uom"

// Many2One convert UomUom to *Many2One.
func (uu *UomUom) Many2One() *Many2One {
	return NewMany2One(uu.Id.Get(), "")
}

// CreateUomUom creates a new uom.uom model and returns its id.
func (c *Client) CreateUomUom(uu *UomUom) (int64, error) {
	ids, err := c.CreateUomUoms([]*UomUom{uu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateUomUom creates a new uom.uom model and returns its id.
func (c *Client) CreateUomUoms(uus []*UomUom) ([]int64, error) {
	var vv []interface{}
	for _, v := range uus {
		vv = append(vv, v)
	}
	return c.Create(UomUomModel, vv, nil)
}

// UpdateUomUom updates an existing uom.uom record.
func (c *Client) UpdateUomUom(uu *UomUom) error {
	return c.UpdateUomUoms([]int64{uu.Id.Get()}, uu)
}

// UpdateUomUoms updates existing uom.uom records.
// All records (represented by ids) will be updated by uu values.
func (c *Client) UpdateUomUoms(ids []int64, uu *UomUom) error {
	return c.Update(UomUomModel, ids, uu, nil)
}

// DeleteUomUom deletes an existing uom.uom record.
func (c *Client) DeleteUomUom(id int64) error {
	return c.DeleteUomUoms([]int64{id})
}

// DeleteUomUoms deletes existing uom.uom records.
func (c *Client) DeleteUomUoms(ids []int64) error {
	return c.Delete(UomUomModel, ids)
}

// GetUomUom gets uom.uom existing record.
func (c *Client) GetUomUom(id int64) (*UomUom, error) {
	uus, err := c.GetUomUoms([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*uus)[0]), nil
}

// GetUomUoms gets uom.uom existing records.
func (c *Client) GetUomUoms(ids []int64) (*UomUoms, error) {
	uus := &UomUoms{}
	if err := c.Read(UomUomModel, ids, nil, uus); err != nil {
		return nil, err
	}
	return uus, nil
}

// FindUomUom finds uom.uom record by querying it with criteria.
func (c *Client) FindUomUom(criteria *Criteria) (*UomUom, error) {
	uus := &UomUoms{}
	if err := c.SearchRead(UomUomModel, criteria, NewOptions().Limit(1), uus); err != nil {
		return nil, err
	}
	return &((*uus)[0]), nil
}

// FindUomUoms finds uom.uom records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUomUoms(criteria *Criteria, options *Options) (*UomUoms, error) {
	uus := &UomUoms{}
	if err := c.SearchRead(UomUomModel, criteria, options, uus); err != nil {
		return nil, err
	}
	return uus, nil
}

// FindUomUomIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUomUomIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(UomUomModel, criteria, options)
}

// FindUomUomId finds record id by querying it with criteria.
func (c *Client) FindUomUomId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UomUomModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
