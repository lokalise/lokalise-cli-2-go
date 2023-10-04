## lokalise2 payment-card

Manage payment cards

### Synopsis

Credit cards are used to pay for translation orders. Each user has their own cards, that are not shared with other users. We do not store credit card details. Once the card is added, we send the details to Stripe and receive the card token, which can only be used for order purchases at Lokalise.

### Options

```
  -h, --help   help for payment-card
```

### Options inherited from parent commands

```
      --config string   config file (default is ./config.yml)
  -t, --token string    API token. You can create API tokens at https://app.lokalise.com/profile.
```

### SEE ALSO

* [lokalise2](lokalise2.md)	 - Lokalise CLI v2.6.11. Read the docs at https://github.com/lokalise/lokalise-cli-2-go
* [lokalise2 payment-card create](lokalise2_payment-card_create.md)	 - Create a card
* [lokalise2 payment-card delete](lokalise2_payment-card_delete.md)	 - Delete a card
* [lokalise2 payment-card list](lokalise2_payment-card_list.md)	 - Lists all cards
* [lokalise2 payment-card retrieve](lokalise2_payment-card_retrieve.md)	 - Retrieve a card

