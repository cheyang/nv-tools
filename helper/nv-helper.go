package helper

import

//nvdiaDocker "github.com/NVIDIA/nvidia-docker/tools/src/docker"
(
	"fmt"

	"github.com/NVIDIA/nvidia-docker/tools/src/nvidia"
)

type NvidiaHelper struct {
	Devices []nvidia.Device
	NumaMap map[uint]uint // dev path: numa node
}

func NewNvHelper() (*NvidiaHelper, error) {

	if err := nvidia.LoadUVM(); err != nil {
		return nil, err
	}
	if err := nvidia.Init(); err != nil {
		return nil, err
	}

	nv := &NvidiaHelper{}
	if err := nv.detect(); err != nil {
		return nil, err
	}

	return nv, nil
}

func (this *NvidiaHelper) detect() (err error) {

	if len(this.Devices) == 0 {
		this.Devices, err = nvidia.LookupDevices()

		if err != nil {
			// log.Warningf("Cannot detect the nvidia gpu %q", err)
			return fmt.Errorf("Cannot detect the nvidia gpu %q", err)
		}

		for i, dev := range this.Devices {
			this.NumaMap[uint(i)] = *dev.CPUAffinity
		}
	}

	return nil
}
