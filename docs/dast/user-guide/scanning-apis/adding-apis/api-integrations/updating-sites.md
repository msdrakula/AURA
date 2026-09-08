> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/updating-sites

DAST

# Updating your API sites

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 When your integration discovers new endpoints in an API, a banner appears on the site's **Endpoints** tab to alert you. You can then select the new endpoints for inclusion in future scans.

## When new endpoints are detected

 When your integration discovers new endpoints that aren't currently in scope for a site, a banner appears on the site's **Endpoints** tab. New endpoints are excluded from scans by default until you select them.

## Including new endpoints in scans

 To include new endpoints in scans:

1.
 Go to the site in the site tree and open the **Endpoints** tab.

1.
 Click **Edit**.

1.
 Select the checkboxes next to the new endpoints you want to include in scans.

1.
 Click **Save**.


 The selected endpoints are now in scope and will be included in future scans of this site.

#### Related pages

-

[Managing your APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations)
-

[Creating sites for added APIs](https://portswigger.net/burp/documentation/dast/user-guide/scanning-apis/adding-apis/api-integrations/creating-sites)
