> Source: https://portswigger.net/burp/documentation/desktop/tools/repeater

ProfessionalCommunity Edition

# Burp Repeater

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Burp Repeater is a tool that enables you to modify and send an interesting HTTP or WebSocket message over and over.


 You can use Repeater for all kinds of purposes, for example to:


- Send a request with varying parameter values to test for input-based vulnerabilities.
- Send a series of HTTP requests in a specific sequence to test for vulnerabilities in multi-step processes, or vulnerabilities that rely on manipulating the connection state.
- Manually verify issues reported by Burp Scanner.
- Use AI-powered prompts to investigate behavior or look for security issues.

 Repeater enables you to work on multiple messages simultaneously, each in its own tab. Any modifications you make to a message are saved in the tab's history. You can easily manage large numbers of open tabs with the grouping function. For HTTP requests, you can also [add notes to each tab](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages#adding-notes-for-http-repeater-tabs).


 Burp AT can use Repeater for you as part of a Burp AT task, sending and analyzing requests on its own. For more information, see [Burp AT](https://portswigger.net/burp/documentation/desktop/burp-at).


#### Related pages

-  [Getting started: Reissuing requests with Burp Repeater](https://portswigger.net/burp/documentation/desktop/getting-started/reissuing-http-requests).

-

[Working with HTTP messages with Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages).


  -  [Using Burp AI in Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages/burp-ai-in-repeater).

  -  [Generating AI-powered explanations](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages/ai-explainer).

  -  [Custom actions](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages/custom-actions).

  -  [Sending HTTP requests in sequence](https://portswigger.net/burp/documentation/desktop/tools/repeater/send-group).


-  [Working with WebSocket messages with Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater/websocket-messages).

-

[Managing Burp Repeater tabs](https://portswigger.net/burp/documentation/desktop/tools/repeater/managing-tabs).


  - [Managing tab groups](https://portswigger.net/burp/documentation/desktop/tools/repeater/groups).

-

[Burp Repeater settings](https://portswigger.net/burp/documentation/desktop/settings/tools/repeater).


  - [Configuring tab-specific settings](https://portswigger.net/burp/documentation/desktop/tools/repeater/tab-settings).
