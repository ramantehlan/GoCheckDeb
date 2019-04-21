# GoCheckDeb

Tool and package to get Dependencies in Go for Debian packaging. :cyclone: :page_with_curl:

# Index

- [About](#about)
- [Usage](#usage)
- [Development](#development)
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
`gocheckdeb` library.

## Usage
  - Installation
  - Commands

## Development
  - Pre-Requisites
  - Development-Environment
  - FileStructure
  - Build

## Contribution
  - Community
  - Guideline
  - Workflow
  - Best-Practices

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
