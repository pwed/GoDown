start=`date +%s`


cd gui

ng build --prod

cd ..

go-bindata-assetfs static/...


go build -o build/Current-GoDown


end=`date +%s`
runtime=$((end-start))
echo "Build finished after $runtime seconds."
