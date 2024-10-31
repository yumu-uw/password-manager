package ymuwutil

import (
	"os"
	"os/user"
	"path"
)

var configDir string

func SetupConfigDir() error {
	user, err := user.Current()
	if err != nil {
		return err
	}
	home_dir := user.HomeDir
	configDir = path.Join(home_dir, "Library/Application Support/YmuwApps/password-manager")

	// 設定ファイル保存ディレクトリの作成
	os.MkdirAll(configDir, os.ModePerm)

	// if err = createSubnetListFile(); err != nil {
	// 	return err
	// }

	return nil
}

func GetConfigDir() string {
	return configDir
}

// func createSubnetListFile() error {
// 	p := path.Join(configDir, "subnet_list.json")
// 	//  ファイルがあればリターン
// 	if _, err := os.Stat(p); err == nil {
// 		return nil
// 	}

// 	// ファイル作成
// 	file, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// 初期値書き込み
// 	encoder := json.NewEncoder(file)
// 	var nd model.SubnetListData
// 	nd.SubnetList = []model.Subnet{}
// 	if err := encoder.Encode(nd); err != nil {
// 		return err
// 	}
// 	return nil
// }
