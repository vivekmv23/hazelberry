# Hazelberry

Command line utility for testing REST APIs and organizing them leveraging [collections](https://learning.postman.com/docs/collections/collections-overview/).

## Initial scope
1. Capable to adopting exisitng collection exports
2. Abilty to create, read, update and delete collections
3. List and invoke the APIs from a collection

## Unsupported 
### 1. Types
- Http methods: ["COPY","LINK","UNLINK","PURGE","LOCK","UNLOCK","PROPFIND","VIEW"]
- Body types: ["urlencoded","formdata","file","graphql"]
- Auth types: [ "apikey", "awsv4", "bearer", "digest", "edgegrid", "hawk", "ntlm", "oauth1", "oauth2"]
- Request Certificates
- Request Proxies
- Request Cookies
- Header of type string is not supported
- Scripts
- Variables
- Descriptions
- Responses
- Events
- Request of type string
- Item Group