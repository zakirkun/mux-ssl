mkdir -p certs
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout certs/localhost.key -out certs/localhost.crt \
    -subj "/C=NL/ST=Limburg/L=Geleen/O=Marco Franssen/OU=Development/CN=localhost/emailAddress=marco.franssen@gmail.com"

go mod init mux-ssl
go mod tidy
go build -o mux-ssl && ./mux-ssl