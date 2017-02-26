// Copyright (C) 2016-2017 Anonymous
// RWTXliExJCquO54R+qP94i4V+X8bQegE6L9EjhKIH23ePweJG8u7dqDK
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU Affero General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version. See COPYING for the full text of the License.

package musiclocker

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	c := newState()
	http.HandleFunc("/v1/tracks/", c.tracks)
	http.HandleFunc("/tunes", c.index)
}

type State struct {
	Markup    *template.Template
	License   template.HTML
	TrackList TrackList
}

type TrackList struct {
	XMLName xml.Name `xml:"tracklist"`
	Tracks  []Track  `xml:"track"`
}

type Track struct {
	XMLName xml.Name `xml:"track"`
	Artist  string   `xml:"artist,attr"`
	Album   string   `xml:"album,attr"`
	Title   string   `xml:"title,attr"`
	Year    string   `xml:"year,attr"`
	Tags    []Tag    `xml:"tag"`
}

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Text    string   `xml:"text,attr"`
}

func newState() *State {
	markup, err := template.ParseFiles("templates/index.template.html")
	if err != nil {
		log.Fatal(err)
	}

	l, err := ioutil.ReadFile("templates/license.template.html")
	if err != nil {
		log.Fatal(err)
	}
	license := template.HTML(string(l))

	t, err := ioutil.ReadFile("testdata/tracklist.xml")
	if err != nil {
		log.Fatal(err)
	}

	var tracks TrackList
	err = xml.Unmarshal(t, &tracks)
	if err != nil {
		log.Fatal(err)
	}

	return &State{markup, license, tracks}
}

func (c *State) tracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(w, xml.Header)

	out, err := xml.Marshal(c.TrackList)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(out)
}

func (c *State) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := c.Markup.Execute(w, c); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
	}
}
