> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/ldap

DAST

# Configuring LDAP single sign-on for Burp Suite DAST

-

**Last updated: ** September 3, 2026
-

**Read time: ** 2 Minutes

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you have a self-hosted instance of Burp Suite DAST, you can configure LDAP-based single sign-on (SSO). This enables your users to log in with their existing Active Directory credentials.


 To configure the LDAP connection between Burp Suite DAST and your Active Directory server:


1.
 Log in to Burp Suite DAST as an administrator.

1.
 From the settings menu , select **Integrations**.

1.
 On the **LDAP** tile, click **Configure**.

1.
 Under **Connection details**, select **LDAP** or **LDAPS**. We recommend using LDAPS wherever possible.

1.

In the **Server** field, enter the IP address or hostname of your Active Directory server.

#### Note

 The port updates automatically. By default, LDAP uses port 389 and LDAPS uses port 636.

1.
 Under **Service account details**, enter the username and password for a valid Active Directory service account. This is used to query your Active Directory when authenticating users.

1.
 Specify the base distinguished name where Burp Suite DAST should search for users. All the users that you want to manage must be children of this base distinguished name.

1.
 Select a **Login method**. This determines whether users log in with their `UserPrincipalName` or their `sAMAccountName`.

1.
 When you are happy with your entries, click **Check Connection**.

1.
 To use a self-signed certificate for LDAPS, upload the root certificate when prompted if necessary.


## Testing your configuration

 Once the connection is successfully established, you can test your configuration by logging in to Burp Suite DAST. If the configuration was successful, you will see a message that you have logged in, but you don't yet have permission to do anything.


 You can now configure user groups and permissions for your users. For more information, see [Creating local groups for SAML or LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).


#### Note

 Burp Suite DAST matches an LDAP group on the group's common name (CN), not its full distinguished name. If two Active Directory groups share a CN, then they both match the same Burp Suite DAST group.
