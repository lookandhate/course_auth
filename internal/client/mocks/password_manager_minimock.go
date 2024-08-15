// Code generated by http://github.com/gojuno/minimock (v3.3.13). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/lookandhate/course_auth/internal/client.PasswordManager -o password_manager_minimock.go -n PasswordManagerMock -p mocks

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// PasswordManagerMock implements client.PasswordManager
type PasswordManagerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcComparePassword          func(hash string, password string) (err error)
	inspectFuncComparePassword   func(hash string, password string)
	afterComparePasswordCounter  uint64
	beforeComparePasswordCounter uint64
	ComparePasswordMock          mPasswordManagerMockComparePassword

	funcHashPassword          func(password string) (s1 string, err error)
	inspectFuncHashPassword   func(password string)
	afterHashPasswordCounter  uint64
	beforeHashPasswordCounter uint64
	HashPasswordMock          mPasswordManagerMockHashPassword
}

// NewPasswordManagerMock returns a mock for client.PasswordManager
func NewPasswordManagerMock(t minimock.Tester) *PasswordManagerMock {
	m := &PasswordManagerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ComparePasswordMock = mPasswordManagerMockComparePassword{mock: m}
	m.ComparePasswordMock.callArgs = []*PasswordManagerMockComparePasswordParams{}

	m.HashPasswordMock = mPasswordManagerMockHashPassword{mock: m}
	m.HashPasswordMock.callArgs = []*PasswordManagerMockHashPasswordParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mPasswordManagerMockComparePassword struct {
	optional           bool
	mock               *PasswordManagerMock
	defaultExpectation *PasswordManagerMockComparePasswordExpectation
	expectations       []*PasswordManagerMockComparePasswordExpectation

	callArgs []*PasswordManagerMockComparePasswordParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// PasswordManagerMockComparePasswordExpectation specifies expectation struct of the PasswordManager.ComparePassword
type PasswordManagerMockComparePasswordExpectation struct {
	mock      *PasswordManagerMock
	params    *PasswordManagerMockComparePasswordParams
	paramPtrs *PasswordManagerMockComparePasswordParamPtrs
	results   *PasswordManagerMockComparePasswordResults
	Counter   uint64
}

// PasswordManagerMockComparePasswordParams contains parameters of the PasswordManager.ComparePassword
type PasswordManagerMockComparePasswordParams struct {
	hash     string
	password string
}

// PasswordManagerMockComparePasswordParamPtrs contains pointers to parameters of the PasswordManager.ComparePassword
type PasswordManagerMockComparePasswordParamPtrs struct {
	hash     *string
	password *string
}

// PasswordManagerMockComparePasswordResults contains results of the PasswordManager.ComparePassword
type PasswordManagerMockComparePasswordResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmComparePassword *mPasswordManagerMockComparePassword) Optional() *mPasswordManagerMockComparePassword {
	mmComparePassword.optional = true
	return mmComparePassword
}

// Expect sets up expected params for PasswordManager.ComparePassword
func (mmComparePassword *mPasswordManagerMockComparePassword) Expect(hash string, password string) *mPasswordManagerMockComparePassword {
	if mmComparePassword.mock.funcComparePassword != nil {
		mmComparePassword.mock.t.Fatalf("PasswordManagerMock.ComparePassword mock is already set by Set")
	}

	if mmComparePassword.defaultExpectation == nil {
		mmComparePassword.defaultExpectation = &PasswordManagerMockComparePasswordExpectation{}
	}

	if mmComparePassword.defaultExpectation.paramPtrs != nil {
		mmComparePassword.mock.t.Fatalf("PasswordManagerMock.ComparePassword mock is already set by ExpectParams functions")
	}

	mmComparePassword.defaultExpectation.params = &PasswordManagerMockComparePasswordParams{hash, password}
	for _, e := range mmComparePassword.expectations {
		if minimock.Equal(e.params, mmComparePassword.defaultExpectation.params) {
			mmComparePassword.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmComparePassword.defaultExpectation.params)
		}
	}

	return mmComparePassword
}

