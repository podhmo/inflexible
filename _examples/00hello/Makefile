export ADDR := :44444

00:
	go run cmd/hello-api/main.go

# check
x00:
#	http --json POST $(ADDR)/
	echo '{"name": "foo"}' | http --json POST $(ADDR)/Hello
x01:
	echo '{"name": "foo"}' | http --json POST $(ADDR)/Hello short==true

x10:
	go run cmd/hello-cli/main.go Hello --name foo
x11:
	SHORT=1 go run cmd/hello-cli/main.go Hello --name foo
x12:
	go run cmd/hello-cli/main.go Hello --name foo --short
