> Source: https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/saml/access-okta-groups

DAST

# Enabling Burp Suite DAST to access your Okta groups

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

![ps-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/burp-icon-enterprise.svg)

Cloud

![self-hosted-lozenge](https://portswigger.net/burp/documentation/images/icons/person.svg)

Self-hosted

 If you're not using SCIM, you can match local Burp Suite DAST groups to your Okta groups. This enables you to manage these groups locally.


 To configure your Okta Group Attribute statements in a way that Burp Suite DAST can recognize:


1.
 From the Okta admin console, go to SAML settings for your Burp Suite DAST integration.

1.

 Create Group Attribute Statements with the following values:


  -  **Name**: `http://schemas.xmlsoap.org/claims/Group`
  -  **Name format**: `Unspecified`
  -  **Filter**: `Matches regex`
  -  **Value**: `.*`

 The filter value determines the groups that Okta sends. The regex in this example makes sure that all groups are sent. If you want to limit the selection to a particular subset of groups, refer to the Okta documentation.


 Your Okta integration may already send group membership under a different attribute name. In this case, enter that name in Burp Suite DAST rather than adding a statement to Okta. For more information, see [Creating local groups for SAML or LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).


## Adding your groups to Burp Suite DAST

 To grant permissions, set the **Identity provider group** value on each group that you create in Burp Suite DAST. For more information, see [Creating local groups for SAML or LDAP](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/sso/permissions).
