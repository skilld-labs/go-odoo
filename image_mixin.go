package odoo

// ImageMixin represents image.mixin model.
type ImageMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
	Image1024   *String `xmlrpc:"image_1024,omitempty"`
	Image128    *String `xmlrpc:"image_128,omitempty"`
	Image1920   *String `xmlrpc:"image_1920,omitempty"`
	Image256    *String `xmlrpc:"image_256,omitempty"`
	Image512    *String `xmlrpc:"image_512,omitempty"`
}

// ImageMixins represents array of image.mixin model.
type ImageMixins []ImageMixin

// ImageMixinModel is the odoo model name.
const ImageMixinModel = "image.mixin"

// Many2One convert ImageMixin to *Many2One.
func (im *ImageMixin) Many2One() *Many2One {
	return NewMany2One(im.Id.Get(), "")
}

// CreateImageMixin creates a new image.mixin model and returns its id.
func (c *Client) CreateImageMixin(im *ImageMixin) (int64, error) {
	ids, err := c.CreateImageMixins([]*ImageMixin{im})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateImageMixin creates a new image.mixin model and returns its id.
func (c *Client) CreateImageMixins(ims []*ImageMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range ims {
		vv = append(vv, v)
	}
	return c.Create(ImageMixinModel, vv, nil)
}

// UpdateImageMixin updates an existing image.mixin record.
func (c *Client) UpdateImageMixin(im *ImageMixin) error {
	return c.UpdateImageMixins([]int64{im.Id.Get()}, im)
}

// UpdateImageMixins updates existing image.mixin records.
// All records (represented by ids) will be updated by im values.
func (c *Client) UpdateImageMixins(ids []int64, im *ImageMixin) error {
	return c.Update(ImageMixinModel, ids, im, nil)
}

// DeleteImageMixin deletes an existing image.mixin record.
func (c *Client) DeleteImageMixin(id int64) error {
	return c.DeleteImageMixins([]int64{id})
}

// DeleteImageMixins deletes existing image.mixin records.
func (c *Client) DeleteImageMixins(ids []int64) error {
	return c.Delete(ImageMixinModel, ids)
}

// GetImageMixin gets image.mixin existing record.
func (c *Client) GetImageMixin(id int64) (*ImageMixin, error) {
	ims, err := c.GetImageMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ims)[0]), nil
}

// GetImageMixins gets image.mixin existing records.
func (c *Client) GetImageMixins(ids []int64) (*ImageMixins, error) {
	ims := &ImageMixins{}
	if err := c.Read(ImageMixinModel, ids, nil, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindImageMixin finds image.mixin record by querying it with criteria.
func (c *Client) FindImageMixin(criteria *Criteria) (*ImageMixin, error) {
	ims := &ImageMixins{}
	if err := c.SearchRead(ImageMixinModel, criteria, NewOptions().Limit(1), ims); err != nil {
		return nil, err
	}
	return &((*ims)[0]), nil
}

// FindImageMixins finds image.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImageMixins(criteria *Criteria, options *Options) (*ImageMixins, error) {
	ims := &ImageMixins{}
	if err := c.SearchRead(ImageMixinModel, criteria, options, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindImageMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImageMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ImageMixinModel, criteria, options)
}

// FindImageMixinId finds record id by querying it with criteria.
func (c *Client) FindImageMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImageMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
