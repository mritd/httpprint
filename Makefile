all:
	gox -osarch="darwin/amd64 linux/386 linux/amd64" \
        -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"

docker: all
	docker build -t mritd/httpprint .

clean:
	rm -rf dist

install:
	go install

.PHONY : all docker clean install
