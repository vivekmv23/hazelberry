# Hazelberry

Command line utility for testing REST APIs and organizing them leveraging [collections](https://learning.postman.com/docs/collections/collections-overview/).

## Initial scope
1. Capable to adopting exisitng collection exports
2. Abilty to create, read, update and delete collections
3. List and invoke the APIs from a collection in an intuitive way without leaving terminal


## Unsupported 
### 1. Types
1. Header of type string is not supported
2. Http methods: ["COPY","LINK","UNLINK","PURGE","LOCK","UNLOCK","PROPFIND","VIEW"]
3. Body types: ["urlencoded","formdata","file","graphql"]
4. Request Certificates
5. Request Proxies
6. Request Cookies
7. Scripts
8. Variables
9. Descriptions
10. Responses
11. Events
12. Request of type string
13. Item Group