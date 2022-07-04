package valid

type DbTTLAddParam struct {
	DSN   string `json:"dsn" binding:"required"`
	DB    string `json:"db" binding:"required"`
	Table string `json:"table" binding:"required"`
	Field string `json:"field" binding:"required"`
	Cron  string `json:"cron" binding:"required"`
	TTL   int    `json:"ttl" binding:"required"`
	Limit int    `json:"limit" binding:"required"`
	Desc  string `json:"desc"`
}

type DbTTLEditParam struct {
	ID     string `json:"id" binding:"required"`
	Field  string `json:"field"`
	Cron   string `json:"cron"`
	TTL    int    `json:"ttl"`
	Limit  int    `json:"limit"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

type DbTTLQParam struct {
	DB    string `form:"db"`
	Table string `form:"table"`
}
