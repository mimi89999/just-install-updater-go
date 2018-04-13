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
			"a[hRef$='32bit.msi']",
			"a[hRef$='64bit.msi']",
			"href",
			"href",
			nil,
			nil,
		),
	)
}
