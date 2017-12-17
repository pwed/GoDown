rm bindata_assetfs.go

cd Angular/GoDown/


go-bindata-assetfs dist/...

mv bindata_assetfs.go ../..

cd ../..


go build
