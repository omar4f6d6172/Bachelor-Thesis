.PHONY: sync

MSG ?= sync

sync:
	git add -A
	@if ! git diff --cached --quiet; then \
		git commit -m "$(MSG)"; \
	else \
		echo "nothing to commit"; \
	fi
	git pull --rebase
	git push
