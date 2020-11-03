windows:
	echo "--------------------------------------------------------------"
	echo "|---------   First make 386 version for windows...  ---------|"
	echo "--------------------------------------------------------------"
	GOOS=windows GOARCH=386  CC=/usr/bin/i686-w64-mingw32-gcc CXX=/usr/bin/i686-w64-mingw32-g++ CGO_ENABLED=0 go  build -ldflags="-w -s " -o "dist/windows/nhweb_上网助手(32位).exe" 
	upx "dist/windows/nhweb_上网助手(32位).exe"

	echo "--------------------------------------------------------------"
	echo "|--------- Second make amd64 version for windows... ---------|"
	echo "--------------------------------------------------------------"
	GOOS=windows GOARCH=amd64  CC=/usr/bin/x86_64-w64-mingw32-gcc CXX=/usr/bin/x86_64-w64-mingw32-g++ CGO_ENABLED=0 go  build  -ldflags="-w -s" -o "dist/windows/nhweb_上网助手(64位).exe" 
	upx "dist/windows/nhweb_上网助手(64位).exe"


build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "dist/nhweb_linux" 
	upx "dist/nhweb_linux"

deb: build 
	echo "--------------------------------------------------------------"
	echo "|----------- Making deb package based on debian... ----------|"
	echo "|------------After Install,modify nhweb.service ------------|"
	echo "--------------------------------------------------------------"
	chmod 0755 dist/nhweb_linux
	mkdir -p dist/nhweb/usr/local/bin/
	cp dist/nhweb_linux dist/nhweb/usr/local/bin/nhweb
	dpkg-deb -b dist/nhweb dist/nhweb.deb