.PHONY: sync auto-sync test-db

MSG ?= sync
INTERVAL ?= 5

sync:
	git add -A
	@if ! git diff --cached --quiet; then \
		git commit -m "$(MSG)"; \
	else \
		echo "nothing to commit"; \
	fi
	git pull --rebase --autostash
	git push

auto-sync:
	@echo "auto-sync every $(INTERVAL)s — Ctrl+C to stop"
	@while true; do \
		git add -A; \
		if ! git diff --cached --quiet; then \
			git commit -m "$(MSG) @ $$(date +%H:%M:%S)" || true; \
		fi; \
		git pull --rebase --autostash || { echo "rebase conflict — resolve manually"; exit 1; }; \
		git push || echo "push failed, will retry"; \
		sleep $(INTERVAL); \
	done

test-db:
	go test -v -run TestStartPostgres ./internal/testutil/...
