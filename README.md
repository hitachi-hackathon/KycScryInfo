# kyc
<img align="right" width="159px" src="http://arcosxblog.oss-cn-beijing.aliyuncs.com/hackathon/941553402259_.pic_hd.jpg">

This Project made by **"Babe I still love you"** for **hitachi-hackathon**

include:

* our software design
* demo
* how to run

## Design

### User:
* send user information data on IPFS,
* get zero proof + private key + digital certificate
* sign zero proof + digital certificate with private key to generate a digital signature
* send zero-proof + certificate + digital signature to exchange

### Authority:
* receive data from the user
* generate a digital certificate using its own private key + user public key
* generate zero-knowledge proof to user
* upload the data to IPFS

### Financial-Institute:
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

## Demo
* At first people send info to our service
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/8838860f-efd0-4c0e-b8f1-b74bf883de36.jpg)
* authority then verify the integrity of data
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/56f210b8-b350-4fa3-b89c-8fdb5184eadd.jpg)
* institute issue certificate for user
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/1bf8f014-4585-4d5e-a0a9-bfe69e9ff82c.jpg)
* user send to institute
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/e7e39059-f270-4376-829a-fb9ff0b0e112.jpg)
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/c813cf14-28f3-4e28-ad2b-65e293473281.jpg)
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/a2c62a26-a475-40b2-9cdf-30b0ebbe136f.jpg)
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/d4e0454c-3d1c-40e6-8f56-7e280f47b099.jpg)
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/bb779e07-20f4-4f09-986e-a6a8a47a1f23.jpg)
![](https://arcosxblog.oss-cn-beijing.aliyuncs.com/img/217cb9ca-e5f9-46bc-ad2e-9d77cbf85ba2.jpg)


## How to run
1. begin the backend at first
    1. go run main.go
2. begin the frontend
    1. npm install
    2. npm start
