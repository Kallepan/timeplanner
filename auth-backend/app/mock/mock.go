/* This file contains the Mock interface which provides useful testing functionality
Functions:
	- On(function_name string) Mock --> This function is used to set the function name to be mocked
	- Return(mockData interface{}, error error) --> This function is used to set the mock data and error to be returned
*/

package mock

type Mock interface {
	On(functionName string) Mock
	Return(mockData interface{}, errorData error) Mock
}

/*** Example Implementation ***/
type MockImpl struct {
	mockDataContainer map[string]interface{}
	errorDataContainer    map[string]error
	primedFunctionName    string
}

// Return implements Mock.
func (m *MockImpl) Return(mockData interface{}, errorData error) Mock {
	m.mockDataContainer[m.primedFunctionName] = mockData
	m.errorDataContainer[m.primedFunctionName] = errorData

	return m
}

// On implements Mock.
func (m *MockImpl) On(functionName string) Mock {
	m.primedFunctionName = functionName

	m.mockDataContainer[functionName] = nil
	m.errorDataContainer[functionName] = nil
	return m
}
/*** End of Example Implementation ***/