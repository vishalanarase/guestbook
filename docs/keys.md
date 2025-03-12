# Generate RSA Key Pair (Private & Public Keys)

You can generate the RSA key pair using the OpenSSL command-line tool or by writing a Go program.

Using OpenSSL to Generate RSA Keys

1. Generate a private key:

```bash
openssl genpkey -algorithm RSA -out auth-private-key.pem -aes256
```

2. Verify the private key:

```bash
openssl rsa -in auth-private-key.pem -check
```

3. Extract the public key from the private key:

```bash
openssl rsa -pubout -in auth-private-key.pem -out auth-public-key.pem
```

Now you have:

`auth-private-key.pem ` (used for signing JWTs)
`auth-public-key.pem` (used for validating JWTs)
