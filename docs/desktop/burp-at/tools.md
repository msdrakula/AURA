> Source: https://portswigger.net/burp/documentation/desktop/burp-at/tools

Professional

# Tools

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Burp AT has direct access to Burp's tools, with no setup or integrations to configure. Because it uses the same tried-and-tested tools you use manually, Burp AT can cope with edge cases, malformed traffic, and unusual application behavior that could derail an improvised script.


## Tool list

 Burp AT's tooling is organized into groups. This helps it to work efficiently, and enables you to set fine-grained permissions over what it can do autonomously.


 The tools are grouped as follows:


-

**Scanning and auditing** - crawl a site, audit requests for vulnerabilities, run a full scan, and pause or resume a running task.
-

**Scope and site map** - add to or remove from scope, inspect the current scope, add site map entries, import API definitions, and list the issues found for a URL.
-

**Requests and responses** - send and replay requests, inspect requests and responses, search across your captured traffic, and compare two messages.
-

**Issues** - raise, inspect, edit, and delete issues.
-

**Scan checks** - create, update, delete, and list custom scan checks.
-

**Datasets** - retrieve, search, and summarize Burp's data, such as proxy history, the site map, scan issues and definitions, scan tasks and configurations, and Burp Organizer entries. Burp AT can also add notes and status updates to the rows it retrieves, and stop a bulk request send while it's running.
-

**Configuration** - create and edit named scan configurations, and view saved scan or login configurations.
-

**Scripting** - run custom or packaged skill scripts via Burp, update a saved script, and inspect a script's result.
-

**Fuzzing** - fuzz a request with Burp Intruder and manage payload lists.
-

**Collaborator** - generate Burp Collaborator payloads and poll or inspect interactions.
-

**Utilities** - encode, decode, or hash data, look up the Montoya API, and send messages to another Burp tool.

## Controlling tool autonomy

 Burp AT chooses which tools it uses at each step, but you control what those tools are allowed to do without asking for permission. To manage tool autonomy, go to **Settings > Tools** to open **Tool settings**, which has two tabs:


-

**Enabled tools** - turn individual tools on or off. Burp AT never uses a tool you've disabled in any of its autonomy modes.
-

**Manual mode** - set each tool to **Ask** or **Act**. These per-tool permissions apply only when Burp AT is in **Manual** mode.

 For more information on autonomy modes and how approvals work, see [Configuring autonomy](https://portswigger.net/burp/documentation/desktop/burp-at/permissions).


#### Next step

-

[Skills](https://portswigger.net/burp/documentation/desktop/burp-at/skills)
