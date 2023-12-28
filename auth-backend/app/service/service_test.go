package service

/**
 * Struct to be used for testing service
 * @param queryParams map[string]string --> Query params to be used for testing
 * @param mockValue interface{} --> Mock value to be returned by repository
 * @param mockError error --> Mock Error to be returned by repository
 * @param expectedValue interface{} --> Expected value to be returned by service
 * @param expectedStatusCode int --> Expected status code to be returned by service
 */
type ServiceTestGET struct {
	mockValue interface{}
	mockError error
	expectedValue interface{}
	expectedStatusCode int
}

type ServiceTestPOST struct {
	data map[string]interface{}
	mockValue interface{}
	mockError error
	expectedValue interface{}
	expectedStatusCode int
}
