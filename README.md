# DNSMORPH

[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)
[![GitHub release](https://img.shields.io/github/release/dnsmorph/releases.svg)](https://github.com/netevert/dnsmorph/releases)
[![license](https://img.shields.io/github/license/netevert/dnsmorph.svg)](https://github.com/netevert/dnsmorph/blob/master/LICENSE)
[![Maintenance](https://img.shields.io/maintenance/yes/2018.svg)]()
[![](https://img.shields.io/github/issues-raw/netevert/dnsmorph.svg)](https://github.com/netevert/dnsmorph/issues)
[![](https://img.shields.io/github/issues-closed-raw/netevert/dnsmorph.svg)](https://github.com/netevert/dnsmorph/issues?q=is%3Aissue+is%3Aclosed)
[![GitHub last commit](https://img.shields.io/github/last-commit/errantbot/dnsmorph.svg)](https://github.com/netevert/dnsmorph/commit/master)
[![Donations](https://img.shields.io/badge/donate-bitcoin-orange.svg?logo=bitcoin)](https://github.com/netevert/dnsmorph#donations)


DNSMORPH is a domain name permutation engine in the works. It is being written in [Go](https://golang.org/) to ensure that the end result is a robust, small and fast tool ideal for everyday use.

Installation
============
To install DNSMORPH run the following commands:

```go get github.com/netevert/dnsmorph```

`cd /$GOPATH/src/github.com/netevert/dnsmorph`

`go build`

Features
========
Usage:

    ./dnsmorph -d amazon.com

The following domain permutation attacks are included:
- Homograph attack (both on single and duplicate characters)
- Bitsquat attack
- Hyphenation attack
- Omission attack
- Repetition attack
- Replacement attack
- Subdomain attack
- Transposition attack
- Vowel swap attack
- Addition attack

**Please note that DNSMORPH is work in progress**, consult the [issues page](https://github.com/netevert/dnsmorph/issues) to see what's in the [roadmap](https://github.com/netevert/dnsmorph/milestone/1) to v.1.0.0 and how the project is progressing.

License
=======

Distributed under the terms of the [MIT](http://www.linfo.org/mitlicense.html) license, DNSMORPH is free and open
source software.

Versioning
==========

This project adheres to [Semantic Versioning](https://semver.org/).

Donations
=========

<details><summary>If you like DNSMORPH please consider donating</summary>
<p>
    
    Bitcoin:  13i3hFGN1RaQqdeWqmPTMuYEj9FiJWuMWf
    Litecoin: LZqLoRNHvJyuKz99mNAgVUj6M8iyEQuio9
</p>
</details>