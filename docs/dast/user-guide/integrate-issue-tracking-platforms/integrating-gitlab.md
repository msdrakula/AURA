> Source: https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-gitlab

DAST

# Integrating Burp Suite DAST with GitLab

-

**Last updated: ** September 3, 2026
-

**Read time: ** 3 Minutes

 If you or your teams use GitLab, you can integrate it with Burp Suite DAST. Once configured, this enables you to raise GitLab issues from directly within Burp Suite DAST for any security issues found by your scans.


## Prerequisites

1.
 You have access to Burp Suite DAST as an administrator.

1.
 You have access to your GitLab instance as an administrator.

1.
 You have the **Maintainer** or **Owner** role for any GitLab projects you want to create issues on.


## (Recommended) Create a new GitLab user for the integration

 To integrate with GitLab, Burp Suite DAST must be linked to a specific GitLab user.


 We recommend creating a new GitLab user specifically for the integration. This allows you to control which projects are available for use in Burp Suite DAST - simply by adding your user as a **Reporter** to those projects.


## Generate a GitLab impersonation token

 A GitLab impersonation token allows Burp Suite DAST to raise GitLab issues as a specific user.


#### Note

 The latest version of GitLab adds a prefix to the personal access token, which means the token now exceeds our 20-character limit. You can use the following workaround to fix this:


1.
 Go to the following link: `https://gitlab.example.com/admin/application_settings/general#application_setting_personal_access_token_prefix`
1.
 Delete all the contents from the **Personal Access Token prefix** field.

1.
 Click **Save changes**.

1.
 Create a new impersonation token.


 To generate a GitLab impersonation token:

1.
 Sign into GitLab with administrator privileges.

1.
 In the **Admin** area, select the user you want Burp Suite DAST to use to raise GitLab issues.

1.
 Click **Impersonation Tokens**.

1.
 Give the impersonation token a name (e.g. "Burp Suite DAST"), and check the box to give the impersonation token `api` scope.


![Creating a GitLab impersonation token](https://portswigger.net/burp/documentation/dast/images/gitlab/creating-a-gitlab-impersonation-token-for-burp-suite-dast.png)

1.
 Click the **Create impersonation token** button.

1.
 Copy the impersonation token to your clipboard.


## Connect Burp Suite DAST to GitLab

 If you're logged in as an administrator, you can connect Burp Suite DAST to GitLab:


1.
 Go to **Settings > Integrations**.

1.
 On the GitLab tile, click **Configure**. This takes you to the GitLab integration screen.

1.

 In the provided field, enter your GitLab API URL, for example, `https://gitlab.example.com`.


![Connecting Burp Suite DAST to GitLab](https://portswigger.net/burp/documentation/dast/images/gitlab/connecting-burp-suite-dast-to-gitlab.jpg)

1.
 Enter the GitLab personal access token you created earlier. For more information, see [Generating a GitLab impersonation token](https://portswigger.net/burp/documentation/dast/user-guide/integrate-issue-tracking-platforms/integrating-gitlab#generate-a-gitlab-impersonation-token).

1.
 Click **Connect**.


 If Burp Suite DAST successfully connects to GitLab, you'll be presented with options to configure how issues are raised both manually and automatically.


#### Note

 You must enable at least one of these in order to complete the GitLab configuration.


### Enable GitLab issues to be raised manually

 To enable users to [raise GitLab issues manually](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/gitlab-issues) from within Burp Suite DAST, you need to configure the list of GitLab projects and issue types that they can choose from:


1.
 Select a project from the **Project** drop-down list.

1.
 Select an issue type from the **Issue type** drop-down list.

1.
 Click the **+** symbol.

1.

 If necessary, repeat these steps to add more projects and issue types.


#### Note

 You need to add separate entries for each issue type, even when adding multiple issue types from the same project.

1.
 Click **Save**.


![Enabling manual issue creation for GitLab](https://portswigger.net/burp/documentation/dast/images/gitlab/manual-gitlab-issue-permit.png)

### Enable GitLab issues to be raised automatically

 You can configure Burp Suite DAST to raise GitLab issues automatically. Issues are created if they meet the minimum severity and confidence levels that you specify.


#### Note

 To avoid inadvertently flooding your GitLab backlog with an overwhelming number of issues, we recommend setting high severity and confidence levels initially. You can then lower these once you have a better understanding of how many issues are raised as a result of your scans.


1.
 Click **Enable**.

1.
 Select a project from the **Project** drop-down list.

1.
 Select an issue type from the **Issue type** drop-down list.

1.
 Use the sliders to set the minimum issue severity and confidence levels that trigger GitLab issue
 creation.

1.
 Click **Save**.


![Enable automatic issue creation for GitLab](https://portswigger.net/burp/documentation/dast/images/gitlab/auto-gitlab-issue-create.png)

## Raising GitLab issues from within Burp Suite DAST

 For information on how users can manually raise GitLab issues, refer to [Raise GitLab issues from within Burp Suite DAST](https://portswigger.net/burp/documentation/dast/user-guide/work-with-scan-results/raising-tickets/gitlab-issues).
