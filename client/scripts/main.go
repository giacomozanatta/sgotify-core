package main

import (
	"bytes"
	"encoding/json"
	"honnef.co/go/js/dom"
	"net/http"
)

import "Sgotify/sgotify"

func addToQueueEvent(element dom.Element) {
	element.AddEventListener("click", false, func(event dom.Event) {
		go func() {
			body, err := json.Marshal(sgotify.Song{
				Title:  element.GetAttribute("data-title"),
				Author: element.GetAttribute("data-author"),
				Id:     element.GetAttribute("data-id"),
			})
			if err != nil {
				println(err)
			}

			resp, err := http.Post("/queue/add", "application/json", bytes.NewReader(body))

			if err != nil {
				// handle error
				println(err)
				return
			}
			defer resp.Body.Close()
			element.(*dom.HTMLAnchorElement).Style().SetProperty("display", "none", "")
		}()
	})
}

func main() {
	d := dom.GetWindow().Document()

	addToQueue := d.GetElementsByClassName("add-to-queue")
	for _, a := range addToQueue {
		go addToQueueEvent(a)
	}
}
