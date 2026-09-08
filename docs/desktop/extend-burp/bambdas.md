> Source: https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas

ProfessionalCommunity Edition

# Bambdas

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

Bambdas are scripts that you can run directly from Burp Suite's interface. They enable you to quickly personalize various tasks, such as creating custom match-and-replace rules, table columns, and filters.

 You can create your own scripts and save them to your Bambda library for easy reuse. You can also import existing scripts shared by others or downloaded from our [ Bambdas GitHub repository](https://github.com/PortSwigger/bambdas). Once a script is in your Bambda library, you can easily reuse it across Burp and in different projects.

#### Warning

 Bambda scripts can run arbitrary code. For security reasons, please be cautious when using scripts from unverified sources.


#### In this section

-

[Managing scripts in your Bambda library](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/managing)
-

[Importing scripts into your Bambda library](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/importing)
-

[Creating scripts](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating)

## Feature specific instructions

 Many tools in Burp enable you to apply scripts directly. For more information, see the feature-specific instructions.

### Filtering tables

For instructions on how to use scripts for filtering tables, see the following pages:

-

[Filtering the HTTP history with scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/scripts)
-

[Filtering the WebSockets history with scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/scripts)
-

[Configuring the Logger view filter with scripts](https://portswigger.net/burp/documentation/desktop/tools/logger/filter/view-with-scripts)
-

[Configuring the Logger capture filter with scripts](https://portswigger.net/burp/documentation/desktop/tools/logger/filter/capture-with-scripts)
-

[Filtering the site map with scripts](https://portswigger.net/burp/documentation/desktop/tools/target/site-map/scripts)

### Adding custom columns

Professional For instructions on how to use scripts for adding custom columns to tables, see the following pages:

-

[Adding custom columns in the HTTP history](https://portswigger.net/burp/documentation/desktop/tools/proxy/http-history/custom-columns)
-

[Adding custom columns in the WebSockets history](https://portswigger.net/burp/documentation/desktop/tools/proxy/websockets-history/custom-columns)
-

[Adding custom columns in Logger](https://portswigger.net/burp/documentation/desktop/tools/logger/custom-columns)

### Adding custom scan checks

Professional For instructions on how to use scripts for creating and adding custom scan checks, see the following pages:

- [Custom scan checks](https://portswigger.net/burp/documentation/desktop/extend-burp/custom-scan-checks)
- [Adding custom scan checks to scans](https://portswigger.net/burp/documentation/desktop/running-scans/custom-checks)

### Adding custom actions

Professional Custom actions are tasks that you can apply to HTTP messages in Burp Repeater to extract, transform, and analyze data.

 For instructions on how to create custom actions, see [Custom actions](https://portswigger.net/burp/documentation/desktop/tools/repeater/http-messages/custom-actions).

 For guidance on writing custom actions, see [Writing custom actions](https://portswigger.net/burp/documentation/desktop/extend-burp/bambdas/creating/writing-custom-actions).

### Adding match and replace rules

Professional Match and replace rules automatically replace parts of HTTP messages as they pass through the proxy.

For instructions on how to create HTTP match and replace rules with scripts, see [Creating HTTP match and replace rules with scripts](https://portswigger.net/burp/documentation/desktop/tools/proxy/match-and-replace/scripts).

#### Related pages

-
 To learn what a Bambda is and see a couple of filter script examples, watch the [Burp Suite Shorts | Bambdas](https://www.youtube.com/watch?v=neQpukwW43g) video on YouTube.

-

To learn more about how to use the different script types, see the following videos on YouTube:

  -  [Burp Suite #Shorts | Bambda table filters](https://www.youtube.com/watch?v=EYSsd2I7qcs)
  -  [Burp Suite #Shorts | Bambda table customization](https://www.youtube.com/watch?v=QyME5blj3e4)
  -  [Burp Suite: Introducing Custom actions](https://www.youtube.com/watch?v=u3GX4LgMdHQ)

-

Join our [PortSwigger Discord](https://discord.com/invite/portswigger) to chat with the
 community in our `#bambdas` channel - get tips, share ideas, and stay up-to-date with the
 latest Bambdas developments.
