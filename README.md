# kyc
### User:
* send user information data on IPFS,
* get zero proof + private key + digital certificate
* sign zero proof + digital certificate with private key to generate a digital signature
* send zero-proof + certificate + digital signature to exchange

### Validator:
* receive data from the user
* generate a digital certificate using its own private key + user public key
* generate zero-knowledge proof to user
* upload the data to IPFS

### Exchange:
* save zero-knowledge proof,
* send all to the validator for verification.
* receive true or false

### Front:
* save user private key, do the digital signature
* send data to server
* receive proof + digital certificate

### back:
* upload the data on chain
* get the meta id for IPFS data
* generate digital certificate, do the zero knowledge proof.
* send back the data to the user.
