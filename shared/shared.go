package shared

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
	return createElement(dom, "div", class)
}

func CreatePre(dom js.Value, class string) js.Value {
	return createElement(dom, "pre", class)
}

func CreateText(dom js.Value, tag, class, content string) js.Value {
	element := createElement(dom, tag, class)
	element.Set("innerText", content)
	return element
}

func CreateTitle(dom js.Value, size int, class, content string) js.Value {
	return createText(dom, "h" + strconv.Itoa(size), class, content)
}

func CreateParagraph(dom js.Value, class, content string) js.Value {
	return createText(dom, "p", class, content)
}

func CreateAnchor(dom js.Value, class, content, url string) js.Value {
	element := createText(dom, "a", class, content)
	element.Set("href", url)
	return element
}

func AppendToDiv(div, child js.Value) {
	div.Call("appendChild", child)
}
