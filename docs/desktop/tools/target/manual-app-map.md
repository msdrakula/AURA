> Source: https://portswigger.net/burp/documentation/desktop/tools/target/manual-app-map

ProfessionalCommunity Edition

# Manual application mapping

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You can choose between manual and automated application mapping. Consider the type of application and how you intend to use the results.


 For some use cases, Burp's [automated crawler](https://portswigger.net/burp/documentation/scanner/crawling) is superior to manual mapping. The crawler captures the navigational paths in a way that lets Burp Scanner automatically maintain session when it [audits](https://portswigger.net/burp/documentation/scanner/auditing) the application.


 If you map the application manually, you can guide the process and avoid potentially dangerous functionality. You can also make sure that navigational actions work as you expect, and familiarize yourself with the application.


 To manually map the application:


1.
 Launch [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser).

1.
 Browse the entire application manually.

1.
 Follow every link, submit every form, step through every multi-stage process, and log in to all protected areas.


 This manual mapping process populates the Target site map with the content requested via the Proxy. In addition, you can use [live passive crawling](https://portswigger.net/burp/documentation/desktop/running-scans/live-tasks) to map content that can be inferred from application responses, such as from links or forms. This process builds up a fairly complete record of all the visible application content.
