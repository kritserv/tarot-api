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
| `/api/v1/` or `/api/v1/cards` | return all cards (Done)                        |                                                                                                                 |
| `/api/v1/cards/:name_short`   | return card with specified `name_short` (Done) | **minors:** `/swac`, `/wa02`, ..., `/cupa`, `/pequ`, `/waqu`, `/swki`, **majors** `/ar01`, `/ar02`, ...`/ar[n]` |
| `/api/v1/cards/search`        | search all cards (WIP)                         | `q={text}`, `meaning={text}`, `meaning_rev={text}`                                                              |
| `/api/v1/cards/random`        | get random card(s) (WIP)                       | _optional_ `n={integer <= 78}`                                                                                  |

---

Examples:

```
go run main.go
```

Get the Knight of Wands:

http://localhost:8080/api/v1/cards/wakn

