> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/access-controls/using-match-and-replace

ProfessionalCommunity Edition

# Spoofing your IP address using Burp Proxy match and replace

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

[Burp Proxy](https://portswigger.net/burp) allows you to configure [match and replace rules](https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace) that automatically modify your requests and responses while you explore the target application as normal using Burp's browser. This enables you to add, remove, or modify headers in requests or responses, for example.


 There are a number of uses for this, including potentially spoofing your IP address. In some cases, this may allow you to trick a server into believing that you belong to its local network, which could enable you to communicate with internal infrastructure that is otherwise inaccessible.


 In this tutorial you'll learn how to:


- Set match and replace rules in Burp Suite.
- Use match and replace rules to add a header to your requests.
- Spoof your IP address to compromise a vulnerable server that uses a form of IP-based authentication.

#### Note

 Burp's browser is an easy way to proxy HTTP traffic - even over the encrypted HTTPS protocol. There is no setup required - simply go to the **Proxy** tab, click **Open Browser**, and make sure that the intercept toggle is set to **Intercept off**.


## Step 1: Open the lab

 Open Burp's browser and access the following lab:
 `https://portswigger.net/web-security/information-disclosure/exploiting/lab-infoleak-authentication-bypass`

## Step 2: Attempt to access the admin panel

 Try to access the admin panel at `/admin`.


 Note that you are prevented from doing so as this is only accessible to local users.


![A Web Security Admin page which you are unable to access](https://portswigger.net/burp/documentation/desktop/images/using-match-and-replace/unable-to-access-admin-interface.png)

 For simplicity, let's assume that you've subsequently identified that the server is using a custom HTTP header, `X-Custom-IP-Authorization`, to determine your IP address.


#### Note

 In the wild, the de-facto standard header `X-Forwarded-For` is often used for this purpose, but you may encounter websites that use different custom headers. These are normally applied to your requests by an intermediary server, such as a load balancer or other reverse proxy belonging to a CDN, for example.


## Step 3: Add a custom match and replace rule

 To add a custom match and replace rule:

1.
 Go to **Proxy > Match and replace**.

1.
 Under **HTTP match and replace rules**, click **Add**. The **Add match/replace rule** dialog opens.

1.
 Under **Type**, make sure that **Request header** is selected.

1.
 Leave the **Match** field empty. This ensures that Burp appends a new header to requests rather than replacing an existing one.

1.

In the **Replace** field, enter the following: `X-Custom-IP-Authorization: 127.0.0.1`
1.
 Click **Test**.

1.
 Under **Auto-modified request**, notice that Burp has added the `X-Custom-IP-Authorization` header to the modified request.

1.
 Click **OK**.


 Burp Proxy will now add this header to every request you make in Burp's browser.


![Match and replace rule](https://portswigger.net/burp/documentation/desktop/images/using-match-and-replace/adding-a-custom-match-and-replace-rule.png)

## Step 4: Try to access the admin panel again

 In Burp's browser, try browsing to `/admin` again. Observe that you can now access the admin page and delete `carlos` to solve the lab.


![Deleting Carlos from the admin panel](https://portswigger.net/burp/documentation/desktop/images/using-match-and-replace/delete-carlos-to-complete-the-web-security-academy-lab.png)

 In Burp, you can confirm that the header was added to your requests by checking them in the **Logger** tab:


![Viewing the automatically modified request in the Logger tab](https://portswigger.net/burp/documentation/desktop/images/using-match-and-replace/viewing-the-modified-request-in-logger.png)

 Alternatively, on the **Proxy > HTTP history** tab, you can use the drop-down menu to toggle between the original request that was sent by the browser, and the modified one that Burp Proxy forwarded to the server.


![Viewing the automatically modified request in the HTTP history tab](https://portswigger.net/burp/documentation/desktop/images/using-match-and-replace/viewing-a-custom-http-header-in-burp-proxy.png)

#### Note

 Although we've manually added a custom header in this case, Burp Suite provides a number of built-in match and replace rules to cover some of the most common use cases. You just need to enable them under **Proxy > Match and replace**.


## Summary and next steps

 Congratulations - now you know how to use Burp Proxy's match and replace rules, and have used them to spoof your IP address.


 To learn how you could have discovered the custom header we used to solve the lab, check out the [learning materials](https://portswigger.net/web-security/information-disclosure/exploiting.html#information-disclosure-due-to-insecure-configuration) on the Web Security Academy.


 For more general information on authentication, as well as other types of attack you can carry out using Burp Suite, see the [Authentication](https://portswigger.net/web-security/authentication/index.html) topic in the Web Security Academy.
