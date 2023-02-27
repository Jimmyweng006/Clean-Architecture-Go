/*
 * Digimon Service API
 *
 * 提供孵化數碼蛋與培育等數碼寶貝養成服務
 *
 * API version: 1.0.0
 * Contact: 123@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DigimonInfo struct {
	// 數碼蛋的唯一識別碼，格式為uuid v4
	Id string `json:"id,omitempty"`
	// 數碼蛋的名稱
	Name string `json:"name,omitempty"`
	// 數碼蛋此時的狀態
	Status string `json:"status,omitempty"`
}