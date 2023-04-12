module local.com/go-api

go 1.14

require (
	github.com/bdwilliams/go-jsonify v0.0.0-20141020182238-48749139e742
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/sys v0.0.0-20201211090839-8ad439b19e0f // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	local.com/db-module/model v0.0.0-00010101000000-000000000000
	local.com/db-module/val v0.0.0-00010101000000-000000000000
	local.com/go-api/skmdb v0.0.0-00010101000000-000000000000
)

replace (
	local.com/db-module/model => ./skmdb/model
	local.com/db-module/val => ./skmdb/val
	local.com/go-api/server => ./server
	local.com/go-api/skmdb => ./skmdb
)