// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/warehouse-13/hammertime/pkg/client"
	"github.com/weaveworks-liquidmetal/flintlock/api/services/microvm/v1alpha1"
	"github.com/weaveworks-liquidmetal/flintlock/api/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FakeFlintlockClient struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(*types.MicroVMSpec) (*v1alpha1.CreateMicroVMResponse, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *types.MicroVMSpec
	}
	createReturns struct {
		result1 *v1alpha1.CreateMicroVMResponse
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *v1alpha1.CreateMicroVMResponse
		result2 error
	}
	DeleteStub        func(string) (*emptypb.Empty, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 *emptypb.Empty
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *emptypb.Empty
		result2 error
	}
	GetStub        func(string) (*v1alpha1.GetMicroVMResponse, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 *v1alpha1.GetMicroVMResponse
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *v1alpha1.GetMicroVMResponse
		result2 error
	}
	ListStub        func(string, string) (*v1alpha1.ListMicroVMsResponse, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listReturns struct {
		result1 *v1alpha1.ListMicroVMsResponse
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *v1alpha1.ListMicroVMsResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFlintlockClient) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeFlintlockClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeFlintlockClient) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeFlintlockClient) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFlintlockClient) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFlintlockClient) Create(arg1 *types.MicroVMSpec) (*v1alpha1.CreateMicroVMResponse, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *types.MicroVMSpec
	}{arg1})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlintlockClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeFlintlockClient) CreateCalls(stub func(*types.MicroVMSpec) (*v1alpha1.CreateMicroVMResponse, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeFlintlockClient) CreateArgsForCall(i int) *types.MicroVMSpec {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFlintlockClient) CreateReturns(result1 *v1alpha1.CreateMicroVMResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *v1alpha1.CreateMicroVMResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeFlintlockClient) CreateReturnsOnCall(i int, result1 *v1alpha1.CreateMicroVMResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.CreateMicroVMResponse
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *v1alpha1.CreateMicroVMResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeFlintlockClient) Delete(arg1 string) (*emptypb.Empty, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlintlockClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeFlintlockClient) DeleteCalls(stub func(string) (*emptypb.Empty, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeFlintlockClient) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFlintlockClient) DeleteReturns(result1 *emptypb.Empty, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *emptypb.Empty
		result2 error
	}{result1, result2}
}

func (fake *FakeFlintlockClient) DeleteReturnsOnCall(i int, result1 *emptypb.Empty, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *emptypb.Empty
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *emptypb.Empty
		result2 error
	}{result1, result2}
}

func (fake *FakeFlintlockClient) Get(arg1 string) (*v1alpha1.GetMicroVMResponse, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlintlockClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeFlintlockClient) GetCalls(stub func(string) (*v1alpha1.GetMicroVMResponse, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeFlintlockClient) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFlintlockClient) GetReturns(result1 *v1alpha1.GetMicroVMResponse, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1alpha1.GetMicroVMResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeFlintlockClient) GetReturnsOnCall(i int, result1 *v1alpha1.GetMicroVMResponse, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.GetMicroVMResponse
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1alpha1.GetMicroVMResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeFlintlockClient) List(arg1 string, arg2 string) (*v1alpha1.ListMicroVMsResponse, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFlintlockClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeFlintlockClient) ListCalls(stub func(string, string) (*v1alpha1.ListMicroVMsResponse, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeFlintlockClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFlintlockClient) ListReturns(result1 *v1alpha1.ListMicroVMsResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *v1alpha1.ListMicroVMsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeFlintlockClient) ListReturnsOnCall(i int, result1 *v1alpha1.ListMicroVMsResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.ListMicroVMsResponse
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *v1alpha1.ListMicroVMsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeFlintlockClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFlintlockClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ client.FlintlockClient = new(FakeFlintlockClient)
