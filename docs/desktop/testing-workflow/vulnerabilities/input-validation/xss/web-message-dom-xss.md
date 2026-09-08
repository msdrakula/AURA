> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/web-message-dom-xss

ProfessionalCommunity Edition

# Testing for web message DOM XSS with DOM Invader

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Web message DOM XSS occurs if the destination origin for a web message trusts the sender not to transmit malicious data in the message, and handles the data in an unsafe way by passing it into a sink.


 You can use DOM Invader to test applications for web message DOM XSS. DOM Invader enables you to log any messages that are sent via the `postMessage()` method, and modify and resend web messages.


 To learn more about sources and sinks, see [DOM-based vulnerabilities](https://portswigger.net/web-security/dom-based).


#### Note

 DOM Invader is pre-installed in Burp's browser. It's disabled by default as some of its features may interfere with your other testing activities.


## Before you start

-
 Enable DOM Invader. For more information, see [Enabling DOM Invader](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/enabling).

-
 Enable **Postmessage interception** from the DOM Invader settings menu. For more information, see [Main DOM Invader settings](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/settings/main).


## Steps

 You can follow the processes below using the lab [DOM XSS using web messages](https://portswigger.net/web-security/dom-based/controlling-the-web-message-source/lab-dom-xss-using-web-messages).


1.
 Use Burp's browser to visit your target website.

1.
 Right-click the browser window and select **Inspect**.

1.
 Select the **DOM Invader** tab and then select **Messages** from the right-hand panel. You can see the messages that DOM invader has flagged as exploitable.

1.

 Click each message to review it, and see if the `origin`, `data`, or `source` properties of the message are accessed by the client-side JavaScript:


  -
 If the `origin` property isn't accessed, it's likely that the origin isn't being validated.

  -
 If the `data` property isn't accessed, the message can't be exploited.

  -
 If the `source` property isn't accessed, it's likely the source (usually an iframe) isn't being validated.


 You can use the message information to craft an exploit. Use DOM Invader to send a modified web message:


1.
 From the **Messages** view, click on any message to open the message details dialog.

1.
 Review the message information to identify the type of sink the data ends up in.

1.

 Edit the **Data** field with an exploit that matches the sink type.


![Reviewing messages in DOM Invader](https://portswigger.net/burp/documentation/desktop/images/tutorials/dom-invader-web-message-xss.png)

1.
 Click **Send**.

1.

 If you find an exploitable vulnerability, use DOM Invader to generate a proof of concept:


  -
 Select the vulnerable message to open the message details dialog.

  -
 Modify the values as required for your exploit.

  -
 Click **Build PoC** to save the HTML to your clipboard.


#### Related pages

- [Testing for DOM XSS using web messages](https://portswigger.net/burp/documentation/desktop/tools/dom-invader/web-messages.html)
- [DOM Invader](https://portswigger.net/burp/documentation/desktop/tools/dom-invader)
