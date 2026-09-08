> Source: https://portswigger.net/burp/documentation/dast/user-guide/managing-your-sites/adding-tags-to-sites

DAST

# Adding tags to sites

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 You can create and add tags to organize your sites or folders. Once you've added tags to your site, you can do the following:

- Filter your sites and folders by tags
- Add descriptions to your tags, such as region, owner, or importance

#### Note

 If you apply a tag to a folder, you also apply it automatically to all the folder's sites and subfolders.


## Permissions required to manage tags

 To create, edit, or delete tags, you need to be assigned a role that has **Manage tags** permission. You need **Edit sites and folders** permission to add or remove tags from sites and folders.

 To view tags, you only need permission to view the site. No extra permissions are required.

 For more information about role-based permissions, see [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control).

## Creating tags

When you create a tag, you can choose the color and name. You can also add a detailed description.

1. On the **Sites** page, select the tick box for a site or folder in the site tree. The bottom menu appears.
1. From the bottom menu, select **Tags**.
1. Click **Create new tag**.
1. Enter a **Tag name**. You can also add a **Description**.
1. Select a color for your tag.
1. Click **Save**.

 The tag now appears alphabetically in the list of tags.

## Adding tags to a single site or folder

1. Go to the **Sites** page.
1. In the **Tags** column for your site or folder, click the  button.
1. Select one or more tags to add to your site or folder.
1. Click **Confirm**.

## Adding tags to multiple sites or folders

1. Go to the **Sites** page.
1. Select the tick box for the sites and folders that you want to tag. The bottom menu appears.
1. From the bottom menu, select **Tags**.
1. Select one or more tags.
1. Click **Confirm**.

 The tags now appear in the **Tags** column for your selected sites and folders.

## Filtering by tags

 If you use tags to filter sites and folders, you only see sites and folders that you have permission to view. You don't need specific permissions to be able to filter by tags.

1. Go to the **Sites** page.
1. Click the  **Filter** menu and select one or more tags to filter by.

The site list is now filtered by tagged sites, and folders that contain tagged sites.

#### Note

 You can schedule scans for your filtered sites and folders. All the sites in a tagged folder will be scanned, even if you have untagged individual sites.


## Untagging a single site or folder

 Untagging a site of folder removes the tag from that site, but doesn't remove it from your library of tags. Use this method If you want to untag a folder, but leave the tag on its sites and subfolders.

1. Go to the **Sites** page.
1. In the **Tags** column, click the tag you want to remove from the site or folder. If the site has more than one tag, click the **+n** label to see more tags.
1. Click  and then click **Untag site**.

## Untagging multiple sites or folders

You can untag multiple sites or folders. Use this method if you want to untag a folder, and all its sites and subfolders.

To untag multiple sites and folders:

1. Go to the **Sites** page.
1. Select the sites and folders that you want to untag.
1. From the bottom menu, click **Tags**.
1. Deselect your chosen tags.
1. Click **Confirm**.

## Editing tags

 You can edit the name and description for tags, and change the color. If you edit a tag, you change it for every site that has the tag.

1. Go to the **Sites** page.
1. In the **Tags** column, click the tag you want to edit. If the site has more than one tag, click the **+n** label to see more tags.
1. Click  and then click **Edit tag**.
1. Edit the tag.
1. Click **Save**.

## Deleting tags

If you delete a tag, you remove it from every site and folder. You also remove it from your library.

1. Go to the **Sites** page.
1. In the **Tags** column, click the tag you want to delete. If the site or folder has more than one tag, click the **+n** label to see more tags.
1. Click , then click **Delete tag**. A dialog box appears.
1. To confirm, click **Delete tag**.
