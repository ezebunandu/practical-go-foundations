test:
	staticcheck .
	go test -v .

fuzz:
	go test -run NONE -fuzz . -fuzztime 30s