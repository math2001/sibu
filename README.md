# Sibu

> A simplistic SQL request builder for Postgresql

## The philosophy

> Be simple.

 It's what I love about go.

> There should be one obvious way to do it

Thanks Tim

> Reflect simplicity the result

We're building human-writable strings. There is no reason to write a complex
chaining system.

> Do YOUR job

This is a SQL builder. Sibu *builds* SQL requests.

If the requests are invalid (syntactically or logically), the *SQL database*
will report it, not sibu.

> Use the power of strong types

`interface{}` should be avoided as much as possible. Types are great hints.
