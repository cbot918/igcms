run:
	go run .

test:
	echo $(@ipy testpsql)

# database
psqlc:
	docker run -it --rm --network web_default postgres psql -h 172.19.0.2 -U postgres -W

createdb: 
	docker exec -it postgres createdb --username=postgres --owner=postgres igcms 
migrateup: 
	migrate -path internal/db/migrations -database 'postgresql://postgres:12345@localhost:5432/igcms?sslmode=disable' up 

migratedown:
	migrate -path internal/db/migrations -database "postgresql://postgres:12345@localhost:5432/igcms?sslmode=disable" -verbose down


.PHONY: psqlc createdb migrateup
.SILENT: psql createdb migrateup


