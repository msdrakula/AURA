> Source: https://portswigger.net/burp/documentation/desktop/tools/dom-invader

ProfessionalCommunity Edition

# DOM Invader

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 DOM Invader is a browser-based tool that helps you test for DOM XSS vulnerabilities using a variety of sources and sinks, including both web message and prototype pollution vectors. It is available exclusively via Burp's built-in browser, where it comes preinstalled as an extension.


![DOM Invader tab for the browser DevTools](https://portswigger.net/burp/documentation/desktop/images/dom-invader-innerhtml-sink.png)

## Key features

 When enabled, DOM Invader adds a new tab to the browser's DevTools panel. This enables you to perform the following key tasks:


-

 Test for DOM XSS like reflected XSS. The augmented DOM view enables you to instantly identify controllable sinks on the page, showing you both the XSS context and how your input is being sanitized. For more information, see [Testing for DOM XSS](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/dom-xss).

-

 Log, modify, and resend web messages that are sent on a page via the `postMessage()` method. This enables you to test for DOM XSS via web messages. You can also let DOM Invader probe for vulnerabilities on your behalf by sending its own specially crafted web messages. For more information, see [Testing for DOM XSS using web messages](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/web-messages).

-

 Automatically identify sources of client-side prototype pollution and scan for controllable gadgets that are passed to dangerous sinks. For more information, see [Testing for client-side prototype pollution](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/prototype-pollution).

-

 Automatically identify [DOM clobbering](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/dom-clobbering) vulnerabilities.


 For more information on how to enable DOM Invader, see [Enabling DOM Invader](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/enabling).


 DOM Invader is highly configurable, so you can fine-tune its behavior to suit different websites and use cases. For more information, see [DOM Invader settings](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/settings).
