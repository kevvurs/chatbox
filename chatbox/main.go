package main

import (
	"fmt"
	"syscall/js"

	"github.com/kevvurs/chatbox/html"	
)

func main() {
	fmt.Println("Loading chatbox v1.0")

	// create a floating icon to open the chat when idle
	document := js.Global().Get("document")
	container := html.CreateDiv(document, "fixed-bottom")
	openChatButton := html.CreateButton(document, "btn btn-light btn-lg btn-circle")
	icon := html.CreateItalic(document, "glyphicon glyphicon-comment", "")

	html.Append(openChatButton, icon)
	html.Append(container, openChatButton)

	// create a floating window for chat
	chatWindow := html.CreateDiv(document, "fixed-bottom chatwindow")

	menubar := html.CreateDiv(document, "d-flex justify-content-between " +
		"align-items-center bg-light p-1")
	chatTitle := html.CreateTitle(document, 5, "mb-0", "Chatbox")
	dismissButton := html.CreateButton(document, "btn bg-transparent" +
		"btn-sm border-0")
	dismissIcon := html.CreateItalic(document, "bi bi-x-lg", "")
	html.Append(menubar, chatTitle)
	html.Append(dismissButton, dismissIcon)
	html.Append(menubar, dismissButton)

	chat := html.CreateDiv(document, "bg-light p-1 border " +
		"border-start-0 border-end-0 chatstream")

	sendbar := html.CreateDiv(document, "input-group bg-light px-1 py-2")
	textEntry := html.CreateInput(document, "text", "form-control bg-light " +
		"border-1 border-secondary focus-ring focus-ring-secondary")
	textEntry.Set("placeholder", "Aa")
	sendChatButton := html.CreateButton(document, "btn btn-outline-secondary")
	sendIcon := html.CreateItalic(document, "bi bi-send", "")
	html.Append(sendbar, textEntry)
	html.Append(sendChatButton, sendIcon)
	html.Append(sendbar, sendChatButton)
	
	html.Append(chatWindow, menubar)
	html.Append(chatWindow, chat)
	html.Append(chatWindow, sendbar)

	// initialize click handlers
	
	document.Get("body").Call("appendChild", chatWindow)
	fmt.Println("chatbox loaded")
}
