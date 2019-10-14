#!/bin/bash
if [[ $1 == "windows" ]]; then

echo "--------------------------------------------------------------"
echo "|---------   First make 386 version for windows...  ---------|"
echo "--------------------------------------------------------------"

GOOS=windows GOARCH=386  CC=/usr/bin/i686-w64-mingw32-gcc CXX=/usr/bin/i686-w64-mingw32-g++ CGO_ENABLED=1 go  build -ldflags="-w -s " -o "dist/windows/nhweb_上网助手(32位).exe" 

upx "dist/windows/nhweb_上网助手(32位).exe"

# build amd64

echo "--------------------------------------------------------------"
echo "|--------- Second make amd64 version for windows... ---------|"
echo "--------------------------------------------------------------"

GOOS=windows GOARCH=amd64  CC=/usr/bin/x86_64-w64-mingw32-gcc CXX=/usr/bin/x86_64-w64-mingw32-g++ CGO_ENABLED=1 go  build  -ldflags="-w -s" -o "dist/windows/nhweb_上网助手(64位).exe" 

upx "dist/windows/nhweb_上网助手(64位).exe"

fi

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "dist/nhweb_linux" 

if [[ $1 == "prod" ]];then

echo "--------------------------------------------------------------"
echo "|----------- Making deb package based on debian... ----------|"
echo "|------------After Install,modify nhweb.service ------------|"
echo "--------------------------------------------------------------"
upx "dist/nhweb_linux"
cd dist
chmod 0755 nhweb_linux
mkdir -p nhweb/usr/local/bin/
cp nhweb_linux nhweb/usr/local/bin/nhweb
dpkg-deb --build nhweb

fi
