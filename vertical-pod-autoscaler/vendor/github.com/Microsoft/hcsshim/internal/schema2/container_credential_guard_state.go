/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type ContainerCredentialGuardState struct {

	//  Authentication cookie for calls to a Container Credential Guard instance.
	Cookie string `json:"Cookie,omitempty"`

	//  Name of the RPC endpoint of the Container Credential Guard instance.
	RpcEndpoint string `json:"RpcEndpoint,omitempty"`

	//  Transport used for the configured Container Credential Guard instance.
	Transport string `json:"Transport,omitempty"`

	//  Credential spec used for the configured Container Credential Guard instance.
	CredentialSpec string `json:"CredentialSpec,omitempty"`
}