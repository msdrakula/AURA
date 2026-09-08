> Source: https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/create-api-user

DAST

# Creating an API user for the CI/CD integration

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

 Regardless of which CI/CD platform you use, the first step of the integration process is to create a dedicated API user in Burp Suite DAST. This is used by the CI/CD platform to communicate with the DAST server in order to create scans, access the results, and so on.


## Create a role and group for CI/CD users

 Before creating the API user, you need to create a new role and group to make sure that your user has all of the permissions required to initiate scans from a CI/CD system.


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the top navigation menu, go to **Team > Roles** and then click the **New role** button.

1.
 Enter a name for the role, such as **CI/CD scan initiator**.

1.
 From the list of permissions, select both the **View sites** and **View site details** permissions, then click **Save**.

1.
 Go to the **Groups** tab and click the **New group** button.

1.
 Enter a name for the group, such as **CI/CD scan initiators**.

1.
 From the list of roles, select both the built-in **Scan initiators** role and the new role that you just created, then click **Save**.


## Create the CI/CD API user

1.
 While logged in to Burp Suite DAST as an administrator, go to **Team > Add a new user**.

1.
 In the corresponding fields, enter a name and username to help you identify this user later. This can be anything you want, for example, **Jenkins User**.

1.
 Enter an email address for the user. This can be any email address you want, but please use an address that you monitor regularly. Burp Suite DAST may occasionally send important notifications to this address.

1.
 Select the login type **API Key**.

1.
 Assign the user to the new **CI/CD scan initiators** group that you created earlier, then click **Save**.

1.
 A dialog will appear prompting you to save your API key and API link. Copy these using the buttons provided and save them somewhere secure **before** closing the dialog.


#### Warning

 Once you close this dialog, you cannot retrieve the API key for an existing user. If you lose it, you will need to generate a new key and manually update this in any other applications that use the old one.


 Now that you've created the API user, you can use it to configure the integration with your preferred CI/CD platform. Please follow the relevant instructions below.


-  [Integrating with Jenkins](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/jenkins)
-  [Integrating with TeamCity](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/teamcity)
-  [Integrating with other CI/CD platforms](https://portswigger.net/burp/documentation/dast/user-guide/ci-cd/plugins/generic)
