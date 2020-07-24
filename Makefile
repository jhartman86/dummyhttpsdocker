SHELL = /bin/bash
MAKEFLAGS += --no-print-directory --silent
APP_ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

runssl:
	docker-compose up; true && make down
rebuild:
	docker-compose build
down:
	docker-compose down