// This file is auto-generated by @hey-api/openapi-ts

import { createClient, createConfig, type Options } from '@hey-api/client-fetch';
import type { ListClabernetesContainerlabDevV1Alpha1ConnectivityForAllNamespacesData, ListClabernetesContainerlabDevV1Alpha1ConnectivityForAllNamespacesError, ListClabernetesContainerlabDevV1Alpha1ConnectivityForAllNamespacesResponse, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConnectivityData, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConnectivityError, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConnectivityResponse, ListClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ListClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ListClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, CreateClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, CreateClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, CreateClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, DeleteClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, DeleteClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, DeleteClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, ReadClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ReadClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ReadClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, PatchClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, PatchClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, PatchClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, ListClabernetesContainerlabDevV1Alpha1ImagerequestForAllNamespacesData, ListClabernetesContainerlabDevV1Alpha1ImagerequestForAllNamespacesError, ListClabernetesContainerlabDevV1Alpha1ImagerequestForAllNamespacesResponse, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedImagerequestData, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedImagerequestError, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedImagerequestResponse, ListClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ListClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ListClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, CreateClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, CreateClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, CreateClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, DeleteClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, DeleteClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, DeleteClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, ReadClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ReadClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ReadClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, PatchClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, PatchClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, PatchClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, ListClabernetesContainerlabDevV1Alpha1ConfigForAllNamespacesData, ListClabernetesContainerlabDevV1Alpha1ConfigForAllNamespacesError, ListClabernetesContainerlabDevV1Alpha1ConfigForAllNamespacesResponse, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConfigData, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConfigError, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConfigResponse, ListClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ListClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ListClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, CreateClabernetesContainerlabDevV1Alpha1NamespacedConfigData, CreateClabernetesContainerlabDevV1Alpha1NamespacedConfigError, CreateClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, DeleteClabernetesContainerlabDevV1Alpha1NamespacedConfigData, DeleteClabernetesContainerlabDevV1Alpha1NamespacedConfigError, DeleteClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, ReadClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ReadClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ReadClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, PatchClabernetesContainerlabDevV1Alpha1NamespacedConfigData, PatchClabernetesContainerlabDevV1Alpha1NamespacedConfigError, PatchClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, ListClabernetesContainerlabDevV1Alpha1TopologyForAllNamespacesData, ListClabernetesContainerlabDevV1Alpha1TopologyForAllNamespacesError, ListClabernetesContainerlabDevV1Alpha1TopologyForAllNamespacesResponse, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedTopologyData, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedTopologyError, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedTopologyResponse, ListClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ListClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ListClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, CreateClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, CreateClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, CreateClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, DeleteClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, DeleteClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, DeleteClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, ReadClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ReadClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ReadClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, PatchClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, PatchClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, PatchClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse } from './types.gen';

export const client = createClient(createConfig());

/**
 * list objects of kind Connectivity
 */
export const listClabernetesContainerlabDevV1Alpha1ConnectivityForAllNamespaces = <ThrowOnError extends boolean = false>(options?: Options<ListClabernetesContainerlabDevV1Alpha1ConnectivityForAllNamespacesData, ThrowOnError>) => { return (options?.client ?? client).get<ListClabernetesContainerlabDevV1Alpha1ConnectivityForAllNamespacesResponse, ListClabernetesContainerlabDevV1Alpha1ConnectivityForAllNamespacesError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/connectivities'
}); };

/**
 * delete collection of Connectivity
 */
export const deleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConnectivity = <ThrowOnError extends boolean = false>(options: Options<DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConnectivityData, ThrowOnError>) => { return (options?.client ?? client).delete<DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConnectivityResponse, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConnectivityError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/connectivities'
}); };

/**
 * list objects of kind Connectivity
 */
export const listClabernetesContainerlabDevV1Alpha1NamespacedConnectivity = <ThrowOnError extends boolean = false>(options: Options<ListClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ThrowOnError>) => { return (options?.client ?? client).get<ListClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, ListClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/connectivities'
}); };

/**
 * create a Connectivity
 */
export const createClabernetesContainerlabDevV1Alpha1NamespacedConnectivity = <ThrowOnError extends boolean = false>(options: Options<CreateClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ThrowOnError>) => { return (options?.client ?? client).post<CreateClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, CreateClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/connectivities'
}); };

/**
 * delete a Connectivity
 */
export const deleteClabernetesContainerlabDevV1Alpha1NamespacedConnectivity = <ThrowOnError extends boolean = false>(options: Options<DeleteClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ThrowOnError>) => { return (options?.client ?? client).delete<DeleteClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, DeleteClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/connectivities/{name}'
}); };

/**
 * read the specified Connectivity
 */
