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
type ServiceTestGET struct {
	// mock value to be returned by repository
	mockValue interface{}
	// mock error to be returned by repository
	mockError error

	// expected value to be returned by service
	expectedResponse interface{}
	// expected status code to be returned by service
	expectedStatusCode int
}

type ServiceTestPOST struct {
	// data to be used for update
	mockRequestData map[string]interface{}
	// mock error to be returned by repository
	mockError error

	// expected value to be returned by service
	expectedValue interface{}
	// expected status code to be returned by service
	expectedStatusCode int
}

type ServiceTestPUT struct {
	// data to be used for update
	mockRequestData map[string]interface{}

	// mock value to be returned by repository
	mockValue interface{}
	// mock error to be returned by repository
	mockError error

	// expected value to be returned by service
	expectedValue interface{}
	// expected status code to be returned by service
	expectedStatusCode int

	queryParams map[string]string
}

func (s *ServiceTestPUT) QueryParamsToGinParams() gin.Params {
	var params gin.Params
	for key, value := range s.queryParams {
		params = append(params, gin.Param{Key: key, Value: value})
	}
	return params
}

type ServiceTestDELETE struct {
	/**
	 * Struct to be used for testing service
	 * It is used for testing service with DELETE method
	 * It has a mockValue and mockError to be returned by repository
	 * It has an expectedStatusCode to be returned by service
	 **/
	// mock value to be returned by repository
	mockValue interface{}
	// mock error to be returned by repository
	mockError error

	// expected status code to be returned by service
	expectedStatusCode int

	queryParams map[string]string
}

func (s *ServiceTestDELETE) QueryParamsToGinParams() gin.Params {
	var params gin.Params
	for key, value := range s.queryParams {
		params = append(params, gin.Param{Key: key, Value: value})
	}
	return params
}
