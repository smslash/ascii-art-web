# ascii-art-web

Ascii-art-web involves setting up a server to run a web-based GUI version of my previous project, ascii-art.

## Author

* [@smslash](https://github.com/smslash)

## Run 

```bash
go run ./cmd/web
```

Then open link <a href="http://127.0.0.1:4000" style="color:blue">http://127.0.0.1:4000</a> in your browser.

## Objectives

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, [ascii-art](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art).

Your webpage must allow the use of the different banners:

&nbsp; · [shadow](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
&nbsp; · [standard](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
&nbsp; · [thinkertoy](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)

Implement the following HTTP endpoints:

&nbsp;1. GET `/`: Sends HTML response, the main page.
&nbsp;&nbsp;1.1 GET Tip: [go templates](https://pkg.go.dev/html/template) to receive and display data from the server.

&nbsp;2. POST `/ascii-art`: that sends data to Go server (text and a banner)
&nbsp;&nbsp;2.1 POST Tip: use form and other types of [tags](https://developer.mozilla.org/en-US/docs/Web/HTML/Element) to make the post request.\

The way you display the result from the POST is up to you. What we recommend are one of the following :

&nbsp; · Display the result in the route `/ascii-art` after the POST is completed. So going from the home page to another page.
&nbsp; · Or display the result of the POST in the home page. This way appending the results in the home page.

The main page must have:

&nbsp; · text input
&nbsp; · radio buttons, select object or anything else to switch between banners
&nbsp; · button, which sends a POST request to `/ascii-art` and outputs the result on the page.

## HTTP status code

Your endpoints must return appropriate HTTP status codes.

&nbsp; · OK (200), if everything went without errors.
&nbsp; · Not Found, if nothing is found, for example templates or banners.
&nbsp; · Bad Request, for incorrect requests.
&nbsp; · Internal Server Error, for unhandled errors.

In the root project directory create a `README.MD` file with the following sections and contents:

&nbsp; · Description
&nbsp; · Authors
&nbsp; · Usage: how to run
&nbsp; · Implementation details: algorithm

## Instructions

&nbsp; · HTTP server must be written in **Go**.
&nbsp; · HTML templates must be in the project root directory *templates*.
&nbsp; · The code must respect the good [practices](https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md).

## Allowed packages

&nbsp; · Only the [standard go](https://pkg.go.dev/std) packages are allowed

## Usage

&nbsp; · [Here's an example](http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20).
