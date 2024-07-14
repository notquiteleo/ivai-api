package models

import "encoding/json"

type Filters struct {
	ID           int    `json:"code"`
	Name         string `json:"name"`
	ParentID     int    `json:"-"`
	FilterType   int    `json:"-"`
	TemplateType int    `json:"-"`

	SubList []*Filters `json:"sub_list,omitempty" gorm:"foreignKey:ParentID"`
}

var FilterMap = map[int]string{
	0: "experience",
	1: "job",
	2: "position",
	3: "school",
}

func (f Filters) MarshalJSON() ([]byte, error) {
	type Alias Filters
	temp := struct {
		Alias
		ParentID *int `json:"parentID,omitempty"`
	}{
		Alias: (Alias)(f),
	}
	if f.ParentID > 0 {
		temp.ParentID = &f.ParentID
	}
	return json.Marshal(temp)
}

func GetTemplateFilters() (map[string][]*Filters, error) {
	var filters []*Filters
	err := DB.Find(&filters).Error
	if err != nil {
		return nil, err
	}

	groupFilters := make(map[string][]*Filters)
	for _, f := range filters {
		if f.ParentID == 0 {
			groupFilters[FilterMap[f.FilterType]] = append(groupFilters[FilterMap[f.FilterType]], f)
		}
	}

	for _, rootFilters := range groupFilters {
		for _, rootFilter := range rootFilters {
			fillSubList(rootFilter, filters)
		}
	}

	return groupFilters, nil
}

func fillSubList(parent *Filters, filters []*Filters) {
	for _, filter := range filters {
		if filter.ParentID == parent.ID {
			parent.SubList = append(parent.SubList, filter)
			fillSubList(filter, filters)
		}
	}
}
