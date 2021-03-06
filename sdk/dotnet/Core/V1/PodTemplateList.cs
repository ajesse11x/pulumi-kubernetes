// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Core.V1
{
    /// <summary>
    /// PodTemplateList is a list of PodTemplates.
    /// </summary>
    public partial class PodTemplateList : KubernetesResource
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
        /// List of pod templates
        /// </summary>
        [Output("items")]
        public Output<ImmutableArray<Types.Outputs.Core.V1.PodTemplate>> Items { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers
        /// may infer this from the endpoint the client submits requests to. Cannot be updated. In
        /// CamelCase. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard list metadata. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ListMeta> Metadata { get; private set; } = null!;


        /// <summary>
        /// Create a PodTemplateList resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PodTemplateList(string name, Types.Inputs.Core.V1.PodTemplateListArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:core/v1:PodTemplateList", name, SetAPIKindAndVersion(args), MakeOptions(options))
        {
        }

        internal PodTemplateList(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:core/v1:PodTemplateList", name, dictionary, options)
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.Core.V1.PodTemplateListArgs? args)
        {
            args ??= new Types.Inputs.Core.V1.PodTemplateListArgs();
            args.ApiVersion = "v1";
            args.Kind = "PodTemplateList";
            return args;
        }

        private static CustomResourceOptions? MakeOptions(CustomResourceOptions? options)
        {
            return options;
        }

        /// <summary>
        /// Get an existing PodTemplateList resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PodTemplateList Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PodTemplateList(name, default(Types.Inputs.Core.V1.PodTemplateListArgs),
                CustomResourceOptions.Merge(options, new CustomResourceOptions {Id = id}));
        }
    }
}
