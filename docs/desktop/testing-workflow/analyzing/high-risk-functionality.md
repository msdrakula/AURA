> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/analyzing/high-risk-functionality

ProfessionalCommunity Edition

# Identifying high-risk functionality with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can use the Target site map and the Proxy history to identify high-risk functionality as you browse your target website. This enables you to plan how you will audit the application. Some key areas to investigate include:


- The application's core functionality.
- Security mechanisms such as authentication functions.
- Application behavior such as error messages, administrative functions, and off-site links.

## Before you start

 Start mapping the target website. For more information, see [Mapping the visible attack surface](https://portswigger.net/burp/documentation/desktop/testing-workflow/mapping/visible-attack-surface).


## Steps

 You can follow along with the process below using [ginandjuice.shop](https://ginandjuice.shop), our deliberately vulnerable demonstration site. To identify high-risk functionality:


1. Open Burp's browser and browse your target website.
1.

As you browse, use either or both the Target site map or HTTP history to help you identify high-risk functionality:

  - Review the **URL view** in the **Target** > **Site map** tab. For example, look for the colored circles on the nodes. Burp Scanner adds these when it detects an issue. By default, Burp Scanner passively audits traffic as you browse.
  - Review a list of the HTTP traffic in the **Proxy** > **HTTP history** tab.

![Using site map and browser](https://portswigger.net/burp/documentation/desktop/images/tutorials/high-risk-functionality-1.png)

1. When you identify functionality that you want to investigate further, keep the browser open and go to
 **Proxy** > **Intercept** in Burp.
1.

Configure your interception settings to intercept responses as well as requests:

  1. Click on  **Settings** in the top toolbar to open the **Settings** dialog.
  1. Go to **Tools** > **Proxy**.
  1. Under **Response interception rules**, select **Intercept responses based on the following rules**. By default, responses with a text content type header will now be intercepted.
  1. Close the **Settings** dialog.

1. In the **Intercept** tab, set the intercept toggle to **Intercept on**.
1. In Burp's browser, click through the function. Burp intercepts all requests and responses. The message details are
 shown in the **Intercept** tab.
1. Review each intercepted message for interesting content. You can also modify and send messages to see how this
 impacts the website.
1. To flag and store a message for further investigation, send it to Burp Organizer. Right-click the message and select
 **Send to Organizer**.

#### Note

 You can also use the Proxy HTTP history to investigate high-risk functionality. Click on any item in the HTTP history to view the request and response. You can't directly modify the message, but you can still send it to Burp Repeater to modify and investigate further.


#### Related pages

- [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser)
- [HTTP history](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history)
- [Target site map](https://portswigger.net/burp/documentation/desktop/tools/target/site-map)
- [Site map - URL view icons](https://portswigger.net/burp/documentation/desktop/tools/target/site-map#url-view-icons)
- [Proxy intercept](https://portswigger.net/burp/documentation/desktop/tools/proxy/intercept-messages)
- [Proxy settings - Request and response interception rules](https://portswigger.net/burp/documentation/desktop/settings/tools/proxy#request-and-response-interception-rules)
- [Burp Organizer](https://portswigger.net/burp/documentation/desktop/tools/organizer)
