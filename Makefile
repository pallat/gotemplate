.PHONY: init db image up vegeta reload load plot hist
init:
	pre-commit install --config ./githook/pre-commit-config.yaml
db:
	sqlite3 school.sqlite < ./scripts/create_table_courses.sql
	sqlite3 school.sqlite < ./scripts/insert_into_courses.sql
postgres:
	docker run --name some-postgres -e POSTGRES_PASSWORD=2009 -e POSTGRES_USER=postgres -e POSTGRES_DB=school -d -p 5432:5432 postgres
image:
	docker build -t todo:latest -f Dockerfile .
up:
	docker run --init -p:8080:8080 --env-file ./.env --name myapp todo:latest
vegeta:
	go install github.com/tsenart/vegeta@latest
reload:
	echo "GET http://:8081/limitz" | vegeta attack -rate=10/s -duration=1s | vegeta report
load:
	echo "GET http://:8081/limitz" | vegeta attack -rate=10/s -duration=1s > results.10qps.bin
plot:
	 cat results.10qps.bin | vegeta plot > plot.10qps.html
hist:
	cat results.10qps.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
