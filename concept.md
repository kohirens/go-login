# SSO Multiple login Support

We need to change how accounts track providers and profiles. Currently, they
track a single client ID and profile per provider. Change it so they track
multiple.

SSO we want the client to be able to tie multiple profiles to their account.
It can be theirs or someone they want to share the account with. So the account
is NOT specific to a user. But rather the account acts as a stand-a-lone
entity that many users can access. This will allow for accounts for individuals,
families, small or large organizations of various kinds. A profile will take the
place of what we'll call a traditional account. And an account will one to many
profiles.

We can have the account owner login and send links that invite someone to add
themselves (as an additional profile) to the account. Ergo, that link takes
them to a page where they log in with their preferred provider (or sign up
without a provider) and ties them to the account that sent the invitation link.

## Account Structure Hierarchy

* An Account stores Profiles.
* A single Profile stores:
  * A users personally identifiable information.
  * OIDC providers can be used to log in to the account.
    * Each provider list client app IDs where it was used to log in.
  * A map of applications the client has logged in from.
* A ClientApp is:
  * a way to track when the client successfully logs in from an app.
  * stored as data in the storage of the applicatoin the client used to login.
  * only stored when one is not already present in the clients application's
    secure storage.
  * recorded in the profile list of client app IDs.
  * cleared from the secure storage, then it is considered orphaned and has to
    be deleted (manually) from the users profile.
  * is deleted from the profile's provider list on logout.
  * disposable, it merely represents a login to a particular client app.

## Login Flow Described

We store the path or just the filename to the account in the cookie, saving
query time.

* An account stores user profiles.
  * Each profile stores user info.
    * User info stores providers. They also store apps where a client has logged
      in from. A client app ID is generated when the client logs in from that
      app. If the app's storage is cleared, then the client app ID is orphaned
      and has to be deleted (manually) from the users profile.
      * The client app ID is also deleted on logout.
      * So client app IDs are disposable, and merely represents a login from
        a client's app.
      * The client is successfully logged in.
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

## Account Instantiation Described

I just opened my browser and went to the site.

1. It does not find a cookie.
   1. So it generates an ID for the client's app and stores it in there.
   2. I am presented with login options:
      1. Make an Account:
         1. It does not recognize my email address, so it lets me proceed.
            1. A new profile is made with:
               1. Set the user info.
               2. Add this client app ID to the client apps list.
            2. Make a new Account:
               1. Generate a new ID.
               2. Set the owner to the ID of the first profile.
            3. End flow.
         2. It recognized my email, so I'm sent to recover.
            1. End flow.
      2. Start an Account with Google as OIDC:
         1. Click the Login with Google Button.
            1. The client has not logged in with Google from this app:
                1. The OIDC flow begins:
                   1. The user agrees to consent.
                   2. User is returned to the site.
                   3. A token is pulled for the user.
                   4. The user info is returned from Google servers using the
                      token.
                   5. A new account is started:
                      1. A new profile is made with:
                          1. Set the user info.
                          2. Add this client app ID to the client apps list.
                      2. Make a new Account:
                          1. Generate a new ID.
                          2. Set the owner to the ID of the first profile.
                      3. End flow.
            2. There is an existing login for the client from this app with
               Google.
               1. End flow.
      3. Start an Account with Apple as OIDC:
         1. Unknown.