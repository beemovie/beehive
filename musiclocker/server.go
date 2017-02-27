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
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"html/template"
	"io/ioutil"
	"net/http"
)

func init() {
	c := newState()
	http.HandleFunc("/tunes", c.index)
	http.HandleFunc("/v1/tracks/", c.tracks)
	http.HandleFunc("/v1/upload/", c.upload)
}

func (c *State) index(w http.ResponseWriter, r *http.Request) {
	c.initRequest(w, r)

	w.Header().Set("Content-Type", "text/html")
	if err := c.Markup.Execute(w, c); err != nil {
		c.error(err)
	}
}

func (c *State) tracks(w http.ResponseWriter, r *http.Request) {
	c.initRequest(w, r)

	out, err := xml.Marshal(c.TrackList)
	if err != nil {
		c.error(err)
		return
	}

	w.Header().Set("Content-Type", "text/xml")
	fmt.Fprintf(w, xml.Header)

	w.Write(out)
}

func (c *State) upload(w http.ResponseWriter, r *http.Request) {
	c.initRequest(w, r)

	const maxMem = 1024 * 1024 * 32
	err := r.ParseMultipartForm(maxMem)
	if err != nil {
		c.error(err)
		return
	}
}

type State struct {
	Markup     *template.Template
	License    template.HTML
	TrackList  TrackList
	GAEContext context.Context
	Request    *http.Request
	Writer     http.ResponseWriter
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
		panic(err)
	}

	l, err := ioutil.ReadFile("templates/license.template.html")
	if err != nil {
		panic(err)
	}
	license := template.HTML(string(l))

	t, err := ioutil.ReadFile("testdata/tracklist.xml")
	if err != nil {
		panic(err)
	}

	var tracks TrackList
	err = xml.Unmarshal(t, &tracks)
	if err != nil {
		panic(err)
	}

	return &State{
		Markup:    markup,
		License:   license,
		TrackList: tracks,
	}
}

func (c *State) initRequest(w http.ResponseWriter, r *http.Request) {
	c.GAEContext = appengine.NewContext(r)
	c.Request = r
	c.Writer = w
}

func (c *State) error(err error) {
	log.Errorf(c.GAEContext, "%v", err)
	http.Error(c.Writer, err.Error(), 500)
	return
}
