# MAILDIE PARSER
- Subscribe to topic `mailstation:INCOMING_EMAIL`
- Observe any incoming email event:
  - Extracting email in `text_body` to get the OPT
  - Extracting email `from` and `text_body` to specify the OPT service name.
  - Extracting email `to` to specify the OPT receiver.
- Publish the incoming email information parsed above to topic `mailstation:NEW_OTP_PARSED`