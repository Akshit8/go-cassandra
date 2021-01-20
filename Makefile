git:
	git add .
	git commit -m "$(msg)"
	git push origin master

cassandra: docker-compose up -d

server: go run main.go

.PHONY: git cassandra server