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
- Variable value types: ["boolean","any","number"]
- Request Certificates
- Request Proxies
- Request Cookies
- Header of type string is not supported
- Scripts
- Descriptions
- Responses
- Events
- Request of type string
- Item Group

## Command Guidelines
1. There are hazelberry commands and REPL commands
2. Each command MAY support flags
3. Supported formats of flags SHOULD be : 
    ```sh
    --flag value
    -flag value
    --flag=value
    -f value
    -fvalue
    ```
4. All comands SHOULD support help
5. Help SHOULD print command, usage, and flags with their defaults

## Commands

```sh
## Default behaviour is to load a "default" collection
## User MAY use the env variable HZ_DEFAULT_COLLECTION_PATH to change this default
## default HZ_DEFAULT_COLLECTION_PATH = ~/hazelberry/default-collection.json
~ $ hazelberry
default>
```

### `load`
```sh
## Loads a collection from specified path
~ $ hazelberry load path/to/my-collection.json
my-collection>
```

### `exit`
```sh
## Exits out of the REPL
~ $ hazelberry
default> exit
~ $
```

### `ls`
```sh
## Lists all the items in a collection
~ $ hazelberry
default> ls
Selection   Item                Method      Url
1           my-get-request      GET         https://my-host.com/some/resources/id
default>
```
### `invoke`
```sh
## Invokes the selected item by selection id or name
~ $ hazelberry
default> invoke 1
HTTP response 200:OK
{
    "ack": true
}
default> invoke my-get-request
HTTP response 200:OK
{
    "ack": true
}
```
