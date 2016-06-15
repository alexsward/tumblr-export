# tumblr-export
quick and dirty utility to export tumblr posts to disk. doesn't handle gifs.

### install
`go get github.com/alexsward/tumblr-export`

`go build`

`./tumblr-export -blog yourblog -key 1234 -output /path/to/files`

### usage
all parameters are required

**-blog** *string*: tumblr blog name without tumblr.com - blogname.tumblr.com

**-key** *string*: API key for tumblr

**-output** *string*: output directory for the saved files
