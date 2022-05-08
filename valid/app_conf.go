package valid

type EnvAddParam struct {
	Name       string `json:"name" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Sign       string `json:"sign" binding:"required"`
	Desc       string `json:"desc"`
	OwnerName  string `json:"owner"`
	OwnerPhone string `json:"owner_phone"`
}
