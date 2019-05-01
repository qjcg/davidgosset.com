# davidgosset.com

Personal website built with [Hugo](https://gohugo.io).


## Theme Configuration

It all happens in `config.toml`. See the theme's [wiki] configuration parameters.

[wiki]: https://github.com/luizdepra/hugo-coder/wiki/Configurations


## Local Development

To work on the site, run:

```sh
hugo serve
```

To serve on all LAN IPs (eg. to test on your phone):

```sh
hugo serve --bind 0.0.0.0 --baseUrl http://YOUR.LAN.IP.ADDRESS/
```


## Building the Static Site

To build static files ready for deployment:

```sh
hugo
```

The resulting static files can then be found under the `public/` directory.

To then test serving the static files with a local web server:

```sh
hugo --baseUrl http://localhost:SERVER_PORT ./public
```
