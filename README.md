# beehive

beehive is a personal music locker server which runs on [Google App Engine][]
and uses other Google Cloud services.

It has some unique features:

  * Efficiently browse giant libraries, even with a poor internet connection,
    using the "omnibox" which generates a mix from hashtag-based query language:

        #sleep
        #pop #new
        #footwork #2016

  * Tags replace genres, playlists and likes, with an expanding vocabulary of
    smart tags like #new, artist tags, year tags and others.

  * Jump to individual songs or reorder the current mix as desired. Tags that
    end with "mix" preserve the ordering when searched.

  * Store all your music on Google Cloud for pennies per day (or even free if
    qualified for the free quotas!)

[Google App Engine]: https://cloud.google.com/appengine/
[Google Cloud Shell]: https://cloud.google.com/shell/


## Query language

Queries are specified as ASCII text, using only alphanumerics and spaces. They
have no context; they cannot refer to past queries or recent play history.


## Running

This software runs on the Google App Engine SDK for Go or equivalent
environment. One easy way to get access to an authenticated client is to run the
below commands on [Google Cloud Shell][]:

    git clone https://github.com/beemovie/beehive.git
    cd beehive
    goapp serve

(Alternatively, all source code can be obtained directly from a running instance
if the above repository is no longer available using [manifest.txt][] included
in each installation.)

Signatures with [signify][] are made on each commit for authorship and checksum
verification; these signatures can be verified after obtaining the public key
from a trusted source with:

    signify -C -p beehive.pub -x SHA256.sig

[Google Cloud Shell]: https://cloud.google.com/shell/
[manifest.txt]: static/manifest.txt
[signify]: http://man.openbsd.org/signify


## Development

This project is developed as an anonymous work. If you release patches or fork,
you are encouraged to explicitly assign copyright to the original author,
Anonymous, for easy inclusion into upstream.

It is released under the AGPLv3+ with the hope that users will always be free to
run it and tweak it.


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
