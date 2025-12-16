# Arcane CRD
This repository contains Custom Resource Definitions for [Arcane](https://github.com/SneaksAndData/arcane-operator), packed in a Helm Chart.
This chart must be installed prior to installing Arcane components that depend on these CRDs.

## Quickstart

Install the chart directly from GHCR.

```shell
helm install arcane-crd oci://ghcr.io/sneaksanddata/helm/arcane-crd
```

## Updating the CRD
CRDs are generated from [arcane-operator](https://github.com/SneaksAndData/arcane-operator).
Assuming you have a local copy of that repo, follow these steps to update the CRDs in this repo:
```shell
ARCANE_LOCATION=../arcane-operator go generate generate.go
```

Once the script completes, the updated CRDs will be in the `.helm/templates` directory.
