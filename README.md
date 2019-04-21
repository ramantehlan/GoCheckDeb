# GoCheckDeb

Tool and package to get Dependencies in Go for Debian packaging. :cyclone: :page_with_curl:

# Index

- [About](#about)
- [Usage](#usage)
	- [Tool](#tool)
	- [Package](#package)
- [Development](#development)
	- [Pre-Requisites](#pre-requisites)
	- [Development Environment](#development-environment)
	- [File Structure](#file-structure)
- [Contribution](#contribution)
- [Resources](#resources)
- [Gallery](#gallery)
- [License](#license)

## About

Packaging a Golang package for Debian requires all the Golang dependencies to be
already packaged for Debian, so we need to know all the dependencies and
sub-dependencies for the Golang project, and then we need to know which out of
them is not yet packaged for Debian.

GoCheckDeb is a tool to do the same, it will get all the Golang project
dependencies and sub-dependencies and filter out the ones which are not packaged
yet. You can also use this tool in your project, you just have to import the
**gocheckdeb** library.

## Usage

You can either use it as a tool or as a package for your own project. Instructions to download and use it are given below.

#### Tool

1. Download the Binary from [here](https://github.com/ramantehlan/GoCheckDeb/raw/master/GoCheckDeb)
2. Make the Binary executable. `$ chmod +x ./GoCheckDeb`

Once you have downloaded the tool, you can use following commands to use it.

#### Package

1. Download the package. `$ go get github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb`
2. Import it into you go program. `import github.com/ramantehlan/GoCheckDeb/pkg/gocheckdeb`

Once you have imported the package, you can use following document to use it.

 ```go

// Global Variables inclure
GoDebBinaryStruct


 ```

## Development

#### Pre-Requisites

To work on this project, you need to have the following pre-requisites.

1. Little experience of Linux and Terminal.
2. Familiarity with Golang and Golang packages.
3. Familiarity with Debian packaging.

Also, you need to have Go Environment setup in your system.

#### Development Environment

To contribute to this project, or build it for other platforms use the following steps.

1. Clone this project in your $GOPATH. `$ git clone https://github.com/ramantehlan/GoCheckDeb`
2. Get all the dependencies. `$ go get ./...`
3. Build the project. `$ go build`

#### File Structure

```
.
├── GoCheckDeb
├── LICENSE
├── main.go
├── README.md
├── .gitignore
└── pkg
    └── gocheckdeb
        └── gocheckdeb.go
```

## Contribution

Your contributions are always welcome and appreciated. Following are the things you can do to contribute to this project.

1. **Report a bug** <br>
If you think you have encountered a bug, and I should know about it, feel free to report it [here](https://github.com/ramantehlan/GoCheckDeb/issues/new) and I will take care of it.

2. **Request a feature** <br>
You can also request for a feature [here](https://github.com/ramantehlan/GoCheckDeb/issues/new), and if it will viable, it will be picked for development.  

3. **Create a pull request** <br>
It can't get better then this, your pull request will be really appreciated by the community. You can get started by picking up any open issues from [here](https://github.com/ramantehlan/GoCheckDeb/issues) and make a pull request.

> If you are new to open-source, make sure to check read more about it [here](https://www.digitalocean.com/community/tutorial_series/an-introduction-to-open-source) and learn more about creating a pull request [here](https://www.digitalocean.com/community/tutorials/how-to-create-a-pull-request-on-github).

## Resources

- [Golang Packages](https://golang.org/pkg/)
- [The Debian Go Packaging Team](https://go-team.pages.debian.net/)
- [dh-make-golang](https://github.com/Debian/dh-make-golang)

## Gallery

`To be Updated`

## License

MIT License

Copyright (c) 2019 Raman Tehlan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
