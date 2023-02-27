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

type ModelError struct {
	// 錯誤訊息
	Message string `json:"message"`
	// 錯誤代碼:  * `3000` - Internal error 
	Code float64 `json:"code"`
}
