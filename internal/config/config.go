package config

type Apiuser struct{
	Api_port string
	Api_secret string
	Save_dir string
}
type Apiadmin struct{
	Api_port string
	Api_secret string
	Save_dir string
}
type Wxconfig struct{
	App_id string
	App_secret string
	Oauth_url string
	Grant_type string
}
type Alisms struct{
	Template_code string
	Sign_name string
	Access_key_id string
	Access_key_secret string
	Region_id string
}
type S3cnf struct{
	Bucket string
	Access_key string
	Secret_key string
	Region string
	Host string
}
type Wxpay struct{
	App_id string
	Mch_id string
	Key string
	Notify_url string
	Trade_type string
}
type Alipay struct{
	App_id string
	Private_key string
	Ali_pay_public_key string
	Notify_url string
	Return_url string
}
type Mysql struct{
	Db_name string
	Driver string
	User string
	Pass_word string
	Port string
	Loc string
	Host string
	Charset string
	Show_sql string
	Parsetime string
}
type Redis struct{
	Pass_word string
	Db float64
	Host string
	Port string
}
type Config struct{
	host string
	Apiuser Apiuser
	Apiadmin Apiadmin
	Wxconfig Wxconfig
	Alisms Alisms
	S3cnf S3cnf
	Wxpay Wxpay
	Alipay Alipay
	Mysql Mysql
	Redis Redis
}