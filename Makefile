md-update-deps:
	cd docGen && go get github.com/khulnasoft-lab/defsec \
	&& go mod tidy

md-build: 
	cd docGen && go build -o ../generator .

md-test:
	cd docGen && go test -v ./...

md-clean:
	rm -f ./generator

md-clone-all:
	# git clone git@github.com:khulnasoft-lab/cvedb.git cvedb-repo/
	git clone git@github.com:khulnasoft-lab/vuln-list.git cvedb-repo/vuln-list
	git clone git@github.com:khulnasoft-lab/kube-hunter.git cvedb-repo/kube-hunter-repo
	git clone git@github.com:khulnasoft-lab/kube-bench.git cvedb-repo/kube-bench-repo
	git clone git@github.com:khulnasoft-lab/oss-chain-bench.git cvedb-repo/chain-bench-repo
	git clone git@github.com:khulnasoft-lab/cloud-security-remediation-guides.git cvedb-repo/remediations-repo
	git clone git@github.com:khulnasoft-lab/tracker.git cvedb-repo/tracker-repo
	git clone git@github.com:khulnasoft-lab/defsec.git cvedb-repo/defsec-repo
	git clone git@github.com:khulnasoft-lab/cloudsploit.git cvedb-repo/cloudsploit-repo

update-all-repos:
	cd cvedb-repo/vuln-list && git pull
	cd cvedb-repo/kube-hunter-repo && git pull
	cd cvedb-repo/kube-bench-repo && git pull
	cd cvedb-repo/chain-bench-repo && git pull
	cd cvedb-repo/remediations-repo && git pull
	cd cvedb-repo/tracker-repo && git pull
	cd cvedb-repo/defsec-repo && git pull
	cd cvedb-repo/cloudsploit-repo && git pull

sync-all:
	rsync -av ./ cvedb-repo/ --exclude=.idea --exclude=go.mod --exclude=go.sum --exclude=nginx.conf --exclude=main.go --exclude=main_test.go --exclude=README.md --exclude=cvedb-repo --exclude=.git --exclude=.gitignore --exclude=.github --exclude=content --exclude=docs --exclude=Makefile --exclude=goldens

md-generate:
	cd cvedb-repo && ./generator

nginx-start:
	-cd cvedb-repo/docs && nginx -p . -c ../../nginx.conf

nginx-stop:
	-cd cvedb-repo/docs && nginx -s stop -p . -c ../../nginx.conf

nginx-restart:
	make nginx-stop nginx-start

hugo-devel:
	hugo server -D --debug

hugo-clean:
	cd cvedb-repo && rm -rf docs

hugo-generate: hugo-clean
	cd cvedb-repo && hugo --destination=docs
	echo "cvedb.khulnasoft.com" > cvedb-repo/docs/CNAME

simple-host:
	cd cvedb-repo && python3 -m http.server

copy-assets:
	cp -R cvedb-repo/remediations-repo/resources cvedb-repo/docs/resources
	touch cvedb-repo/docs/.nojekyll

build-all-no-clone: md-clean md-build sync-all md-generate hugo-generate copy-assets nginx-restart
	echo "Build Done, navigate to http://localhost:9011/ to browse"

build-all: md-clean md-build md-clone-all sync-all md-generate hugo-generate copy-assets nginx-restart
	echo "Build Done, navigate to http://localhost:9011/ to browse"

compile-theme-sass:
	cd themes/khulnasoftblank/static/sass && sass cvedbblank.scss:../css/cvedbblank.css && sass cvedbblank.scss:../css/cvedbblank.min.css --style compressed
