git-add:
	git add .
	git commit -am '${cmt}'

test:
	echo ${cmt}

entermysql:
	docker exec -it mysql_simple_login mysql -u ADMIN -pSECRET rakamin_intern

entermysqlroot:
	docker exec -it mysql_simple_login mysql -u root -pSECRET_ROOT rakamin_intern

runenv:
	docker compose up -d

run:
	docker compose up -d
	go run app/main.go

commit:
	git add .
	git commit -am '${cmt}'

struct:
	gomodifytags -file ${file} -struct ${struct} -add-tags ${tags}

stop:
	docker compose stop

down:
	docker compose down -v

logs:
	docker compose logs -f

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  ./dist/example ./app/main.go

dockerbuild:
	docker build --rm -t simple_login .
	docker image prune --filter label=stage=dockerbuilder -f

dockerun:
	docker run --name simple_login  -p 8080:8080 simple_login 

dockerrm:
	docker rm simple_login -f
	docker rmi simple_login

dockeenter:
	docker exec -it simple_login bash