> Source: https://portswigger.net/burp/documentation/desktop/testing-workflow/mapping/hidden-content/automated-discovery

Professional

# Automated content discovery with Burp Suite

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 You can use Burp Suite Professional's automated content discovery tool to discover hidden directories, files, and other endpoints. The tool uses lists of common file and directory names to guess the names of hidden functionality. It also derives a naming scheme from the resources already identified and uses this to search for similarly named items. This enables you to discover additional attack surface.


## Before you start

 You need to populate the site map with some content before you use the content discovery tool. For a tutorial on how to map the target website, see
 [Mapping the visible attack surface](https://portswigger.net/burp/documentation/desktop/testing-workflow/mapping/visible-attack-surface).


## Steps

 You can follow along with the process below using [ginandjuice.shop](https://ginandjuice.shop), our deliberately vulnerable demonstration site. To scan for hidden content:


1. Go to **Target** > **Site map**.
1. Right-click on the root node for the domain.
1. Click **Engagement tools** > **Discover content**. The **Content discovery** dialog opens.
1. Click **Session is not running** to start the discovery session with default settings.

 By default, the tool adds discovered content to the Target site map. You can also view the site map that the content discovery tool builds in the dialog's
 **Site map** tab.


#### Note

 You can alternatively use Burp Intruder to discover hidden directories, files, and other endpoints. This gives you more manual control over the process.


#### Related pages

- [Content discovery](https://portswigger.net/burp/documentation/desktop/tools/engagement-tools/content-discovery).
- [Site map](https://portswigger.net/burp/documentation/desktop/tools/target/site-map).
- [Typical uses for Burp Intruder](https://portswigger.net/burp/documentation/desktop/tools/intruder/uses).
