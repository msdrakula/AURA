> Source: https://portswigger.net/burp/documentation/dast/user-guide/scanning-web-apps/importing-sites-in-bulk

DAST

# Importing sites in bulk

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 Burp Suite DAST's bulk site upload feature makes it easier for you to add large numbers of sites to the system.


#### Note

Your scanning machines must be able to access the sites you want to scan. For information on allowing access, see [Configuring your environment network and firewall settings](https://portswigger.net/burp/documentation/dast/setup/self-hosted/standard/network-firewall-config).

 When uploading sites in bulk, you first need to prepare a CSV template with the site information, which you can then upload to Burp Suite DAST. To help you, we provide a sample template that you can edit.


#### Note

 Bulk upload is only available for web app sites.


## Preparing the import CSV file

 To prepare the upload CSV file:


1. Click **Sites** in the menu bar to display the site tree.
1. Click **Import sites**. A dialog box is displayed.
1. Click  **Download CSV template** to save the `sites-template.csv` file to your browser's default download location.
1. Open the `sites-template.csv` file in a spreadsheet or text editor.
1. Add your sites to the template in the same format as the example.
1. Delete the first three rows of the file (i.e. the rows containing the field titles, the instruction text, and the sample site).
1. Save the file in `.csv` format.

## Uploading the CSV file

 To upload the CSV file:


1. Click **Sites** in the menu bar to display the site tree.
1. Click **Import sites**. A dialog box is displayed.
1. Click **Choose file** to and select the file you want to upload.
1. Click **Continue**.
1. From the **Add sites to existing folder** drop-down menu, select the folder that you want to add the sites to.
1. Click **Import**.

 Burp Suite DAST imports the sites in the file and adds them to the selected folder.


#### Related pages

- [Adding new sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites) - explains how to add new sites individually.
- [Editing existing sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/edit-existing-sites).
