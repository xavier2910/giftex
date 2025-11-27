# giftex

A simple program for setting up a gift exchange via email. Basically,
me and some friends wanted to do a gift exchange, but we couldn't find
a time to get together to draw names out of a hat. So I whipped this
up to do it via email.

## inputs

The program expects a file called `people.json` in the current directory.
[People.json](people.json) is an example of how the json should be
formatted. The program randomly assigns each person in `people.json` to
another person, and sends an email to the latter telling him he is getting
a gift for the former. You can specify a different file via the `-input` flag.

## environment

The program gets info about the email account to send from via two environment
variables: `EMAIL_SENDER` and `EMAIL_PASSWORD`. `EMAIL_SENDER` should be set to
a gmail email address, and `EMAIL_PASSWORD` to an "App Password".
