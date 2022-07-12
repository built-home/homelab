
sync:
	@git submodule sync
	@git submodule update --init --recursive --remote

docs:
	mkdocs serve