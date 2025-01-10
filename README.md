# Setup

1. Build gotty-core image from source:
```bash
./build-image.sh
```
2. Create otp secret using gen2fa username
3. Update credentials.json
4. Build gotty docker image and run:
```bash
srv/create.sh
```

# ![](https://raw.githubusercontent.com/sorenisanerd/gotty/master/resources/favicon.ico) GoTTY - Share your terminal as a web application
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-57-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![GitHub release](http://img.shields.io/github/release/sorenisanerd/gotty.svg?style=flat-square)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Maintainer streaming](https://twitch-status.soren.tools/sorencodes)][twitch]

[release]: https://github.com/sorenisanerd/gotty/releases
[license]: https://github.com/sorenisanerd/gotty/blob/master/LICENSE
[twitch]: https://twitch.tv/sorencodes

GoTTY is a simple command line tool that turns your CLI tools into web applications.

[Original work](https://github.com/yudai/gotty) by [Iwasaki Yudai](https://github.com/yudai). There would be no GoTTY without him. ❤️

![Screenshot](https://raw.githubusercontent.com/sorenisanerd/gotty/master/screenshot.gif)

# Installation

## From release page

You can download the latest stable binary file from the [Releases](https://github.com/sorenisanerd/gotty/releases) page. Note that the release marked `Pre-release` is built for testing purpose, which can include unstable or breaking changes. Download a release marked [Latest release](https://github.com/sorenisanerd/gotty/releases/latest) for a stable build.

(Files named with `darwin_amd64` are for Mac OS X users)

## Homebrew Installation

You can install GoTTY with [Homebrew](http://brew.sh/) as well.

```sh
$ brew install sorenisanerd/gotty/gotty
```

## `go get` Installation (Development)

If you have a Go language environment, you can install GoTTY with the `go get` command. However, this command builds a binary file from the latest master branch, which can include unstable or breaking changes. GoTTY requires go1.9 or later.

```sh
$ go get github.com/sorenisanerd/gotty
```

# Usage

```
Usage: gotty [options] <command> [<arguments...>]
```

Run `gotty` with your preferred command as its arguments (e.g. `gotty top`).

By default, GoTTY starts a web server at port 8080. Open the URL on your web browser and you can see the running command as if it were running on your terminal.

## Options
```sh
   --address value, -a value     IP address to listen (default: "0.0.0.0") [$GOTTY_ADDRESS]
   --port value, -p value        Port number to liten (default: "8080") [$GOTTY_PORT]
   --path value, -m value        Base path (default: "/") [$GOTTY_PATH]
   --permit-write, -w            Permit clients to write to the TTY (BE CAREFUL) (default: false) [$GOTTY_PERMIT_WRITE]
   --credential value, -c value  Credential for Basic Authentication (ex: user:pass, default disabled) [$GOTTY_CREDENTIAL]
   --random-url, -r              Add a random string to the URL (default: false) [$GOTTY_RANDOM_URL]
   --random-url-length value     Random URL length (default: 8) [$GOTTY_RANDOM_URL_LENGTH]
   --tls, -t                     Enable TLS/SSL (default: false) [$GOTTY_TLS]
   --tls-crt value               TLS/SSL certificate file path (default: "~/.gotty.crt") [$GOTTY_TLS_CRT]
   --tls-key value               TLS/SSL key file path (default: "~/.gotty.key") [$GOTTY_TLS_KEY]
   --tls-ca-crt value            TLS/SSL CA certificate file for client certifications (default: "~/.gotty.ca.crt") [$GOTTY_TLS_CA_CRT]
   --index value                 Custom index.html file [$GOTTY_INDEX]
   --title-format value          Title format of browser window (default: "{{ .command }}@{{ .hostname }}") [$GOTTY_TITLE_FORMAT]
   --reconnect                   Enable reconnection (default: false) [$GOTTY_RECONNECT]
   --reconnect-time value        Time to reconnect (default: 10) [$GOTTY_RECONNECT_TIME]
   --max-connection value        Maximum connection to gotty (default: 0) [$GOTTY_MAX_CONNECTION]
   --once                        Accept only one client and exit on disconnection (default: false) [$GOTTY_ONCE]
   --timeout value               Timeout seconds for waiting a client(0 to disable) (default: 0) [$GOTTY_TIMEOUT]
   --permit-arguments            Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB) (default: false) [$GOTTY_PERMIT_ARGUMENTS]
   --width value                 Static width of the screen, 0(default) means dynamically resize (default: 0) [$GOTTY_WIDTH]
   --height value                Static height of the screen, 0(default) means dynamically resize (default: 0) [$GOTTY_HEIGHT]
   --ws-origin value             A regular expression that matches origin URLs to be accepted by WebSocket. No cross origin requests are acceptable by default [$GOTTY_WS_ORIGIN]
   --enable-webgl                Enable WebGL renderer (default: true) [$GOTTY_ENABLE_WEBGL]
   --close-signal value          Signal sent to the command process when gotty close it (default: SIGHUP) (default: 1) [$GOTTY_CLOSE_SIGNAL]
   --close-timeout value         Time in seconds to force kill process after client is disconnected (default: -1) (default: -1) [$GOTTY_CLOSE_TIMEOUT]
   --config value                Config file path (default: "~/.gotty") [$GOTTY_CONFIG]
   --help, -h                    show help (default: false)
   --version, -v                 print the version (default: false)
```
### Config File
You can customize default options and your terminal by providing a config file to the `gotty` command. GoTTY loads a profile file at `~/.gotty` by default when it exists.

```
// Listen at port 9000 by default
port = "9000"

// Enable TSL/SSL by default
enable_tls = true

```

See the [`.gotty`](https://github.com/sorenisanerd/gotty/blob/master/.gotty) file in this repository for the list of configuration options.

### Security Options

By default, GoTTY doesn't allow clients to send any keystrokes or commands except terminal window resizing. When you want to permit clients to write input to the TTY, add the `-w` option. However, accepting input from remote clients is dangerous for most commands. When you need interaction with the TTY for some reasons, consider starting GoTTY with tmux or GNU Screen and run your command on it (see "Sharing with Multiple Clients" section for detail).

To restrict client access, you can use the `-c` option to enable the basic authentication. With this option, clients need to input the specified username and password to connect to the GoTTY server. Note that the credentials will be transmitted between the server and clients in plain text. For more strict authentication, consider the SSL/TLS client certificate authentication described below.

The `-r` option is a little bit more casual way to restrict access. With this option, GoTTY generates a random URL so that only people who know the URL can get access to the server.

All traffic between the server and clients are NOT encrypted by default. When you send secret information through GoTTY, we strongly recommend you use the `-t` option which enables TLS/SSL on the session. By default, GoTTY loads the crt and key files placed at `~/.gotty.crt` and `~/.gotty.key`. You can overwrite these file paths with the `--tls-crt` and `--tls-key` options. When you need to generate a self-signed certification file, you can use the `openssl` command.

```sh
openssl req -x509 -nodes -days 9999 -newkey rsa:2048 -keyout ~/.gotty.key -out ~/.gotty.crt
```

(NOTE: For Safari uses, see [how to enable self-signed certificates for WebSockets](http://blog.marcon.me/post/24874118286/secure-websockets-safari) when use self-signed certificates)

For additional security, you can use the SSL/TLS client certificate authentication by providing a CA certificate file to the `--tls-ca-crt` option (this option requires the `-t` or `--tls` to be set). This option requires all clients to send valid client certificates that are signed by the specified certification authority.

## Sharing with Multiple Clients

GoTTY starts a new process with the given command when a new client connects to the server. This means users cannot share a single terminal with others by default. However, you can use terminal multiplexers for sharing a single process with multiple clients.
### Screen
After installing GNU screen, start a new session with `screen -S name-for-session` and connect to it with gotty in another terminal window/tab through `screen -x name-for-session`. All commands and activities being done in the first terminal tab/window will now be broadcasted by gotty.
### Tmux
For example, you can start a new tmux session named `gotty` with `top` command by the command below.

```sh
$ gotty tmux new -A -s gotty top
```

This command doesn't allow clients to send keystrokes, however, you can attach the session from your local terminal and run operations like switching the mode of the `top` command. To connect to the tmux session from your terminal, you can use following command.

```sh
$ tmux new -A -s gotty
```

By using terminal multiplexers, you can have the control of your terminal and allow clients to just see your screen.

### Quick Sharing on tmux

To share your current session with others by a shortcut key, you can add a line like below to your `.tmux.conf`.

```
# Start GoTTY in a new window with C-t
bind-key C-t new-window "gotty tmux attach -t `tmux display -p '#S'`"
```

## Playing with Docker

When you want to create a jailed environment for each client, you can use Docker containers like following:

```sh
$ gotty -w docker run -it --rm busybox
```

## Development

You can build a binary by simply running `make`. go1.16 is required.

To build the frontend part (JS files and other static files), you need `npm`.

## Architecture

GoTTY uses [xterm.js](https://xtermjs.org/) to run a JavaScript based terminal on web browsers. GoTTY itself provides a websocket server that simply relays output from the TTY to clients and receives input from clients and forwards it to the TTY. This xterm + websocket idea is inspired by [Wetty](https://github.com/krishnasrinivas/wetty).
