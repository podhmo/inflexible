export ADDR := :44444

00:
	go run main.go

# check
x00:
#	http --json POST $(ADDR)/
	echo '{"name": "foo"}' | http --json POST $(ADDR)/Hello

x01:
	go run cmd/hello-cli/main.go Hello --name foo
