package odoo

// BillToPoWizard represents bill.to.po.wizard model.
type BillToPoWizard struct {
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	PartnerId       *Many2One `xmlrpc:"partner_id,omitempty"`
	PurchaseOrderId *Many2One `xmlrpc:"purchase_order_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BillToPoWizards represents array of bill.to.po.wizard model.
type BillToPoWizards []BillToPoWizard

// BillToPoWizardModel is the odoo model name.
const BillToPoWizardModel = "bill.to.po.wizard"

// Many2One convert BillToPoWizard to *Many2One.
func (btpw *BillToPoWizard) Many2One() *Many2One {
	return NewMany2One(btpw.Id.Get(), "")
}

// CreateBillToPoWizard creates a new bill.to.po.wizard model and returns its id.
func (c *Client) CreateBillToPoWizard(btpw *BillToPoWizard) (int64, error) {
	ids, err := c.CreateBillToPoWizards([]*BillToPoWizard{btpw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBillToPoWizard creates a new bill.to.po.wizard model and returns its id.
func (c *Client) CreateBillToPoWizards(btpws []*BillToPoWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range btpws {
		vv = append(vv, v)
	}
	return c.Create(BillToPoWizardModel, vv, nil)
}

// UpdateBillToPoWizard updates an existing bill.to.po.wizard record.
func (c *Client) UpdateBillToPoWizard(btpw *BillToPoWizard) error {
	return c.UpdateBillToPoWizards([]int64{btpw.Id.Get()}, btpw)
}

// UpdateBillToPoWizards updates existing bill.to.po.wizard records.
// All records (represented by ids) will be updated by btpw values.
func (c *Client) UpdateBillToPoWizards(ids []int64, btpw *BillToPoWizard) error {
	return c.Update(BillToPoWizardModel, ids, btpw, nil)
}

// DeleteBillToPoWizard deletes an existing bill.to.po.wizard record.
func (c *Client) DeleteBillToPoWizard(id int64) error {
	return c.DeleteBillToPoWizards([]int64{id})
}

// DeleteBillToPoWizards deletes existing bill.to.po.wizard records.
func (c *Client) DeleteBillToPoWizards(ids []int64) error {
	return c.Delete(BillToPoWizardModel, ids)
}

// GetBillToPoWizard gets bill.to.po.wizard existing record.
func (c *Client) GetBillToPoWizard(id int64) (*BillToPoWizard, error) {
	btpws, err := c.GetBillToPoWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*btpws)[0]), nil
}

// GetBillToPoWizards gets bill.to.po.wizard existing records.
func (c *Client) GetBillToPoWizards(ids []int64) (*BillToPoWizards, error) {
	btpws := &BillToPoWizards{}
	if err := c.Read(BillToPoWizardModel, ids, nil, btpws); err != nil {
		return nil, err
	}
	return btpws, nil
}

// FindBillToPoWizard finds bill.to.po.wizard record by querying it with criteria.
func (c *Client) FindBillToPoWizard(criteria *Criteria) (*BillToPoWizard, error) {
	btpws := &BillToPoWizards{}
	if err := c.SearchRead(BillToPoWizardModel, criteria, NewOptions().Limit(1), btpws); err != nil {
		return nil, err
	}
	return &((*btpws)[0]), nil
}

// FindBillToPoWizards finds bill.to.po.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBillToPoWizards(criteria *Criteria, options *Options) (*BillToPoWizards, error) {
	btpws := &BillToPoWizards{}
	if err := c.SearchRead(BillToPoWizardModel, criteria, options, btpws); err != nil {
		return nil, err
	}
	return btpws, nil
}

// FindBillToPoWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBillToPoWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BillToPoWizardModel, criteria, options)
}

// FindBillToPoWizardId finds record id by querying it with criteria.
func (c *Client) FindBillToPoWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BillToPoWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
