### Description
URL checker allows to validate HTTP status and time taken to complete request for give url. Output is handled by net/http as HTTP server

### Installation 
```
git clone https://github.com/arturgut/urlchecker
cd urlchecker
git build *.go
```

Expected output in Prometheus format:
```
curl http://localhost:8090/metrics
# HELP Version: 0.1 Alpha - https://github.com/arturgut/urlchecker
# HELP url_checker. Label: HTTP Response code. Value: Request duration in Ms
url_checker{ url='http://ebay.com', http_status_code=200 } 529
url_checker{ url='http://google.com', http_status_code=200 } 104
url_checker{ url='http://facebook.com', http_status_code=200 } 520
url_checker{ url='http://amazon.com', http_status_code=200 } 698
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

### Roadmap

######0.1 

* use Go routines to periodically check URL
* implement Prometheus Go library to display metrics 
* use log library to handle logs

######0.2 
* add /check endpoint to allow dynamically check URL. Eg. /check?url=http:\/\/google.com
* add Dockerfile 
* add CI/CD pipeline

######0.3
* * allow dynmically manipulate list of URL's with API. Eg. /api/add/{url} /api/remove/{url}

