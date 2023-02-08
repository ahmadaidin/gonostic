package model

type Search struct {
	Search string `json:"search" query:"search"`
}

type Sort struct {
	Sort []string `json:"-"`
}

func (s Sort) UnmarshalParam(param string) error {
	return nil
}
