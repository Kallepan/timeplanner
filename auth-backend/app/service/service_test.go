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
	// mock value to be returned by repository
	mockValue interface{}
	// mock error to be returned by repository
	mockError error

	// expected value to be returned by service
	expectedValue interface{}
	// expected status code to be returned by service
	expectedStatusCode int
}

type ServiceTestPOST struct {
	// data to be used for update
	data map[string]interface{}
	// mock value to be returned by repository

	mockValue interface{}
	// mock error to be returned by repository
	mockError error

	// expected value to be returned by service
	expectedValue interface{}
	// expected status code to be returned by service
	expectedStatusCode int
}

type ServiceTestPUT struct {
	// data to be used for update
	data map[string]interface{}

	// mock value to be returned by repository
	mockValue interface{}
	// mock error to be returned by repository
	mockError error

	// expected value to be returned by service
	expectedValue interface{}
	// expected status code to be returned by service
	expectedStatusCode int
}
