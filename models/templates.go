package models

type Templates struct {
	ID           int    `json:"code"`
	Name         string `json:"name"`
	Description  int    `json:"-"`
	TemplateType int    `json:"-"`
	PreviewURL   int    `json:"preview_url"`
}

func GetTemplates(templateType int64) ([]*Templates, error) {
	var templates []*Templates
	err := DB.Where("template_type = ?", templateType).Find(&templates).Error
	return templates, err
}
