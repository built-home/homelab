
sync:
	@git submodule sync
	@git submodule update --init --recursive

docs:
	mkdocs serve