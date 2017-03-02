FORMAT: 1A

# Encryption API

A simple API for demonstrating AES-256 bit encryption methods and base64 encoding in Golang!

# Group Encryption

## Encryption Collection [/encrypt]

### Encrypt [POST]
Encrypts plain text using AES-256 and base64 encoding

+ key (required, string) - The password used for encrypting
+ text (required, string) - The text to be encoded

+ Request (application/json)

        {
            "text": "Hello World!",
            "key": "password"
        }

+ Response 200 (application/json)

        {
            "text": "Kl9LmR7dn0KxBTYzIXwOTyE2JAPgVUwHR8p4kjx0ZrEgICAgNDA5MiAzMiAxNlJ+/SeVzBq6j8ObY3Zjuwq4cU3eH7sLOSZEghI=",
            "timestamp": "2017-02-25T21:07:50Z"
        }

# Group Decryption

## Decryption Collection [/decrypt]

### Decrypt [POST]
Decrypts plain text using AES-256 and base64 encoding

+ key (required, string) - The password used for decrypting
+ text (required, string) - The text to be decoded

+ Request (application/json)

        {
            "text": "Kl9LmR7dn0KxBTYzIXwOTyE2JAPgVUwHR8p4kjx0ZrEgICAgNDA5MiAzMiAxNlJ+/SeVzBq6j8ObY3Zjuwq4cU3eH7sLOSZEghI=",
            "key": "password"
        }

+ Response 200 (application/json)

        {
            "text": "Hello World!",
            "timestamp": "2017-02-25T21:07:50Z"
        }
