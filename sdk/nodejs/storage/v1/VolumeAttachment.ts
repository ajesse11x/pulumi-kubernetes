// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { core } from "../..";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import { getVersion } from "../../version";

    /**
     * VolumeAttachment captures the intent to attach or detach the specified volume to/from the
     * specified node.
     * 
     * VolumeAttachment objects are non-namespaced.
     */
    export class VolumeAttachment extends pulumi.CustomResource {
      /**
       * APIVersion defines the versioned schema of this representation of an object. Servers should
       * convert recognized schemas to the latest internal value, and may reject unrecognized
       * values. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
       */
      public readonly apiVersion: pulumi.Output<"storage.k8s.io/v1">;

      /**
       * Kind is a string value representing the REST resource this object represents. Servers may
       * infer this from the endpoint the client submits requests to. Cannot be updated. In
       * CamelCase. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
       */
      public readonly kind: pulumi.Output<"VolumeAttachment">;

      /**
       * Standard object metadata. More info:
       * https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
       */
      public readonly metadata: pulumi.Output<outputs.meta.v1.ObjectMeta>;

      /**
       * Specification of the desired attach/detach volume behavior. Populated by the Kubernetes
       * system.
       */
      public readonly spec: pulumi.Output<outputs.storage.v1.VolumeAttachmentSpec>;

      /**
       * Status of the VolumeAttachment request. Populated by the entity completing the attach or
       * detach operation, i.e. the external-attacher.
       */
      public readonly status: pulumi.Output<outputs.storage.v1.VolumeAttachmentStatus>;

      /**
       * Get the state of an existing `VolumeAttachment` resource, as identified by `id`.
       * The ID is of the form `[namespace]/<name>`; if `namespace` is omitted, then (per
       * Kubernetes convention) the ID becomes `default/<name>`.
       *
       * Pulumi will keep track of this resource using `name` as the Pulumi ID.
       *
       * @param name _Unique_ name used to register this resource with Pulumi.
       * @param id An ID for the Kubernetes resource to retrieve. Takes the form `[namespace]/<name>`.
       * @param opts Uniquely specifies a CustomResource to select.
       */
      public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VolumeAttachment {
          return new VolumeAttachment(name, undefined, { ...opts, id: id });
      }

      /** @internal */
      private static readonly __pulumiType = "kubernetes:storage.k8s.io/v1:VolumeAttachment";

      /**
       * Returns true if the given object is an instance of VolumeAttachment.  This is designed to work even
       * when multiple copies of the Pulumi SDK have been loaded into the same process.
       */
      public static isInstance(obj: any): obj is VolumeAttachment {
          if (obj === undefined || obj === null) {
              return false;
          }

          return obj["__pulumiType"] === VolumeAttachment.__pulumiType;
      }

      /**
       * Create a storage.v1.VolumeAttachment resource with the given unique name, arguments, and options.
       *
       * @param name The _unique_ name of the resource.
       * @param args The arguments to use to populate this resource's properties.
       * @param opts A bag of options that control this resource's behavior.
       */
      constructor(name: string, args?: inputs.storage.v1.VolumeAttachment, opts?: pulumi.CustomResourceOptions) {
          const props: pulumi.Inputs = {};
          props["spec"] = args && args.spec || undefined;

          props["apiVersion"] = "storage.k8s.io/v1";
          props["kind"] = "VolumeAttachment";
          props["metadata"] = args && args.metadata || undefined;

          props["status"] = undefined;

          if (!opts) {
              opts = {};
          }

          if (!opts.version) {
              opts.version = getVersion();
          }
          super(VolumeAttachment.__pulumiType, name, props, opts);
      }
    }
