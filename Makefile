all: 
	@echo "check the README.md file for instructions on how to run the program"

bin:
	cd pkg/core && go build -o hibernate ./

release_patch:
	hack/version_bump.sh ./version.txt patch
	VERSION=$(cat ./version.txt)
	git tag $VERSION
	git push origin $VERSION 

release_minor:
	hack/version_bump.sh ./version.txt minor
	VERSION=$(cat ./version.txt)
	git tag $VERSION
	git push origin $VERSION 

release_major:
	hack/version_bump.sh ./version.txt major
	VERSION=$(cat ./version.txt)
	git tag $VERSION
	git push origin $VERSION 
