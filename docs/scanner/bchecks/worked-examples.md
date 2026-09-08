> Source: https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples

DASTProfessional

# BChecks worked examples

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 BChecks are defined by importing `.bcheck` files into Burp Suite Professional. This section provides some example definitions that correspond to real-world use cases, and breaks down how each definition works to help you design your own scan checks.


 These examples are also available as templates when you create a BCheck in Burp Suite Professional. You can use them as the basis for your new check, and edit them as required.


#### Note

 The examples in this section use Python-style indentation for readability. However, indentation is entirely optional when writing BCheck definitions.


## In this section

- [Host check](https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples/host)
- [Path check](https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples/path)
- [Response-level (passive) check](https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples/passive)
- [Insertion point check](https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples/insertion-point)
- [Collaborator-based check](https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples/collaborator)
- [Log4Shell check](https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples/log4shell)
- [Server-side prototype pollution check](https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples/server-side-prototype-pollution)

#### Related pages

-

[BCheck definition reference](https://portswigger.net/burp/documentation/scanner/bchecks/bcheck-definition-reference)

-

[Adding custom checks to scans [Burp Suite Professional]](https://portswigger.net/burp/documentation/desktop/running-scans/custom-checks)
