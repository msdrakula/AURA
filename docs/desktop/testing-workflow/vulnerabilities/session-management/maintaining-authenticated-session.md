> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/vulnerabilities/session-management/maintaining-authenticated-session

ProfessionalCommunity Edition

# Maintaining an authenticated session

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 When testing, some actions may result in an application terminating your session. For example, an application may automatically log you out if you submit suspicious input. This may prevent you from performing actions such as fuzzing with Burp Intruder.


 Burp enables you to configure a session handling rule to automatically log back into an application. The session handling rule determines whether a session is valid. If it's invalid, it will run a macro to update the session cookies and log back in.


 You can follow along with the process below using `ginandjuice.shop`, our deliberately vulnerable demonstration site. The process consists of three steps:


1.

Identifying a valid login expression.
1.

Configuring a session handling rule.
1.

Checking the session handling rule.

## Identifying an invalid login expression

 Before you configure a session handling rule, you need to identify how the target site behaves when the session is invalid.


1.

In Burp's browser, log in to the target site using valid credentials. If you're using `ginandjuice.shop`, the credentials are `carlos:hunter2`.
1.

Go to a page that requires you to be logged in to access it. If you're using `ginandjuice.shop`, visit **My Account**.
1.

Log out.
1.

Try to get back to **My Account** without logging in. If you're using `ginandjuice.shop`, notice that you are redirected to the login page instead.
1.

In Burp, go to the **Proxy > HTTP history** tab to identify the behavior of the target site when an unauthorized user tries to access a restricted page. If you're using `ginandjuice.shop`, trying to access **My Account** when you're not logged in results in a 302 redirect to `/login`.

## Configuring a session handling rule

 To configure a session handling rule that enables you to maintain an authenticated session:


1.

Click **Settings** to open the **Settings** dialog.
1.

Under **Sessions > Session handling rules**, click **Add**. The session handling rule editor opens.
1.

Go to the **Scope** tab. Select the tools and URLs that you want the rule to apply to. The default tool scope and the suite URL scope are suitable for most use cases.
1.

Go to the **Details** tab. Add a unique rule description.
1.

Under **Rule actions**, click **Add**, then select **Check session is valid** from the drop-down menu. The session handling action editor opens.

![Open session handling editor](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/maintain-active-session/active-session-1.png)

1.

Under **Inspect response to determine session validity**, specify the expression that is found in an invalid login response. This should be the expression you identified earlier. Also, specify the aspects of each in-scope response that Burp should inspect for the expression:

  -

**Location(s)** - Select the locations in the response that you want Burp to inspect. If you're using `ginandjuice.shop`, select **URL of redirection target**.
  -

**Look for expression** - Specify the expression that is found in a valid login response. If you're using `ginandjuice.shop`, enter `login`.
  -

**Match type** - Select whether the expression is a literal string or regex. If you're using `ginandjuice.shop`, select **Literal string**.
  -

**Case-sensitivity** - Select whether the expression is case-sensitive or insensitive. If you're using `ginandjuice.shop`, select **Insensitive**.
  -

**Match indicates** - Select whether a match indicates that the session is valid or invalid. If you're using `ginandjuice.shop`, select **Invalid session**.

1.

Under **Define behavior dependent on session validity**, select **If session is invalid, perform the action below > Run a macro**.
1.

Click **Add**. The **Macro editor** and **Macro recorder** dialogs open.
1.

In the **Macro recorder** dialog, select the login requests, then click **OK**. If you're using `ginandjuice.shop`, select the `GET /login` and the two `POST /login` requests.

![Select macro requests](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/maintain-active-session/active-session-2.png)

1.

Click **OK** to close all open dialogs. The rule is added to the list of session handling rules.

## Checking the session handling rule

 It's a good idea to check that the session handling rule works. To do this:


1.

In Burp's browser, log out of the website.
1.

In **Proxy > HTTP history**, identify a request for a page that you need to be logged in to access. For example, if you're using `ginandjuice.shop`, you can use a `GET /my-account` request. The page should contain a session cookie that is now invalid.
1.

Right-click the request and select **Send to Repeater**.
1.

Go to the **Repeater** tab and send the request. Notice that the session cookies automatically update.
1.

Review the response to confirm that you've logged in successfully.

![Review the response to check valid session](https://portswigger.net/burp/documentation/desktop/testing-workflow/images/maintain-active-session/active-session-3.png)

#### Note

 If Repeater is set to never follow redirects you will need to click **Follow redirect to complete the login sequence**.


 For more information on configuring redirects in Repeater, see [Repeater settings - Redirects](https://portswigger.net/burp/documentation/desktop/settings/tools/repeater#redirects).


#### Related pages

- [Sessions settings](https://portswigger.net/burp/documentation/desktop/settings/sessions)
- [Session handling rule editor](https://portswigger.net/burp/documentation/desktop/settings/sessions/session-handling-rules)
- [Macro editor](https://portswigger.net/burp/documentation/desktop/settings/sessions/macros)
- [Session handling tracer](https://portswigger.net/burp/documentation/desktop/settings/sessions#session-handling-tracer) - Use the session handling tracer to troubleshoot your session handling configuration.
