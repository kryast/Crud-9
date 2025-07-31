CRUD ke-9

POST
curl -X POST http://localhost:8080/articles \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Golang",
    "content": "Golang itu bahasa pemrograman yang cepat dan efisien.",
    "author": "Ahmad Syarifuddin"
}'


GET
curl http://localhost:8080/articles
curl http://localhost:8080/articles/1

PUT
curl -X PUT http://localhost:8080/articles/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Golang Update",
    "content": "Update: Golang sangat cocok untuk backend services.",
    "author": "Ahmad"
}'

DELETE
curl -X DELETE http://localhost:8080/articles/1