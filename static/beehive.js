// Copyright (C) 2016-2017 Anonymous
// RWTXliExJCquO54R+qP94i4V+X8bQegE6L9EjhKIH23ePweJG8u7dqDK
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU Affero General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version. See COPYING for the full text of the License.

'use strict';

const query = document.getElementById('query');
const playlist = document.getElementById('playlist');

let timer;

function handleResponse() {
  if (this.status != 200) return;
  const tracks = this.responseXML.documentElement.children;

  while (playlist.firstChild) {
    playlist.removeChild(playlist.firstChild);
  }

  for (let i = 0; i < tracks.length; i++) {
    const track = tracks[i];

    const artist = document.createElement('p');
    artist.innerHTML = track.getAttribute('artist');
    const album = document.createElement('p');
    album.innerHTML = track.getAttribute('album');
    const title = document.createElement('p');
    title.innerHTML = track.getAttribute('title');

    const trackElement = document.createElement('li');
    trackElement.appendChild(title);
    trackElement.appendChild(artist);
    trackElement.appendChild(album);
    playlist.appendChild(trackElement);
  }
}

function handleTimerExpiration() {
  if (!query.value) return;
  timer = undefined;
  const xhr = new XMLHttpRequest();
  const path = '/api/query?q=' + encodeURIComponent(query.value);
  xhr.addEventListener('load', handleResponse);
  xhr.open('GET', path, true);
  xhr.responseType = 'document';
  xhr.send();
}

function handleQueryChange() {
  if (timer) clearTimeout(timer);
  timer = setTimeout(handleTimerExpiration, 450);
}

query.addEventListener('input', handleQueryChange);
window.addEventListener('load', function() {
  query.value = 'liked this year';
  handleQueryChange();
});
