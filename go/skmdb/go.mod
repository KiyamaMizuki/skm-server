module local.com/db-module

go 1.14

require (
	github.com/jinzhu/gorm v1.9.16
	local.com/db-module/model v0.0.0-00010101000000-000000000000
	local.com/db-module/val v0.0.0-00010101000000-000000000000
)

replace (
	local.com/db-module/model => ./model
	local.com/db-module/val => ./val
)