export const readClabernetesContainerlabDevV1Alpha1NamespacedConnectivity = <ThrowOnError extends boolean = false>(options: Options<ReadClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ThrowOnError>) => { return (options?.client ?? client).get<ReadClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, ReadClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/connectivities/{name}'
}); };

/**
 * partially update the specified Connectivity
 */
export const patchClabernetesContainerlabDevV1Alpha1NamespacedConnectivity = <ThrowOnError extends boolean = false>(options: Options<PatchClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ThrowOnError>) => { return (options?.client ?? client).patch<PatchClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, PatchClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/connectivities/{name}'
}); };

/**
 * replace the specified Connectivity
 */
export const replaceClabernetesContainerlabDevV1Alpha1NamespacedConnectivity = <ThrowOnError extends boolean = false>(options: Options<ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConnectivityData, ThrowOnError>) => { return (options?.client ?? client).put<ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConnectivityResponse, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConnectivityError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/connectivities/{name}'
}); };

/**
 * list objects of kind Imagerequest
 */
export const listClabernetesContainerlabDevV1Alpha1ImagerequestForAllNamespaces = <ThrowOnError extends boolean = false>(options?: Options<ListClabernetesContainerlabDevV1Alpha1ImagerequestForAllNamespacesData, ThrowOnError>) => { return (options?.client ?? client).get<ListClabernetesContainerlabDevV1Alpha1ImagerequestForAllNamespacesResponse, ListClabernetesContainerlabDevV1Alpha1ImagerequestForAllNamespacesError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/imagerequests'
}); };

/**
 * delete collection of Imagerequest
 */
export const deleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedImagerequest = <ThrowOnError extends boolean = false>(options: Options<DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedImagerequestData, ThrowOnError>) => { return (options?.client ?? client).delete<DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedImagerequestResponse, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedImagerequestError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/imagerequests'
}); };

/**
 * list objects of kind Imagerequest
 */
export const listClabernetesContainerlabDevV1Alpha1NamespacedImagerequest = <ThrowOnError extends boolean = false>(options: Options<ListClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ThrowOnError>) => { return (options?.client ?? client).get<ListClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, ListClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/imagerequests'
}); };

/**
 * create a Imagerequest
 */
export const createClabernetesContainerlabDevV1Alpha1NamespacedImagerequest = <ThrowOnError extends boolean = false>(options: Options<CreateClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ThrowOnError>) => { return (options?.client ?? client).post<CreateClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, CreateClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/imagerequests'
}); };

/**
 * delete a Imagerequest
 */
export const deleteClabernetesContainerlabDevV1Alpha1NamespacedImagerequest = <ThrowOnError extends boolean = false>(options: Options<DeleteClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ThrowOnError>) => { return (options?.client ?? client).delete<DeleteClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, DeleteClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/imagerequests/{name}'
}); };

/**
 * read the specified Imagerequest
 */
export const readClabernetesContainerlabDevV1Alpha1NamespacedImagerequest = <ThrowOnError extends boolean = false>(options: Options<ReadClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ThrowOnError>) => { return (options?.client ?? client).get<ReadClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, ReadClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/imagerequests/{name}'
}); };

/**
 * partially update the specified Imagerequest
 */
export const patchClabernetesContainerlabDevV1Alpha1NamespacedImagerequest = <ThrowOnError extends boolean = false>(options: Options<PatchClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ThrowOnError>) => { return (options?.client ?? client).patch<PatchClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, PatchClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/imagerequests/{name}'
}); };

/**
 * replace the specified Imagerequest
 */
export const replaceClabernetesContainerlabDevV1Alpha1NamespacedImagerequest = <ThrowOnError extends boolean = false>(options: Options<ReplaceClabernetesContainerlabDevV1Alpha1NamespacedImagerequestData, ThrowOnError>) => { return (options?.client ?? client).put<ReplaceClabernetesContainerlabDevV1Alpha1NamespacedImagerequestResponse, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedImagerequestError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/imagerequests/{name}'
}); };

/**
 * list objects of kind Config
 */
export const listClabernetesContainerlabDevV1Alpha1ConfigForAllNamespaces = <ThrowOnError extends boolean = false>(options?: Options<ListClabernetesContainerlabDevV1Alpha1ConfigForAllNamespacesData, ThrowOnError>) => { return (options?.client ?? client).get<ListClabernetesContainerlabDevV1Alpha1ConfigForAllNamespacesResponse, ListClabernetesContainerlabDevV1Alpha1ConfigForAllNamespacesError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/configs'
}); };

/**
 * delete collection of Config
 */
export const deleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConfig = <ThrowOnError extends boolean = false>(options: Options<DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConfigData, ThrowOnError>) => { return (options?.client ?? client).delete<DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConfigResponse, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedConfigError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/configs'
}); };

/**
 * list objects of kind Config
 */
