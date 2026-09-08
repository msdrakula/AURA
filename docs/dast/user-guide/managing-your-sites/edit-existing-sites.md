> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/edit-existing-sites

DAST

# Editing existing sites

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can edit a site's details after it has been added to Burp Suite DAST.


#### Note

 In order to modify authentication details for an API site after the site has been saved, you need both the `View site application login details` and `Edit site application logins` permissions. This includes changing the specification upload method from a URL to a local file or vice versa. Note that admin users do not have these permissions by default.


 Users who have the `Edit site application logins` permission but not the `View site application login details` permission can see details of the authentication methods used in the specification but cannot see any details of the credentials provided.


 To edit an existing site:


1. Select **Sites** on the top menu to display the site tree page.
1. Select the site that you want to edit.
1. Select the **Details** tab.
1. Click **Edit**.
1. Make the required changes and click **Save**.

 The fields available when editing an existing site are the same as those available when adding a new site.


 If you edit an API site that uses a Postman Collection, you can add or update the environment file. Burp Suite DAST merges the environment variables with your Postman Collection, to help speed up your API testing setup.


#### Related pages

- [Adding a new site](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites).
- [Importing sites in bulk](https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/importing-sites-in-bulk).
