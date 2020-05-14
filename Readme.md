### Description
URL checker allows to validate HTTP status and time taken to complete request for give url. Output is handled by net/http as HTTP server

---

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

---

### Configuration 
The configuration is stored in a single `config.yaml` file
```
server: 
  port: 8091
client: 
  skipSSL: true
  timeout: 30 # http.client timeout
  period: 5 # How often to scan a URL. In seconds.
urls:
  - https://google.com
  - http://facebook.com
  - https://amazon.com
  - http://ebay.com
  - http://fake.local # This is a fake one. Handy for testing
```

---

### Roadmap

##### 0.1 

* *[Completed]* - use Go routines to periodically check URL
* *[Completed]* - Add http.client timeout 
* *[Completed]* - Add dedicated log library
* *[Completed]* - Add support for environment variables for port, log level, period
* *[Todo]* - add /check endpoint to allow dynamically check URL. Eg. /check?url=http:\/\/google.com



##### 0.2 

* add Dockerfile 
* *[Todo]* - allow dynmically manipulate list of URL's with API. Eg. /api/add/{url} /api/remove/{url}


##### 0.3

* *[Todo]* - add tests - https://github.com/stretchr/testify
* *[Todo]* - command line attributes for port, log level 
