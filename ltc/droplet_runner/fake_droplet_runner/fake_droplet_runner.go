// This file was generated by counterfeiter
package fake_droplet_runner

import (
	"io"
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/app_runner"
	"github.com/cloudfoundry-incubator/lattice/ltc/droplet_runner"
)

type FakeDropletRunner struct {
	UploadBitsStub        func(dropletName, uploadPath string) error
	uploadBitsMutex       sync.RWMutex
	uploadBitsArgsForCall []struct {
		dropletName string
		uploadPath  string
	}
	uploadBitsReturns struct {
		result1 error
	}
	BuildDropletStub        func(taskName, dropletName, buildpackUrl string, environment map[string]string) error
	buildDropletMutex       sync.RWMutex
	buildDropletArgsForCall []struct {
		taskName     string
		dropletName  string
		buildpackUrl string
		environment  map[string]string
	}
	buildDropletReturns struct {
		result1 error
	}
	LaunchDropletStub        func(appName, dropletName, startCommand string, startArgs []string, appEnvironmentParams app_runner.AppEnvironmentParams) error
	launchDropletMutex       sync.RWMutex
	launchDropletArgsForCall []struct {
		appName              string
		dropletName          string
		startCommand         string
		startArgs            []string
		appEnvironmentParams app_runner.AppEnvironmentParams
	}
	launchDropletReturns struct {
		result1 error
	}
	ListDropletsStub        func() ([]droplet_runner.Droplet, error)
	listDropletsMutex       sync.RWMutex
	listDropletsArgsForCall []struct{}
	listDropletsReturns     struct {
		result1 []droplet_runner.Droplet
		result2 error
	}
	RemoveDropletStub        func(dropletName string) error
	removeDropletMutex       sync.RWMutex
	removeDropletArgsForCall []struct {
		dropletName string
	}
	removeDropletReturns struct {
		result1 error
	}
	ExportDropletStub        func(dropletName string) (io.ReadCloser, io.ReadCloser, error)
	exportDropletMutex       sync.RWMutex
	exportDropletArgsForCall []struct {
		dropletName string
	}
	exportDropletReturns struct {
		result1 io.ReadCloser
		result2 io.ReadCloser
		result3 error
	}
}

func (fake *FakeDropletRunner) UploadBits(dropletName string, uploadPath string) error {
	fake.uploadBitsMutex.Lock()
	fake.uploadBitsArgsForCall = append(fake.uploadBitsArgsForCall, struct {
		dropletName string
		uploadPath  string
	}{dropletName, uploadPath})
	fake.uploadBitsMutex.Unlock()
	if fake.UploadBitsStub != nil {
		return fake.UploadBitsStub(dropletName, uploadPath)
	} else {
		return fake.uploadBitsReturns.result1
	}
}

func (fake *FakeDropletRunner) UploadBitsCallCount() int {
	fake.uploadBitsMutex.RLock()
	defer fake.uploadBitsMutex.RUnlock()
	return len(fake.uploadBitsArgsForCall)
}

func (fake *FakeDropletRunner) UploadBitsArgsForCall(i int) (string, string) {
	fake.uploadBitsMutex.RLock()
	defer fake.uploadBitsMutex.RUnlock()
	return fake.uploadBitsArgsForCall[i].dropletName, fake.uploadBitsArgsForCall[i].uploadPath
}

func (fake *FakeDropletRunner) UploadBitsReturns(result1 error) {
	fake.UploadBitsStub = nil
	fake.uploadBitsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDropletRunner) BuildDroplet(taskName string, dropletName string, buildpackUrl string, environment map[string]string) error {
	fake.buildDropletMutex.Lock()
	fake.buildDropletArgsForCall = append(fake.buildDropletArgsForCall, struct {
		taskName     string
		dropletName  string
		buildpackUrl string
		environment  map[string]string
	}{taskName, dropletName, buildpackUrl, environment})
	fake.buildDropletMutex.Unlock()
	if fake.BuildDropletStub != nil {
		return fake.BuildDropletStub(taskName, dropletName, buildpackUrl, environment)
	} else {
		return fake.buildDropletReturns.result1
	}
}

func (fake *FakeDropletRunner) BuildDropletCallCount() int {
	fake.buildDropletMutex.RLock()
	defer fake.buildDropletMutex.RUnlock()
	return len(fake.buildDropletArgsForCall)
}

func (fake *FakeDropletRunner) BuildDropletArgsForCall(i int) (string, string, string, map[string]string) {
	fake.buildDropletMutex.RLock()
	defer fake.buildDropletMutex.RUnlock()
	return fake.buildDropletArgsForCall[i].taskName, fake.buildDropletArgsForCall[i].dropletName, fake.buildDropletArgsForCall[i].buildpackUrl, fake.buildDropletArgsForCall[i].environment
}

