// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package hotkey

var (
	MockRegister   func(fsModifers, vk uint32, handle func()) (Id, error)
	MockUnregister func(id int32)
	MockStop       func()
	MockIsStop     func() bool
)

func init() {
	newServer = func() server {
		svr := mockServer{MockRegister, MockUnregister, MockStop, MockIsStop}
		MockRegister = nil
		MockUnregister = nil
		MockStop = nil
		MockIsStop = nil
		return &svr
	}
}

type mockServer struct {
	mockRegister   func(fsModifiers, vk uint32, handle func()) (Id, error)
	mockUnregister func(id int32)
	mockStop       func()
	mockIsStop     func() bool
}

func (mock *mockServer) register(fsModifiers, vk uint32, handle func()) (id Id, err error) {
	if mock.mockRegister != nil {
		id, err = mock.mockRegister(fsModifiers, vk, handle)
	}
	return
}

func (mock *mockServer) unregister(id int32) {
	if mock.mockRegister != nil {
		mock.mockUnregister(id)
	}
}

func (mock *mockServer) stop() {
	if mock.mockStop != nil {
		mock.mockStop()
	}
}

func (mock *mockServer) isStop() bool {
	if mock.mockIsStop != nil {
		return mock.mockIsStop()
	}
	return false
}

func (mock *mockServer) useDebugLog() {
	// nothing
}
