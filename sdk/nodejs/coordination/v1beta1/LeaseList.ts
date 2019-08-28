// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { core } from "../..";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import { getVersion } from "../../version";

    /**
     * LeaseList is a list of Lease objects.
     */
    export class LeaseList extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"coordination.k8s.io/v1beta1">;

      /**
       * Items is a list of schema objects.
       */
      public readonly items: pulumi.Output<outputs.coordination.v1beta1.Lease[]>;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"LeaseList">;

      /**
       * Standard list metadata. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputs.meta.v1.ListMeta>;

      /**
       * Get the state of an existing `LeaseList` resource, as identified by `id`.
       * The ID is of the form `[namespace]/<name>`; if `namespace` is omitted, then (per
       * Kubernetes convention) the ID becomes `default/<name>`.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form `[namespace]/<name>`.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LeaseList {
          return new LeaseList(name, undefined, { ...opts, id: id });
      }

      /** @internal */
      private static readonly __pulumiType = "kubernetes:coordination.k8s.io/v1beta1:LeaseList";

      /**
       * Returns true if the given object is an instance of LeaseList.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is LeaseList {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === LeaseList.__pulumiType;
      }

      /**
       * Create a coordination.v1beta1.LeaseList resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputs.coordination.v1beta1.LeaseList, opts?: pulumi.CustomResourceOptions) {
          const props: pulumi.Inputs = {};
          props["items"] = args && args.items || undefined;

          props["apiVersion"] = "coordination.k8s.io/v1beta1";
          props["kind"] = "LeaseList";
          props["metadata"] = args && args.metadata || undefined;

          props["status"] = undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }
          super(LeaseList.__pulumiType, name, props, opts);
      }
    }
