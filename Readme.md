### Description
URL checker allows to validate HTTP status and time taken to complete request for give url. Output is handled by net/http as HTTP server

### Installation 
```
git clone https://github.com/arturgut/urlchecker
cd urlchecker
git build *.go
```

### Configuration 
The configuration is stored in a single `config.yaml` file
```
server: 
  port: 9940 # <- Port (Currently not implemented )
client: 
  skipSSL: true # (Currently not implemented )
  timeout: 30
urls: # <- Array of URL to check
  - http://google.com
  - http://facebook.com
  - http://amazon.com
  - http://ebay.com
```