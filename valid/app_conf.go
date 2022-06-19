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

type AppAddParam struct {
	Name       string `json:"name" binding:"required"`
	Sign       string `json:"sign" binding:"required"`
	Desc       string `json:"desc"`
	OwnerName  string `json:"owner"`
	OwnerPhone string `json:"owner_phone"`
}

type AppEditParam struct {
	Sign       string `json:"sign" binding:"required"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	OwnerName  string `json:"owner"`
	OwnerPhone string `json:"owner_phone"`
}

type AppFileAddParam struct {
	AppSign     string `json:"app_sign" binding:"required"`
	EnvSign     string `json:"env_sign" binding:"required"`
	Header      string `json:"header"`
	HeaderPhone string `json:"header_phone"`
	FileName    string `json:"file_name" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Type        string `json:"type"`
}

type AppFileParam struct {
	App  string `form:"app" json:"app" binding:"required"`
	Env  string `form:"env" json:"env" binding:"required"`
	File string `form:"file" json:"file" binding:"required"`
}

type AppConfParam struct {
	App string `form:"app" binding:"required"`
	Env string `form:"env" binding:"required"`
}
