// server.go
// Copyright (C) 2016 Beehive Authors
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU Affero General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version.
//
// See COPYING for a full copy of the GNU Affero General Public License.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", "127.0.0.1:8080",
	"Port and address on which to bind")

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

func query(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		http.Error(w, "missing field", 422)
	}
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(w, response)
}

func main() {
	flag.Parse()
	http.HandleFunc("/_beehive/query", query)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
