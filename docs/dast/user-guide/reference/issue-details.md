> Source: https://portswigger.net/burp/documentation/dast/user-guide/reference/issue-details

DAST

# Issue details

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 When you [view the details of a particular issue](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/viewing-issues), the **Issues** window displays several tabs. Depending on the type of issue, the tabs displayed change. However, the **Advisory** tab is always present.


## Advisory tab

 This tab shows key information about the issue:


-
 The **Severity** of the issue.

-
 The **Confidence** that the issue is present.

-
 The host and URL path where the issue was found.

-
 The Advisory tab shows whether the issue was found by an [extension](https://portswigger.net/burp/documentation/dast/user-guide/extensions).


 The collapsible headings contain more detailed information about the issue. Note that only headings that apply to the particular issue are shown:


-
 Issue description.

-
 Issue detail.

-
 Issue background.

-
 Issue remediation.

-
 References.

-
 Vulnerability classifications.


## Request and response tabs

 The **Request** and **Response** tabs show a snippet of the HTTP requests and responses in which the issue was found. There might only be one request and one response, or there might be a series of interconnected requests and responses that lead to the issue.


 To help you to analyze the issue, key parts of each request and response are highlighted in red. These include payloads injected by the scanner and the string or regex in the response that confirms the vulnerability.


## Dynamic analysis

 For [DOM-based vulnerabilities](https://portswigger.net/web-security/dom-based), the **Dynamic analysis** tab shows the results of Burp Scanner's dynamic analysis of JavaScript on the site using an embedded headless browser. It loads HTTP responses into the browser, injects payloads into the DOM at locations that are potentially controllable by an attacker, and executes the JavaScript within the response.


 Burp Scanner also interacts with the page by creating mouse events to achieve as much code coverage as possible. It monitors dangerous sinks that could be used to perform an attack in order to identify any injected payloads that reach them.


 The tab shows:


-
 The values that were injected into a given source.

-
 The values that subsequently reached a sink.

-
 A stack trace at both the source and sink are also included.


 Wherever possible, the dynamic analysis also generates a proof of concept that you can use to reproduce the issue manually.


## Timeline tab

 The **Timeline** tab shows a chronological record of everything that has happened to an issue, so you can see its full history in one place. It records:


- Status changes, made by the scanner, by a user, or through the API.
- Re-tests.
- Comments.
- Integrations, such as tickets raised in Jira or GitLab.
- AI investigations.

 You can filter the timeline by event type: **All activity**, **Status changes**, **Re-tests**, **Comments**, **Integrations**, or **AI investigation**.


#### Note

 The timeline is only available on an issue record, which you can open at the global, folder, or site level. An issue viewed from a scan is a snapshot of that issue as it was found during the scan, so it does not have a timeline.


#### Related pages

-  [Viewing issue details](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/viewing-issues).
