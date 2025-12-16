package arcane_crd

// run controller-gen to generate CRDs into Helm chart templates
//go:generate controller-gen crd paths=$ARCANE_LOCATION/pkg/apis/... output:crd:dir=./.helm/templates output:stdout
