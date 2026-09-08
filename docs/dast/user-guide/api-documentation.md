> Source: https://portswigger.net/burp/documentation/dast/user-guide/api-documentation

DAST

# API overview

-

**Last updated: ** September 3, 2026
-

**Read time: ** 1 Minute

 Burp Suite DAST provides two APIs that you can use to interact with the system from other third-party software. The [GraphQL API](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/graphql-api) offers the broadest range of functionality and is recommended for new integrations, while the [REST API](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/rest) offers a simple migration for users who are familiar with the Burp Suite Professional API.


## Using the APIs

 In order to use either of Burp Suite DAST's APIs, you will need to set up an API user. API users each have a unique API key that enables them to authenticate when making requests.


 Note that Burp Suite DAST's user roles apply to API users in the same way as UI users. You can only use the APIs to perform those tasks that the user permissions associated with your role allow. As such, you should make sure that any API users you set up have the correct roles applied.


#### Related pages

-

[Creating API users.](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/create-api-user)
-  [GraphQL API](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/graphql-api)
-  [REST API](https://portswigger.net/burp/documentation/dast/user-guide/api-documentation/rest)
-  [MCP server](https://portswigger.net/burp/documentation/dast/user-guide/using-ai/mcp-server)
-

For more information on configuring roles in Burp Suite DAST, see [Role-based access control](https://portswigger.net/burp/documentation/dast/user-guide/users-and-permissions/role-based-access-control).
