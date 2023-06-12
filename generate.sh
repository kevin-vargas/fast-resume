
openssl req -newkey rsa:2048 \
  -new -nodes -x509 \
  -days 3650 \
  -out cert.pem \
  -keyout key.pem \
  -subj "/C=AR/ST=BuenosAires/L=CABA View/O=FAST/OU=FAST DEV/CN=localhost"