## moactl create idp

Add IDP for cluster

### Synopsis

Add an Identity providers to determine how users log into the cluster.

```
moactl create idp [flags]
```

### Examples

```
  # Add a GitHub identity provider to a cluster named "mycluster"
  moactl create idp --type=github --cluster=mycluster

  # Add an identity provider following interactive prompts
  moactl create idp --cluster=mycluster --interactive
```

### Options

```
  -c, --cluster string               Name or ID of the cluster to add the IdP to (required).
  -t, --type string                  Type of identity provider. Options are [github gitlab google ldap openid].
      --name string                  Name for the identity provider.
                                     
      --mapping-method string        Specifies how new identities are mapped to users when they log in. (default "claim")
      --client-id string             Client ID from the registered application.
      --client-secret string         Client Secret from the registered application.
      --ca string                    Path to PEM-encoded certificate file to use when making requests to the server.
                                     
      --hostname string              GitHub: Optional domain to use with a hosted instance of GitHub Enterprise.
      --organizations string         GitHub: Only users that are members of at least one of the listed organizations will be allowed to log in.
      --teams string                 GitHub: Only users that are members of at least one of the listed teams will be allowed to log in. The format is <org>/<team>.
                                     
      --host-url string              GitLab: The host URL of a GitLab provider. (default "https://gitlab.com")
      --hosted-domain string         Google: Restrict users to a Google Apps domain.
                                     
      --url string                   LDAP: An RFC 2255 URL which specifies the LDAP search parameters to use.
      --insecure                     LDAP: Do not make TLS connections to the server.
      --bind-dn string               LDAP: DN to bind with during the search phase.
      --bind-password string         LDAP: Password to bind with during the search phase.
      --id-attributes string         LDAP: The list of attributes whose values should be used as the user ID. (default "dn")
      --username-attributes string   LDAP: The list of attributes whose values should be used as the preferred username. (default "uid")
      --name-attributes string       LDAP: The list of attributes whose values should be used as the display name. (default "cn")
      --email-attributes string      LDAP: The list of attributes whose values should be used as the email address.
                                     
      --issuer-url string            OpenID: The URL that the OpenID Provider asserts as the Issuer Identifier. It must use the https scheme with no URL query parameters or fragment.
      --email-claims string          OpenID: List of claims to use as the email address.
      --name-claims string           OpenID: List of claims to use as the display name.
      --username-claims string       OpenID: List of claims to use as the preferred username when provisioning a user.
      --extra-scopes string          OpenID: List of scopes to request, in addition to the 'openid' scope, during the authorization token request.
                                     
  -h, --help                         help for idp
```

### Options inherited from parent commands

```
      --debug            Enable debug mode.
  -i, --interactive      Enable interactive mode.
      --profile string   Use a specific AWS profile from your credential file.
  -v, --v Level          log level for V logs
```

### SEE ALSO

* [moactl create](moactl_create.md)	 - Create a resource from stdin

