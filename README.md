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

## Usage

```go
var args = []interface{}{limit, }
var q strings.Builder
q.Write("SELECT p.title, p.content, u.username, u.avatar")
q.Write("FROM posts AS p")
q.Write("JOIN users AS u")
q.Write("ON p.userid=u.id")

if userid != -1 {
    q.Add("WHERE userid={{ p }}", userid)
}
if limit != -1 {
    q.Add("LIMIT {{ p }}", limit)
}
```
