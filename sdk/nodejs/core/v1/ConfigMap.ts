// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { core } from "../..";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import { getVersion } from "../../version";

    /**
     * ConfigMap holds configuration data for pods to consume.
     */
    export class ConfigMap extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"v1">;

      /**
       * BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-',
       * '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys
       * stored in BinaryData must not overlap with the ones in the Data field, this is enforced
       * during validation process. Using this field will require 1.10+ apiserver and kubelet.
       */
      public readonly binaryData: pulumi.Output<object>;

      /**
       * Data contains the configuration data. Each key must consist of alphanumeric characters,
       * '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The
       * keys stored in Data must not overlap with the keys in the BinaryData field, this is
       * enforced during validation process.
       */
      public readonly data: pulumi.Output<{[key: string]: pulumi.Output<string>}>;

      /**
       * Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated
       * (only object metadata can be modified). If not set to true, the field can be modified at
       * any time. Defaulted to nil. This is an alpha field enabled by ImmutableEphemeralVolumes
       * feature gate.
       */
      public readonly immutable: pulumi.Output<boolean>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"ConfigMap">;

      /**
       * Standard object's metadata. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputs.meta.v1.ObjectMeta>;

      /**
       * Get the state of an existing `ConfigMap` resource, as identified by `id`.
       * The ID is of the form `[namespace]/<name>`; if `namespace` is omitted, then (per
       * Kubernetes convention) the ID becomes `default/<name>`.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form `[namespace]/<name>`.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConfigMap {
          return new ConfigMap(name, undefined, { ...opts, id: id });
      }

      /** @internal */
      private static readonly __pulumiType = "kubernetes:core/v1:ConfigMap";

      /**
       * Returns true if the given object is an instance of ConfigMap.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is ConfigMap {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === ConfigMap.__pulumiType;
      }

      /**
       * Create a core.v1.ConfigMap resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputs.core.v1.ConfigMap, opts?: pulumi.CustomResourceOptions) {
          const props: pulumi.Inputs = {};

          props["apiVersion"] = "v1";
          props["binaryData"] = args?.binaryData;
          props["data"] = args?.data;
          props["immutable"] = args?.immutable;
          props["kind"] = "ConfigMap";
          props["metadata"] = args?.metadata;

          props["status"] = undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }

          super(ConfigMap.__pulumiType, name, props, opts);
      }
    }