// ExpectHashParam1 sets up expected param hash for PasswordManager.ComparePassword
func (mmComparePassword *mPasswordManagerMockComparePassword) ExpectHashParam1(hash string) *mPasswordManagerMockComparePassword {
	if mmComparePassword.mock.funcComparePassword != nil {
		mmComparePassword.mock.t.Fatalf("PasswordManagerMock.ComparePassword mock is already set by Set")
	}

	if mmComparePassword.defaultExpectation == nil {
		mmComparePassword.defaultExpectation = &PasswordManagerMockComparePasswordExpectation{}
	}

	if mmComparePassword.defaultExpectation.params != nil {
		mmComparePassword.mock.t.Fatalf("PasswordManagerMock.ComparePassword mock is already set by Expect")
	}

	if mmComparePassword.defaultExpectation.paramPtrs == nil {
		mmComparePassword.defaultExpectation.paramPtrs = &PasswordManagerMockComparePasswordParamPtrs{}
	}
	mmComparePassword.defaultExpectation.paramPtrs.hash = &hash

	return mmComparePassword
}

// ExpectPasswordParam2 sets up expected param password for PasswordManager.ComparePassword
func (mmComparePassword *mPasswordManagerMockComparePassword) ExpectPasswordParam2(password string) *mPasswordManagerMockComparePassword {
	if mmComparePassword.mock.funcComparePassword != nil {
		mmComparePassword.mock.t.Fatalf("PasswordManagerMock.ComparePassword mock is already set by Set")
	}

	if mmComparePassword.defaultExpectation == nil {
		mmComparePassword.defaultExpectation = &PasswordManagerMockComparePasswordExpectation{}
	}

	if mmComparePassword.defaultExpectation.params != nil {
		mmComparePassword.mock.t.Fatalf("PasswordManagerMock.ComparePassword mock is already set by Expect")
	}

	if mmComparePassword.defaultExpectation.paramPtrs == nil {
		mmComparePassword.defaultExpectation.paramPtrs = &PasswordManagerMockComparePasswordParamPtrs{}
	}
	mmComparePassword.defaultExpectation.paramPtrs.password = &password

	return mmComparePassword
}

// Inspect accepts an inspector function that has same arguments as the PasswordManager.ComparePassword
func (mmComparePassword *mPasswordManagerMockComparePassword) Inspect(f func(hash string, password string)) *mPasswordManagerMockComparePassword {
	if mmComparePassword.mock.inspectFuncComparePassword != nil {
		mmComparePassword.mock.t.Fatalf("Inspect function is already set for PasswordManagerMock.ComparePassword")
	}

	mmComparePassword.mock.inspectFuncComparePassword = f

	return mmComparePassword
}

// Return sets up results that will be returned by PasswordManager.ComparePassword
func (mmComparePassword *mPasswordManagerMockComparePassword) Return(err error) *PasswordManagerMock {
	if mmComparePassword.mock.funcComparePassword != nil {
		mmComparePassword.mock.t.Fatalf("PasswordManagerMock.ComparePassword mock is already set by Set")
	}

	if mmComparePassword.defaultExpectation == nil {
		mmComparePassword.defaultExpectation = &PasswordManagerMockComparePasswordExpectation{mock: mmComparePassword.mock}
	}
	mmComparePassword.defaultExpectation.results = &PasswordManagerMockComparePasswordResults{err}
	return mmComparePassword.mock
}

// Set uses given function f to mock the PasswordManager.ComparePassword method
func (mmComparePassword *mPasswordManagerMockComparePassword) Set(f func(hash string, password string) (err error)) *PasswordManagerMock {
	if mmComparePassword.defaultExpectation != nil {
		mmComparePassword.mock.t.Fatalf("Default expectation is already set for the PasswordManager.ComparePassword method")
	}

	if len(mmComparePassword.expectations) > 0 {
		mmComparePassword.mock.t.Fatalf("Some expectations are already set for the PasswordManager.ComparePassword method")
	}

	mmComparePassword.mock.funcComparePassword = f
	return mmComparePassword.mock
}

// When sets expectation for the PasswordManager.ComparePassword which will trigger the result defined by the following
// Then helper
func (mmComparePassword *mPasswordManagerMockComparePassword) When(hash string, password string) *PasswordManagerMockComparePasswordExpectation {
	if mmComparePassword.mock.funcComparePassword != nil {
		mmComparePassword.mock.t.Fatalf("PasswordManagerMock.ComparePassword mock is already set by Set")
	}

	expectation := &PasswordManagerMockComparePasswordExpectation{
		mock:   mmComparePassword.mock,
		params: &PasswordManagerMockComparePasswordParams{hash, password},
	}
	mmComparePassword.expectations = append(mmComparePassword.expectations, expectation)
	return expectation
}

