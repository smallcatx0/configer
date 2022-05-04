package mdb

type Contact struct {
	Name  string
	Phone string
}

type AppEnv struct {
	Name      string
	Sign      string
	Desc      string
	Principal Contact
}

type AppConf struct {
	AppName   string
	AppSign   string
	Env       string
	FileName  string
	Content   string
	CreateAt  string
	Principal Contact
}
