package dirs

import (
	"os"
	"path/filepath"

	"github.com/amyadzuki/amygolib/onfail"
)

type Dirs struct {
	vendor, application string

	sExecutableDirectory string
	sSystemCache         string
	sSystemConfig        string
	sSystemData          string
	sUserCache           string
	sUserConfig          string
	sUserData            string
	sUserDesktop         string
	sUserDocuments       string
	sUserDownloads       string
	sUserHome            string
	sUserPictures        string
	sUserScreenshots     string
}

func New(vendor, application string, onFail ...onfail.OnFail) *Dirs {
	return new(Dirs).Init(vendor, application, onFail...)
}

func (d *Dirs) Application() string {
	return d.application
}

func (d *Dirs) ExecutableDirectory() string {
	return sExecutableDirectory
}

func (d *Dirs) Init(vendor, application string, onFail ...onfail.OnFail) *Dirs {
	d.vendor, d.application = vendor, application
	exedir, err := os.Executable()
	if err == nil {
		exedir, err = filepath.EvalSymlinks(exedir)
	}
	if err == nil {
		d.sExecutableDirectory = exedir
	} else {
		onfail.Fail(err, nil, onfail.Panic, onFail...)
	}
	initDirs(d, vendor, application)
	return d
}

func (d *Dirs) SystemCache() string {
	return sSystemCache
}

func (d *Dirs) SystemConfig() string {
	return sSystemConfig
}

func (d *Dirs) SystemData() string {
	return sSystemData
}

func (d *Dirs) UserCache() string {
	return sUserCache
}

func (d *Dirs) UserConfig() string {
	return sUserConfig
}

func (d *Dirs) UserData() string {
	return sUserData
}

func (d *Dirs) UserDesktop() string {
	return sUserDesktop
}

func (d *Dirs) UserDocuments() string {
	return sUserDocuments
}

func (d *Dirs) UserDownloads() string {
	return sUserDownloads
}

func (d *Dirs) UserHome() string {
	return sUserHome
}

func (d *Dirs) UserPictures() string {
	return sUserPictures
}

func (d *Dirs) UserScreenshots() string {
	return sUserScreenshots
}

func (d *Dirs) Vendor() string {
	return d.vendor
}

var initDirs = func(d *Dirs, vendor, application string) {
	return
}