// Then sets up PasswordManager.ComparePassword return parameters for the expectation previously defined by the When method
func (e *PasswordManagerMockComparePasswordExpectation) Then(err error) *PasswordManagerMock {
	e.results = &PasswordManagerMockComparePasswordResults{err}
	return e.mock
}

// Times sets number of times PasswordManager.ComparePassword should be invoked
func (mmComparePassword *mPasswordManagerMockComparePassword) Times(n uint64) *mPasswordManagerMockComparePassword {
	if n == 0 {
		mmComparePassword.mock.t.Fatalf("Times of PasswordManagerMock.ComparePassword mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmComparePassword.expectedInvocations, n)
	return mmComparePassword
}

func (mmComparePassword *mPasswordManagerMockComparePassword) invocationsDone() bool {
	if len(mmComparePassword.expectations) == 0 && mmComparePassword.defaultExpectation == nil && mmComparePassword.mock.funcComparePassword == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmComparePassword.mock.afterComparePasswordCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmComparePassword.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// ComparePassword implements client.PasswordManager
func (mmComparePassword *PasswordManagerMock) ComparePassword(hash string, password string) (err error) {
	mm_atomic.AddUint64(&mmComparePassword.beforeComparePasswordCounter, 1)
	defer mm_atomic.AddUint64(&mmComparePassword.afterComparePasswordCounter, 1)

	if mmComparePassword.inspectFuncComparePassword != nil {
		mmComparePassword.inspectFuncComparePassword(hash, password)
	}

	mm_params := PasswordManagerMockComparePasswordParams{hash, password}

	// Record call args
	mmComparePassword.ComparePasswordMock.mutex.Lock()
	mmComparePassword.ComparePasswordMock.callArgs = append(mmComparePassword.ComparePasswordMock.callArgs, &mm_params)
	mmComparePassword.ComparePasswordMock.mutex.Unlock()

	for _, e := range mmComparePassword.ComparePasswordMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmComparePassword.ComparePasswordMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmComparePassword.ComparePasswordMock.defaultExpectation.Counter, 1)
		mm_want := mmComparePassword.ComparePasswordMock.defaultExpectation.params
		mm_want_ptrs := mmComparePassword.ComparePasswordMock.defaultExpectation.paramPtrs

		mm_got := PasswordManagerMockComparePasswordParams{hash, password}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.hash != nil && !minimock.Equal(*mm_want_ptrs.hash, mm_got.hash) {
				mmComparePassword.t.Errorf("PasswordManagerMock.ComparePassword got unexpected parameter hash, want: %#v, got: %#v%s\n", *mm_want_ptrs.hash, mm_got.hash, minimock.Diff(*mm_want_ptrs.hash, mm_got.hash))
			}

			if mm_want_ptrs.password != nil && !minimock.Equal(*mm_want_ptrs.password, mm_got.password) {
				mmComparePassword.t.Errorf("PasswordManagerMock.ComparePassword got unexpected parameter password, want: %#v, got: %#v%s\n", *mm_want_ptrs.password, mm_got.password, minimock.Diff(*mm_want_ptrs.password, mm_got.password))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmComparePassword.t.Errorf("PasswordManagerMock.ComparePassword got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmComparePassword.ComparePasswordMock.defaultExpectation.results
		if mm_results == nil {
			mmComparePassword.t.Fatal("No results are set for the PasswordManagerMock.ComparePassword")
		}
		return (*mm_results).err
	}
	if mmComparePassword.funcComparePassword != nil {
		return mmComparePassword.funcComparePassword(hash, password)
	}
	mmComparePassword.t.Fatalf("Unexpected call to PasswordManagerMock.ComparePassword. %v %v", hash, password)
	return
}

// ComparePasswordAfterCounter returns a count of finished PasswordManagerMock.ComparePassword invocations
func (mmComparePassword *PasswordManagerMock) ComparePasswordAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmComparePassword.afterComparePasswordCounter)
}

// ComparePasswordBeforeCounter returns a count of PasswordManagerMock.ComparePassword invocations
func (mmComparePassword *PasswordManagerMock) ComparePasswordBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmComparePassword.beforeComparePasswordCounter)
}

// Calls returns a list of arguments used in each call to PasswordManagerMock.ComparePassword.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmComparePassword *mPasswordManagerMockComparePassword) Calls() []*PasswordManagerMockComparePasswordParams {
	mmComparePassword.mutex.RLock()

	argCopy := make([]*PasswordManagerMockComparePasswordParams, len(mmComparePassword.callArgs))
	copy(argCopy, mmComparePassword.callArgs)

	mmComparePassword.mutex.RUnlock()

	return argCopy
}

