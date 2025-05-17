package main

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/kevvurs/chatbox/html"	
)


var document js.Value;
var container js.Value;
var chatWindow js.Value;
var chat js.Value;
var textEntry js.Value;

func main() {
	fmt.Println("Loading chatbox v1.0")

	// create a floating icon to open the chat when idle
	document = js.Global().Get("document")
	container = html.CreateDiv(document, "fixed-bottom openchat collapse show")
	openChatButton := html.CreateButton(document, "btn btn-light btn-lg btn-circle p-2")
	icon := html.CreateItalic(document, "bi bi-chat-fill", "")

	html.Append(openChatButton, icon)
	html.Append(container, openChatButton)

	// create a floating window for chat
	chatWindow = html.CreateDiv(document, "fixed-bottom chatwindow collapse")

	menubar := html.CreateDiv(document, "d-flex justify-content-between " +
		"align-items-center bg-light p-1")
	chatTitle := html.CreateTitle(document, 5, "mb-0", "Chatbox")
	dismissButton := html.CreateButton(document, "btn bg-transparent" +
		"btn-sm border-0")
	dismissIcon := html.CreateItalic(document, "bi bi-x-lg", "")
	html.Append(menubar, chatTitle)
	html.Append(dismissButton, dismissIcon)
	html.Append(menubar, dismissButton)

	chat = html.CreateDiv(document, "d-flex flex-column bg-light p-1 " +
		"border border-start-0 border-end-0 overflow-auto chatstream")

	sendbar := html.CreateDiv(document, "input-group bg-light px-1 py-2")
	textEntry = html.CreateInput(document, "text", "form-control bg-light " +
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
	textEntry.Call("addEventListener", "keydown", js.FuncOf(sendChatOnEnter))
	sendChatButton.Call("addEventListener", "click", js.FuncOf(sendChat))
	openChatButton.Call("addEventListener", "click", js.FuncOf(openChat))
	dismissButton.Call("addEventListener", "click", js.FuncOf(dismissChat))

	document.Get("body").Call("appendChild", chatWindow)
	document.Get("body").Call("appendChild", container)
	fmt.Println("chatbox loaded")
	<-make(chan bool)
}

func sendChat(this js.Value, args []js.Value) any {
	message := textEntry.Get("value").String()
	if strings.TrimSpace(message) == "" {
		return ""
	}
	messageBubble := html.CreateParagraph(document, "bg-primary p-1 mw-75 align-self-end " +
		"flex-shrink-0 text-light rounded shadow-lg",	message)
	html.Append(chat, messageBubble)
	chat.Set("scrollTop", chat.Get("scrollHeight"))
	textEntry.Set("value", "")
	return ""
}

func sendChatOnEnter(this js.Value, args []js.Value) any {
	if len(args) > 0 && args[0].Get("key").String() != "Enter" {
		return ""
	}
	return sendChat(this, args)
}


func openChat(this js.Value, args []js.Value) any {
  openClassList := container.Get("classList")
	openClassList.Call("remove", "show")
	chatClassList := chatWindow.Get("classList")
	chatClassList.Call("add", "show")
	return ""
}


func dismissChat(this js.Value, args []js.Value) any {
	classList := chatWindow.Get("classList")
	classList.Call("remove", "show")
	openClassList := container.Get("classList")
	openClassList.Call("add", "show")
	return ""
}
