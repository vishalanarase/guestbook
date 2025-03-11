
runauth:
	@echo "Running the server with authentication"
	cd services/auth && go run cmd/main.go

runguestbook:
	@echo "Running the server with authentication"
	cd services/guestbook && go run cmd/main.go

bufgenauth:
	@echo "Generating the buf for the auth service"
	cd apis/auth && buf generate

bufgenguestbook:
	@echo "Generating the buf for the guestbook service"
	cd apis/guestbook && buf generate