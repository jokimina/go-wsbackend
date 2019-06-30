module go-wsbackend

go 1.12

require (
	github.com/bradfitz/gomemcache v0.0.0-20190329173943-551aad21a668 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/jinzhu/gorm v1.9.8
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mozillazg/go-pinyin v0.15.0
	github.com/silenceper/wechat v1.0.1-0.20190522143304-894b1972d710
)

replace github.com/silenceper/wechat v1.0.1-0.20190522143304-894b1972d710 => github.com/jokimina/wechat v1.0.1-0.20190630160513-172728bcdb10
