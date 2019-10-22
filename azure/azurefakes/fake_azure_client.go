// Code generated by counterfeiter. DO NOT EDIT.
package azurefakes

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/pivotal-cf/azure-blobstore-resource/azure"
)

type FakeAzureClient struct {
	CreateSnapshotStub        func(string) (time.Time, error)
	createSnapshotMutex       sync.RWMutex
	createSnapshotArgsForCall []struct {
		arg1 string
	}
	createSnapshotReturns struct {
		result1 time.Time
		result2 error
	}
	createSnapshotReturnsOnCall map[int]struct {
		result1 time.Time
		result2 error
	}
	DownloadBlobToFileStub        func(string, *os.File, int64, *time.Time) error
	downloadBlobToFileMutex       sync.RWMutex
	downloadBlobToFileArgsForCall []struct {
		arg1 string
		arg2 *os.File
		arg3 int64
		arg4 *time.Time
	}
	downloadBlobToFileReturns struct {
		result1 error
	}
	downloadBlobToFileReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string, time.Time) ([]byte, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 time.Time
	}
	getReturns struct {
		result1 []byte
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetBlobSizeInBytesStub        func(string, time.Time) (int64, error)
	getBlobSizeInBytesMutex       sync.RWMutex
	getBlobSizeInBytesArgsForCall []struct {
		arg1 string
		arg2 time.Time
	}
	getBlobSizeInBytesReturns struct {
		result1 int64
		result2 error
	}
	getBlobSizeInBytesReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetBlobURLStub        func(string) (string, error)
	getBlobURLMutex       sync.RWMutex
	getBlobURLArgsForCall []struct {
		arg1 string
	}
	getBlobURLReturns struct {
		result1 string
		result2 error
	}
	getBlobURLReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ListBlobsStub        func(storage.ListBlobsParameters) (storage.BlobListResponse, error)
	listBlobsMutex       sync.RWMutex
	listBlobsArgsForCall []struct {
		arg1 storage.ListBlobsParameters
	}
	listBlobsReturns struct {
		result1 storage.BlobListResponse
		result2 error
	}
	listBlobsReturnsOnCall map[int]struct {
		result1 storage.BlobListResponse
		result2 error
	}
	UploadFromStreamStub        func(string, int, io.Reader) error
	uploadFromStreamMutex       sync.RWMutex
	uploadFromStreamArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 io.Reader
	}
	uploadFromStreamReturns struct {
		result1 error
	}
	uploadFromStreamReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAzureClient) CreateSnapshot(arg1 string) (time.Time, error) {
	fake.createSnapshotMutex.Lock()
	ret, specificReturn := fake.createSnapshotReturnsOnCall[len(fake.createSnapshotArgsForCall)]
	fake.createSnapshotArgsForCall = append(fake.createSnapshotArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CreateSnapshot", []interface{}{arg1})
	fake.createSnapshotMutex.Unlock()
	if fake.CreateSnapshotStub != nil {
		return fake.CreateSnapshotStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createSnapshotReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAzureClient) CreateSnapshotCallCount() int {
	fake.createSnapshotMutex.RLock()
	defer fake.createSnapshotMutex.RUnlock()
	return len(fake.createSnapshotArgsForCall)
}

func (fake *FakeAzureClient) CreateSnapshotCalls(stub func(string) (time.Time, error)) {
	fake.createSnapshotMutex.Lock()
	defer fake.createSnapshotMutex.Unlock()
	fake.CreateSnapshotStub = stub
}

func (fake *FakeAzureClient) CreateSnapshotArgsForCall(i int) string {
	fake.createSnapshotMutex.RLock()
	defer fake.createSnapshotMutex.RUnlock()
	argsForCall := fake.createSnapshotArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAzureClient) CreateSnapshotReturns(result1 time.Time, result2 error) {
	fake.createSnapshotMutex.Lock()
	defer fake.createSnapshotMutex.Unlock()
	fake.CreateSnapshotStub = nil
	fake.createSnapshotReturns = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) CreateSnapshotReturnsOnCall(i int, result1 time.Time, result2 error) {
	fake.createSnapshotMutex.Lock()
	defer fake.createSnapshotMutex.Unlock()
	fake.CreateSnapshotStub = nil
	if fake.createSnapshotReturnsOnCall == nil {
		fake.createSnapshotReturnsOnCall = make(map[int]struct {
			result1 time.Time
			result2 error
		})
	}
	fake.createSnapshotReturnsOnCall[i] = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) DownloadBlobToFile(arg1 string, arg2 *os.File, arg3 int64, arg4 *time.Time) error {
	fake.downloadBlobToFileMutex.Lock()
	ret, specificReturn := fake.downloadBlobToFileReturnsOnCall[len(fake.downloadBlobToFileArgsForCall)]
	fake.downloadBlobToFileArgsForCall = append(fake.downloadBlobToFileArgsForCall, struct {
		arg1 string
		arg2 *os.File
		arg3 int64
		arg4 *time.Time
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DownloadBlobToFile", []interface{}{arg1, arg2, arg3, arg4})
	fake.downloadBlobToFileMutex.Unlock()
	if fake.DownloadBlobToFileStub != nil {
		return fake.DownloadBlobToFileStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadBlobToFileReturns
	return fakeReturns.result1
}

func (fake *FakeAzureClient) DownloadBlobToFileCallCount() int {
	fake.downloadBlobToFileMutex.RLock()
	defer fake.downloadBlobToFileMutex.RUnlock()
	return len(fake.downloadBlobToFileArgsForCall)
}

func (fake *FakeAzureClient) DownloadBlobToFileCalls(stub func(string, *os.File, int64, *time.Time) error) {
	fake.downloadBlobToFileMutex.Lock()
	defer fake.downloadBlobToFileMutex.Unlock()
	fake.DownloadBlobToFileStub = stub
}

func (fake *FakeAzureClient) DownloadBlobToFileArgsForCall(i int) (string, *os.File, int64, *time.Time) {
	fake.downloadBlobToFileMutex.RLock()
	defer fake.downloadBlobToFileMutex.RUnlock()
	argsForCall := fake.downloadBlobToFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeAzureClient) DownloadBlobToFileReturns(result1 error) {
	fake.downloadBlobToFileMutex.Lock()
	defer fake.downloadBlobToFileMutex.Unlock()
	fake.DownloadBlobToFileStub = nil
	fake.downloadBlobToFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAzureClient) DownloadBlobToFileReturnsOnCall(i int, result1 error) {
	fake.downloadBlobToFileMutex.Lock()
	defer fake.downloadBlobToFileMutex.Unlock()
	fake.DownloadBlobToFileStub = nil
	if fake.downloadBlobToFileReturnsOnCall == nil {
		fake.downloadBlobToFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadBlobToFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAzureClient) Get(arg1 string, arg2 time.Time) ([]byte, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 time.Time
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAzureClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeAzureClient) GetCalls(stub func(string, time.Time) ([]byte, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeAzureClient) GetArgsForCall(i int) (string, time.Time) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAzureClient) GetReturns(result1 []byte, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) GetReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) GetBlobSizeInBytes(arg1 string, arg2 time.Time) (int64, error) {
	fake.getBlobSizeInBytesMutex.Lock()
	ret, specificReturn := fake.getBlobSizeInBytesReturnsOnCall[len(fake.getBlobSizeInBytesArgsForCall)]
	fake.getBlobSizeInBytesArgsForCall = append(fake.getBlobSizeInBytesArgsForCall, struct {
		arg1 string
		arg2 time.Time
	}{arg1, arg2})
	fake.recordInvocation("GetBlobSizeInBytes", []interface{}{arg1, arg2})
	fake.getBlobSizeInBytesMutex.Unlock()
	if fake.GetBlobSizeInBytesStub != nil {
		return fake.GetBlobSizeInBytesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBlobSizeInBytesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAzureClient) GetBlobSizeInBytesCallCount() int {
	fake.getBlobSizeInBytesMutex.RLock()
	defer fake.getBlobSizeInBytesMutex.RUnlock()
	return len(fake.getBlobSizeInBytesArgsForCall)
}

func (fake *FakeAzureClient) GetBlobSizeInBytesCalls(stub func(string, time.Time) (int64, error)) {
	fake.getBlobSizeInBytesMutex.Lock()
	defer fake.getBlobSizeInBytesMutex.Unlock()
	fake.GetBlobSizeInBytesStub = stub
}

func (fake *FakeAzureClient) GetBlobSizeInBytesArgsForCall(i int) (string, time.Time) {
	fake.getBlobSizeInBytesMutex.RLock()
	defer fake.getBlobSizeInBytesMutex.RUnlock()
	argsForCall := fake.getBlobSizeInBytesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAzureClient) GetBlobSizeInBytesReturns(result1 int64, result2 error) {
	fake.getBlobSizeInBytesMutex.Lock()
	defer fake.getBlobSizeInBytesMutex.Unlock()
	fake.GetBlobSizeInBytesStub = nil
	fake.getBlobSizeInBytesReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) GetBlobSizeInBytesReturnsOnCall(i int, result1 int64, result2 error) {
	fake.getBlobSizeInBytesMutex.Lock()
	defer fake.getBlobSizeInBytesMutex.Unlock()
	fake.GetBlobSizeInBytesStub = nil
	if fake.getBlobSizeInBytesReturnsOnCall == nil {
		fake.getBlobSizeInBytesReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.getBlobSizeInBytesReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) GetBlobURL(arg1 string) (string, error) {
	fake.getBlobURLMutex.Lock()
	ret, specificReturn := fake.getBlobURLReturnsOnCall[len(fake.getBlobURLArgsForCall)]
	fake.getBlobURLArgsForCall = append(fake.getBlobURLArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetBlobURL", []interface{}{arg1})
	fake.getBlobURLMutex.Unlock()
	if fake.GetBlobURLStub != nil {
		return fake.GetBlobURLStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBlobURLReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAzureClient) GetBlobURLCallCount() int {
	fake.getBlobURLMutex.RLock()
	defer fake.getBlobURLMutex.RUnlock()
	return len(fake.getBlobURLArgsForCall)
}

func (fake *FakeAzureClient) GetBlobURLCalls(stub func(string) (string, error)) {
	fake.getBlobURLMutex.Lock()
	defer fake.getBlobURLMutex.Unlock()
	fake.GetBlobURLStub = stub
}

func (fake *FakeAzureClient) GetBlobURLArgsForCall(i int) string {
	fake.getBlobURLMutex.RLock()
	defer fake.getBlobURLMutex.RUnlock()
	argsForCall := fake.getBlobURLArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAzureClient) GetBlobURLReturns(result1 string, result2 error) {
	fake.getBlobURLMutex.Lock()
	defer fake.getBlobURLMutex.Unlock()
	fake.GetBlobURLStub = nil
	fake.getBlobURLReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) GetBlobURLReturnsOnCall(i int, result1 string, result2 error) {
	fake.getBlobURLMutex.Lock()
	defer fake.getBlobURLMutex.Unlock()
	fake.GetBlobURLStub = nil
	if fake.getBlobURLReturnsOnCall == nil {
		fake.getBlobURLReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getBlobURLReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) ListBlobs(arg1 storage.ListBlobsParameters) (storage.BlobListResponse, error) {
	fake.listBlobsMutex.Lock()
	ret, specificReturn := fake.listBlobsReturnsOnCall[len(fake.listBlobsArgsForCall)]
	fake.listBlobsArgsForCall = append(fake.listBlobsArgsForCall, struct {
		arg1 storage.ListBlobsParameters
	}{arg1})
	fake.recordInvocation("ListBlobs", []interface{}{arg1})
	fake.listBlobsMutex.Unlock()
	if fake.ListBlobsStub != nil {
		return fake.ListBlobsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listBlobsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAzureClient) ListBlobsCallCount() int {
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	return len(fake.listBlobsArgsForCall)
}

func (fake *FakeAzureClient) ListBlobsCalls(stub func(storage.ListBlobsParameters) (storage.BlobListResponse, error)) {
	fake.listBlobsMutex.Lock()
	defer fake.listBlobsMutex.Unlock()
	fake.ListBlobsStub = stub
}

func (fake *FakeAzureClient) ListBlobsArgsForCall(i int) storage.ListBlobsParameters {
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	argsForCall := fake.listBlobsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAzureClient) ListBlobsReturns(result1 storage.BlobListResponse, result2 error) {
	fake.listBlobsMutex.Lock()
	defer fake.listBlobsMutex.Unlock()
	fake.ListBlobsStub = nil
	fake.listBlobsReturns = struct {
		result1 storage.BlobListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) ListBlobsReturnsOnCall(i int, result1 storage.BlobListResponse, result2 error) {
	fake.listBlobsMutex.Lock()
	defer fake.listBlobsMutex.Unlock()
	fake.ListBlobsStub = nil
	if fake.listBlobsReturnsOnCall == nil {
		fake.listBlobsReturnsOnCall = make(map[int]struct {
			result1 storage.BlobListResponse
			result2 error
		})
	}
	fake.listBlobsReturnsOnCall[i] = struct {
		result1 storage.BlobListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeAzureClient) UploadFromStream(arg1 string, arg2 int, arg3 io.Reader) error {
	fake.uploadFromStreamMutex.Lock()
	ret, specificReturn := fake.uploadFromStreamReturnsOnCall[len(fake.uploadFromStreamArgsForCall)]
	fake.uploadFromStreamArgsForCall = append(fake.uploadFromStreamArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 io.Reader
	}{arg1, arg2, arg3})
	fake.recordInvocation("UploadFromStream", []interface{}{arg1, arg2, arg3})
	fake.uploadFromStreamMutex.Unlock()
	if fake.UploadFromStreamStub != nil {
		return fake.UploadFromStreamStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uploadFromStreamReturns
	return fakeReturns.result1
}

func (fake *FakeAzureClient) UploadFromStreamCallCount() int {
	fake.uploadFromStreamMutex.RLock()
	defer fake.uploadFromStreamMutex.RUnlock()
	return len(fake.uploadFromStreamArgsForCall)
}

func (fake *FakeAzureClient) UploadFromStreamCalls(stub func(string, int, io.Reader) error) {
	fake.uploadFromStreamMutex.Lock()
	defer fake.uploadFromStreamMutex.Unlock()
	fake.UploadFromStreamStub = stub
}

func (fake *FakeAzureClient) UploadFromStreamArgsForCall(i int) (string, int, io.Reader) {
	fake.uploadFromStreamMutex.RLock()
	defer fake.uploadFromStreamMutex.RUnlock()
	argsForCall := fake.uploadFromStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAzureClient) UploadFromStreamReturns(result1 error) {
	fake.uploadFromStreamMutex.Lock()
	defer fake.uploadFromStreamMutex.Unlock()
	fake.UploadFromStreamStub = nil
	fake.uploadFromStreamReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAzureClient) UploadFromStreamReturnsOnCall(i int, result1 error) {
	fake.uploadFromStreamMutex.Lock()
	defer fake.uploadFromStreamMutex.Unlock()
	fake.UploadFromStreamStub = nil
	if fake.uploadFromStreamReturnsOnCall == nil {
		fake.uploadFromStreamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadFromStreamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAzureClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createSnapshotMutex.RLock()
	defer fake.createSnapshotMutex.RUnlock()
	fake.downloadBlobToFileMutex.RLock()
	defer fake.downloadBlobToFileMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getBlobSizeInBytesMutex.RLock()
	defer fake.getBlobSizeInBytesMutex.RUnlock()
	fake.getBlobURLMutex.RLock()
	defer fake.getBlobURLMutex.RUnlock()
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	fake.uploadFromStreamMutex.RLock()
	defer fake.uploadFromStreamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAzureClient) recordInvocation(key string, args []interface{}) {
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

var _ azure.AzureClient = new(FakeAzureClient)