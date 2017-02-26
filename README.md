# beehive

beehive is a personal music locker service with some unique features.

  * Store all your music on Google Cloud for pennies per day (or even free if
    qualified for the free quotas!)

  * Browse giant libraries in browser, even with a poor internet connection.

  * Tags replace genres and likes, with trivial playlists based on tags.

  * Playlist generation engine based on simple text queries.

  * Todo lists for continuous improvement of tags and recommendations.

beehive is powered by several Google Cloud Platform services,
including [Google App Engine][] and [Google Cloud Storage][].

[Google App Engine]: https://cloud.google.com/appengine/
[Google Cloud Storage]: https://cloud.google.com/storage/


## Example queries

Queries are specified as ASCII text, using only alphanumerics and spaces. They
have no context; they cannot refer to past queries or recent play history.

    similar to always by erasure
    similar to aphex twin
    tagged #pop and loved
    loved in 2014
    released in early 2016 and loved
    tagged #rave and loved in 1997

(`loved` describes songs which were often played during a given time period, as
opposed to `released` which was when it was first publically available.)

beehive uses only a few heuristics to determine similar tracks:

  * Tracks released in the same time period.

  * Tracks loved in the same time windows.

  * Tracks which share tags.


## Running

This software runs on the Google App Engine SDK for Go or equivalent
environment.

The source code repository is on [GitHub](https://github.com/beemovie/beehive):

    git clone git@github.com/beemovie/beehive.git
    cd beehive
    goapp serve

Altenatively, all source code can be obtained directly from a running instance
if the above repository is no longer available or valid, using `wget`:

    mkdir beehive
    cd beehive
    wget -x -nH -i https://example.appspot.com/static/manifest.txt

Signatures with [signify](http://man.openbsd.org/signify) are made on each
commit for authorship and checksum verification; you can verify after obtaining
the public key from a trusted source.

    signify -C -p beehive.pub -x SHA256.sig


## Development

This project is developed as an anonymous work. If you release patches or fork,
you are encouraged to explicitly assign copyright to the original author,
Anonymous, for easy inclusion into upstream.

It is released under the AGPLv3+ with the hope that users will be free
to modify any future variant.


## Copying

Copyright (C) 2016-2017 Anonymous

RWTXliExJCquO54R+qP94i4V+X8bQegE6L9EjhKIH23ePweJG8u7dqDK

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU Affero General Public License as published by the Free
Software Foundation, either version 3 of the License, or (at your option) any
later version. See [COPYING](COPYING) for the full text of the License.

The Roboto font by Google is provided with this source code for zero external
dependencies. It is available under the Apache License, Version 2.0;
see [COPYING.APL](COPYING.APL) for a copy of that License.
