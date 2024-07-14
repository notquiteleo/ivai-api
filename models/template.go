package models

type Templates struct {
	ID           int    `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"-"`
	TemplateType int    `json:"-"`
	PreviewURL   string `json:"preview_url"`
}

func GetTemplates(templateType int64) ([]*Templates, error) {
	var templates []*Templates
	err := DB.Where("template_type = ?", templateType).Find(&templates).Error
	return templates, err
}
