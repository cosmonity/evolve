# Proto

This directory contains the protobuf files used by v2.

Its buf module is called `buf.build/mods/evolve` and is published to the [buf registry](https://buf.build/mods/evolve).
However, the fork makes use of `cosmossdk.io/api/cosmos/app/runtime/v2` go package instead of the generated package from the buf registry (`buf.build/gen/go/mods/evolve/protocolbuffers/go/cosmos/app/runtime/v2`). This is because `cosmos.app.runtime.v2` has a dependency on other packages that are not part of the v2 core layer, namely `cosmos.app.v1alpha1`. That package defines depinject module, which is a concept shared between `runtime` and `runtime/v2`.
