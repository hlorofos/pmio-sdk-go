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

type EventAttributes struct {

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	ProcessId string `json:"process_id"`

	MessageId int32 `json:"message_id,omitempty"`

	Type_ string `json:"type"`

	Definition string `json:"definition"`

	Interrupting bool `json:"interrupting,omitempty"`

	Condition string `json:"condition,omitempty"`

	Time string `json:"time,omitempty"`

	Duration string `json:"duration,omitempty"`

	Cycle string `json:"cycle,omitempty"`

	AttachedToTaskId string `json:"attached_to_task_id,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`
}