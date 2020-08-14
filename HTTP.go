package wiz

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Exposed types and methods:
//			Client
//				Wrapper for http client with this file's methods
//			Client.Get(url string) ([]byte, error)
//				Classic HTTP Get, returning response body (error if status code outside 200-299 range)
//			Client.GetStruct(url string, responseVessel interface{}) error
//				Get, but immediately unmarshals response body into a struct if 200-299 status.
//				Note that json.Unmarshal requires the responsevessel be passed as a pointer to function properly
//			Client.Post(url string, requestBody []byte) ([]byte, error)
//				Classic HTTP Post, returning response body (error if status code outside 200-299 range)
//			Client.PostStruct(url string, requestPayload interface{}, responseVessel interface{}) error
//				Post, but automatically marshals a struct as request body, and immediately unmarshals response body into a struct if 200-299 status.
//				Note that json.Unmarshal requires the responsevessel be passed as a pointer to function properly

//		Exposed functions:
//			NewClient(c *http.Client, timeout int) Client
//				Create a new Client from an http.Client, with a request timeout in seconds.
//				Changes any negative or zero timeout to 1 second.
//				Passing nil client causes the function to use the default &http.Client{} (then add the timeout)
//			ServeSimple(ln net.Listener, getter func([]string) (int, []byte), poster func([]string, []byte) (int, []byte)) error
//				Using a provided http.Listener, serve and forward all requests to a getHandler and postHandler function,
//				each of which only receives request url (split into a slice by '/') and/or body, and returns a
//				status code and response body (no headers and other request/response features supported)
//				Note: Block until server fails (or forever) - start serve in a goroutine if program should continue while serving

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

//		Simple client example:
//		c := NewClient(nil, 10)
//		resp, err := c.Get("http://example.com/whatever")

//		Advanced client example (uses tls, post, structs)
//		type SomeStruct struct {
//			Field1 string
//			Field2 uint64
//		}
//		c := NewClient(existingTLSclient, 10)
//		s1 := SomeStruct{Field1:"A",Field2:99} //Will send this in POST request
//		s2 := SomeStruct{} //Empty - will receive response into this
//		err := c.PostStruct("https://api.example.com/methodfoo",s1,&s2)
//		//Now if err == nil, s2 contains the response from the api, as type SomeStruct

//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*
//		*	*	*	*	*	*	*	*	*	*	*	*	*	*	*	*

type Client struct {
	client *http.Client
}

func NewClient(c *http.Client, timeout int) Client {
	if c == nil {
		c = &http.Client{}
	}
	if timeout <= 0 {
		timeout = 1
	}
	c.Timeout = time.Second * time.Duration(timeout)
	return Client{client: c}
}

func (c *Client) Get(url string) ([]byte, error) {
	if c == nil || c.client == nil {
		return []byte{}, errors.New("Client.Get: nil client")
	}
	r, err := c.client.Get(url)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Client.Get")
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Client.Get")
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		//Note: response body might contain useful information, but only printable
		// strings should be added to the error message
		info := "Status code " + strconv.Itoa(r.StatusCode)
		//Check body before adding it to info
		printable := true
		for _, c := range body {
			if c > unicode.MaxASCII || !unicode.IsGraphic(rune(c)) {
				printable = false
			}
		}
		if printable {
			info += ", Response body: " + string(body)
		} else {
			info += ", Response body contained unprintable characters."
		}
		return []byte{}, errors.New("Client.Get: " + info)
	}
	return body, nil
}

func (c *Client) GetStruct(url string, responseVessel interface{}) error {
	if c == nil || c.client == nil {
		return errors.New("Client.GetStruct: nil client")
	}
	response, err := c.Get(url)
	if err == nil {
		err = json.Unmarshal(response, responseVessel)
	}
	return errors.Wrap(err, "Client.GetStruct")
}

func (c *Client) Post(url string, requestBody []byte) ([]byte, error) {
	if c == nil || c.client == nil {
		return []byte{}, errors.New("Client.Post: nil client")
	}
	r, err := c.client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return []byte{}, errors.Wrap(err, "Client.Post")
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Client.Post")
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		//Note: response body might contain useful information, but only printable
		// strings should be added to the error message
		info := "Status code " + strconv.Itoa(r.StatusCode)
		//Check body before adding it to info
		printable := true
		for _, c := range body {
			if c > unicode.MaxASCII || !unicode.IsGraphic(rune(c)) {
				printable = false
			}
		}
		if printable {
			info += ", Response body: " + string(body)
		} else {
			info += ", Response body contained unprintable characters."
		}
		return []byte{}, errors.New("Client.Post: " + info)
	}
	return body, nil
}

func (c *Client) PostStruct(url string, requestPayload interface{}, responseVessel interface{}) error {
	if c == nil || c.client == nil {
		return errors.New("Client.PostStruct: nil client")
	}
	body, err := json.Marshal(requestPayload)
	if err != nil {
		return errors.Wrap(err, "Client.PostStruct")
	}
	response, err := c.Post(url, body)
	if err == nil {
		err = json.Unmarshal(response, responseVessel)
	}
	return errors.Wrap(err, "Client.PostStruct")
}

//
//
//
//
//

//Serve on a given address, and forward GET and POST requests to the separate handlers provided
func ServeSimple(ln net.Listener, getter func([]string) (int, []byte), poster func([]string, []byte) (int, []byte)) error {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Begin root handler function
		url := strings.Split(r.URL.Path, "/")
		method := r.Method
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500) //Internal server error
			w.Write([]byte("ServeSimple root handler: " + err.Error()))
			return
		}
		status, response := 0, []byte("Initializing response")
		switch method {
		case "GET":
			status, response = getter(url)
		case "POST":
			status, response = poster(url, body)
		default:
			status, response = 405, []byte("ServeSimple root handler: Method "+method+" not allowed")
		}
		w.WriteHeader(status)
		w.Write(response)
		//End root handler function
	})
	return http.Serve(ln, http.HandlerFunc(handler))
}
