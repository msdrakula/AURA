> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/session-management/session-timeout

ProfessionalCommunity Edition

# Determining the session timeout

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 When a user doesn't use an application for a certain amount of time, most applications will automatically log out the user and destroy their session.


 To determine how long it takes for a session to timeout, you can use Burp Intruder to issue the same request multiple times with increasing delays. This enables you to test compliance with security standards that require applications to timeout within a specified period. A longer timeout gives an attacker more time to use or guess a session token.


## Steps

 You can follow along with the process below using [ginandjuice.shop](https://ginandjuice.shop), our deliberately vulnerable demonstration site.


 To determine the session timeout:


1. In Burp's browser, log in to your target website. If you're using `ginandjuice.shop`, the correct credentials are `carlos:hunter2`.
1. Go to **Proxy > HTTP history**. Identify a logged-in request and send it to Intruder.
1. Go to **Intruder**.
1. In the **Payloads** side panel, under **Payload type**, select **Null payloads**.
1. Under **Payload configuration**, select **Continue indefinitely**.
1. Click on the ** Resource pool** tab to open the **Resource pool** side panel.
1. Select **Create new resource pool**.
1.

In the resource pool settings, select **Delay between requests**, then **Increase
 delay in
 increments of ___ milliseconds**. Add a delay time.

![Resource pool for session timeout](https://portswigger.net/burp/documentation/desktop/images/tutorials/determining-session-timeout-1.png)

1. Click ** Start attack**. The attack starts running in a new dialog. Intruder repeatedly sends the requests, with an
 increasing delay between requests.
1. With the attack results dialog open, click **Columns** on the top-level menu and select **Time of day**. A **Time of day** column is added to the results table. Sort the results by this column.
1. Review the responses as the attack progresses. Identify the first request that indicates the session is invalid. For example, look for a redirection to a login page.
1. To determine the session timeout, identify the time difference between the logged-out request and the previous request.

#### Note

 This attack may take some time. To continue the attack in the background, close the results dialog and click
 **Continue attack in background**. The attack is added to the **Tasks** panel on the
 **Dashboard**.


#### Related pages

[Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder)
