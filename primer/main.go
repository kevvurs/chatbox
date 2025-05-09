package main

import (
	"fmt"
	"syscall/js"

	"github.com/kevvurs/chatbox/html"
)

func main() {
	fmt.Println("Running WASM script...")

	// create the elements
	document := js.Global().Get("document")
	background := html.CreateDiv(document, "bg-secondary d-flex " +
	  "align-items-center justify-content-center vh-100")
	foreground := html.CreateDiv(document, "card bg-dark text-light p-4 " +
		"mx-4 my-auto rounded shadow-lg w3-card-4")
	title := html.CreateTitle(document, 1, "",
		`Welcome to WebAssembly UI!`)
	description := html.CreateParagraph(document, "",
	"This page was created using Go WASM. All the DOM elements are" +
	" generated with WASM and assembled into the DOM at runtime. A" +
	" driver script runs a .wasm binary file in the <head> tag to " +
	" execute the Go code responsible for this page.")
	code := html.CreatePre(document, "p-4 text-dark bg-light")
	code.Set("innerText",
		`document := js.Global().Get("document")
background := createDiv(document,
	"bg-dark d-flex align-items-center justify-content-center vh-100")
foreground := createDiv(document,
	"card bg-secondary text-light p-4 rounded shadow-lg, mx-4")
title := createTitle(document, 1, "",
	"Welcome to WebAssembly UI!"")
description := createText(document, "",
	"This page was created using Go WASM. All the DOM elements are
	generated with WASM and assembled into the DOM at runtime. A driver
	script runs a .wasm binary file in the <head> tag to execute the Go
	code responsible for this page."")`)
	link := html.CreateAnchor(document, "my-auto link-danger", "Learn more at seedshare.io",
		"https://seedshare.io/blog/wasm")

	// assemble the structure
	html.Append(foreground, title)
	html.Append(foreground, description)
	html.Append(foreground, code)
	html.Append(foreground, link)
	html.Append(background, foreground)

	// manipulate the dom
	document.Get("body").Call("appendChild", background)
	fmt.Println("Web = Assembled!")
}
