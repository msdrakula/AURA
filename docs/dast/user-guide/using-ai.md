> Source: https://portswigger.net/burp/documentation/dast/user-guide/using-ai

DAST

# Using AI

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp Suite DAST works with AI in two separate ways. You can use either one independently of the other. Their availability differs, so check the edition and deployment labels on each page.


## Burp AI features

 Burp AI is a set of AI-powered features built into Burp Suite DAST. These features help you to triage issues faster and to simplify scan setup. For more information, see [Burp AI](https://portswigger.net/burp/documentation/dast/user-guide/using-ai/burp-ai).


## MCP server

 The MCP server lets you connect your own AI client, such as Claude Code or Cursor, so that it can work with your scans and issues on your behalf. For the connection steps, see [Connecting your AI client to the MCP server](https://portswigger.net/burp/documentation/dast/user-guide/using-ai/mcp-server).


### Using your assistant with DAST

 You can use your own AI client to work with your scans and issues without opening the Burp Suite DAST interface. Your assistant chooses the right action based on what you ask, so you do not need to know the underlying tools. For example, you can ask it to do the following:


- Report which scans ran today, or whether a particular scan is still running.
- Show the high-severity issues for a site or folder, or summarize them for triage.
- Start a scan of a site, or as part of an AI-driven development or CI workflow.
- Mark an issue as a false positive or accepted risk.
- Report on your agent pool scanning capacity.

 Some actions are only partly available:


- Your assistant can generate an HTML report for a scan. PDF and XML reports are not available through the assistant.
- Your assistant can update a site's simple scope (seed URLs and in-scope or out-of-scope prefixes), but not advanced scope rules, and not sites whose targets come from an attached API definition.

 The set of actions grows over time. To see what is currently available, ask your assistant what it can do with Burp Suite DAST.


#### Note

 Your assistant can only use the same actions you have permission to perform in the Burp Suite DAST interface.


#### Related pages

-  [Configuring AI-enhanced scanning](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/configuring-ai-enhanced-scanning)
-  [Viewing AI-enhanced scan results](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/ai-enhanced-results)
-  [Using recorded logins](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/configure-authentication/recorded-logins)
-  [AI features](https://portswigger.net/burp/documentation/ai-features)
