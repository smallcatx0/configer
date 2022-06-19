package valid

type DbTTLParam struct {
	DSN   string `json:"dsn" binding:"required"`
	DB    string `json:"db" binding:"required"`
	Table string `json:"table" binding:"required"`
	Field string `json:"field" binding:"required"`
	Cron  string `json:"cron" binding:"required"`
	TTL   int    `json:"ttl" binding:"required"`
	Limit int    `json:"limit" binding:"required"`
	Desc  string `json:"desc"`
}
