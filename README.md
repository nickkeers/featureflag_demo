# OpenFeature demo

This is a small demo of how to use the OpenFeature library to evaluate feature flags in your program, this example uses `flagd`  as the provider in the main application (in [cmd/main.go](./cmd/main.go)) but there are many other providers ([see a list here for Go](https://openfeature.dev/ecosystem?instant_search%5BrefinementList%5D%5Btype%5D%5B0%5D=Provider&instant_search%5BrefinementList%5D%5Btechnology%5D%5B0%5D=Go)).

## Running flagd in Docker locally

To run flagd I use this command:

```bash
docker run \
  --rm -it \
  --name flagd \
  -p 8013:8013 \
  -v .:/data \
  ghcr.io/open-feature/flagd:latest start \
  --uri file:///data/flags.flagd.json
```

It mounts the root directory into the container under `/data`, when the program is running you can edit [the flags json file](./flags.flagd.json) to change the `"defaultVariant"` to "on" or "off" and you should see the message printed out changes accordingly.

## Note on testing

For a simple approach to unit testing there is a package called "github.com/open-feature/go-sdk/openfeature/memprovider" in the main library which provides and "InMemoryProvider", I set up the provider in the test and create a new client to pass to my code to test which lets us bypass using `flagd` for our unit testing purposes. Doing it this way makes our test cases more explicit as we have to set up the flags for each test-case.
