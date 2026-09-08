> Source: https://portswigger.net/burp/documentation/desktop/tools/dom-invader/dom-clobbering

ProfessionalCommunity Edition

# Testing for DOM clobbering with DOM Invader

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 DOM Invader can automatically test for DOM clobbering vulnerabilities on your behalf. DOM clobbering is a technique in which you inject HTML into a page to manipulate the DOM in a way that enables you to change the behavior of JavaScript on the page.


#### Web Security Academy

 For more information about DOM clobbering, as well as some interactive, deliberately vulnerable labs, check out the related topic on the Web Security Academy.
  [DOM clobbering](https://portswigger.net/web-security/dom-based/dom-clobbering)

## Enabling DOM clobbering

 To avoid interfering with your target site's functionality, DOM clobbering is disabled by default. To enable these checks:


1.

 Go to the DOM Invader settings menu.

1.

 Under **Attack types**, toggle the switch so that **DOM clobbering** is on.

1.

 Click **Reload** to refresh the browser. This is necessary for your changes to take effect.


![Enabling DOM clobbering in DOM Invader](https://portswigger.net/burp/documentation/desktop/images/dom-invader-dom-clobbering-enabling.png)

 DOM Invader now scans for DOM clobbering vulnerabilities as you browse.
