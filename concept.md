# SSO Multiple login Support

We need to change how accounts track providers and profiles. Currently, they
track a single client ID and profile per provider. Change it so they track
multiple.

SSO we want the client to be able to tie multiple logins to their account.
It can be theirs or someone they want to share the account with. So the account
is NOT specific to a user. But rather the accounts act as stand-a-lone entities
that users can access. This will allow for accounts for individuals, families,
small or large organizations of various kinds. A profile will take the place
of what we'll call a traditional account.

We can have the account owner login and send links that invite someone to add
a profile to account. Ergo, that link takes them to a page where they log in
with their preferred provider (or sign up without a provider) and ties them to
the account that sent the invitation link.

## Return Flow

We store the path or just the filename to the account in the cookie, saving
query time.

* An account stores user profiles.
  * Each profile stores user info.
    * User info stores providers. They also store devices logged into. A devices
      ID is generated when the client logs in on that device. If the cookie is
      cleared on that device, then the Device ID is orphaned and has to be
      deleted (manually) from the users profile.
      * The device is also deleted on logout.
      * So device GUIDs are disposable. It merely represents a login to a device.
      * The client is successfully logged in. Now they have returned to the site
        entry page (index.html).
      * This ends this flow.

1. Lookup the Encrypted Value Cookie:
    1. Not present, they get sent/stay on the login page.
    2. Present:
        1. Load the cookie,
        2. Check the cookie is valid:
           1. Valid:
              1. Check for login token:
                 1. Valid
                    1. Restore the token in the provider, which should then save it to the session.
                    2. Send/Show the page the client has requested.
                 2. Invalid:
                    1. Try to refresh the token:
                       1. Success:
                          1. Restore the token in the provider, which should then save it to the session.
                          2. Send/Show the page the client has requested.
                       2. Unsuccessful:
                          1. Send them to a page where they can re-consent.
