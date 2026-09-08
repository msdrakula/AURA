> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/mapping/hidden-content/hostname-discovery

ProfessionalCommunity Edition

# Discovering supported hostnames with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 You can use Burp Intruder to discover additional, hidden hosts that are in scope but aren't explicitly linked from the initial set of domains you're testing. This technique involves sending requests to the same IP address while modifying the `Host` header to look for different back-end systems. It enables you to discover additional access points and attack surface, including:


- Internal systems that shouldn't be publicly accessible.
- Hosts that may have been forgotten about and therefore aren't regularly updated.

## Steps

 You can follow along with the process below using [portswigger-labs.net](https://portswigger-labs.net), our deliberately vulnerable sandbox domain. To enumerate additional supported hostnames:


1. Send a request for the main domain you want to investigate to Burp Intruder. For example, `http://portswigger-labs.net`.
1. Go to **Intruder**. The request is displayed in a new attack tab.
1. At the beginning of the `Host` header value, add a placeholder subdomain. For example, `x.portswigger-labs.net`.
1.

 Highlight the placeholder subdomain and click **Add §** to mark it as a payload position.

1.

 Uncheck **Update Host header to match target**.


![Set payload position](https://portswigger.net/burp/documentation/desktop/images/tutorials/enumerating-subdomains-1.png)

1. In the **Payloads** side panel, under **Payload configuration**, add a list of potential subdomain names. If you are
 using Burp Suite Professional, you can select from a list of built-in wordlists. The **Directories - short** list is suitable in
 this case.
1. Click ** Start attack**. An attack results window opens. Intruder sends a request for each payload in the list, with the
 payload in place of the placeholder subdomain.
1.

 Click the column headers to sort the responses. Identify any inconsistent items. For example, the response for `http://staff.portswigger-labs.net` has a different length from the other responses.


![Identifying inconsistent results](https://portswigger.net/burp/documentation/desktop/images/tutorials/enumerating-subdomains-2.png)

1. Right-click the subdomain you enumerated. Select **Request in browser** > **In original session** to copy a URL for the request.
1. Open Burp's browser and access the URL. For example, `http://staff.portswigger-labs.net` renders a login form.

#### Related pages

- [Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder).
- [Predefined payload lists](https://portswigger.net/burp/documentation/desktop/tools/intruder/configure-attack/payload-lists).
- To learn about an advanced approach that uses response timing to reveal whether a reverse proxy routes requests to internal hostnames, see [Listen
 to the whispers: web timing attacks that actually work](https://portswigger.net/research/listen-to-the-whispers-web-timing-attacks-that-actually-work#reverse)
