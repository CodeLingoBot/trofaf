// trofaf is a super-simple *live* static blog engine.
//
// Install using: `go get github.com/PuerkitoBio/trofaf`
//
// ## Description
//
// It generates a static website from *markdown* files and front matter, and requires only a simple
// 3-directories structure to get going. It favors simplicity over features.
//
// To get started, create the 3 subdirectories (you can look at the `example/` subdirectory
// for... an example):
// * posts
// * public
// * templates
//
// trofaf only cares about `*.md` files in the posts directory, and about `*.amber` (Amber templates)
// or `*.html` (native Go templates) files in the templates directory. It will watch for changes,
// creates or deletes on those files in these directories, and will re-generate automatically
// the website when required. This is the *live* part.
//
// All files in the public directory are exposed by the web server. Posts in markdown format get
// translated to static html files at the root of the public directory. The html file name is
// an URL-friendly slug generated from the original markdown file name. There is no extension, so
// the URL looks clean and, uh, *modern*?
//
// An RSS feed is automatically generated from a number of recent posts, and saved as a static
// XML file in the public directory.
//
// There is no special template for an index page, the most recent post (based on the publication
// date found in the front matter of the markdown files) is saved twice - once under its own
// html file, once under the index.html file, so that this is the page displayed when the root
// of the web server is requested.
//
// When the site is (re-)generated, the public directory must be cleaned, because some posts may
// have been deleted. Subdirectories are left untouched (so that `css/` or `js/` directories can
// coexist peacefully), as well as hidden (dot) files, and some special files are also graced
// from the destruction (robots.txt, favicon.ico, etc., see gen.go).
//
// ## Command-line Options
//
// The following options can be set at the command-line:
// * Port (-p) : the port number for the web server, defaults to 9000.
// * NoGen (-G) : prevents watching and live-generating the site. This is equivalent to running the static public directory.
// * SiteName (-n) : the name of the web site, passed to the template.
// * TagLine (-t) : a tag line for the web site, passed to the template.
// * RecentPostsCount (-r) : the number of posts in the recent posts list, passed to the template and used for the RSS feed.
// * BaseURL (-b) : the base URL of the web site, most likely the host name (i.e. http://www.mysite.com).
package main