package html

import (
	"strconv"
	"syscall/js"
)

func CreateElement(dom js.Value, tag, class string) js.Value {
	div := dom.Call("createElement", tag)
	div.Set("className", class)
	return div
}

func CreateDiv(dom js.Value, class string) js.Value {
	return CreateElement(dom, "div", class)
}

func CreatePre(dom js.Value, class string) js.Value {
	return CreateElement(dom, "pre", class)
}

func CreateText(dom js.Value, tag, class, content string) js.Value {
	element := CreateElement(dom, tag, class)
	element.Set("innerText", content)
	return element
}

func CreateTitle(dom js.Value, size int, class, content string) js.Value {
	return CreateText(dom, "h" + strconv.Itoa(size), class, content)
}

func CreateParagraph(dom js.Value, class, content string) js.Value {
	return CreateText(dom, "p", class, content)
}

func CreateItalic(dom js.Value, class, content string) js.Value {
	return CreateText(dom, "i", class, content)
}

func CreateAnchor(dom js.Value, class, content, url string) js.Value {
	element := CreateText(dom, "a", class, content)
	element.Set("href", url)
	return element
}

func CreateButton(dom js.Value, class string) js.Value {
	element := CreateElement(dom, "button", class)
	element.Set("type", "button")
	return element
}

func CreateInput(dom js.Value, _type, class string) js.Value {
	element := CreateElement(dom, "input", class)
	element.Set("type", _type)
	return element
}

func CreateSpan(dom js.Value, class string) js.Value {
	element := CreateElement(dom, "span", class)
	return element
}

func CreateLine(dom js.Value, class string) js.Value {
	element := CreateElement(dom, "hr", class)
	return element
}

func Append(parent, child js.Value) {
	parent.Call("appendChild", child)
}
