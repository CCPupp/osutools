# https://www.osutools.com

## A static site attempting to maintain a list of all 3rd party resources commonly used alongside <a target="_blank" href="https://osu.ppy.sh"> osu! </a>

<br>

# How to add items

Please make a pull request for any additional resources that should be added to the site. I want things to stay pretty basic for this, so just make changes to the html page for now.

If you aren't git-savy, feel free to request an item in my discord server or directly to me ponpar#0001.

<br>

# What am I looking for?

Anything that isn't created by peppy, or is not available through the main osu site. Things like skins, beatmaps, or tournaments, are not intended for this list.

<br>

# How do I test my addition?

-   Download and install golang
-   Clone the repository locally
-   Run the server using `go run server.go`
-   Navigate to `localhost:3000` in your browser
-   To set up SASS functionality, run `sass --watch .\web\sass\:.\web\css` (must have [sass](https://sass-lang.com/install) installed)
