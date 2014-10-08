# paperlesspost/cef2go

*This is a fork of [https://github.com/CzarekTomczak/cef2go/](https://github.com/CzarekTomczak/cef2go/).*

@CzarekTomczak has done a ton of amazing work in wrapping [Chromium Embedded Framework](https://code.google.com/p/chromiumembedded/) in Go and making it cross platform. If you're interested in cross platform CEF please visit the main repo.

This fork includes a lot of changes very specific to our needs and is not attempting to be anything other than a tool for our use case. Specifically we:

- Are only running in Linux
- Want/need to use cef2go as a pkg from another project
- Use offscreen rendering and render handlers
- Want callbacks from JS into Go


