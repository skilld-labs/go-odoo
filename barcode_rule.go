package odoo

// BarcodeRule represents barcode.rule model.
type BarcodeRule struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	Alias                 *String    `xmlrpc:"alias,omptempty"`
	BarcodeNomenclatureId *Many2One  `xmlrpc:"barcode_nomenclature_id,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	Encoding              *Selection `xmlrpc:"encoding,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	Pattern               *String    `xmlrpc:"pattern,omptempty"`
	Sequence              *Int       `xmlrpc:"sequence,omptempty"`
	Type                  *Selection `xmlrpc:"type,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// BarcodeRules represents array of barcode.rule model.
type BarcodeRules []BarcodeRule

// BarcodeRuleModel is the odoo model name.
const BarcodeRuleModel = "barcode.rule"

// Many2One convert BarcodeRule to *Many2One.
func (br *BarcodeRule) Many2One() *Many2One {
	return NewMany2One(br.Id.Get(), "")
}

// CreateBarcodeRule creates a new barcode.rule model and returns its id.
func (c *Client) CreateBarcodeRule(br *BarcodeRule) (int64, error) {
	ids, err := c.CreateBarcodeRules([]*BarcodeRule{br})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBarcodeRule creates a new barcode.rule model and returns its id.
func (c *Client) CreateBarcodeRules(brs []*BarcodeRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range brs {
		vv = append(vv, v)
	}
	return c.Create(BarcodeRuleModel, vv, nil)
}

// UpdateBarcodeRule updates an existing barcode.rule record.
func (c *Client) UpdateBarcodeRule(br *BarcodeRule) error {
	return c.UpdateBarcodeRules([]int64{br.Id.Get()}, br)
}

// UpdateBarcodeRules updates existing barcode.rule records.
// All records (represented by ids) will be updated by br values.
func (c *Client) UpdateBarcodeRules(ids []int64, br *BarcodeRule) error {
	return c.Update(BarcodeRuleModel, ids, br, nil)
}

// DeleteBarcodeRule deletes an existing barcode.rule record.
func (c *Client) DeleteBarcodeRule(id int64) error {
	return c.DeleteBarcodeRules([]int64{id})
}

// DeleteBarcodeRules deletes existing barcode.rule records.
func (c *Client) DeleteBarcodeRules(ids []int64) error {
	return c.Delete(BarcodeRuleModel, ids)
}

// GetBarcodeRule gets barcode.rule existing record.
func (c *Client) GetBarcodeRule(id int64) (*BarcodeRule, error) {
	brs, err := c.GetBarcodeRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*brs)[0]), nil
}

// GetBarcodeRules gets barcode.rule existing records.
func (c *Client) GetBarcodeRules(ids []int64) (*BarcodeRules, error) {
	brs := &BarcodeRules{}
	if err := c.Read(BarcodeRuleModel, ids, nil, brs); err != nil {
		return nil, err
	}
	return brs, nil
}

// FindBarcodeRule finds barcode.rule record by querying it with criteria.
func (c *Client) FindBarcodeRule(criteria *Criteria) (*BarcodeRule, error) {
	brs := &BarcodeRules{}
	if err := c.SearchRead(BarcodeRuleModel, criteria, NewOptions().Limit(1), brs); err != nil {
		return nil, err
	}
	return &((*brs)[0]), nil
}

// FindBarcodeRules finds barcode.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBarcodeRules(criteria *Criteria, options *Options) (*BarcodeRules, error) {
	brs := &BarcodeRules{}
	if err := c.SearchRead(BarcodeRuleModel, criteria, options, brs); err != nil {
		return nil, err
	}
	return brs, nil
}

// FindBarcodeRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBarcodeRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BarcodeRuleModel, criteria, options)
}

// FindBarcodeRuleId finds record id by querying it with criteria.
func (c *Client) FindBarcodeRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BarcodeRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
