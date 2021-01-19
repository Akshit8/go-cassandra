git:
	git add .
	git commit -m "$(msg)"
	git push origin master

cassandra: docker-compose up -d

.PHONY: git cassandra