# ascii-art-web

Ascii-art-web involves setting up a server to run a web-based GUI version of my previous project, ascii-art.

## Author

* [@smslash](https://github.com/smslash)

## Run 

```bash
go run ./cmd/web
```

Then open link [http://127.0.0.1:4000](http://127.0.0.1:4000) in your browser.

## Objectives

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, [ascii-art](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art).

Your webpage must allow the use of the different banners:

- [shadow](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
- [standard](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
- [thinkertoy](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)

Implement the following HTTP endpoints:

- .1. GET `/`: Sends HTML response, the main page.
- 1.1 GET Tip: [go templates](https://pkg.go.dev/html/template) to receive and display data from the server.

- .2. POST `/ascii-art`: that sends data to Go server (text and a banner)
- 2.1 POST Tip: use form and other types of [tags](https://developer.mozilla.org/en-US/docs/Web/HTML/Element) to make the post request.\

The way you display the result from the POST is up to you. What we recommend are one of the following :

-  Display the result in the route `/ascii-art` after the POST is completed. So going from the home page to another page.
- Or display the result of the POST in the home page. This way appending the results in the home page.

The main page must have:

- text input
- radio buttons, select object or anything else to switch between banners
- button, which sends a POST request to `/ascii-art` and outputs the result on the page.

## HTTP status code

Your endpoints must return appropriate HTTP status codes.

- OK (200), if everything went without errors.
- Not Found, if nothing is found, for example templates or banners.
- Bad Request, for incorrect requests.
- Internal Server Error, for unhandled errors.

In the root project directory create a `README.MD` file with the following sections and contents:

- Description
- Authors
- Usage: how to run
- Implementation details: algorithm

## Instructions

- HTTP server must be written in **Go**.
- HTML templates must be in the project root directory *templates*.
- The code must respect the good [practices](https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md).

## Allowed packages

- Only the [standard go](https://pkg.go.dev/std) packages are allowed

## Usage

- [Here's an example](http://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20).
