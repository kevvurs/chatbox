
package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	fmt.Println("Running WASM script...")

	// create the elements
	document := js.Global().Get("document")
	background := createDiv(document, "bg-secondary d-flex " +
	  "align-items-center justify-content-center vh-100")
	foreground := createDiv(document, "card bg-dark text-light p-4 " +
		"mx-4 my-auto rounded shadow-lg w3-card-4")
	title := createTitle(document, 1, "",
		`Welcome to WebAssembly UI!`)
	description := createParagraph(document, "",
	"This page was created using Go WASM. All the DOM elements are" +
	" generated with WASM and assembled into the DOM at runtime. A" +
	" driver script runs a .wasm binary file in the <head> tag to " +
	" execute the Go code responsible for this page.")
	code := createPre(document, "p-4 text-dark bg-light")
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
	link := createAnchor(document, "my-auto link-danger", "Learn more at seedshare.io",
		"https://seedshare.io/blog/wasm")

	// assemble the structure
	appendToDiv(foreground, title)
	appendToDiv(foreground, description)
	appendToDiv(foreground, code)
	appendToDiv(foreground, link)
	appendToDiv(background, foreground)

	// manipulate the dom
	document.Get("body").Call("appendChild", background)
	fmt.Println("Web = Assembled!")
}

func createElement(dom js.Value, tag, class string) js.Value {
	div := dom.Call("createElement", tag)
	div.Set("className", class)
	return div
}

func createDiv(dom js.Value, class string) js.Value {
	return createElement(dom, "div", class)
}

func createPre(dom js.Value, class string) js.Value {
	return createElement(dom, "pre", class)
}

func createText(dom js.Value, tag, class, content string) js.Value {
	element := createElement(dom, tag, class)
	element.Set("innerText", content)
	return element
}

func createTitle(dom js.Value, size int, class, content string) js.Value {
	return createText(dom, "h" + strconv.Itoa(size), class, content)
}

func createParagraph(dom js.Value, class, content string) js.Value {
	return createText(dom, "p", class, content)
}

func createAnchor(dom js.Value, class, content, url string) js.Value {
	element := createText(dom, "a", class, content)
	element.Set("href", url)
	return element
}

func appendToDiv(div, child js.Value) {
	div.Call("appendChild", child)
}
