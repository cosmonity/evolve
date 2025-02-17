<div align="center">
  <h1> SDK Evolve </h1>
</div>

Community maintained fork of original [Cosmos SDK v2](https://github.com/cosmos/cosmos-sdk/releases/tag/server%2Fv2.0.0-beta.1).

### Why?

v2 has been postpned indefinitely by the Cosmos SDK team.
The software was however in a good state, having feature parity and ready for release.

### How to use?

Unfortunately, due to the cancellation of v0.52.x by the new stewarding team of the Cosmos SDK.
v2, as it is now, is not easy to use in a chain. While the core layer is independant of the SDK, the modules are not.
Without a proper release of v0.52.x, Cosmonity would have to maintain a fork of the SDK, IBC and Cosmwasm to be able to use those mdules in a chain. This is not a viable solution.

### Changes

* Remove all packages that aren't part of the v2 core layer
* Change vanity url to `go.cosmonity.xyz`
* Simplify the build process and CI configuration
