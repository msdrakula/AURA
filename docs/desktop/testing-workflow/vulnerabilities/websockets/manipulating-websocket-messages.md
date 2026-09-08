> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/websockets/manipulating-websocket-messages

ProfessionalCommunity Edition

# Manipulating WebSocket messages with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

Finding vulnerabilities in WebSockets generally involves manipulating messages in ways that the application doesn't expect. For example, if you can modify existing messages or create new ones you may be able to deliver SQL injection or cross-site scripting exploits.

This tutorial explains how to modify and resend WebSocket messages in Repeater. You can follow along with the steps below using the [Manipulating WebSocket messages to exploit vulnerabilities](https://portswigger.net/web-security/websockets/lab-manipulating-messages-to-exploit-vulnerabilities) Web Security Academy lab.

## Steps

To modify and re-send WebSocket messages:

1. Browse around your target application to map its attack surface.
1. Go to **Proxy > WebSockets history**. This tab displays a table of any WebSocket messages that Burp's browser has exchanged with the target host.
1. Right-click a message that you want to re-send or modify (for example, an outbound chat message) and select **Send to Repeater**. Burp creates a new WebSocket tab in Repeater.
1. Go to Repeater, select the new tab, and click **Send**.
1. Check the **History** panel to confirm that the message was re-sent.
1. In the **Send WebSocket message** panel, modify the message. For example, you could send a proof-of-concept XSS attack at this point.
1. Click **Send** again.
1. Confirm that the modified message appears in the **History** panel as sent.

#### Related pages

- [Web Security Academy: Testing for WebSockets security vulnerabilities](https://portswigger.net/web-security/websockets)
- [Working with WebSocket messages in Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater/websocket-messages)
- [WebSockets history](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history)
- [Proxy settings: WebSocket interception rules](https://portswigger.net/burp/documentation/desktop/settings/tools/proxy#websocket-interception-rules)
