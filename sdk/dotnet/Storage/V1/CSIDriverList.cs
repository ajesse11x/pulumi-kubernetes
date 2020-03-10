// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Storage.V1
{
    /// <summary>
    /// CSIDriverList is a collection of CSIDriver objects.
    /// </summary>
    public partial class CSIDriverList : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers
        /// should convert recognized schemas to the latest internal value, and may reject
        /// unrecognized values. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// items is the list of CSIDriver
        /// </summary>
        [Output("items")]
        public Output<ImmutableArray<Types.Outputs.Storage.V1.CSIDriver>> Items { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers
        /// may infer this from the endpoint the client submits requests to. Cannot be updated. In
        /// CamelCase. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard list metadata More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ListMeta> Metadata { get; private set; } = null!;


        /// <summary>
        /// Create a CSIDriverList resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CSIDriverList(string name, Types.Inputs.Storage.V1.CSIDriverListArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1:CSIDriverList", name, SetAPIKindAndVersion(args), options)
        {
        }

        internal CSIDriverList(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:storage.k8s.io/v1:CSIDriverList", name, dictionary, options)
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.Storage.V1.CSIDriverListArgs? args)
        {
            args ??= new Types.Inputs.Storage.V1.CSIDriverListArgs();
            args.ApiVersion = "storage.k8s.io/v1";
            args.Kind = "CSIDriverList";
            return args;
        }

        /// <summary>
        /// Get an existing CSIDriverList resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CSIDriverList Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CSIDriverList(name, default(Types.Inputs.Storage.V1.CSIDriverListArgs),
                CustomResourceOptions.Merge(options, new CustomResourceOptions {Id = id}));
        }

    }
}
