A golang-based digital signage server.
===

The goal of this project is a simple, easy-to-implement digital signage server that can be utilized with nothing more than a web browser on the client side.

Usage
===
The server runs, by default, on port 3030 of your host.

The configuration page lives at <server root>/config.html and allows you to enter a series of newline-delimited URLS. There'll a pause up to a minute long before the server begins providing these URLs.

The client lives at the server root, and is simply a full-page iframe that serves up whatever the current page delivered by the server is. The border and scrollbar has been removed - pages that do not require scrolling are currently required.

Page Layout DSL
===============

The DSL we've created has three special characters:

- the ` | ` (space pipe space), which is used to separate columns
- the `\n` (newline character), which is used to separate rows
- the `=` (equals), which is used to separate pages

A display has one or more pages, each of which has one or more rows with one or more columns of urls to load. 

Example: Let's say you want two pages. The first page will show a list of events and weather side-by-side, and traffic on a new line. The second page will show only your website. Your configuration would look like this:

	http://list-of-events.com/yourOrg | http://forecast.io/#/f/your_location
	http://traffic.com/your_location
	=
	http://yourOrg.org

Parsing this DSL is incredibly non-fault-tolerant right now, so don't mess up.