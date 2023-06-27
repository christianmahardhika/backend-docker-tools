.PHONY: deploy tools dockerize

run-tools:
	@echo deploying tools..
	docker compose up -d

clean-tools:
	@echo tearing down tools...
	docker compose down --rmi all
