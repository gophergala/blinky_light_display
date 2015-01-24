A golang-based digital signage server.
===

The goal of this project is a simple, easy-to-implement digital signage server that can be utilized with nothing more than a web browser on the client side.

Usage
===
The server runs, by default, on port 3030 of your host.

The configuration page lives at <server root>/config.html and allows you to enter a series of newline-delimited URLS. There'll a pause up to a minute long before the server begins providing these URLs.

The client lives at the server root, and is simply a full-page iframe that serves up whatever the current page delivered by the server is. The border and scrollbar has been removed - pages that do not require scrolling are currently required.
