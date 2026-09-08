> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/testing-for-stored-xss

ProfessionalCommunity Edition

# Testing for stored XSS with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 Stored XSS arises when an application receives data from an untrusted source and includes that data within its later HTTP responses in an unsafe way. In a stored XSS attack, the attacker places their exploit into the application, but the exploit only executes when a user visits a particular location. For example, an attacker may place an exploit in a blog post comment that later executes when a victim user visits the blog post.


 While Burp Scanner can detect stored XSS, you can also use Burp to manually identify linked input and output points in the application, then test these links to determine whether a stored XSS vulnerability is present.


## Steps

 You can follow along with the steps below using the [Stored XSS into HTML context with nothing encoded](https://portswigger.net/web-security/cross-site-scripting/stored/lab-html-context-nothing-encoded) Web Security Academy lab.


### Identifying linked input and output points

 To test for stored XSS with Burp Suite, you first need to identify points where user input is stored and then later displayed by the application:


1. Go to **Proxy > Intercept** and set the intercept toggle to **Intercept on**. Burp will now intercept each request and display the message details in the **Intercept** tab.
1.

 In Burp's browser, click around the website. As you browse, review each request in the **Intercept**
 tab:


  1. Identify data entry points that may be vulnerable to stored XSS.
  1. Submit a unique value into each of the identified data entry points.
  1. Click **Forward** to send the request.

1.

In **Proxy > HTTP history**, filter the messages to identify any responses that include the unique values:

  1. Click on the filter bar.
  1. Under **Filter** by search term, enter the unique value.
  1. Click **Apply**. The **HTTP history** tab now only shows messages that include the unique value.
  1. Review the messages in the **HTTP history** tab to identify any linked input and output points.

### Testing for stored XSS

 Once you've found linked input and output points, you can test these for stored XSS:


1. In **Proxy > HTTP history**, right-click the request that results in the input being
 stored. Select **Send to Repeater**. Do the same for the later request that results in the
 input being displayed.
1. Go to the **Repeater** tab.
1.

Create a group in Repeater that includes the two messages:

  1. Right-click a message tab in Repeater and select **Add tab to group >  Create tab group**.
  1. Add both messages to the folder.
  1. Click **Create**.
  1. Drag the messages to place the input message before the output message.

![Stored XSS Repeater group](https://portswigger.net/burp/documentation/desktop/images/tutorials/stored-xss-1.png)

1. In the input request, change the data entry point to a proof-of-concept XSS payload. For example, `<script>alert(1)</script>`.
1. Click the **Send** drop-down menu, then select **Send group in sequence (separate connections)**.
1. Click **Send group (separate connections)** to send the request.
1. Right-click on the output and select **Show response in browser**. Burp Suite displays a dialog containing a URL.
1. Copy and paste this URL into your browser to see if the proof of concept ran successfully. If you used the `alert()`
 function in your payload, you should see an alert pop-up.

#### Related pages

- Academy: [Stored XSS](https://portswigger.net/web-security/cross-site-scripting/stored)
- [Proxy intercept](https://portswigger.net/burp/documentation/desktop/tools/proxy/intercept-messages)
- [Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater)
- [Bypassing XSS filters by enumerating permitted tags and attributes](https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/bypassing-filters)