export const listClabernetesContainerlabDevV1Alpha1NamespacedConfig = <ThrowOnError extends boolean = false>(options: Options<ListClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ThrowOnError>) => { return (options?.client ?? client).get<ListClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, ListClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/configs'
}); };

/**
 * create a Config
 */
export const createClabernetesContainerlabDevV1Alpha1NamespacedConfig = <ThrowOnError extends boolean = false>(options: Options<CreateClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ThrowOnError>) => { return (options?.client ?? client).post<CreateClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, CreateClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/configs'
}); };

/**
 * delete a Config
 */
export const deleteClabernetesContainerlabDevV1Alpha1NamespacedConfig = <ThrowOnError extends boolean = false>(options: Options<DeleteClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ThrowOnError>) => { return (options?.client ?? client).delete<DeleteClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, DeleteClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/configs/{name}'
}); };

/**
 * read the specified Config
 */
export const readClabernetesContainerlabDevV1Alpha1NamespacedConfig = <ThrowOnError extends boolean = false>(options: Options<ReadClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ThrowOnError>) => { return (options?.client ?? client).get<ReadClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, ReadClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/configs/{name}'
}); };

/**
 * partially update the specified Config
 */
export const patchClabernetesContainerlabDevV1Alpha1NamespacedConfig = <ThrowOnError extends boolean = false>(options: Options<PatchClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ThrowOnError>) => { return (options?.client ?? client).patch<PatchClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, PatchClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/configs/{name}'
}); };

/**
 * replace the specified Config
 */
export const replaceClabernetesContainerlabDevV1Alpha1NamespacedConfig = <ThrowOnError extends boolean = false>(options: Options<ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConfigData, ThrowOnError>) => { return (options?.client ?? client).put<ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConfigResponse, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedConfigError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/configs/{name}'
}); };

/**
 * list objects of kind Topology
 */
export const listClabernetesContainerlabDevV1Alpha1TopologyForAllNamespaces = <ThrowOnError extends boolean = false>(options?: Options<ListClabernetesContainerlabDevV1Alpha1TopologyForAllNamespacesData, ThrowOnError>) => { return (options?.client ?? client).get<ListClabernetesContainerlabDevV1Alpha1TopologyForAllNamespacesResponse, ListClabernetesContainerlabDevV1Alpha1TopologyForAllNamespacesError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/topologies'
}); };

/**
 * delete collection of Topology
 */
export const deleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedTopology = <ThrowOnError extends boolean = false>(options: Options<DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedTopologyData, ThrowOnError>) => { return (options?.client ?? client).delete<DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedTopologyResponse, DeleteClabernetesContainerlabDevV1Alpha1CollectionNamespacedTopologyError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/topologies'
}); };

/**
 * list objects of kind Topology
 */
export const listClabernetesContainerlabDevV1Alpha1NamespacedTopology = <ThrowOnError extends boolean = false>(options: Options<ListClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ThrowOnError>) => { return (options?.client ?? client).get<ListClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, ListClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/topologies'
}); };

/**
 * create a Topology
 */
export const createClabernetesContainerlabDevV1Alpha1NamespacedTopology = <ThrowOnError extends boolean = false>(options: Options<CreateClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ThrowOnError>) => { return (options?.client ?? client).post<CreateClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, CreateClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/topologies'
}); };

/**
 * delete a Topology
 */
export const deleteClabernetesContainerlabDevV1Alpha1NamespacedTopology = <ThrowOnError extends boolean = false>(options: Options<DeleteClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ThrowOnError>) => { return (options?.client ?? client).delete<DeleteClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, DeleteClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/topologies/{name}'
}); };

/**
 * read the specified Topology
 */
export const readClabernetesContainerlabDevV1Alpha1NamespacedTopology = <ThrowOnError extends boolean = false>(options: Options<ReadClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ThrowOnError>) => { return (options?.client ?? client).get<ReadClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, ReadClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/topologies/{name}'
}); };

/**
 * partially update the specified Topology
 */
export const patchClabernetesContainerlabDevV1Alpha1NamespacedTopology = <ThrowOnError extends boolean = false>(options: Options<PatchClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ThrowOnError>) => { return (options?.client ?? client).patch<PatchClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, PatchClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/topologies/{name}'
}); };

/**
 * replace the specified Topology
 */
export const replaceClabernetesContainerlabDevV1Alpha1NamespacedTopology = <ThrowOnError extends boolean = false>(options: Options<ReplaceClabernetesContainerlabDevV1Alpha1NamespacedTopologyData, ThrowOnError>) => { return (options?.client ?? client).put<ReplaceClabernetesContainerlabDevV1Alpha1NamespacedTopologyResponse, ReplaceClabernetesContainerlabDevV1Alpha1NamespacedTopologyError, ThrowOnError>({
    ...options,
    url: '/apis/clabernetes.containerlab.dev/v1alpha1/namespaces/{namespace}/topologies/{name}'
}); };