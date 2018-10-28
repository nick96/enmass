Command line tool to easily send out texts to an entire group in you
Google Contacts.

```
Usage: enmass <group> [arguments] [options]
  group is the name of the contact group we will be operating on.

  Arguments:
    check: Check that all of the members of <group> have at least one
           email and phone number
    stat: Print out the stats of a group (nuber of members, emails and
          phone numbers, etc.)
    get: Perform a get operation on the group (this requires a second
         argument of "email" or "text")
    send: Perform a send operation on the group (this requires a
          second argument of "email" or "text")

  Options:
    --credentials/-c: Path to the credentials file
    --Token/-t:       Path to the Token file

  The usage of the "get" and "send" arguments requires an additional
  argument of "email" or "text". This is used to specify what
  operation is to be performed. "email" indicates the operation is to
  be on emails and "text" on phone numbers. The possible combinations
  are:

    get email:  Get all the emails for the group
    get text:   Get all the phone numbers for the group
    send email: Send an email to all the members of the group
    send text:  Send a text message to all the members of the group
```

The environment variable `GOOGLE_APPLICATION_CREDENTIALS` must be set to point to the file with the application's credentials from the google developer console. These credentials must include access to the people API. This file's default name when downloaded from google is `credentials.json`. These credentials will then be exchanged for an OAuth2 token. To do so you will be prompted to go to a URL and copy the given code, then paste this back into the terminal. This token will be saved for future use.

To utilise Twilio for sending text messages. You will need to set the environment variables `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, and `TWILIO_URL`. These are just the corresponding account details you get when setting up your Twilio account. The phone number that Twilio sends the text messages from is set by `ENMASS_PHONE`. This can be the phone number given to you, or a an alphanumeric number.

Finally, to add a signature to your text messages use `ENMASS_SIGNATURE`. This signature will be included in the message that comes up in the message editor.
