
runauth:
	@echo "Running the server with authentication"
	cd services/auth && go run cmd/main.go

runguestbook:
	@echo "Running the server with authentication"
	cd services/guestbook && go run cmd/main.go