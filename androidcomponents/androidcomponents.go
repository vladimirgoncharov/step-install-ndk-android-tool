package androidcomponents

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitrise-io/go-android/sdk"
	"github.com/bitrise-io/go-android/sdkmanager"
	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-io/go-utils/pathutil"
)

// InstallLicences ...
func InstallLicences(androidSdk *sdk.Model) error {
	sdkManager, err := sdkmanager.New(androidSdk)
	if err != nil {
		return err
	}

	licencesDir, licenceMap := filepath.Join(androidSdk.GetAndroidHome(), "licenses"), map[string]string{
		"android-sdk-license":           "\n24333f8a63b6825ea9c5514f83c2829b004d1fee",
		"android-googletv-license":      "\n601085b94cd77f0b54ff86406957099ebe79c4d6",
		"android-sdk-preview-license":   "\n84831b9409646a918e30573bab4c9c91346d8abd",
		"intel-android-extra-license":   "\nd975f751698a77b662f1254ddbeed3901e976f5a",
		"google-gdk-license":            "\n33b6a2b64607f11b759f320ef9dff4ae5c47d97a",
		"mips-android-sysimage-license": "\ne9acab5b5fbb560a72cfaecce8946896ff6aab9d",
	}

	if !sdkManager.IsLegacySDK() {
		cmdLineToolsPath, err := androidSdk.CmdlineToolsPath()
		if err != nil {
			return err
		}
		licensesCmd := command.New(filepath.Join(cmdLineToolsPath, "sdkmanager"), "--licenses")
		licensesCmd.SetStdin(bytes.NewReader([]byte(strings.Repeat("y\n", 1000))))
		if err := licensesCmd.Run(); err != nil {
			log.Warnf("Failed to install licenses using $(sdkmanager --licenses) command")
			log.Printf("Continue using legacy license installation...")
			log.Printf("")
		} else {
			sdkLicencePath, oldLicenceHash := filepath.Join(licencesDir, "android-sdk-license"), "d56f5187479451eabf01fb78af6dfcb131a6481e"
			if content, err := fileutil.ReadStringFromFile(sdkLicencePath); err == nil && strings.Contains(content, oldLicenceHash) {
				if err := fileutil.WriteStringToFile(sdkLicencePath, licenceMap[filepath.Base(sdkLicencePath)]); err != nil {
					return err
				}
			}
			return nil
		}
	}

	if exist, err := pathutil.IsDirExists(licencesDir); err != nil {
		return err
	} else if !exist {
		if err := os.MkdirAll(licencesDir, os.ModePerm); err != nil {
			return err
		}
	}

	for name, content := range licenceMap {
		pth := filepath.Join(licencesDir, name)
		if err := fileutil.WriteStringToFile(pth, content); err != nil {
			return err
		}
	}

	return nil
}