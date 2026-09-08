> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/bypassing-filters

ProfessionalCommunity Edition

# Bypassing XSS filters by enumerating permitted tags and attributes

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Reflected cross-site scripting (XSS) arises when an application receives data in an HTTP request, then includes that data in its response in an unsafe way.


 Applications use a range of processing and input validation methods to protect against common XSS payloads. You can use Burp Intruder to enumerate tags and attributes that are permitted by the application. This enables you to craft an XSS payload that will be executed by the application, and is a useful next step if your attempts to test using proof-of-concept payloads were not successful.


## Before you start

 Identify a request / response pair with reflected input. For more information, see [Identifying reflected
 input](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/reflected-input).


## Steps

 You can follow the processes below using the lab [Reflected XSS into HTML context with most tags and attributes blocked](https://portswigger.net/web-security/cross-site-scripting/contexts/lab-html-context-with-most-tags-and-attributes-blocked).


1. In **Proxy > HTTP history**, right-click the request with a reflected input that you want to investigate. Select **Send to Intruder**.
1.

Identify whether any tags are permitted:

  1. In Intruder, replace the value of the input with: `<>`.
  1.

Click inside the angle brackets, then click **Add §** to add a payload
 position.

![Payload for identifying whether tags are permitted](https://portswigger.net/burp/documentation/desktop/images/tutorials/bypassing-filters-1.png)

  1. In the **Payloads** side panel, under **Payload configuration**, add a list of tags that you want to test. For example, use the tags in the
 [XSS cheat sheet](https://portswigger.net/web-security/cross-site-scripting/cheat-sheet).
  1. Click ** Start attack**. The attack starts running in a new dialog. Intruder sends a request for each tag on the list.
  1. When the attack is finished, look for any responses with a `200` status code. This indicates that the tag is permitted. If a tag is filtered out, it has a
 `400` status code instead.

1.

Identify whether any attributes are permitted:

  1.

In the **Intruder** tab, update the payload position. Add a tag that you enumerated in the previous step, click **Add §** to add a payload position to test different attributes.

![Payload for identifying whether attributes are permitted](https://portswigger.net/burp/documentation/desktop/images/tutorials/bypassing-filters-2.png)

  1. In the **Payloads** side panel, under **Payload configuration**, click **Clear** to remove the list of tags that you tested in the previous step.
  1. Add a list of attributes that you want to test. For example, use the events listed in the [XSS cheat sheet](https://portswigger.net/web-security/cross-site-scripting/cheat-sheet).
  1. Click ** Start attack**. The attack starts running in a new dialog. Intruder sends a request for each attribute on the list.
  1. When the attack is finished, look for any responses with a `200` status code. This indicates that an attribute is permitted.

 You can use the permitted tags and attributes that you identified to construct an attack string. For more information, see [Testing for reflected XSS manually](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/testing-for-reflected-xss).


#### Related pages

- Academy: [Reflected XSS](https://portswigger.net/web-security/cross-site-scripting/reflected)
- [Cross-site scripting (XSS) cheat sheet](https://portswigger.net/web-security/cross-site-scripting/cheat-sheet)
- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
- [Testing for reflected XSS manually with Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/testing-for-reflected-xss)
- [Testing for stored XSS with Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/testing-for-stored-xss)
