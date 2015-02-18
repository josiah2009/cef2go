# paperlesspost/cef2go

*This is a fork of [https://github.com/CzarekTomczak/cef2go/](https://github.com/CzarekTomczak/cef2go/).*

@CzarekTomczak has done a ton of amazing work in wrapping [Chromium Embedded Framework](https://code.google.com/p/chromiumembedded/) in Go and making it cross platform. If you're interested in cross platform CEF please visit the main repo.

This fork includes a lot of changes very specific to our needs and is not attempting to be anything other than a tool for our use case. Specifically we:

- Are only running in Linux
- Want/need to use cef2go as a pkg from another project
- Use offscreen rendering and render handlers
- Want callbacks from JS into Go

## Known issues/constraints

- This fork is only tested on linux and currently only tested on Ubuntu (13 & 14).
- This fork is specifically geared towards building headless applications that need offscreen rendering and JS integration.
- JS callbacks into the application have some interesting properties when running in multi-process mode as they callback to the inidividual browser process, not the main process (where your go code is actually running). We're working on better ways around this, but currently we run in single-process mode to get around this constraint.
- If you're running on a machine that does not have a Graphics processor, it is best to run with the `--disable-gpu` flag to avoid warnings and potential errors.
- There is an issue where in certain circumstances an application based on cef2go does not boot properly and then can not create browsers. There is a bunch of discussion about this here <https://github.com/CzarekTomczak/cef2go/issues/16> and here <https://code.google.com/p/chromiumembedded/issues/detail?id=1362>

## CEF Version and Compatibility

This fork is compiled against CEF 3 branch 2272 revision 1998 (Chrome 41) from <http://cefbuilds.com>

## Building and running an application

Building an application based on cef2go on linux requires some special setup and dependencies.

After downloading the specific cef build outlined above and extracting it to a directory you need to move the libraries and resources to their proper locations.

``` bash
cd cef_binary_3.2272.1998_linux64
# Move libcef.so to the shared lib directory (/usr/lib) 
# so it can be found by the linker (ld) (will probably require sudo)
mv Release/*.so /usr/lib
# Make a shared directory to hold the resources files and copy the resources there
mkdir /var/lib/cef
cp Resources/* /var/lib/cef/
```

See the `extract_cef` script for a working example of extraction.

Because of hardcoded paths in Chromium itself, every application needs the resource files in the same directory as the compiled application binary. The easiest thing to do is move all the resources to a shared directory (as in the script above) and then symlink each file into your binary working directory.

See the `link_cef` script for a full example of linking.

Messages like:

```
[0218/214957:ERROR:icu_util.cc(152)] Couldn't mmap /opt/src/paperlesspost/go/src/github.com/paperlesspost/cef2go/icudtl.dat
[0218/214957:FATAL:content_main_runner.cc(749)] Check failed: base::i18n::InitializeICU(). 
```

mean that the resources are not correctly linked

## Updating the CEF Version

The process for updating CEF to a newer version is relatively simple:


