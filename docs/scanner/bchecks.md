> Source: https://portswigger.net/burp/documentation/scanner/bchecks.html

DASTProfessional

# BCheck definitions

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 BChecks are custom scan checks that you can create and import. Burp Scanner runs these checks in addition to its built-in scanning routine, helping you to target your scans and make your testing workflow as efficient as possible.


 Each BCheck is defined as a plain text file with a `.bcheck` file extension. These files use a custom definition language to specify the behavior of the check.


 This section explains how the BCheck definition language works, with a reference of all available keywords and some worked examples.


#### Note

 You can share your BChecks and download new ones via the [BChecks GitHub repository](https://github.com/PortSwigger/BChecks). This includes example checks created by PortSwigger, as well as BChecks developed by the Burp Suite community.


 For information on contributing to the BChecks repository, see the [Contributing](https://github.com/PortSwigger/BChecks/blob/main/CONTRIBUTING.md) readme page.


## Managing BChecks in Burp Suite Professional

 For information on how to manage BChecks in Burp Suite Professional, including importing definition files and creating new definitions from a template, see [Managing custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks/managing).


## Managing BChecks in Burp Suite DAST

 For information on how to manage BChecks, and other extensions, in Burp Suite DAST, see [Extensions in Burp Suite DAST](https://portswigger.net/burp/documentation/dast/user-guide/extensions).

#### In this section

-

[BCheck definition reference](https://portswigger.net/burp/documentation/scanner/bchecks/bcheck-definition-reference).
-

[BCheck worked examples](https://portswigger.net/burp/documentation/scanner/bchecks/worked-examples).
-

[Submitting BChecks to the community](https://portswigger.net/burp/documentation/scanner/bchecks/contribute-bchecks).
