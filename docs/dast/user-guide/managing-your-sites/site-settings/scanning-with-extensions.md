> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/site-settings/scanning-with-extensions

DAST

# Scanning with extensions in Burp Suite DAST

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Once your administrator has [added an extension](https://portswigger.net/burp/documentation/dast/user-guide/extensions/adding-extensions) to your library, you can apply it to one or more sites. The extension is used whenever a scan runs on that site.

## Applying extensions to sites

Apply extensions to your sites to have Burp Suite DAST use them whenever it runs a scan on that site.

### Applying extensions to an existing site

To apply extensions to an existing site:

1.

 From the **Sites** page, select the site you want to apply the extension to.

1.

 On the **Details** tab, click **Edit** .

1.

In **Scan settings**, go to the **Extensions** tab, then:

  -  **For BChecks:** go to the **BChecks** tab.

  -  **For BApps and custom extensions:** go to the **BApps & custom extensions** tab.


Extensions that your system administrator has added to Burp Suite DAST are listed on these tabs.
1.

 Select the extensions you want to apply to the site.

1.

 Click **Save**.


The selected extensions are applied to your site.

#### Note

Using extensions can increase the duration of your scans.

### Applying extensions to new sites

You can also apply extensions when you are [creating a new site in Burp Suite DAST](https://portswigger.net/burp/documentation/dast/user-guide/extensions/adding-extensions).

To apply extensions to a new site:

1.

 On the **Create a new site** page, in **Site settings**, go to the **Extensions** tab.

1.

 On the **Details** tab, click **Edit** .

1.

In **Scan settings**, go to the **Extensions** tab, then:

  -  **For BChecks:** go to the **BChecks** tab.

  -  **For BApps and custom extensions:** go to the **BApps & custom extensions** tab.


Extensions that your system administrator has added to Burp Suite DAST are listed on these tabs.
1.

 Select the extensions you want to apply to the site.

1.

 Finish creating your new site, then click **Save**.


The selected extensions are applied to your site.

#### Note

Using extensions can increase the duration of your scans.

### Removing extensions from sites

To remove an extension from a site:

1.

 From the **Sites** page, select the site you want to remove the extension from.

1.

 On the **Details** tab, click **Edit** .

1.

In **Scan settings**, go to the **Extensions** tab, then:

  -  **For BChecks:** go to the **BChecks** tab.

  -  **For BApps and custom extensions:** go to the **BApps & custom extensions** tab.


1.

 Remove the extensions you no longer want applied to the site.

1.

 Click **Save**.


The selected extensions are removed from your site.

## Applying extensions to folders

You can apply extensions at folder-level in Burp Suite DAST.
 These are inherited by any subfolders and sites inside the folder, meaning these extensions are used whenever scans are run on sites within this folder.

#### Note

It's easy to identify inherited extensions by the information banner that appears at the top of the **Details** tab for sites,
 and at the top of the **Scan settings** tab for folders.

Inherited extensions can be managed from the **Scan settings > Extensions** tab of the parent folder they are inherited from.

To apply an extension to a folder:

1.

 From the **Sites** page, select the folder you want to apply the extension to.

1.

In **Scan settings**, go to the **Extensions** tab, then:

  -  **For BChecks:** go to the **BChecks** tab.

  -  **For BApps and custom extensions:** go to the **BApps & custom extensions** tab.


 Extensions that your system administrator has added to Burp Suite DAST are listed on these tabs.

1.

 Select the extensions you want to apply to the folder.

1.

 Click **Save**.


The selected extensions are applied to your folder.

#### Note

Using extensions can increase the duration of your scans.

### Removing extensions from folders

To remove an extension from a folder:

1.

 From the **Sites** page, select the folder you want to remove the extension from.

1.

In **Scan settings**, go to the **Extensions** tab, then:

  -  **For BChecks:** go to the **BChecks** tab.

  -  **For BApps and custom extensions:** go to the **BApps & custom extensions** tab.


1.

 Remove the extensions you no longer want applied to the folder.

1.

 Click **Save**.


The selected extensions are removed from your folder.
