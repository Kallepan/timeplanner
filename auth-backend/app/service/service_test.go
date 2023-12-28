package service

import "github.com/gin-gonic/gin"

/**
 * Struct to be used for testing service
 * @param queryParams map[string]string --> Query params to be used for testing
 * @param mockValue interface{} --> Mock value to be returned by repository
 * @param mockError error --> Mock Error to be returned by repository
 * @param expectedValue interface{} --> Expected value to be returned by service
 * @param expectedStatusCode int --> Expected status code to be returned by service
 */
type ServiceTest struct {
	queryParams map[string]string
	mockValue interface{}
	mockError error
	expectedValue interface{}
	expectedStatusCode int
}
/**
 * Interface to be implemented by service test
 */
type ServiceTestInterface interface {
	queryParamsToGinParams(queryParams map[string]string) gin.Params 
}

/**
 * Function to convert queryParams to gin.Params
 * @param queryParams map[string]string --> Query params to be converted from struct
 * @return gin.Params
 */
func (s ServiceTest) queryParamsToGinParams() gin.Params {
	var ginParams []gin.Param
	for key, value := range s.queryParams {
		ginParams = append(ginParams, gin.Param{Key: key, Value: value})
	}
	return ginParams
}