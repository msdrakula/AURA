> Source: https://portswigger.net/burp/documentation/desktop/tools/inspector/getting-started-inspector

ProfessionalCommunity Edition

# Getting started with the Inspector

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 The Inspector makes it easy to view and edit interesting items in HTTP messages. It automatically groups items from the selected request and response pair by category, such as HTTP headers, cookies, and parameters.


 The Inspector has the following useful features:


-

 It automatically decodes values so that you can read them more easily.

-

 It lets you edit encoded values in their plain text form, and then automatically re-applies the relevant sequence of encodings to the request.

-

 It lets you easily add, remove, and reorder items, without having to manually clean up the rest of the request afterward.


 You can find the **Inspector** in the side panel on many of the tabs in Burp Suite, next to the message editor.


![The message editor and Inspector](https://portswigger.net/burp/documentation/desktop/images/inspector.png)

 We recommend following the tutorial below to learn how to use the Inspector.


 For more detailed information, please see the [Inspector documentation](https://portswigger.net/burp/documentation/desktop/tools/inspector).


## Tutorial

 In this tutorial, you'll use the Inspector to work with encoded values more easily.


### Step 1: Access the lab

 Open Burp's browser and use it to access the following URL:
 `https://portswigger.net/web-security/deserialization/exploiting/lab-deserialization-modifying-serialized-objects`

 Click **Access the lab** and log in to your PortSwigger account if prompted. This opens your own instance of a deliberately vulnerable blog website.


### Step 2: Log in to a user account

 Click **My account** and log in using the following credentials: `wiener:peter`

### Step 3: Use the Inspector to examine the request

 Go to the **Proxy > HTTP history** tab, and select the `GET /my-account` request.


 In the **Inspector**, expand the **Request Cookies** section.


 To drill down into the `session` cookie, click **>**. Notice that the Inspector automatically performs the correct sequence of decoding steps to the value of the cookie. In this case, it shows the result after URL decoding, and then after the subsequent Base64 decoding.


![Decoding a cookie](https://portswigger.net/burp/documentation/desktop/images/inspector-cookie-decoding.png)

 Right-click the `GET /my-account` request and select **Send to Repeater**.


 Go to the **Repeater** tab, and drill down into the `session` cookie again.


### Step 4: Use the Inspector to edit the cookie

 In Repeater, you can modify the decoded value. For example, in the **Decoded from Base64** field, change `wiener` to `administrator`, then click
 **Apply changes**.


 Notice that the Inspector automatically re-applies the correct sequence of encodings when inserting the modified value into the request.


![Using the inspector to edit a cookie](https://portswigger.net/burp/documentation/desktop/images/inspector-edit-cookie.png)

### Step 5: Using the selection widget

 You can also select arbitrary strings in the message editor to examine them in the Inspector.


 Go back to the Inspector and highlight the value of the `session` cookie manually. Notice that the Inspector automatically decodes your selection in the same way as before.


 This is especially useful for working with substrings and non-standard data structures that the Inspector is unable to parse automatically.


![Inspector selection widget](https://portswigger.net/burp/documentation/desktop/images/inspector-selector-widget.png)

## Learn more about the Inspector

 You have now learned how to use the Inspector to work with encoded values more easily. To learn more, refer to the links below:


-

[Side panel settings](https://portswigger.net/burp/documentation/desktop/settings/ui/side-panel).

-

[Full documentation for the Inspector](https://portswigger.net/burp/documentation/desktop/tools/inspector).


 For more advanced users, the Inspector also provides some unique features for [working with HTTP/2 requests](https://portswigger.net/burp/documentation/desktop/http2) in a way that lets you exploit a number of HTTP/2-exclusive request smuggling vulnerabilities. We've covered these extensively on the [Web Security Academy](https://portswigger.net/web-security/request-smuggling/advanced).
