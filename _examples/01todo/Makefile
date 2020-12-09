SHELL := bash
export ADDR := :44444

00:
	go run cmd/todo-api/main.go

openapi.json: cmd/todo-gendoc/main.go $(wildcard design/*.go) $(wildcard todogenerated/*.go)
	go run cmd/todo-gendoc/main.go > $@ || rm $@

# check
x00:
	http --json POST $(ADDR)/
x01:
	http -b --json POST $(ADDR)/ListTodo
x02:
	http -b --json POST $(ADDR)/ListTodo all==true
x03:
	echo '{"todo": {"title": "foo"}}' | http -b --json POST $(ADDR)/AddTodo
x04:
	echo '{"todo": {"title": "bar", "done": true}}' | http -b --json POST $(ADDR)/AddTodo
x05:
	http -b --json POST $(ADDR)/ListTodo ALL:true

x10:
	go run cmd/todo-cli/main.go ListTodo
x11:
	go run cmd/todo-cli/main.go ListTodo --all
x12:
	go run cmd/todo-cli/main.go AddTodo --todo @todo-sample.json
x13:
	go run cmd/todo-cli/main.go AddTodo --todo @<(echo '{"title": "foo"}')
x14:
	go run cmd/todo-cli/main.go AddTodo --todo '{"title": "bar", "done": true}'
x15:
	go run cmd/todo-cli/main.go AddTodo --POSTDATA '{"todo": {"title": "foo"}}'


clean:
	rm -f todo.json
.PHONY: clean
