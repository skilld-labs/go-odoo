package odoo

// IrModelData represents ir.model.data model.
type IrModelData struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CompleteName *String   `xmlrpc:"complete_name,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DateInit     *Time     `xmlrpc:"date_init,omptempty"`
	DateUpdate   *Time     `xmlrpc:"date_update,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Model        *String   `xmlrpc:"model,omptempty"`
	Module       *String   `xmlrpc:"module,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	Noupdate     *Bool     `xmlrpc:"noupdate,omptempty"`
	Reference    *String   `xmlrpc:"reference,omptempty"`
	ResId        *Int      `xmlrpc:"res_id,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// IrModelDatas represents array of ir.model.data model.
type IrModelDatas []IrModelData

// IrModelDataModel is the odoo model name.
const IrModelDataModel = "ir.model.data"

// Many2One convert IrModelData to *Many2One.
func (imd *IrModelData) Many2One() *Many2One {
	return NewMany2One(imd.Id.Get(), "")
}

// CreateIrModelData creates a new ir.model.data model and returns its id.
func (c *Client) CreateIrModelData(imd *IrModelData) (int64, error) {
	ids, err := c.CreateIrModelDatas([]*IrModelData{imd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModelData creates a new ir.model.data model and returns its id.
func (c *Client) CreateIrModelDatas(imds []*IrModelData) ([]int64, error) {
	var vv []interface{}
	for _, v := range imds {
		vv = append(vv, v)
	}
	return c.Create(IrModelDataModel, vv, nil)
}

// UpdateIrModelData updates an existing ir.model.data record.
func (c *Client) UpdateIrModelData(imd *IrModelData) error {
	return c.UpdateIrModelDatas([]int64{imd.Id.Get()}, imd)
}

// UpdateIrModelDatas updates existing ir.model.data records.
// All records (represented by ids) will be updated by imd values.
func (c *Client) UpdateIrModelDatas(ids []int64, imd *IrModelData) error {
	return c.Update(IrModelDataModel, ids, imd, nil)
}

// DeleteIrModelData deletes an existing ir.model.data record.
func (c *Client) DeleteIrModelData(id int64) error {
	return c.DeleteIrModelDatas([]int64{id})
}

// DeleteIrModelDatas deletes existing ir.model.data records.
func (c *Client) DeleteIrModelDatas(ids []int64) error {
	return c.Delete(IrModelDataModel, ids)
}

// GetIrModelData gets ir.model.data existing record.
func (c *Client) GetIrModelData(id int64) (*IrModelData, error) {
	imds, err := c.GetIrModelDatas([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imds)[0]), nil
}

// GetIrModelDatas gets ir.model.data existing records.
func (c *Client) GetIrModelDatas(ids []int64) (*IrModelDatas, error) {
	imds := &IrModelDatas{}
	if err := c.Read(IrModelDataModel, ids, nil, imds); err != nil {
		return nil, err
	}
	return imds, nil
}

// FindIrModelData finds ir.model.data record by querying it with criteria.
func (c *Client) FindIrModelData(criteria *Criteria) (*IrModelData, error) {
	imds := &IrModelDatas{}
	if err := c.SearchRead(IrModelDataModel, criteria, NewOptions().Limit(1), imds); err != nil {
		return nil, err
	}
	return &((*imds)[0]), nil
}

// FindIrModelDatas finds ir.model.data records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelDatas(criteria *Criteria, options *Options) (*IrModelDatas, error) {
	imds := &IrModelDatas{}
	if err := c.SearchRead(IrModelDataModel, criteria, options, imds); err != nil {
		return nil, err
	}
	return imds, nil
}

// FindIrModelDataIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelDataIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrModelDataModel, criteria, options)
}

// FindIrModelDataId finds record id by querying it with criteria.
func (c *Client) FindIrModelDataId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelDataModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
