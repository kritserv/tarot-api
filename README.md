# Tarot Card API

```
git clone https://github.com/kritserv/tarot-api.git
```

Fork from https://github.com/ekelen/tarot-api

> use go instead of js

---

### Condensed documentation

| GET path                      | Result                                  | Params                                                                                                          |
| :---------------------------- | --------------------------------------- | :-------------------------------------------------------------------------------------------------------------- |
| `/api/v1/` or `/api/v1/cards` | return all cards                        |                                                                                                                 |
| `/api/v1/cards/:name_short`   | return card with specified `name_short` | **minors:** `/swac`, `/wa02`, ..., `/cupa`, `/pequ`, `/waqu`, `/swki`, **majors** `/ar01`, `/ar02`, ...`/ar[n]` |
| `/api/v1/cards/search`        | search all cards                        | `q={text}`, `meaning={text}`, `meaning_rev={text}`                                                              |
| `/api/v1/cards/random`        | get random card(s)                      | _optional_ `n={integer <= 78}`                                                                                  |

---

Examples:

```
go run main.go
```

Get the Knight of Wands:

http://localhost:8080/api/v1/cards/wakn

Get all cards with word "love" in meaning:

http://localhost:8080/api/v1/cards/search/meaning=love

Get 10 Random Cards:

http://localhost:8080/api/v1/cards/random/n=10

Get 1 Random Card:

http://localhost:8080/api/v1/cards/random
