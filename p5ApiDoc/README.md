[scripts]
- go get github.com/swaggo/swag/cmd/swag
- swag init
- go get -u github.com/swaggo/http-swagger
- go get -u github.com/alecthomas/template
- swag init -g controller/TestController.go

[example Reference]
- https://www.soberkoder.com/swagger-go-api-swaggo/


[API Operation annotations]
- **@Success** annotation
  - {array} specifies that the response is an array of type greetingResponse
  - {object} specifies that the response is an object of type greetingResponse
- **@Param** annotation : described request body 
  -  @Param [param_name] [param_type] [data_type] [required/mandatory] [description]
  ```
  The param_type can be one of the following values:
  
  query (indicates a query param)
  path (indicates a path param)
  header (indicates a header param)
  body
  formData
  ```
  
[Http Test]  
- godotenv - to load .env file
- godotenv go test ./test/

