package main

import (
	"os"
	"text/template"
	"fmt"
    "net/http"
	"bytes"
    "github.com/gin-gonic/gin"
)

type ContentfulResponse struct {
	Content string
}

func ReturnContentfulPage(c *gin.Context, skeleton_path string, fragment_path string) {
	// allocate a buffer of bytes to hold the data for our finished template
	buf := new(bytes.Buffer)
	// grab the skeleton of our page, and check errors
	skel_html, err := os.ReadFile(skeleton_path)
	if err != nil {
		fmt.Print(err)
	}
	// grab the fragment that would get handed to HTMX to add to our skeleton
	fragment_html, err := os.ReadFile(fragment_path)
	if err != nil {
		fmt.Print(err)
	}
	// create a new text template named "tmpl" and fill it with the skeleton, which
	// will have the fragment inserted into it when we execute the template.
	tmpl, err := template.New("tmpl").Parse(string(skel_html))
	if err != nil {
		panic(err)
	}
	// add the fragment html to the ContentfulResponse struct
	resp := ContentfulResponse{string(fragment_html)}
	// Execute the template, writing into buf, and using resp as the struct we pull
	// .Content from. We set err equal to this cause this function returns an error
	// if something goes wrong.
	err = tmpl.Execute(buf, resp)
	if err != nil {
		panic(err)
	}

	// Return data using buf, turning the buf into a byte array using buf.Bytes()
	c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())

}

func ReturnFragment(c *gin.Context, fragment_path string) {
	fragment_html, err := os.ReadFile(fragment_path)
	if err != nil {
		fmt.Print(err)
	}

	// return the fragment, converting the fragment string into a byte array
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fragment_html))
}

func main() {
    // Create a default gin router
    router := gin.Default()

    // Define a route for the root path
    // In the case of this tutorial, we're considering content-1 to be the
    // home page
    router.GET("/", func(c *gin.Context) {
		// get the header and store it in hx_header
		hx_header := c.Request.Header.Get("HX-Request")
		// Set the Vary key in the header to HX-Request so the browser will
		// store separate caches for contentful pages and fragments
		c.Header("Vary", "HX-Request")
		if(hx_header == "true") {
			ReturnFragment(c, "./content-1.html")
		} else {
			ReturnContentfulPage(c, "./index.html", "./content-1.html")
		}
    })
    router.GET("/content-2", func(c *gin.Context) {
		// get the header and store it in hx_header
		hx_header := c.Request.Header.Get("HX-Request")
		// Set the Vary key in the header to HX-Request so the browser will
		// store separate caches for contentful pages and fragments
		c.Header("Vary", "HX-Request")
		if(hx_header == "true") {
			ReturnFragment(c, "./content-2.html")
		} else {
			ReturnContentfulPage(c, "./index.html", "./content-2.html")
		}
    })
    router.GET("/content-3", func(c *gin.Context) {
		// get the header and store it in hx_header
		hx_header := c.Request.Header.Get("HX-Request")
		// Set the Vary key in the header to HX-Request so the browser will
		// store separate caches for contentful pages and fragments
		c.Header("Vary", "HX-Request")
		if(hx_header == "true") {
			ReturnFragment(c, "./content-3.html")
		} else {
			ReturnContentfulPage(c, "./index.html", "./content-3.html")
		}
    })

    // Start the server on port 9001
    router.Run(":9001")
}
