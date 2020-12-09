SHELL := bash
export ADDR := :44444

00:
	go run cmd/todo-api/main.go

openapi.json: cmd/todo-gendoc/main.go $(wildcard design/*.go) $(wildcard todogenerated/*.go)
	go run cmd/todo-gendoc/main.go > $@ || rm $@

# check
x00:
#	http --json POST $(ADDR)/
	http -b --json POST $(ADDR)/ListTodo all==true
	echo '{"todo": {"title": "foo"}}' | http -b --json POST $(ADDR)/AddTodo
	echo '{"todo": {"title": "bar", "done": true}}' | http -b --json POST $(ADDR)/AddTodo
	http -b --json POST $(ADDR)/ListTodo all==true
	http -b --json POST $(ADDR)/ListTodo

x01:
	go run cmd/todo-cli/main.go ListTodo
x02:
	go run cmd/todo-cli/main.go ListTodo --all
x03:
	go run cmd/todo-cli/main.go AddTodo --todo @todo-sample.json
x04:
	go run cmd/todo-cli/main.go AddTodo --todo @<(echo '{"title": "foo"}')
x05:
	go run cmd/todo-cli/main.go AddTodo --todo '{"title": "bar", "done": true}'
x06:
	go run cmd/todo-cli/main.go AddTodo --POSTDATA '{"todo": {"title": "foo"}}'


clean:
	rm -f todo.json
.PHONY: clean
