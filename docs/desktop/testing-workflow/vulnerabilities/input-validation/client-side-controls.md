> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/client-side-controls

ProfessionalCommunity Edition

# Bypassing client-side controls

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Some applications rely on measures on the client side to control the data that they submit to the server. This can lead to significant security flaws, as the user has full control over the client. To restrict user input, applications use a range of mechanisms that prevent the user from directly seeing and modifying data.


 You can use Burp Proxy to intercept a request, then modify data in the HTTP request before forwarding it to the server. This enables you to change data in a way that may not have been possible in the application.


## Steps

 You can follow the processes below using the lab [Excessive trust in client-side controls](https://portswigger.net/web-security/logic-flaws/examples/lab-logic-flaws-excessive-trust-in-client-side-controls).


 To view and modify data in Burp in order to bypass client-side controls:


1. In Burp's browser, access the page for the function that you want to test. If you're using the lab,
 access the leather jacket product page.
1. In **Proxy > Intercept**, set the intercept toggle to **Intercept on**.
1. In Burp's browser, click through the function. Burp intercepts each request. The message details are shown in the **Intercept** tab.
1.

In the **Intercept** tab, review each request for interesting content. For example, look
 for parameters that are hidden on the web page. In the lab example, there is a hidden price
 parameter.

![Hidden price parameter](https://portswigger.net/burp/documentation/desktop/images/tutorials/bypassing-client-side-controls-1.png)

1.

Edit the message. If you're following the lab, edit the price.

![Edit the hidden price parameter](https://portswigger.net/burp/documentation/desktop/images/tutorials/bypassing-client-side-controls-2.png)

1. Select the request and click **Forward** to send it.
1.

Determine whether the modification successfully updated on the server-side:

  - In **Proxy > HTTP history**, select the request you edited from the **Edited** column. Review the response.
  - Review the website. In the lab example, view the cart to see that the price has successfully changed to $0.01.

#### Note

 You can also use Burp Repeater to modify data in a request before forwarding it to the server. Send any message to Repeater from the HTTP history or Target site map. You can use Repeater to modify and send the message over and over.


#### Related pages

- [Burp's browser](https://portswigger.net/burp/documentation/desktop/tools/burps-browser)
- [Proxy intercept](https://portswigger.net/burp/documentation/desktop/tools/proxy/intercept-messages)
- [Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater)
- [HTTP history](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history)
