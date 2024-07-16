## lokalise2 payment-card create

Create a card

### Synopsis

Adds new payment card to user cards.

```
lokalise2 payment-card create [flags]
```

### Options

```
      --cvc string      3-digit card CVC code (required).
      --exp-month int   Card expiration month (1-12) (required).
      --exp-year int    Card expiration year (required).
  -h, --help            help for create
      --number string   Card number (required).
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2 payment-card](lokalise2_payment-card.md)	 - Manage payment cards

