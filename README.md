# golang-rest-raml-validation

This project is for a lightning talk I hacked together for [GopherCon 2015](https://youtu.be/jLAuWPls_W0).

### Generate documentation
~~~
$ npm i -g raml2html
$ raml2html -i web/schemas/api.raml -o web/docs/index.html
~~~

### Run the API
~~~
$ go build
$ ./golang-rest-raml-validation
~~~

### Read the Docs
~~~
$ open http://localhost:8080/docs/index.html
$ open http://localhost:8080/schemas/api.raml
$ open http://localhost:8080/schemas/keyvalue.post.body.json
~~~

### Test it out
~~~
$ curl -X POST -d '' http://localhost:8080/api/v0/keys
$ curl -X POST -d '{"foo":"bar"}' http://localhost:8080/api/v0/keys
$ curl -X POST -d '{"value":"something"}' http://localhost:8080/api/v0/keys
~~~
