Ravelin Code Test
=================

## Submission Remarks

[![Go Report Card](https://goreportcard.com/badge/github.com/mindworker/code-test)](https://goreportcard.com/report/github.com/mindworker/code-test) [![Build Status](https://travis-ci.org/mindworker/code-test.svg?branch=master)](https://travis-ci.org/mindworker/code-test)

### General Notes

I tried to provide an 'overkill' solution in order to show off my
implementation skills. Note that I am usually more pragmatic and less
verbose in my code. In this submission I try to show that I am
familiar with the following:

 - Web backend development practices
 - Go concurrency patterns
 - Unit testing
 - Continuous Integration

### Other special remarks

*Travis CI* is set up for Go1.4, Go1.7 and JS environment with the
following:

 * Vetting and Linting for Go and JS
 * Go unit testing
 * Go **race condition test**

A *Middleware* is created to conveniently intercept all incoming
requests and is used for restricting HTTP methods and optional
CORS. See `middleware.go`

Use *environment variables* to change PORT setting or enable CORS

```
$ RAV_PORT=5000 ./code-test
2017/03/21 15:19:46 Listening on port 5000

$ RAV_USE_CORS=true ./code-test
2017/03/21 15:19:57 Using CORS
2017/03/21 15:19:57 Listening on port 8080
```

A SessionManager is created that keeps track of sessions has a
**build-in background routine** that removes sessions form memory
which are older than 60 minutes. Otherwise, server would inevitably
run out of memory. See `sessionmanager.go`.

Unit tests are created for the session manager as this part where it
is the easiest to make mistakes. See `sessionmanager_test.go`.

## Summary
We need an HTTP server that will accept any POST request (JSON) from muliple clients' websites. Each request forms part of a struct (for that particular visitor) that will be printed to the terminal when the struct is fully complete. 

For the JS part of the test please feel free to use any libraries that may help you **but please only use the Go standard library for the backend**.

## Frontend (JS)
Insert JavaScript into the index.html (supplied) that captures and posts data every time one of the below events happens; this means you will be posting multiple times per visitor. Assume only one resize occurs.

  - if the screen resizes, the before and after dimensions
  - copy & paste (for each field)
  - time taken from the 1st character typed to clicking the submit button

### Example JSON Requests
```
{
  "eventType": "copyAndPaste",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "pasted": true,
  "formId": "inputCardNumber"
}

{
  "eventType": "timeTaken",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "time": 72, // seconds
}

...

```

## Backend (Go)

The Backend must to:

1. Create a Server
2. Accept POST requests in JSON format similar to those specified above
3. Map the JSON requests to relevant sections of the data struct (specified below)
4. Print the struct for each stage of its construction
5. Also print the struct when it is complete (i.e. when the form submit button has been clicked)

### Go Struct
```
type Data struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int // Seconds
}

type Dimension struct {
	Width  string
	Height string
}
```




