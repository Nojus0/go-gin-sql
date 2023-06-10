echo "Building..."
GOOS=linux go build -tags lambda
echo "Compressing..."
powershell.exe Compress-Archive -Path go-gin-sql -DestinationPath aws.zip
echo "Uploading Zip File"
aws lambda update-function-code --function-name go-example-api --zip-file fileb://aws.zip --output table --no-cli-pager
rm aws.zip go-gin-sql
echo "Done"
