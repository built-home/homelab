SHELL := /bin/bash

sync:
	@git submodule sync
	@git submodule update --init --recursive --remote

serve_docs:
	mkdocs serve

