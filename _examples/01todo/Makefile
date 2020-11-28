SHELL := bash
export ADDR := :44444

00:
	go run cmd/todo-api/main.go

# check
x00:
#	http --json POST $(ADDR)/
	http -b --json POST $(ADDR)/ListTodo
	echo '{"todo": {"title": "foo"}}' | http -b --json POST $(ADDR)/AddTodo
	http -b --json POST $(ADDR)/ListTodo

x01:
	go run cmd/todo-cli/main.go ListTodo
	go run cmd/todo-cli/main.go AddTodo --todo @todo-sample.json
	go run cmd/todo-cli/main.go AddTodo --todo @<(echo '{"title": "foo"}')
	go run cmd/todo-cli/main.go AddTodo --POSTDATA '{"todo": {"title": "foo"}}'
	go run cmd/todo-cli/main.go ListTodo
