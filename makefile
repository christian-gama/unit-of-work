run:
	@docker compose up -d --remove-orphans
	@go run .