Stash tools
===========

Go package of Stash tools

## Installation

```bash
make
```

## Use

```bash
go get github.com/xoom/stash

import "github.com/xoom/stash"
```

## Development

### Local stash instance

Download and run a development instance of stash via a
[docker image](https://www.docker.com/).

```bash
# pick a directory where to save the data generated by the container
export STASH_DATA="${HOME}/stash/data"

# for a linux host
$ docker run -u root -v $STASH_DATA:/var/atlassian/application-data/stash atlassian/stash chown -R daemon  /var/atlassian/application-data/stash

$ docker run -v $STASH_DATA:/var/atlassian/application-data/stash --name="stash" -d -p 7990:7990 -p 7999:7999 atlassian/stash

# for a MacOs Host via 'boot2docker'
$ docker run -u root -v $STASH_DATA:/var/atlassian/application-data/stash --name=stash -d -p 7990:7990 -p 7999:7999 atlassian/stash
```

Open your browser to `http://localhost:7990` and follow the setup instructions.

** If your are using boot2docker get your IP via `boot2docker ip`
