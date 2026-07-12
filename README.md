# Go Login

Is a login system to apply to web applications. It deviates from the concept
that an account belongs to a single user. But rather an account is linked to
many profiles (users) that have some level of access to manage the account.
Read more about it at [concept].

We will start by explaining the Account schema.

Account scheme contains:
1. A UUID.
2. List of Profiles
3. Most importantly, an owner; with a restriction that there can only be 1
   owner at a time. In hopes of keeping thing simple to manage.

It works better to think of a Profile as a container for specific user
information. The Profile scheme contains:
1. A UUID.
2. A Name, something to call it other than by ID. It can be a nickname or
   online handle.
3. UserInfo structure.
   1. First name.
   2. Email address.
   3. Last name.
   4. Phone number.
4. List of avatars (this is purely for fun) IDs.
5. List of registered client applications used for login.

---

[concept]: /concept.md
