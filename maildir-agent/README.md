# MAILDIR AGENT
- Watching to mailboxes directory, if there are any new mailfile created as the following path format <maildomain>/<mailuser>/new/<mailfile>
  - parse mailfile to information: `from, to, subject, text_body`
  - publish an event to topic `mailstation:INCOMING_EMAIL` with the content is the parsed email information.
  - drop the mailfile
