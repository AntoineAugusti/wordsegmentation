[![Travis CI](http://img.shields.io/travis/antoineaugusti/wordsegmentation/master.svg?style=flat-square)](https://travis-ci.org/antoineaugusti/wordsegmentation)
[![Software License](http://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/antoineaugusti/wordsegmentation/LICENSE.md)

# Word segmentation
Word segmentation is the process of dividing a phrase without spaces back into its constituent parts. For example, consider a phrase like "thisisatest". Humans can immediately identify that the correct phrase should be "this is a test".

## Source and credits
This package is heavily inspired by the Python module [grantjenks/wordsegment](https://github.com/grantjenks/wordsegment).

Copyright (c) 2015 by Grant Jenks under the Apache 2 license

The package is based on code from the chapter [Natural Language Corpus Data](http://norvig.com/ngrams/) by Peter Norvig from the book [Beautiful Data](http://oreilly.com/catalog/9780596157111/) (Segaran and Hammerbacher, 2009).

Copyright (c) 2008-2009 by Peter Norvig

## Getting started
You can grab this package with the following command:
```
go get gopkg.in/antoineaugusti/wordsegmentation.v0
```

## Usage
If you wanna use the default English corpus:
```go
package main

import (
    "fmt"

    "github.com/antoineaugusti/wordsegmentation"
    "github.com/antoineaugusti/wordsegmentation/corpus"
)

func main() {
    // Grab the default English corpus that will be created thanks to TSV files
    englishCorpus := corpus.NewEnglishCorpus()
    fmt.Println(wordsegmentation.Segment(englishCorpus, "thisisatest"))
}
```

## Unigrams and bigrams
> Information: an **n-gram** is a contiguous sequence of n items from a given sequence of text or speech.

This package ships with an English corpus by default that is ready to use. [Data files](https://github.com/antoineaugusti/wordsegmentation/tree/master/data) are derived from the [Google web trillion word corpus](http://googleresearch.blogspot.com/2006/08/all-our-n-gram-are-belong-to-you.html), as described by Thorsten Brants and Alex Franz, and [distributed](https://catalog.ldc.upenn.edu/LDC2006T13) by the Linguistic Data Consortium. This module contains only a subset of that data. The unigram data includes only the most common 333,000 words. Similarly, bigram data includes only the most common 250,000 phrases. Every word and phrase is lowercased with punctuation removed.
