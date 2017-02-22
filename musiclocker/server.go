// Copyright (C) 2016-2017 Anonymous
// RWTXliExJCquO54R+qP94i4V+X8bQegE6L9EjhKIH23ePweJG8u7dqDK
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU Affero General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version. See COPYING for the full text of the License.

package musiclocker

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

const response = `<?xml version="1.0" encoding="UTF-8"?>
<bee:tracklist xmlns:bee="http://example.com/beehive/2016">
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
  <bee:track src="src" artist="artist" album="album" title="title" rating="5" />
</bee:tracklist>
`

func init() {
	markup, err := template.ParseFiles("templates/index.template.html")
	if err != nil {
		log.Panic("could not create template")
	}

	index := IndexHandler{markup: markup}
	http.HandleFunc("/api/query", query)
	http.Handle("/tunes", index)
}

func query(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(w, response)
}

type IndexHandler struct {
	markup *template.Template
}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	h.markup.Execute(w, nil)
}
