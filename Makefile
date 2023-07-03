all: 
	@echo "check the README.md file for instructions on how to run the program"

bin:
	cd pkg/core && go build -o hibernate ./
