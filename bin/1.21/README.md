![sheret-logo](https://cloud.githubusercontent.com/assets/254784/26507827/1d32d530-4220-11e7-8d1c-96d3b0abda36.png)

# Sheret
A tiny, simple static web server written in Go.

[![GPL Licence](https://badges.frapsoft.com/os/gpl/gpl.svg?v=103)](https://opensource.org/licenses/GPL-3.0/)

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [Security](#security)
- [Changelog](#changelog)
- [WIP](#wip)
- [Credits & Acknowledgments](#credits)
- [License](#license)

## Install

Windows binaries are available from the [Sheret Releases](https://github.com/ethanpil/sheret/releases) page on GitHub. Download and extract the binary to anywhere on your filesystem. There are no dependencies.

## Usage

![2017-05-26_144823](https://cloud.githubusercontent.com/assets/254784/26508345/68218c92-4222-11e7-9692-d8d1d21af680.jpg)

```
Usage: sheret [options]

Parameters:
  -d string
        directory to serve from (default ".")
  
  -f string
  	    log to disk path [./sheret.log]
  
  -p string
        port to serve on (default "8100")
  
  -q    suppress all logging
  
  -b    do not open default browser
  
  -h    show usage help
```

After downloading, you can start Sheret simply with `sheret` from the command line. By default, Sheret will begin serving from the directory it is located, on HTTP port 8100. Sheret will log all activity to the console. To terminate Sheret, simply type **CTRL+C** in the console window.

The console window will log all requests, and pretty print both `URL Parameters` as well as `POST` request variables for easier debugging.

```
2020/04/19 22:06:18 Sheret v1.2 serving . on HTTP port: 8100
2020/04/19 22:06:18 -- Press CTRL-C to terminate --
2020/04/19 22:06:19 [::1]:43578 GET /
2020/04/19 22:06:46 [::1]:43578 GET /?urlvar1=abc&urlvar2=def
2020/04/19 22:06:46      URL:   urlvar1  =      abc
2020/04/19 22:06:46      URL:   urlvar2  =      def
2020/04/19 22:06:46 ---- End Data. 2 Values Received ----
2020/04/19 22:07:05 [::1]:43578 POST /?urlvar1=abc&urlvar2=def
2020/04/19 22:07:05      URL:   urlvar1  =      abc
2020/04/19 22:07:05      URL:   urlvar2  =      def
2020/04/19 22:07:05     POST:   selector         =      Option 02
2020/04/19 22:07:05     POST:   texter   =      789
2020/04/19 22:07:05     POST:   radio2   =      on
2020/04/19 22:07:05     POST:   checker  =      on
2020/04/19 22:07:05     POST:   inputer  =      123
2020/04/19 22:07:05 ---- End Data. 7 Values Received ----
```

The **-f** option will enable logging to disk. All console output from Sheret will also be saved in `sheret.log` next to the executable.

## Contribute

PRs accepted. Submit any suggestions or corrections via a GitHub pull request. Support via GitHub issues is limited.

## Security

This is a simple tool meant for developers to test local code. It is 100% based on Golangs included HTTP library. No security related code or testing has been added by the author. This tool is not meant for production use, or even any type of server which connects directly to a public network. No guarantee is provided or implied.

## Changelog

*v1.21 - May 8, 2020
```
  New: Add flag to disable open system browser on launch
``` 

*v1.2 - April 19, 2020
```
  New: Remove folder creation in zipped release by packaging script 
  New: Reduce binary size (2mb!) by stripping debug information during build
  New: Log both URL Params and POST Vars. Simplify output
  New: Open system browser to root on launch
  Internal: Fix some whitespacing in code (Yak shave!)
  Internal: Add notification and pause to release packaging script
```

*v1.1 - April 14, 2020
```
	Remove hardcoded logfile path. -f parameter works now
	Better error handling and messages when logpath fails
	Improve quiet mode for complete quiet
	Release without UPX packing to prevent false positive with AV tools
	Improve default HTML page
	New build/release scripts
```

*v1.0 - May 26, 2017
```
	Initial Release
```

## WIP

Perhaps the next version... perhaps never. :) PR requests will help!

* Make quiet mode not mutually exclusive of logfile/console
* Commandline flag to disable browser open
* Commandline flag to disable data logging (but still log requests)
* Simple CGI passthrough (very unlikely)

## Credits & Acknowledgments

* The Sheret logo icon is provided by [FreePik](http://www.freepik.com/)
* FYI: **שרת** **Sheret** is the Hebrew word for Server :)

## License

GPL v3 © Ethan Piliavin