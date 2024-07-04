
 migrate up 
```bash
go run migrations/migrate.go up
```
 
 migrate down 
 ```bash
go run migrations/migrate.go down
 ```
 in case of version error  ` - fix manually sry`

 to create new migrations 
 ```bash
   migrate create -ext sql -dir migrations/ -seq init_mg
 ```