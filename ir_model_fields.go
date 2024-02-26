package odoo

// IrModelFields represents ir.model.fields model.
type IrModelFields struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	Column1              *String    `xmlrpc:"column1,omptempty"`
	Column2              *String    `xmlrpc:"column2,omptempty"`
	CompleteName         *String    `xmlrpc:"complete_name,omptempty"`
	Compute              *String    `xmlrpc:"compute,omptempty"`
	Copy                 *Bool      `xmlrpc:"copy,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	Depends              *String    `xmlrpc:"depends,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	Domain               *String    `xmlrpc:"domain,omptempty"`
	FieldDescription     *String    `xmlrpc:"field_description,omptempty"`
	Groups               *Relation  `xmlrpc:"groups,omptempty"`
	Help                 *String    `xmlrpc:"help,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	Index                *Bool      `xmlrpc:"index,omptempty"`
	Model                *String    `xmlrpc:"model,omptempty"`
	ModelId              *Many2One  `xmlrpc:"model_id,omptempty"`
	Modules              *String    `xmlrpc:"modules,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	OnDelete             *Selection `xmlrpc:"on_delete,omptempty"`
	Readonly             *Bool      `xmlrpc:"readonly,omptempty"`
	Related              *String    `xmlrpc:"related,omptempty"`
	Relation             *String    `xmlrpc:"relation,omptempty"`
	RelationField        *String    `xmlrpc:"relation_field,omptempty"`
	RelationTable        *String    `xmlrpc:"relation_table,omptempty"`
	Required             *Bool      `xmlrpc:"required,omptempty"`
	Selectable           *Bool      `xmlrpc:"selectable,omptempty"`
	Selection            *String    `xmlrpc:"selection,omptempty"`
	SerializationFieldId *Many2One  `xmlrpc:"serialization_field_id,omptempty"`
	Size                 *Int       `xmlrpc:"size,omptempty"`
	State                *Selection `xmlrpc:"state,omptempty"`
	Store                *Bool      `xmlrpc:"store,omptempty"`
	TrackVisibility      *Selection `xmlrpc:"track_visibility,omptempty"`
	Translate            *Bool      `xmlrpc:"translate,omptempty"`
	Ttype                *Selection `xmlrpc:"ttype,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrModelFieldss represents array of ir.model.fields model.
type IrModelFieldss []IrModelFields

// IrModelFieldsModel is the odoo model name.
const IrModelFieldsModel = "ir.model.fields"

// Many2One convert IrModelFields to *Many2One.
func (imf *IrModelFields) Many2One() *Many2One {
	return NewMany2One(imf.Id.Get(), "")
}

// CreateIrModelFields creates a new ir.model.fields model and returns its id.
func (c *Client) CreateIrModelFields(imf *IrModelFields) (int64, error) {
	ids, err := c.CreateIrModelFieldss([]*IrModelFields{imf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModelFields creates a new ir.model.fields model and returns its id.
func (c *Client) CreateIrModelFieldss(imfs []*IrModelFields) ([]int64, error) {
	var vv []interface{}
	for _, v := range imfs {
		vv = append(vv, v)
	}
	return c.Create(IrModelFieldsModel, vv, nil)
}

// UpdateIrModelFields updates an existing ir.model.fields record.
func (c *Client) UpdateIrModelFields(imf *IrModelFields) error {
	return c.UpdateIrModelFieldss([]int64{imf.Id.Get()}, imf)
}

// UpdateIrModelFieldss updates existing ir.model.fields records.
// All records (represented by ids) will be updated by imf values.
func (c *Client) UpdateIrModelFieldss(ids []int64, imf *IrModelFields) error {
	return c.Update(IrModelFieldsModel, ids, imf, nil)
}

// DeleteIrModelFields deletes an existing ir.model.fields record.
func (c *Client) DeleteIrModelFields(id int64) error {
	return c.DeleteIrModelFieldss([]int64{id})
}

// DeleteIrModelFieldss deletes existing ir.model.fields records.
func (c *Client) DeleteIrModelFieldss(ids []int64) error {
	return c.Delete(IrModelFieldsModel, ids)
}

// GetIrModelFields gets ir.model.fields existing record.
func (c *Client) GetIrModelFields(id int64) (*IrModelFields, error) {
	imfs, err := c.GetIrModelFieldss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imfs)[0]), nil
}

// GetIrModelFieldss gets ir.model.fields existing records.
func (c *Client) GetIrModelFieldss(ids []int64) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.Read(IrModelFieldsModel, ids, nil, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
}

// FindIrModelFields finds ir.model.fields record by querying it with criteria.
func (c *Client) FindIrModelFields(criteria *Criteria) (*IrModelFields, error) {
	imfs := &IrModelFieldss{}
	if err := c.SearchRead(IrModelFieldsModel, criteria, NewOptions().Limit(1), imfs); err != nil {
		return nil, err
	}
	return &((*imfs)[0]), nil
}

// FindIrModelFieldss finds ir.model.fields records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelFieldss(criteria *Criteria, options *Options) (*IrModelFieldss, error) {
	imfs := &IrModelFieldss{}
	if err := c.SearchRead(IrModelFieldsModel, criteria, options, imfs); err != nil {
		return nil, err
	}
	return imfs, nil
}

// FindIrModelFieldsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelFieldsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrModelFieldsModel, criteria, options)
}

// FindIrModelFieldsId finds record id by querying it with criteria.
func (c *Client) FindIrModelFieldsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelFieldsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
