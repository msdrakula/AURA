> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/ssrf/testing-for-blind-ssrf

ProfessionalCommunity Edition

# Testing for blind SSRF with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Blind server-side request forgery (SSRF) is a vulnerability that allows an attacker to induce an application to send HTTP requests to a specified URL, but no response is returned to them.


To detect blind SSRF vulnerabilities with out-of-band testing, you can use Collaborator to inject a domain into a request that attempts to trigger an out-of-band interaction with your target application. Burp then monitors the Collaborator server for any out-of-band interactions with that domain.


 If Collaborator detects that your application has sent a request to the inserted domain, that means it's vulnerable to SSRF.


## Steps

 You can follow the tutorial below by using the [Blind SSRF with out-of-band detection lab](https://portswigger.net/web-security/ssrf/blind/lab-out-of-band-detection) from our Web Security Academy.


 To test for blind SSRF with Burp Suite:


1.
 Go to **Proxy > HTTP history**. Identify a request in which you want to insert a Collaborator payload. For the lab, use a request that includes the `productId` parameter.

1.
 Right-click the request and select **Send to Repeater**.

1.
 Go to the **Repeater** tab.

1.
 Right-click where you want to insert a Collaborator payload and select **Insert Collaborator payload**. In the lab, replace the domain in the `Referer` header with a Collaborator payload.

1.
 Click **Send**.

1.
 Go to the **Collaborator** tab and click **Poll now**. The **Collaborator** tab lists any interactions your target application initiated with the Collaborator server.


#### Related pages

- [What is SSRF?](https://portswigger.net/web-security/ssrf)
- [Burp Collaborator](https://portswigger.net/burp/documentation/collaborator)
- [Deploying a private collaborator server](https://portswigger.net/burp/documentation/collaborator/server/private)
