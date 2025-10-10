package odoo

// ProductDocument represents product.document model.
type ProductDocument struct {
	AccessToken    *String    `xmlrpc:"access_token,omitempty"`
	Active         *Bool      `xmlrpc:"active,omitempty"`
	AttachedOnSale *Selection `xmlrpc:"attached_on_sale,omitempty"`
	Checksum       *String    `xmlrpc:"checksum,omitempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	Datas          *String    `xmlrpc:"datas,omitempty"`
	DbDatas        *String    `xmlrpc:"db_datas,omitempty"`
	Description    *String    `xmlrpc:"description,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	FileSize       *Int       `xmlrpc:"file_size,omitempty"`
	FormFieldIds   *Relation  `xmlrpc:"form_field_ids,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	ImageHeight    *Int       `xmlrpc:"image_height,omitempty"`
	ImageSrc       *String    `xmlrpc:"image_src,omitempty"`
	ImageWidth     *Int       `xmlrpc:"image_width,omitempty"`
	IndexContent   *String    `xmlrpc:"index_content,omitempty"`
	IrAttachmentId *Many2One  `xmlrpc:"ir_attachment_id,omitempty"`
	IsBackground   *Bool      `xmlrpc:"is_background,omitempty"`
	LocalUrl       *String    `xmlrpc:"local_url,omitempty"`
	Mimetype       *String    `xmlrpc:"mimetype,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	OriginalId     *Many2One  `xmlrpc:"original_id,omitempty"`
	Public         *Bool      `xmlrpc:"public,omitempty"`
	Raw            *String    `xmlrpc:"raw,omitempty"`
	ResField       *String    `xmlrpc:"res_field,omitempty"`
	ResId          *Many2One  `xmlrpc:"res_id,omitempty"`
	ResModel       *String    `xmlrpc:"res_model,omitempty"`
	ResName        *String    `xmlrpc:"res_name,omitempty"`
	Sequence       *Int       `xmlrpc:"sequence,omitempty"`
	StoreFname     *String    `xmlrpc:"store_fname,omitempty"`
	Type           *Selection `xmlrpc:"type,omitempty"`
	Url            *String    `xmlrpc:"url,omitempty"`
	VoiceIds       *Relation  `xmlrpc:"voice_ids,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductDocuments represents array of product.document model.
type ProductDocuments []ProductDocument

// ProductDocumentModel is the odoo model name.
const ProductDocumentModel = "product.document"

// Many2One convert ProductDocument to *Many2One.
func (pd *ProductDocument) Many2One() *Many2One {
	return NewMany2One(pd.Id.Get(), "")
}

// CreateProductDocument creates a new product.document model and returns its id.
func (c *Client) CreateProductDocument(pd *ProductDocument) (int64, error) {
	ids, err := c.CreateProductDocuments([]*ProductDocument{pd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductDocument creates a new product.document model and returns its id.
func (c *Client) CreateProductDocuments(pds []*ProductDocument) ([]int64, error) {
	var vv []interface{}
	for _, v := range pds {
		vv = append(vv, v)
	}
	return c.Create(ProductDocumentModel, vv, nil)
}

// UpdateProductDocument updates an existing product.document record.
func (c *Client) UpdateProductDocument(pd *ProductDocument) error {
	return c.UpdateProductDocuments([]int64{pd.Id.Get()}, pd)
}

// UpdateProductDocuments updates existing product.document records.
// All records (represented by ids) will be updated by pd values.
func (c *Client) UpdateProductDocuments(ids []int64, pd *ProductDocument) error {
	return c.Update(ProductDocumentModel, ids, pd, nil)
}

// DeleteProductDocument deletes an existing product.document record.
func (c *Client) DeleteProductDocument(id int64) error {
	return c.DeleteProductDocuments([]int64{id})
}

// DeleteProductDocuments deletes existing product.document records.
func (c *Client) DeleteProductDocuments(ids []int64) error {
	return c.Delete(ProductDocumentModel, ids)
}

// GetProductDocument gets product.document existing record.
func (c *Client) GetProductDocument(id int64) (*ProductDocument, error) {
	pds, err := c.GetProductDocuments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pds)[0]), nil
}

// GetProductDocuments gets product.document existing records.
func (c *Client) GetProductDocuments(ids []int64) (*ProductDocuments, error) {
	pds := &ProductDocuments{}
	if err := c.Read(ProductDocumentModel, ids, nil, pds); err != nil {
		return nil, err
	}
	return pds, nil
}

// FindProductDocument finds product.document record by querying it with criteria.
func (c *Client) FindProductDocument(criteria *Criteria) (*ProductDocument, error) {
	pds := &ProductDocuments{}
	if err := c.SearchRead(ProductDocumentModel, criteria, NewOptions().Limit(1), pds); err != nil {
		return nil, err
	}
	return &((*pds)[0]), nil
}

// FindProductDocuments finds product.document records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductDocuments(criteria *Criteria, options *Options) (*ProductDocuments, error) {
	pds := &ProductDocuments{}
	if err := c.SearchRead(ProductDocumentModel, criteria, options, pds); err != nil {
		return nil, err
	}
	return pds, nil
}

// FindProductDocumentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductDocumentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductDocumentModel, criteria, options)
}

// FindProductDocumentId finds record id by querying it with criteria.
func (c *Client) FindProductDocumentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductDocumentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
