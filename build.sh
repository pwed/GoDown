rm bindata_assetfs.go

cd Angular/GoDown/

ng build


go-bindata-assetfs dist/...

mv bindata_assetfs.go ../..

cd ../..


go build -o build/Current-GoDown
