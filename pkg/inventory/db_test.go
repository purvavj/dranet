/*
Copyright The Kubernetes Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package inventory

import (
	"testing"

	"github.com/jaypipes/ghw"
)

func TestIsAllocatableNetworkDevice(t *testing.T) {
	cases := []struct {
		name   string
		driver string
		want   bool
	}{
		{name: "unbound", driver: "", want: false},
		{name: "vfio passthrough", driver: "vfio-pci", want: false},
		{name: "uio generic", driver: "uio_pci_generic", want: false},
		{name: "dpdk uio", driver: "igb_uio", want: false},
		{name: "pci stub", driver: "pci-stub", want: false},
		{name: "gve", driver: "gve", want: true},
		{name: "mlx5", driver: "mlx5_core", want: true},
		{name: "virtio", driver: "virtio_net", want: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dev := &ghw.PCIDevice{Driver: tc.driver}
			if got := isAllocatableNetworkDevice(dev); got != tc.want {
				t.Errorf("isAllocatableNetworkDevice(driver=%q) = %v, want %v", tc.driver, got, tc.want)
			}
		})
	}
}
