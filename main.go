package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the function
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be convreted to an integer, or the value is less than 1, we return a 404
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	// If it's not, use the w.WriteHeader() method to send a 405 status code and
	// the w.Write() method to write a "Method Not Allowed" response body. We
	// then return from the function so that the subsequent code is not executed.
	if r.Method != "POST" {
		// Use the `Header().Set()` method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second paramenter is the header value.

		// NOTE: Changing the header map after a call to `w.WriteHeader()` or
		// `w.Write()` will have no effect on the response headers that the user receives.
		// You need to make sure that your header map contains all the headers you want
		// BEFORE you call these methods.
		w.Header().Set("Allow", "POST")

		// It's only possible to call w.WriteHeader() ONCE per response.
		// After the status code has been written it can't be changed.
		// If you try to call w.WriteHeader() a second time Go will log a warning message.

		// If you don't call w.WriteHeader() explicitly, then the first call to
		// w.Write() will autmatically send a `200 OK` status code to the user.

		// So if you want to send a non-200 status code, you must call w.WriteHeader()
		// BEFORE any call to w.Write().
		/**
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		*/

		// Alternatively, if you want to send a non-200 status code and a plain-text
		// response body then it's a good opportunity to use the `http.Error()` shortcut.

		// Use the `http.Error()` function to send a 405 status code and "Method Not Allowed"
		// string as the response body.
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Register the two new handler functions and corresponding URL patterns with
	// the servemux, in exactly the same way that we did before.
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
