> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/access-controls/horizontal-access-controls

ProfessionalCommunity Edition

# Testing horizontal access controls

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 When a user logs in to an application, they usually only have access to their own functions and resources. If access controls are incorrectly set, a user can gain access to data and functionality that should only be available to other users.


 If you have credentials for two different accounts with identical privileges, you can test an application's horizontal access controls. You have two options to test access as different users:


- Use Burp Repeater to test individual endpoints.
- Use the Compare site maps function to automatically send multiple requests.

 You can follow along with the processes below using the lab [Using application functionality to exploit
 insecure deserialization](https://portswigger.net/web-security/deserialization/exploiting/lab-deserialization-using-application-functionality-to-exploit-insecure-deserialization) from our Web Security Academy.


## Before you start

 Get credentials for two different users with identical privileges. If you're using the lab, you can use the credentials `wiener:peter` and `gregg:rosebud`.


## Testing a specific endpoint

 To run a quick test on an individual endpoint:


1. Visit the target site and log in.
1. Access the functionality that you want to test.
1. In a new browser window, log in to the target site with a second set of credentials that have identical
 privileges.
1. Go to **Proxy > HTTP history**. Right-click a request that contains the first user's authenticated session cookie and select **Send to Repeater**.
1.

Find the second user's most recent request. Select the request and copy the session cookie.

![Session cookie second user individual endpoint](https://portswigger.net/burp/documentation/desktop/images/tutorials/horizontal-access-1.png)

1. Go to the **Repeater** tab. Paste the second user's cookie into the request, replacing the original session
 cookie.
1. Click **Send**.
1. Review the response to identify if horizontal access is possible. In the example, the account details automatically
 update. This indicates that you can't access the first user's account page using the second user's session.

## Testing across the entire site

 Testing horizontal access controls on a large number of endpoints can be time-consuming. Burp Suite can help you to automate this process across all the requests in the current site map:


1. Log in and map the application.
1. In a new browser window, log in to the target site with a second set of credentials that have identical privileges.
1.

Go to the **Proxy > HTTP history** tab. Select the second user's most recent request and copy the session cookie to use later.

![Session cookie second user](https://portswigger.net/burp/documentation/desktop/images/tutorials/horizontal-access-1.png)

1.

Create a session handling rule that adds the second user's session cookie to all requests sent from the Target tool:

  1. From the **Settings** dialog, go to **Sessions > Session handling rules** and click
 **Add**. The **Session handling rule editor** opens.
  1. Go to the **Scope** tab.
  1. Under **Tools scope**, select **Target** and deselect all other tools.
  1. Under **URL scope**, select **Use custom scope**, click **Add**, then enter the URL of the target site.
  1. Go to the **Details** tab to define the rule.
  1. Under **Rule actions**, click **Add**, then select **Set a specific
 cookie or parameter value**. The
 **Session handling action editor** opens.
  1.

Set the following details to add the second user's session cookie.

**Name**: session

**Value**: The cookie you copied from your second user's request
  1. Keep clicking **OK** to close all open dialogs. The rule is added to the list of session handling rules.

1.

Re-request the entire site map that you mapped as the first user:

  1. Go to the **Target > Site map** tab, right-click the target host, then select **Compare site maps**. The **Compare site maps** dialog opens.
  1. Select **Use current site map**, then click **Next**.
  1. Select **Use only selected branches**.
  1. Select **Request map 1 again in a different session context**.
  1. Keep the default settings for each of the remaining steps.
  1. Review the two site maps. Any differences are highlighted. Look for logged in requests that are identical - this may
 indicate that you accessed the first user's account using the second user's session.

#### Note

 You can also use the Autorize extension from the BApp Store to compare requests. This enables you to browse as one user and mirror the requests as a second user.


#### Related pages

- [Burp Repeater](https://portswigger.net/burp/documentation/desktop/tools/repeater)
- [Mapping the visible attack surface with Burp Suite](https://portswigger.net/burp/documentation/desktop/testing-workflow/mapping/visible-attack-surface)
- [Comparing site maps](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/comparing)
- [Session handling rule editor](https://portswigger.net/burp/documentation/desktop/settings/sessions/session-handling-rules)
