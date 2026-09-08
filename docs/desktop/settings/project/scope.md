> Source: https://portswigger.net/burp/documentation/desktop/settings/project/scope

ProfessionalCommunity Edition

# Scope settings

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 The Scope settings page enables you to configure the following:


- [Target scope](https://portswigger.net/burp/documentation/desktop/settings/project/scope#target-scope).
- [Out-of-scope request handling](https://portswigger.net/burp/documentation/desktop/settings/project/scope#out-of-scope-request-handling).

 The **Scope** settings are project settings. They apply to the current project only.


## Target scope

 These settings determine which hosts and URLs are in scope for your current project. This changes the behavior of tools throughout Burp.


 For more information, see [Setting the target scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope)

## Out-of-scope request handling

 These settings determine whether Burp issues out-of-scope requests.


 Select **Drop all out-of-scope requests** to prevent Burp from issuing any out-of-scope requests. With this setting enabled, out-of-scope outgoing requests are dropped by Burp even if the browser explicitly requests them.


 You can select how Burp defines an out-of-scope request:


- **Use suite scope** - Select this option to use the [suite target scope](https://portswigger.net/burp/documentation/desktop/settings/project/scope#target-scope) to determine which requests to drop.
- **Use custom scope** - Select this option to define a custom scope to determine which requests to drop.

#### Related pages

 For more information on setting the suite-wide scope, see [Setting the target scope](https://portswigger.net/burp/documentation/desktop/tools/target/scope).
