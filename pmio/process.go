/* 
 * example
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package pmio

type Process struct {

	Id string `json:"id,omitempty"`

	Type_ string `json:"type"`

	Attributes ProcessAttributes `json:"attributes,omitempty"`
}
