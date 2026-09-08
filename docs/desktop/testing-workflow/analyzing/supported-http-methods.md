> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/analyzing/supported-http-methods

ProfessionalCommunity Edition

# Identifying supported HTTP methods with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 4 Minutes

 Due to misconfiguration, some websites support additional HTTP methods that may be useful to an attacker in a number of ways. These include:


- Enabling you to perform destructive actions like modifying sensitive data or uploading malicious files.
- Enabling you to bypass flawed access controls and other weak defenses.
- Disclosing additional information about the website and its infrastructure.

 Websites may also ignore the method completely and respond in the same way even if you specify an arbitrary, non-standard method. This could enable you to bypass WAFs and weak input filters, as well as some defenses against cross-site attacks.


 The most reliable approach for identifying supported HTTP methods is to send a request using each potential method to see how the server responds. Burp Intruder provides several ways to automate this time-consuming process.


#### Note

 Although sending an `OPTIONS` request sometimes causes the server to tell you which methods it supports, this information is often misleading or incomplete.


## Before you start

 Map the target website. For more information, see [Mapping the target website with Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/mapping).


## Enumerating supported methods for a single endpoint

 The simplest approach is to test a single endpoint. Supported methods are often fairly consistent across the website as they are usually the result of server-level configuration and the default configuration of back-end frameworks. That said, there may be some exceptions.


 To enumerate supported methods on a single endpoint:


1. Identify a relevant request. If you're testing the general configuration, use a `GET` request for the website's home page.
1. In the request, highlight the request method, then right-click the message and select **Send to
 Intruder**.

1.

Go to **Intruder**. Notice that the request method has been automatically marked as a payload position.

![Set HTTP method payload position](https://portswigger.net/burp/documentation/desktop/images/tutorials/identifying-supported-http-methods-1.png)

1. In the **Payloads** side panel, make sure that the payload type **Simple list** is selected.
1.
 Under **Payload configuration**, add a list of HTTP methods that you want to test:


  - If you're using Burp Suite Professional, open the **Add from list** dropdown menu and select the built-in **HTTP verbs** wordlist.
  -
 If you're using Burp Suite Community Edition, manually add a list of HTTP methods. Make sure you also include an arbitrary, non-existent method like `XYZ` to see how the server responds.


![Add HTTP verbs wordlist](https://portswigger.net/burp/documentation/desktop/images/tutorials/identifying-supported-http-methods-2.png)

1. Click ** Start attack**. The attack starts running in a new dialog.
1. When the attack is finished, study the responses to look for any noteworthy behavior. Don't rely purely on the HTTP status code as this may be misleading.

#### Note

 In Burp Suite Community Edition, the rate at which Burp Intruder sends requests is restricted, so your attack will take significantly longer to run.


## Enumerating supported methods for multiple endpoints

 You may encounter inconsistent behavior in response to the same methods on different parts of the site. This often occurs when websites are maintained by multiple teams of developers, who may configure their own custom request handling. You might also find that specific endpoints behave differently.


 To enumerate supported methods for multiple endpoints at once in order to identify these inconsistencies:


1. In Burp, go to the **Target** > **Site map** tab.
1. From the site map, right-click on the host and select **Copy URLs in this host**.
1. Send a `GET` request for the target host to Burp Intruder.
1.

In Burp Intruder, make the following settings:

  1. Select **Cluster bomb attack** from the attack type drop-down menu.
  1. Select the request method, then click **Add §** to mark this as a payload position.
  1. Select the request path, then click **Add §** to mark this as a second payload position.


![Setting request path and method as payload positions](https://portswigger.net/burp/documentation/desktop/images/tutorials/identifying-supported-http-methods-3.png)

1. In the **Payloads** side panel, make sure that payload position **1** and the payload type **Simple list** are selected.
1.

Under **Payload configuration**, add a list of HTTP methods that you want to test:

  - If you're using Burp Suite Professional, open the **Add from list** dropdown menu and select the
 built-in **HTTP verbs** wordlist.
  - If you're using Burp Suite Community Edition, manually add a list of HTTP methods. Make sure you also include an arbitrary, non-existent method to see how the server responds.

1.

Select position `2` from the **Payload position** drop-down list. Make the following settings:

  1. Under **Payload configuration**, click **Paste** to paste the list of URLs you copied from the site map.
  1.
 Under **Payload processing**, click **Add**. In the dialog that appears, create a **Match/replace** rule that replaces the scheme and hostname of your URLs with an empty string.


![Create a match and replace rule](https://portswigger.net/burp/documentation/desktop/images/tutorials/identifying-supported-http-methods-4.png)

  1. Under **Payload encoding**, deselect the **URL-encode these characters** checkbox. This prevents the forward-slash characters in the path from being encoded.

1. Click ** Start attack**. The attack starts running in a new dialog. Burp Intruder sends a request using each of the specified methods to each of the endpoints you selected.
1. When the attack is finished, study the responses to look for any noteworthy behavior. Don't rely purely on the HTTP status code as this may be misleading.

#### Note

 In Burp Suite Community Edition, the rate at which Burp Intruder sends requests is restricted, so your attack will take significantly longer to run.


#### Related pages

- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
- [LAB: Broken access control resulting from platform misconfiguration](https://portswigger.net/web-security/access-control#broken-access-control-resulting-from-platform-misconfiguration)
- [LAB: Information disclosure due to insecure configuration](https://portswigger.net/web-security/information-disclosure/exploiting#information-disclosure-due-to-insecure-configuration)
- [LAB: Bypass Lax SameSite restrictions using GET requests](https://portswigger.net/web-security/csrf/bypassing-samesite-restrictions#bypassing-samesite-lax-restrictions-using-get-requests)
- [LAB: CSRF validation depends on request method](https://portswigger.net/web-security/csrf/bypassing-token-validation)
