module relation_service

go 1.17

require (
	github.com/cloudwego/hertz v0.6.7
	golang.org/x/crypto v0.6.0
	google.golang.org/protobuf v1.31.0
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.2
	gorm.io/sharding v0.5.3
	user_service v0.0.0
)

require (
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.0.0-20220413063733-65bf48ffb3a7 // indirect
	github.com/bytedance/sonic v1.8.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cloudwego/netpoll v0.3.2 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/longbridgeapp/sqlparser v0.3.1 // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/tidwall/gjson v1.13.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/sys v0.5.0 // indirect
)

replace user_service v0.0.0 => ../UserService
