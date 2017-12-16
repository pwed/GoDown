go-bindata-assetfs static/...

env GOOS=windows GOARCH=amd64 go build -o build/GoDown-Windows-AMD64 -v github.com/Pwed/GoDown
env GOOS=windows GOARCH=386 go build -o build/GoDown-Windows-32x86 -v github.com/Pwed/GoDown
env GOOS=linux GOARCH=amd64 go build -o build/GoDown-Linux-AMD64 -v github.com/Pwed/GoDown
env GOOS=linux GOARCH=386 go build -o build/GoDown-Linux-32x86 -v github.com/Pwed/GoDown
env GOOS=linux GOARCH=arm go build -o build/GoDown-Linux-arm -v github.com/Pwed/GoDown
env GOOS=linux GOARCH=arm64 go build -o build/GoDown-Linux-arm64 -v github.com/Pwed/GoDown
