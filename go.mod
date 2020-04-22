module godemo

go 1.13

replace skgo v0.0.1 => ../skgo

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/mock v1.4.3 // indirect
	github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0
	github.com/google/wire v0.4.0
	github.com/jinzhu/gorm v1.9.12
	github.com/prometheus/common v0.4.0
	google.golang.org/protobuf v1.21.0
	skgo v0.0.1 // indirect
)
