run: compile_css copy_bower_packages
	go run *.go

test:
	curl -X POST -d "{\"body\":\"test body\"}" http://localhost:9021/post

compile_css:
	@lessc assets/css/style.less public/style.css

copy_bower_packages:
	@-mkdir public/js
	@cp ./bower_components/jquery/dist/jquery.min.js public/js/

live:
	git push live master

.PHONY: test compile_css
