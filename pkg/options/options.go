package options

import (
	"fmt"
	"os"
)

type Options struct {
	MachineID     string
	MachineFolder string

	Project     string
	Zone        string
	Network     string
	Subnetwork  string
	Tag         string
	DiskSize    string
	DiskImage   string
	MachineType string
}

func FromEnv(withMachine bool) (*Options, error) {
	retOptions := &Options{}

	var err error
	if withMachine {
		retOptions.MachineID, err = fromEnvOrError("MACHINE_ID")
		if err != nil {
			return nil, err
		}
		// prefix with devpod-
		retOptions.MachineID = "devpod-" + retOptions.MachineID

		retOptions.MachineFolder, err = fromEnvOrError("MACHINE_FOLDER")
		if err != nil {
			return nil, err
		}
	}

	retOptions.Project, err = fromEnvOrError("PROJECT")
	if err != nil {
		return nil, err
	}
	retOptions.Zone, err = fromEnvOrError("ZONE")
	if err != nil {
		return nil, err
	}
	retOptions.DiskSize, err = fromEnvOrError("DISK_SIZE")
	if err != nil {
		return nil, err
	}
	retOptions.DiskImage, err = fromEnvOrError("DISK_IMAGE")
	if err != nil {
		return nil, err
	}
	retOptions.MachineType, err = fromEnvOrError("MACHINE_TYPE")
	if err != nil {
		return nil, err
	}

	retOptions.Network = os.Getenv("NETWORK")
	retOptions.Subnetwork = os.Getenv("SUBNETWORK")
	retOptions.Tag = os.Getenv("TAG")

	return retOptions, nil
}

func fromEnvOrError(name string) (string, error) {
	val := os.Getenv(name)
	if val == "" {
		return "", fmt.Errorf("couldn't find option %s in environment, please make sure %s is defined", name, name)
	}

	return val, nil
}
