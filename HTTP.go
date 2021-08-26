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

//			Client.Post(url string, requestBody []byte) ([]byte, error)
//
//			Client.PostStruct(url string, requestPayload interface{}, responseVessel interface{}) error
//

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

// Wrapper for http client with additional methods
type Client struct {
	client *http.Client
}

// Creates a new Client from an http.Client, with a request timeout in seconds.
// Changes any negative or zero timeout to 1 second.
// Passing nil client causes the function to use the default &http.Client{} (then add the timeout)
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

// Classic HTTP Get, returning response body (error if status code outside 200-299 range)
func (c *Client) Get(url string) ([]byte, error) {
	if c == nil || c.client == nil {
		return []byte{}, errors.New("wiz.Client.Get: nil client")
	}
	r, err := c.client.Get(url)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.Client.Get")
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.Client.Get")
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
		return []byte{}, errors.New("wiz.Client.Get: " + info)
	}
	return body, nil
}

// HTTP Get, but immediately unmarshals response body into a struct if 200-299 status.
// Note that json.Unmarshal requires the responseStruct be passed as a pointer to function properly
func (c *Client) GetStruct(url string, responseStruct interface{}) error {
	if c == nil || c.client == nil {
		return errors.New("wiz.Client.GetStruct: nil client")
	}
	response, err := c.Get(url)
	if err == nil {
		err = json.Unmarshal(response, responseStruct)
	}
	return errors.Wrap(err, "wiz.Client.GetStruct")
}

// Classic HTTP Post, returning response body (error if status code outside 200-299 range)
func (c *Client) Post(url string, requestBody []byte) ([]byte, error) {
	if c == nil || c.client == nil {
		return []byte{}, errors.New("wiz.Client.Post: nil client")
	}
	r, err := c.client.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.Client.Post")
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, errors.Wrap(err, "wiz.Client.Post")
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
		return []byte{}, errors.New("wiz.Client.Post: " + info)
	}
	return body, nil
}

// HTTP Post, but automatically marshals a struct as request body, and immediately unmarshals response body into a struct if 200-299 status.
// Note that json.Unmarshal requires the responseStruct be passed as a pointer to function properly
func (c *Client) PostStruct(url string, requestStruct interface{}, responseStruct interface{}) error {
	if c == nil || c.client == nil {
		return errors.New("wiz.Client.PostStruct: nil client")
	}
	body, err := json.Marshal(requestStruct)
	if err != nil {
		return errors.Wrap(err, "wiz.Client.PostStruct")
	}
	response, err := c.Post(url, body)
	if err == nil {
		err = json.Unmarshal(response, responseStruct)
	}
	return errors.Wrap(err, "wiz.Client.PostStruct")
}

//
//
//
//
//

// Splits a url at '/' characters
func SplitURL(url string) []string {
	u := strings.Split(url, "/")
	//Example:
	// If url was "/view/blocks/1234/"
	// Then path is now ["","view","blocks","1234",""]
	final := []string{}
	for _, item := range u {
		if item != "" {
			final = append(final, strings.ToLower(item))
		}
	}
	//Removed all empty strings from the array.
	return final
}

//Serve on a given address, and forward GET and POST requests to the separate handlers provided
func ServeSimple(ln net.Listener, getter func([]string) (int, []byte), poster func([]string, []byte) (int, []byte)) error {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Begin root handler function
		url := SplitURL(r.URL.Path)
		method := r.Method
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500) //Internal server error
			w.Write([]byte("wiz.ServeSimple root handler: " + err.Error()))
			return
		}
		status, response := 0, []byte("Initializing response")
		switch method {
		case "GET":
			status, response = getter(url)
		case "POST":
			status, response = poster(url, body)
		default:
			status, response = 405, []byte("wiz.ServeSimple root handler: Method "+method+" not allowed")
		}
		w.WriteHeader(status)
		w.Write(response)
		//End root handler function
	})
	return http.Serve(ln, http.HandlerFunc(handler))
}
