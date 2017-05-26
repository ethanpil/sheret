![sheret-logo](https://cloud.githubusercontent.com/assets/254784/26507827/1d32d530-4220-11e7-8d1c-96d3b0abda36.png)

# Sheret
A tiny, simple static web server written in Go.

[![GPL Licence](https://badges.frapsoft.com/os/gpl/gpl.svg?v=103)](https://opensource.org/licenses/GPL-3.0/)

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [Security](#security)
- [Credits & Acknowledgments](#credits)
- [License](#license)

## Install

Windows binaries are available from the (Sheret Releases)[https://github.com/ethanpil/sheret/releases] page on GitHub. Download and extract the binary to anywhere on your filesystem. There are no dependencies.

Note: The binary release is packed with [UPX](https://upx.github.io/) to reduce size.

## Usage

![2017-05-26_144823](https://cloud.githubusercontent.com/assets/254784/26508345/68218c92-4222-11e7-9692-d8d1d21af680.jpg)

```
Usage: sheret [options]

Parameters:
  -d string
        directory to serve from (default ".")
  
  -f    log to disk [sheret.log]
  
  -p string
        port to serve on (default "8100")
  
  -q    suppress all logging
  
  -h    show usage help
```

After downloading, you can start Sheret simply with `sheret` from the command line. By default, Sheret will begin serving from the directory it is located, on HTTP port 8100. Sheret will log all activity to the console. To terminate Sheret, simply type **CTRL+C** in the console window.

The console window will log all requests, and pretty print `POST` request variables for easier debugging.

```
2017/05/26 14:51:41 [::1]:61814 GET /
2017/05/26 14:51:34 [::1]:61814 POST /
2017/05/26 14:51:34 ---- POST Data: ------------------------
2017/05/26 14:51:34 number-one   =       MyValue
2017/05/26 14:51:34 number-two   =       AnotherVal
2017/05/26 14:51:34 number-three =
2017/05/26 14:51:34 ---- End POST Data. 3 Fields Received --
```

The **-f** option will enable logging to disk. All console output from Sheret will also be saved in `sheret.log` next to the executable.

## Contribute

PRs accepted. Submit any suggestions or corrections via a GitHub pull request. Support via GitHub issues is limited.

## Security

## Credits & Acknowledgments

* The Sheret logo icon is provided by [FreePik](http://www.freepik.com/)

## License

GPL v3 © Ethan Piliavin