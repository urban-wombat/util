util_gotables = ../gotables/util.go
util_flattables = ../flattables/util.go
util_gotecho = ../gotables/cmd/gotecho/util.go

all: $(util_gotables) $(util_flattables) $(util_gotecho)

$(util_gotables): util.go
	cat util.go | sed 's/package util/package gotables/' > $(util_gotables)

$(util_flattables): util.go
	cat util.go | sed 's/package util/package flattables/' > $(util_flattables)

$(util_gotecho): util.go
	cat util.go | sed 's/package util/package main/' > $(util_gotecho)