func (fake *FakeDropletRunner) BuildDropletReturns(result1 error) {
	fake.BuildDropletStub = nil
	fake.buildDropletReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDropletRunner) LaunchDroplet(appName string, dropletName string, startCommand string, startArgs []string, appEnvironmentParams app_runner.AppEnvironmentParams) error {
	fake.launchDropletMutex.Lock()
	fake.launchDropletArgsForCall = append(fake.launchDropletArgsForCall, struct {
		appName              string
		dropletName          string
		startCommand         string
		startArgs            []string
		appEnvironmentParams app_runner.AppEnvironmentParams
	}{appName, dropletName, startCommand, startArgs, appEnvironmentParams})
	fake.launchDropletMutex.Unlock()
	if fake.LaunchDropletStub != nil {
		return fake.LaunchDropletStub(appName, dropletName, startCommand, startArgs, appEnvironmentParams)
	} else {
		return fake.launchDropletReturns.result1
	}
}

func (fake *FakeDropletRunner) LaunchDropletCallCount() int {
	fake.launchDropletMutex.RLock()
	defer fake.launchDropletMutex.RUnlock()
	return len(fake.launchDropletArgsForCall)
}

func (fake *FakeDropletRunner) LaunchDropletArgsForCall(i int) (string, string, string, []string, app_runner.AppEnvironmentParams) {
	fake.launchDropletMutex.RLock()
	defer fake.launchDropletMutex.RUnlock()
	return fake.launchDropletArgsForCall[i].appName, fake.launchDropletArgsForCall[i].dropletName, fake.launchDropletArgsForCall[i].startCommand, fake.launchDropletArgsForCall[i].startArgs, fake.launchDropletArgsForCall[i].appEnvironmentParams
}

func (fake *FakeDropletRunner) LaunchDropletReturns(result1 error) {
	fake.LaunchDropletStub = nil
	fake.launchDropletReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDropletRunner) ListDroplets() ([]droplet_runner.Droplet, error) {
	fake.listDropletsMutex.Lock()
	fake.listDropletsArgsForCall = append(fake.listDropletsArgsForCall, struct{}{})
	fake.listDropletsMutex.Unlock()
	if fake.ListDropletsStub != nil {
		return fake.ListDropletsStub()
	} else {
		return fake.listDropletsReturns.result1, fake.listDropletsReturns.result2
	}
}

func (fake *FakeDropletRunner) ListDropletsCallCount() int {
	fake.listDropletsMutex.RLock()
	defer fake.listDropletsMutex.RUnlock()
	return len(fake.listDropletsArgsForCall)
}

func (fake *FakeDropletRunner) ListDropletsReturns(result1 []droplet_runner.Droplet, result2 error) {
	fake.ListDropletsStub = nil
	fake.listDropletsReturns = struct {
		result1 []droplet_runner.Droplet
		result2 error
	}{result1, result2}
}

func (fake *FakeDropletRunner) RemoveDroplet(dropletName string) error {
	fake.removeDropletMutex.Lock()
	fake.removeDropletArgsForCall = append(fake.removeDropletArgsForCall, struct {
		dropletName string
	}{dropletName})
	fake.removeDropletMutex.Unlock()
	if fake.RemoveDropletStub != nil {
		return fake.RemoveDropletStub(dropletName)
	} else {
		return fake.removeDropletReturns.result1
	}
}

func (fake *FakeDropletRunner) RemoveDropletCallCount() int {
	fake.removeDropletMutex.RLock()
	defer fake.removeDropletMutex.RUnlock()
	return len(fake.removeDropletArgsForCall)
}

func (fake *FakeDropletRunner) RemoveDropletArgsForCall(i int) string {
	fake.removeDropletMutex.RLock()
	defer fake.removeDropletMutex.RUnlock()
	return fake.removeDropletArgsForCall[i].dropletName
}

func (fake *FakeDropletRunner) RemoveDropletReturns(result1 error) {
	fake.RemoveDropletStub = nil
	fake.removeDropletReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDropletRunner) ExportDroplet(dropletName string) (io.ReadCloser, io.ReadCloser, error) {
	fake.exportDropletMutex.Lock()
	fake.exportDropletArgsForCall = append(fake.exportDropletArgsForCall, struct {
		dropletName string
	}{dropletName})
	fake.exportDropletMutex.Unlock()
	if fake.ExportDropletStub != nil {
		return fake.ExportDropletStub(dropletName)
	} else {
		return fake.exportDropletReturns.result1, fake.exportDropletReturns.result2, fake.exportDropletReturns.result3
	}
}

func (fake *FakeDropletRunner) ExportDropletCallCount() int {
	fake.exportDropletMutex.RLock()
	defer fake.exportDropletMutex.RUnlock()
	return len(fake.exportDropletArgsForCall)
}

func (fake *FakeDropletRunner) ExportDropletArgsForCall(i int) string {
	fake.exportDropletMutex.RLock()
	defer fake.exportDropletMutex.RUnlock()
	return fake.exportDropletArgsForCall[i].dropletName
}

func (fake *FakeDropletRunner) ExportDropletReturns(result1 io.ReadCloser, result2 io.ReadCloser, result3 error) {
	fake.ExportDropletStub = nil
	fake.exportDropletReturns = struct {
		result1 io.ReadCloser
		result2 io.ReadCloser
		result3 error
	}{result1, result2, result3}
}

var _ droplet_runner.DropletRunner = new(FakeDropletRunner)
