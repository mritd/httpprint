version=${version:-"latest"}

all:
	gox -osarch="darwin/amd64 linux/386 linux/amd64" \
        -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}"

release: all
	ghr -u mritd -t $(GITHUB_RELEASE_TOKEN) -replace -recreate --debug $(version) dist

docker: all
	docker build -t mritd/httpprint:$(version) .

clean:
	rm -rf dist

install:
	go install

.PHONY : all release docker clean install
