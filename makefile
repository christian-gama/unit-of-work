run:
	@docker compose up api --build --force-recreate --remove-orphans

stop:
	@docker compose down