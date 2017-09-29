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

type Pagination struct {

	Total int32 `json:"total"`

	Count int32 `json:"count"`

	PerPage int32 `json:"per_page"`

	CurrentPage int32 `json:"current_page"`

	TotalPages int32 `json:"total_pages"`

	Links PaginationLinks `json:"links,omitempty"`
}