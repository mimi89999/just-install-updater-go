package rules

import (
	"strings"

	. "github.com/just-install/just-install-updater-go/jiup/rules/helpers"
)

func init() {
	AddRule(
		"7zip",
		RegexpVersionExtractor(
			"https://7-zip.org/download.html",
			Re("Download 7-Zip ([0-9][0-9].[0-9][0-9])"),
		),
		func(version string) (string, *string, error) {
			x86dl := "https://www.7-zip.org/a/7z" + strings.Replace(version, ".", "", -1) + ".msi"
			x64dl := "https://www.7-zip.org/a/7z" + strings.Replace(version, ".", "", -1) + "-x64.msi"
			return x86dl, &x64dl, nil
		},
	)
	AddRule(
		"anaconda",
		RegexpVersionExtractor(
			"https://www.anaconda.com/download/",
			Re("Anaconda3-([0-9.]+)-"),
		),
		HTMLDownloadExtractor(
			"https://www.anaconda.com/download/",
			true,
			"a[href*='Windows-x86.exe']",
			"a[href*='Windows-x86_64.exe']",
			"href",
			"href",
			Re("(.+Anaconda3-[0-9.]+-Windows-x86.exe)"),
			Re("(.+Anaconda3-[0-9.]+-Windows-x86_64.exe)"),
		),
	)
	AddRule(
		"android-studio-ide",
		RegexpVersionExtractor(
			"https://developer.android.com/studio/index.html",
			Re("Version: ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://developer.android.com/studio/index.html",
			false,
			"a#win-bundle",
			"",
			"href",
			"",
			Re("(.+android-studio-ide-[0-9.]+-windows.exe)"),
			nil,
		),
	)
	AddRule(
		"arduino",
		RegexpVersionExtractor(
			"https://www.arduino.cc/en/Main/Software",
			Re("arduino-([0-9.]+)-"),
		),
		func(version string) (string, *string, error) {
			return "https://downloads.arduino.cc/arduino-" + version + "-windows.exe", nil, nil
		},
	)
	AddRule(
		"audacity",
		RegexpVersionExtractor(
			"http://www.oldfoss.com/Audacity.html",
			Re("audacity-win-([0-9.]+).exe"),
		),
		func(version string) (string, *string, error) {
			return RegexpDownloadExtractor(
				"http://www.oldfoss.com/Audacity.html",
				Re("\"(http.+audacity-win-"+version+".exe)\""),
				nil,
			)(version)
		},
	)
	AddRule(
		"bleachbit",
		RegexpVersionExtractor(
			"https://www.bleachbit.org/download/windows",
			Re("BleachBit-([0-9.]+)-setup.exe"),
		),
		func(version string) (string, *string, error) {
			return "https://download.bleachbit.org/BleachBit-" + version + "-setup.exe", nil, nil
		},
	)
	AddRule(
		"bcuninstaller",
		GitHubReleaseVersionExtractor(
			"Klocman",
			"Bulk-Crap-Uninstaller",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"Klocman",
			"Bulk-Crap-Uninstaller",
			Re(".*setup.exe"),
			nil,
		),
	)
	AddRule(
		"bitpay",
		GitHubReleaseVersionExtractor(
			"bitpay",
			"copay",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"bitpay",
			"copay",
			Re("BitPay.exe"),
			nil,
		),
	)
	AddRule(
		"brackets",
		GitHubReleaseVersionExtractor(
			"adobe",
			"brackets",
			Re("release-(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"adobe",
			"brackets",
			Re("Brackets.Release.*.msi"),
			nil,
		),
	)
	AddRule(
		"clementine-player",
		GitHubReleaseVersionExtractor(
			"clementine-player",
			"Clementine",
			Re("(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"clementine-player",
			"Clementine",
			Re("ClementineSetup-.*.exe"),
			nil,
		),
	)
	AddRule(
		"conemu",
		GitHubReleaseVersionExtractor(
			"Maximus5",
			"ConEmu",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"Maximus5",
			"ConEmu",
			Re("ConEmuSetup.*.exe"),
			nil,
		),
	)
	AddRule(
		"dbeaver",
		GitHubReleaseVersionExtractor(
			"dbeaver",
			"dbeaver",
			Re("(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"dbeaver",
			"dbeaver",
			Re("dbeaver-ce-.+-x86-setup.exe"),
			Re("dbeaver-ce-.+-x86_64-setup.exe"),
		),
	)
	AddRule(
		"eig",
		GitHubReleaseVersionExtractor(
			"EvilInsultGenerator",
			"c-sharp-desktop",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"EvilInsultGenerator",
			"c-sharp-desktop",
			Re("EvilInsultGenerator_Setup.exe"),
			nil,
		),
	)
	AddRule(
		"etcher",
		GitHubReleaseVersionExtractor(
			"resin-io",
			"etcher",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"resin-io",
			"etcher",
			Re("Etcher-Setup-.+-x86.exe"),
			Re("Etcher-Setup-.+-x64.exe"),
		),
	)
	AddRule(
		"git",
		GitHubReleaseVersionExtractor(
			"git-for-windows",
			"git",
			Re("v([0-9.]+)\\.windows.+"),
		),
		GitHubReleaseDownloadExtractor(
			"git-for-windows",
			"git",
			Re("Git-.+-32-bit.exe"),
			Re("Git-.+-64-bit.exe"),
		),
	)
	AddRule(
		"gitextensions",
		GitHubReleaseVersionExtractor(
			"gitextensions",
			"gitextensions",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"gitextensions",
			"gitextensions",
			Re("GitExtensions-.*-Setup.msi"),
			nil,
		),
	)
	AddRule(
		"git-lfs",
		GitHubReleaseVersionExtractor(
			"git-lfs",
			"git-lfs",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"git-lfs",
			"git-lfs",
			Re("git-lfs-windows-.+.exe"),
			nil,
		),
	)
	AddRule(
		"gow",
		GitHubReleaseVersionExtractor(
			"bmatzelle",
			"gow",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"bmatzelle",
			"gow",
			Re("Gow-.+.exe"),
			nil,
		),
	)
	AddRule(
		"greenshot",
		GitHubReleaseVersionExtractor(
			"greenshot",
			"greenshot",
			Re("Greenshot-RELEASE-([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"greenshot",
			"greenshot",
			Re("Greenshot-INSTALLER-.+-RELEASE.exe"),
			nil,
		),
	)
	AddRule(
		"hashcheck",
		GitHubReleaseVersionExtractor(
			"gurnec",
			"HashCheck",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"gurnec",
			"HashCheck",
			Re("HashCheckSetup-.+.exe"),
			nil,
		),
	)
	AddRule(
		"hugo",
		GitHubReleaseVersionExtractor(
			"gohugoio",
			"hugo",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"gohugoio",
			"hugo",
			Re("hugo_.+_Windows-32bit.zip"),
			Re("hugo_.+_Windows-64bit.zip"),
		),
	)
	AddRule(
		"imageglass",
		GitHubReleaseVersionExtractor(
			"d2phap",
			"ImageGlass",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"d2phap",
			"ImageGlass",
			Re("ImageGlass_.+.exe"),
			nil,
		),
	)
	AddRule(
		"keepassxc",
		GitHubReleaseVersionExtractor(
			"keepassxreboot",
			"keepassxc",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"keepassxreboot",
			"keepassxc",
			Re("KeePassXC-.+-Win32.exe"),
			Re("KeePassXC-.+-Win64.exe"),
		),
	)
	AddRule(
		"keeweb",
		GitHubReleaseVersionExtractor(
			"keeweb",
			"keeweb",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"keeweb",
			"keeweb",
			Re("KeeWeb-.+.win.ia32.exe"),
			Re("KeeWeb-.+.win.x64.exe"),
		),
	)
	AddRule(
		"mumble",
		GitHubReleaseVersionExtractor(
			"mumble-voip",
			"mumble",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"mumble-voip",
			"mumble",
			Re("mumble-.+.msi"),
			nil,
		),
	)
	AddRule(
		"naps2",
		GitHubReleaseVersionExtractor(
			"cyanfish",
			"naps2",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"cyanfish",
			"naps2",
			Re("naps2-.+-setup.msi"),
			nil,
		),
	)
	AddRule(
		"notepad2-mod",
		GitHubReleaseVersionExtractor(
			"XhmikosR",
			"notepad2-mod",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"XhmikosR",
			"notepad2-mod",
			Re("Notepad2-mod..+.exe"),
			nil,
		),
	)
	AddRule(
		"npackd",
		GitHubReleaseVersionExtractor(
			"tim-lebedkov",
			"npackd-cpp",
			Re("version_([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"tim-lebedkov",
			"npackd-cpp",
			Re("Npackd32-.+.msi"),
			Re("Npackd64-.+.msi"),
		),
	)
	AddRule(
		"npackdcl",
		GitHubReleaseVersionExtractor(
			"tim-lebedkov",
			"npackd-cpp",
			Re("version_([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"tim-lebedkov",
			"npackd-cpp",
			Re("NpackdCL-.+.msi"),
			nil,
		),
	)
	AddRule(
		"processhacker",
		GitHubReleaseVersionExtractor(
			"processhacker",
			"processhacker",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"processhacker",
			"processhacker",
			Re("processhacker-.+-setup.exe"),
			nil,
		),
	)
	AddRule(
		"python2-win32",
		GitHubReleaseVersionExtractor(
			"mhammond",
			"pywin32",
			Re("b(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"mhammond",
			"pywin32",
			Re("pywin32-.+.win32-py2.7.exe"),
			Re("pywin32-.+.win-amd64-py2.7.exe"),
		),
	)
	AddRule(
		"ruby",
		GitHubReleaseVersionExtractor(
			"oneclick",
			"rubyinstaller2",
			Re("rubyinstaller-([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"oneclick",
			"rubyinstaller2",
			Re("rubyinstaller-[0-9.]+-.+-x86.exe"),
			Re("rubyinstaller-[0-9.]+-.+-x64.exe"),
		),
	)
	AddRule(
		"sharex",
		GitHubReleaseVersionExtractor(
			"ShareX",
			"ShareX",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"ShareX",
			"ShareX",
			Re("ShareX-.+-setup.exe"),
			nil,
		),
	)
	AddRule(
		"simplenote",
		GitHubReleaseVersionExtractor(
			"Automattic",
			"simplenote-electron",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"Automattic",
			"simplenote-electron",
			Re("Simplenote-windows-.+.exe"),
			nil,
		),
	)
	AddRule(
		"sharpkeys",
		GitHubReleaseVersionExtractor(
			"randyrants",
			"sharpkeys",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"randyrants",
			"sharpkeys",
			Re("sharpkeys.+.msi"),
			nil,
		),
	)
	AddRule(
		"syncthing",
		GitHubReleaseVersionExtractor(
			"syncthing",
			"syncthing",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"syncthing",
			"syncthing",
			Re(".*windows-386.*.zip"),
			Re(".*windows-amd64.*.zip"),
		),
	)
	AddRule(
		"tortoisegit",
		RegexpVersionExtractor(
			"https://tortoisegit.org/download/",
			Re("TortoiseGit-([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://tortoisegit.org/download/",
			true,
			"a[href$='32bit.msi']",
			"a[href$='64bit.msi']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"tortoisesvn",
		RegexpVersionExtractor(
			"https://tortoisesvn.net/downloads.html",
			Re("The current version is ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://tortoisesvn.net/downloads.html",
			true,
			"a[href*='win32-svn']",
			"a[href*='x64-svn']",
			"href",
			"href",
			Re("(.+TortoiseSVN-[0-9.]+-win32-svn-[0-9.]+.msi)"),
			Re("(.+TortoiseSVN-[0-9.]+-x64-svn-[0-9.]+.msi)"),
		),
	)
	AddRule(
		"upx",
		GitHubReleaseVersionExtractor(
			"upx",
			"upx",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"upx",
			"upx",
			Re("upx[0-9]+w.zip"),
			nil,
		),
	)
	AddRule(
		"webtorrent",
		GitHubReleaseVersionExtractor(
			"webtorrent",
			"webtorrent-desktop",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"webtorrent",
			"webtorrent-desktop",
			Re("WebTorrentSetup-v[0-9.]+-ia32.exe"),
			Re("WebTorrentSetup-v[0-9.]+.exe"),
		),
	)
	AddRule(
		"wixedit",
		GitHubReleaseVersionExtractor(
			"WixEdit",
			"WixEdit",
			Re("v([0-9]+\\.[0-9]+\\.[0-9]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"WixEdit",
			"WixEdit",
			Re("wixedit-.+.msi"),
			nil,
		),
	)
	AddRule(
		"workflowy",
		GitHubReleaseVersionExtractor(
			"workflowy",
			"desktop",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"workflowy",
			"desktop",
			Re("WorkFlowy.exe"),
			nil,
		),
	)
	AddRule(
		"wox",
		GitHubReleaseVersionExtractor(
			"Wox-launcher",
			"Wox",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"Wox-launcher",
			"Wox",
			Re("Wox-[0-9.]+.exe"),
			nil,
		),
	)
	AddRule(
		"youtube-dl",
		GitHubReleaseVersionExtractor(
			"rg3",
			"youtube-dl",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"rg3",
			"youtube-dl",
			Re("youtube-dl.exe"),
			nil,
		),
	)
}
