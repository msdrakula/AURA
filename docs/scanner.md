> Source: https://portswigger.net/burp/documentation/scanner

DASTProfessional

# Burp Scanner

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp Scanner is an automated dynamic application security testing (DAST) web vulnerability scanner. Designed to replicate the actions and methodologies of a skilled manual tester, Burp Scanner powers scans in Burp Suite's desktop editions and Burp Suite DAST.


## How do scans work?

 Burp Scanner handles virtually any target. Advanced features such as state management and automated logins enable it to deal with the challenges that scanning modern web applications can pose. Although the actions taken during a scan vary depending on target and configuration, scans generally comprise two key phases:


-  **Crawling** - The scanner catalogs the content of the application and the navigational paths within it. Burp Scanner navigates around the application in largely the same way that a human would. It follows links, submits forms, and logs in where necessary to create a map of the application's content.

-  **Auditing** - The scanner analyzes the application's traffic and behavior to identify security vulnerabilities and other issues. Burp Scanner sends a series of requests to the application and examines the results. It uses the information obtained in the crawl phase to determine the most efficient way to work.


 This section of the site gives more information on Burp Scanner's features and how you can configure scans to best meet your needs.


 Burp AT can run scans and use Burp's other tools for you as part of a Burp AT task. For more information, see [Burp AT](https://portswigger.net/burp/documentation/desktop/burp-at).


#### Read more

-

[Crawling](https://portswigger.net/burp/documentation/scanner/crawling).
-

[Auditing](https://portswigger.net/burp/documentation/scanner/auditing).
-

[Scan configurations](https://portswigger.net/burp/documentation/scanner/scan-configurations).

  -

[Preset scan modes](https://portswigger.net/burp/documentation/scanner/scan-configurations/preset-scan-modes).
  -

[Custom scan configurations](https://portswigger.net/burp/documentation/scanner/scan-configurations/custom-scan-configurations).
  -

[Burp Scanner built-in configurations](https://portswigger.net/burp/documentation/scanner/scan-configurations/burp-scanner-built-in-configs).
  -

[Audit settings](https://portswigger.net/burp/documentation/scanner/scan-configurations/audit-settings).
  -

[Crawl settings](https://portswigger.net/burp/documentation/scanner/scan-configurations/crawl-settings).

-

[Browser-powered scanning](https://portswigger.net/burp/documentation/scanner/browser-powered-scanning).
-

[Authenticated scanning](https://portswigger.net/burp/documentation/scanner/authenticated-scanning).

  -

[Login credentials](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/login-credentials).
  -

[Identifying login and registration forms](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/identifying-login-forms).
  -

[Recorded login sequences](https://portswigger.net/burp/documentation/scanner/authenticated-scanning/recorded-login-sequences).

-

[Requirements for API scanning](https://portswigger.net/burp/documentation/scanner/api-scanning-reqs).
-

[API scanning FAQs](https://portswigger.net/burp/documentation/scanner/api-scanning-faq).
-

[Scanning single-page apps](https://portswigger.net/burp/documentation/scanner/scanning-spas).
-

[Burp Scanner error reference](https://portswigger.net/burp/documentation/scanner/burp-scanner-error-reference).
-

[Vulnerabilities detected by Burp Scanner](https://portswigger.net/burp/documentation/scanner/vulnerabilities-list).
-

[BChecks](https://portswigger.net/burp/documentation/scanner/bchecks)
