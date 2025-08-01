run:
	go run cmd/subscription-service/main.go

up:
	docker-compose up -d --build

down:
	docker-compose down

restart:
	docker-compose down && docker-compose up -d --build

clean:
	docker volume rm effectivemobiletest_postgres_data && docker image rm effectivemobiletest-subscriptions
