# beehive

beehive is a personal music locker service with some unique features.

* Browse million-song libraries even with a poor internet connection.
* Playlist engine based based on simple text or voice queries.
* Smaller footprint by separating metadata from audio data.
* Todo lists for continuous improvement of tags and recommendations.

beehive is powered by several Google Cloud Platform technologies,
including [Google Cloud Storage](https://cloud.google.com/storage/) and
the [Google Cloud Speech API](https://cloud.google.com/speech/).

## Example queries

Queries are specified as case-insensitive ASCII text. They do not have
context (cannot refer to history of tracks already played).

```
similar to ty segall

similar to always by erasure

liked this year

liked and released this year

liked and released in 1990 to 1995 and tagged rave
```

## Development

beehive is developed as an anonymous work. This means we do not typically accept
pull requests, as any other author could later choose to identify themselves or
be identified.

It is released under copyleft terms with the hope that any user will be able to
change any future variant of the software as they see fit.

## Copying

Copyright (C) 2016 Beehive Authors

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option) any
later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along
with this program. If not, see <http://www.gnu.org/licenses/>.
