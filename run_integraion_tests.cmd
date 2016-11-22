echo "Run tests..."
go test .\integration_tests -v --tags=integration -api-path %cd%\integration_tests\api.json

