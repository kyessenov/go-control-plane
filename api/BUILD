load("@org_pubref_rules_protobuf//gogo:rules.bzl", "gogofast_proto_library")

gogofast_proto_library(
    name = "go_default_library",
    importmap = {
        "gogoproto/gogo.proto": "github.com/gogo/protobuf/gogoproto",
        "google/protobuf/timestamp.proto": "github.com/gogo/protobuf/types",
        "google/protobuf/duration.proto": "github.com/gogo/protobuf/types",
    },
    imports = [
        "external/com_google_protobuf/src",
        "external/com_lyft_protoc_gen_validate",
    ],
    inputs = [
        "@com_google_protobuf//:well_known_protos",
        "@com_lyft_protoc_gen_validate//validate:validate_proto",
    ],
    protos = [
        "@envoy_api//api:address.proto",
    ],
    visibility = ["//visibility:public"],
    with_grpc = True,
    deps = [
        "@com_github_gogo_protobuf//gogoproto:go_default_library",
        "@com_github_gogo_protobuf//sortkeys:go_default_library",
        "@com_github_gogo_protobuf//types:go_default_library",
    ],
    verbose = 2,
    output_to_workspace = True,
)
