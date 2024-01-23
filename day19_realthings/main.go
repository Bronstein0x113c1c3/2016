package main

import (
	"io"
	"net/http"
	"strings"
)

// func main() {
// 	for _, p := range Products {
// 		Printfln("Product: %v, Category: %v, Price: $%.2f",
// 			p.Name, p.Category, p.Price)
// 	}
// }

type StringHandler struct {
	message string
}

// func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

// 	/*
// 		Request:
// 			- Context
// 			- URL
// 	*/

// 	//all of the field of the request struct
// 	// Printfln("Method: %v", request.Method)
// 	// Printfln("URL: %v", request.URL)
// 	// Printfln("HTTP Version: %v", request.Proto)
// 	// Printfln("Host: %v", request.Host)
// 	// for name, val := range request.Header {
// 	// 	Printfln("Header: %v, Value: %v", name, val)
// 	// }
// 	// Printfln("---")

// 	// if request.URL.Path == "/favicon.ico" {
// 	// 	Printfln("Request for icon detected - returning 404")
// 	// 	writer.WriteHeader(http.StatusNotFound)
// 	// 	return
// 	// }
// 	// Printfln("Request for %v", request.URL.Path)

// 	//or more sophiscated...

// 	Printfln("Request for %v", request.URL.Path)
// 	switch request.URL.Path {
// 	//if path of the url is favicon.ico
// 	case "/favicon.ico":
// 		http.NotFound(writer, request)
// 	//if we want to find the message?
// 	case "/message":
// 		io.WriteString(writer, sh.message)
// 	// for other cases..., redirect these request to "/message" with same writer and request,
// 	// with the temporary redirect status
// 	default:
// 		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
// 	}

// 	// io.WriteString(writer, sh.message)
// }

func main() {
	// err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
	//divide the whole into handlers.....
	/*
		the signature: w http.ResponseWriter, r *http.Request
		message -> stringhandlers....
		favicon.ico -> 404
		others -> redirect to "/message", with Temporary Redirect Status Code

	*/
	http.Handle("/message", StringHandler{message: "Bonjour!"})
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	//these created the serve mux.....

	go func() {
		err := http.ListenAndServeTLS(":5500", "cert/cert.cer", "cert/cert.pkey", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()
	err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))

	//with nil value, the default serve mux that we have created would be used.

	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

/*but, everything, for each route needs a fuction for each of them, for the ease
of programming, scalability, and maintainance */

/*so, another version....*/

func HTTPSRedirect(writer http.ResponseWriter,
	request *http.Request) {
	host := strings.Split(request.Host, ":")[0]
	target := "https://" + host + ":5500" + request.URL.Path
	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}
	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

//add the HTTPS redirect for the whole thangs....

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Printfln("Request for %v", r.URL.Path)
	io.WriteString(w, sh.message)
}
