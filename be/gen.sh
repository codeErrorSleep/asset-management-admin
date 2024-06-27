# api
goctl api go --api api/doc/admin.api --dir ./api --remote https://github.com/codeErrorSleep/go-zero-template
# 快速格式话import的
# goimports -w .
# mysql
goctl model mysql ddl --src model/mysql/bs.sql --dir model/mysql
# swagger
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api api/doc/*.api -dir api/doc/
