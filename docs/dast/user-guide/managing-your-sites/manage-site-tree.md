> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/manage-site-tree

DAST

# Managing the site tree

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 As you add more sites to Burp Suite DAST, you may find it useful to organize those sites into folders. For example, you could group your sites based on their physical location.


 Each folder has its own dashboard, enabling you to view aggregated charts and statistics for all the sites in the folder. You can also use folders to restrict access to sites, to make sure that users can only access data for sites that are relevant to them.


 You can manage your folders and subfolders using the site tree, located on the **Sites** page.


 You can also add tags to your sites or folders. This enables you to organize sites any way you want. For example, you could create tags for the following:


-
 Geographic locations or time zones

-
 Specific teams

-
 Different levels of criticality


 For more information, see [Adding tags to sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/adding-tags-to-sites).


## Creating folders and subfolders

 To create a new folder:


1.

On the top menu, select **Sites** to display the site tree.
1.

Click **New folder**. The **New Folder** dialog opens.
1.

Enter a **Folder Name**. Make sure there are no other folders in the parent folder with the same name.
1.

Optionally, you can add a **Description** to your folder.
1.

If you want to add the folder to an existing folder, select a parent from the site tree under **Add to existing folder**.
1.

Click **Add** to add your new folder.

## Adding individual sites to a folder

 You can select a folder when you initially add a site using the **Site folder** field on the **Create a new site** page.


 To add an existing site to a folder:


1. On the top menu, select **Sites** to display the site tree.
1. Select the site you want to move.
1. Select the **Details** tab.
1. Click **Edit**.
1. Select the relevant folder from the **Site folder** field.
1. Click **Save** to add the site to the folder.

## Moving sites and folders in bulk

 To move multiple sites or folders at the same time:


1. On the top menu, select **Sites** to display the site tree.
1. Use the checkboxes in the list to select the sites you want to add.
1. In the popup menu, click  **Move**.
1. In the **Select a destination folder** window, select the folder you want to add the sites to.
1. Click **Move** and then click **OK** to close the dialog.

 If you click **Move** on the dialog without selecting a destination folder then Burp Suite DAST moves the selected sites and folders to the root level of the site tree.


## Deleting sites and folders

 To delete an individual site or folder, click its  delete icon.


 To delete multiple sites or folders at the same time:


1. On the top menu, select **Sites** to display the site tree.
1. Use the checkboxes in the list to select the sites or folders you want to delete.
1. In the popup menu, click  **Delete**.
1. In the **Confirm delete** dialog, click **Delete**.

#### Related pages

- [Restricting access to sites](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/restricting-access-to-sites) - explains how to user folders to restrict site access for selected users.
- [Folder-level view](https://portswigger.net/burp/documentation/dast/user-guide/reference/folders) - explains the folder-level dashboard in more detail.
- [Add new sites](https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites) - explains how to add a new site to Burp Suite DAST.
- [Bulk actions reference page](https://portswigger.net/burp/documentation/dast/user-guide/reference/bulk-actions-site-tree).
