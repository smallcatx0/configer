package valid

type EnvAddParam struct {
	Name       string `json:"name" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Sign       string `json:"sign" binding:"required"`
	Desc       string `json:"desc"`
	OwnerName  string `json:"owner"`
	OwnerPhone string `json:"owner_phone"`
}

type EnvEditParam struct {
	Sign       string `json:"sign" binding:"required"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	OwnerName  string `json:"owner"`
	OwnerPhone string `json:"owner_phone"`
}