// MinimockComparePasswordDone returns true if the count of the ComparePassword invocations corresponds
// the number of defined expectations
func (m *PasswordManagerMock) MinimockComparePasswordDone() bool {
	if m.ComparePasswordMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ComparePasswordMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ComparePasswordMock.invocationsDone()
}

// MinimockComparePasswordInspect logs each unmet expectation
func (m *PasswordManagerMock) MinimockComparePasswordInspect() {
	for _, e := range m.ComparePasswordMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PasswordManagerMock.ComparePassword with params: %#v", *e.params)
		}
	}

	afterComparePasswordCounter := mm_atomic.LoadUint64(&m.afterComparePasswordCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ComparePasswordMock.defaultExpectation != nil && afterComparePasswordCounter < 1 {
		if m.ComparePasswordMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PasswordManagerMock.ComparePassword")
		} else {
			m.t.Errorf("Expected call to PasswordManagerMock.ComparePassword with params: %#v", *m.ComparePasswordMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcComparePassword != nil && afterComparePasswordCounter < 1 {
		m.t.Error("Expected call to PasswordManagerMock.ComparePassword")
	}

	if !m.ComparePasswordMock.invocationsDone() && afterComparePasswordCounter > 0 {
		m.t.Errorf("Expected %d calls to PasswordManagerMock.ComparePassword but found %d calls",
			mm_atomic.LoadUint64(&m.ComparePasswordMock.expectedInvocations), afterComparePasswordCounter)
	}
}

type mPasswordManagerMockHashPassword struct {
	optional           bool
	mock               *PasswordManagerMock
	defaultExpectation *PasswordManagerMockHashPasswordExpectation
	expectations       []*PasswordManagerMockHashPasswordExpectation

	callArgs []*PasswordManagerMockHashPasswordParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// PasswordManagerMockHashPasswordExpectation specifies expectation struct of the PasswordManager.HashPassword
type PasswordManagerMockHashPasswordExpectation struct {
	mock      *PasswordManagerMock
	params    *PasswordManagerMockHashPasswordParams
	paramPtrs *PasswordManagerMockHashPasswordParamPtrs
	results   *PasswordManagerMockHashPasswordResults
	Counter   uint64
}

// PasswordManagerMockHashPasswordParams contains parameters of the PasswordManager.HashPassword
type PasswordManagerMockHashPasswordParams struct {
	password string
}

// PasswordManagerMockHashPasswordParamPtrs contains pointers to parameters of the PasswordManager.HashPassword
type PasswordManagerMockHashPasswordParamPtrs struct {
	password *string
}

// PasswordManagerMockHashPasswordResults contains results of the PasswordManager.HashPassword
type PasswordManagerMockHashPasswordResults struct {
	s1  string
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmHashPassword *mPasswordManagerMockHashPassword) Optional() *mPasswordManagerMockHashPassword {
	mmHashPassword.optional = true
	return mmHashPassword
}

// Expect sets up expected params for PasswordManager.HashPassword
func (mmHashPassword *mPasswordManagerMockHashPassword) Expect(password string) *mPasswordManagerMockHashPassword {
	if mmHashPassword.mock.funcHashPassword != nil {
		mmHashPassword.mock.t.Fatalf("PasswordManagerMock.HashPassword mock is already set by Set")
	}

	if mmHashPassword.defaultExpectation == nil {
		mmHashPassword.defaultExpectation = &PasswordManagerMockHashPasswordExpectation{}
	}

	if mmHashPassword.defaultExpectation.paramPtrs != nil {
		mmHashPassword.mock.t.Fatalf("PasswordManagerMock.HashPassword mock is already set by ExpectParams functions")
	}

	mmHashPassword.defaultExpectation.params = &PasswordManagerMockHashPasswordParams{password}
	for _, e := range mmHashPassword.expectations {
		if minimock.Equal(e.params, mmHashPassword.defaultExpectation.params) {
			mmHashPassword.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmHashPassword.defaultExpectation.params)
		}
	}

	return mmHashPassword
}

// ExpectPasswordParam1 sets up expected param password for PasswordManager.HashPassword
func (mmHashPassword *mPasswordManagerMockHashPassword) ExpectPasswordParam1(password string) *mPasswordManagerMockHashPassword {
	if mmHashPassword.mock.funcHashPassword != nil {
		mmHashPassword.mock.t.Fatalf("PasswordManagerMock.HashPassword mock is already set by Set")
	}

	if mmHashPassword.defaultExpectation == nil {
		mmHashPassword.defaultExpectation = &PasswordManagerMockHashPasswordExpectation{}
	}

	if mmHashPassword.defaultExpectation.params != nil {
		mmHashPassword.mock.t.Fatalf("PasswordManagerMock.HashPassword mock is already set by Expect")
	}

	if mmHashPassword.defaultExpectation.paramPtrs == nil {
		mmHashPassword.defaultExpectation.paramPtrs = &PasswordManagerMockHashPasswordParamPtrs{}
	}
	mmHashPassword.defaultExpectation.paramPtrs.password = &password

	return mmHashPassword
}

// Inspect accepts an inspector function that has same arguments as the PasswordManager.HashPassword
func (mmHashPassword *mPasswordManagerMockHashPassword) Inspect(f func(password string)) *mPasswordManagerMockHashPassword {
	if mmHashPassword.mock.inspectFuncHashPassword != nil {
		mmHashPassword.mock.t.Fatalf("Inspect function is already set for PasswordManagerMock.HashPassword")
	}

	mmHashPassword.mock.inspectFuncHashPassword = f

	return mmHashPassword
}

// Return sets up results that will be returned by PasswordManager.HashPassword
func (mmHashPassword *mPasswordManagerMockHashPassword) Return(s1 string, err error) *PasswordManagerMock {
	if mmHashPassword.mock.funcHashPassword != nil {
		mmHashPassword.mock.t.Fatalf("PasswordManagerMock.HashPassword mock is already set by Set")
	}

	if mmHashPassword.defaultExpectation == nil {
		mmHashPassword.defaultExpectation = &PasswordManagerMockHashPasswordExpectation{mock: mmHashPassword.mock}
	}
	mmHashPassword.defaultExpectation.results = &PasswordManagerMockHashPasswordResults{s1, err}
	return mmHashPassword.mock
}

// Set uses given function f to mock the PasswordManager.HashPassword method
func (mmHashPassword *mPasswordManagerMockHashPassword) Set(f func(password string) (s1 string, err error)) *PasswordManagerMock {
	if mmHashPassword.defaultExpectation != nil {
		mmHashPassword.mock.t.Fatalf("Default expectation is already set for the PasswordManager.HashPassword method")
	}

	if len(mmHashPassword.expectations) > 0 {
		mmHashPassword.mock.t.Fatalf("Some expectations are already set for the PasswordManager.HashPassword method")
	}

	mmHashPassword.mock.funcHashPassword = f
	return mmHashPassword.mock
}

// When sets expectation for the PasswordManager.HashPassword which will trigger the result defined by the following
// Then helper
func (mmHashPassword *mPasswordManagerMockHashPassword) When(password string) *PasswordManagerMockHashPasswordExpectation {
	if mmHashPassword.mock.funcHashPassword != nil {
		mmHashPassword.mock.t.Fatalf("PasswordManagerMock.HashPassword mock is already set by Set")
	}

	expectation := &PasswordManagerMockHashPasswordExpectation{
		mock:   mmHashPassword.mock,
		params: &PasswordManagerMockHashPasswordParams{password},
	}
	mmHashPassword.expectations = append(mmHashPassword.expectations, expectation)
	return expectation
}

// Then sets up PasswordManager.HashPassword return parameters for the expectation previously defined by the When method
func (e *PasswordManagerMockHashPasswordExpectation) Then(s1 string, err error) *PasswordManagerMock {
	e.results = &PasswordManagerMockHashPasswordResults{s1, err}
	return e.mock
}

// Times sets number of times PasswordManager.HashPassword should be invoked
func (mmHashPassword *mPasswordManagerMockHashPassword) Times(n uint64) *mPasswordManagerMockHashPassword {
	if n == 0 {
		mmHashPassword.mock.t.Fatalf("Times of PasswordManagerMock.HashPassword mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmHashPassword.expectedInvocations, n)
	return mmHashPassword
}

func (mmHashPassword *mPasswordManagerMockHashPassword) invocationsDone() bool {
	if len(mmHashPassword.expectations) == 0 && mmHashPassword.defaultExpectation == nil && mmHashPassword.mock.funcHashPassword == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmHashPassword.mock.afterHashPasswordCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmHashPassword.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// HashPassword implements client.PasswordManager
func (mmHashPassword *PasswordManagerMock) HashPassword(password string) (s1 string, err error) {
	mm_atomic.AddUint64(&mmHashPassword.beforeHashPasswordCounter, 1)
	defer mm_atomic.AddUint64(&mmHashPassword.afterHashPasswordCounter, 1)

	if mmHashPassword.inspectFuncHashPassword != nil {
		mmHashPassword.inspectFuncHashPassword(password)
	}

	mm_params := PasswordManagerMockHashPasswordParams{password}

	// Record call args
	mmHashPassword.HashPasswordMock.mutex.Lock()
	mmHashPassword.HashPasswordMock.callArgs = append(mmHashPassword.HashPasswordMock.callArgs, &mm_params)
	mmHashPassword.HashPasswordMock.mutex.Unlock()

	for _, e := range mmHashPassword.HashPasswordMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmHashPassword.HashPasswordMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmHashPassword.HashPasswordMock.defaultExpectation.Counter, 1)
		mm_want := mmHashPassword.HashPasswordMock.defaultExpectation.params
		mm_want_ptrs := mmHashPassword.HashPasswordMock.defaultExpectation.paramPtrs

		mm_got := PasswordManagerMockHashPasswordParams{password}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.password != nil && !minimock.Equal(*mm_want_ptrs.password, mm_got.password) {
				mmHashPassword.t.Errorf("PasswordManagerMock.HashPassword got unexpected parameter password, want: %#v, got: %#v%s\n", *mm_want_ptrs.password, mm_got.password, minimock.Diff(*mm_want_ptrs.password, mm_got.password))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmHashPassword.t.Errorf("PasswordManagerMock.HashPassword got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmHashPassword.HashPasswordMock.defaultExpectation.results
		if mm_results == nil {
			mmHashPassword.t.Fatal("No results are set for the PasswordManagerMock.HashPassword")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmHashPassword.funcHashPassword != nil {
		return mmHashPassword.funcHashPassword(password)
	}
	mmHashPassword.t.Fatalf("Unexpected call to PasswordManagerMock.HashPassword. %v", password)
	return
}

// HashPasswordAfterCounter returns a count of finished PasswordManagerMock.HashPassword invocations
func (mmHashPassword *PasswordManagerMock) HashPasswordAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHashPassword.afterHashPasswordCounter)
}

// HashPasswordBeforeCounter returns a count of PasswordManagerMock.HashPassword invocations
func (mmHashPassword *PasswordManagerMock) HashPasswordBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmHashPassword.beforeHashPasswordCounter)
}

// Calls returns a list of arguments used in each call to PasswordManagerMock.HashPassword.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmHashPassword *mPasswordManagerMockHashPassword) Calls() []*PasswordManagerMockHashPasswordParams {
	mmHashPassword.mutex.RLock()

	argCopy := make([]*PasswordManagerMockHashPasswordParams, len(mmHashPassword.callArgs))
	copy(argCopy, mmHashPassword.callArgs)

	mmHashPassword.mutex.RUnlock()

	return argCopy
}

// MinimockHashPasswordDone returns true if the count of the HashPassword invocations corresponds
// the number of defined expectations
func (m *PasswordManagerMock) MinimockHashPasswordDone() bool {
	if m.HashPasswordMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.HashPasswordMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.HashPasswordMock.invocationsDone()
}

// MinimockHashPasswordInspect logs each unmet expectation
func (m *PasswordManagerMock) MinimockHashPasswordInspect() {
	for _, e := range m.HashPasswordMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PasswordManagerMock.HashPassword with params: %#v", *e.params)
		}
	}

	afterHashPasswordCounter := mm_atomic.LoadUint64(&m.afterHashPasswordCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.HashPasswordMock.defaultExpectation != nil && afterHashPasswordCounter < 1 {
		if m.HashPasswordMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PasswordManagerMock.HashPassword")
		} else {
			m.t.Errorf("Expected call to PasswordManagerMock.HashPassword with params: %#v", *m.HashPasswordMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcHashPassword != nil && afterHashPasswordCounter < 1 {
		m.t.Error("Expected call to PasswordManagerMock.HashPassword")
	}

	if !m.HashPasswordMock.invocationsDone() && afterHashPasswordCounter > 0 {
		m.t.Errorf("Expected %d calls to PasswordManagerMock.HashPassword but found %d calls",
			mm_atomic.LoadUint64(&m.HashPasswordMock.expectedInvocations), afterHashPasswordCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PasswordManagerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockComparePasswordInspect()

			m.MinimockHashPasswordInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PasswordManagerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *PasswordManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockComparePasswordDone() &&
		m.MinimockHashPasswordDone()
}
