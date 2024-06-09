package remove

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"nogihen.com/logger"
)

// session_idに紐づくMPDファイルを削除する
func RemoveMpdA(lg logger.Logger, rootdir, tenantId, companyId string) error {

	// MPDファイルのパスを取得
	mpddir := fmt.Sprintf("%s/asset/%s/%s", rootdir, tenantId, companyId)

	// MPDファイルが存在しない場合は処理を終了
	if _, err := os.Stat(mpddir); os.IsNotExist(err) {
		return nil
	}

	// osがwindowsの場合は処理を終了
	if os.Getenv("OS") == "Windows_NT" {
		return nil
	}

	// 排他制御ファイルをチェック
	fi, flerr := os.Stat(mpddir + "/.lock")
	if flerr == nil {
		if fi.ModTime().Add(24 * time.Hour).After(fi.ModTime()) {
			// 排他制御ファイルを削除
			if err := os.Remove(mpddir + "/.lock"); err != nil {
				lg.LW(err, "failed to remove lock file : [%s]", mpddir+"/.lock")
			}
		} else {
			return nil
		}
	} else {
		if !os.IsNotExist(flerr) {
			return lg.LE(flerr, "failed to check lock file : [%s]", mpddir+"/.lock")
		}
	}

	// 排他制御ファイルを作成
	if _, err := os.Create(mpddir + "/.lock"); err != nil {
		return lg.LE(err, "failed to create lock file : [%s]", mpddir+"/.lock")
	}

	// コマンドを実行する
	go func() {

		cmd := exec.Command("find", mpddir, "-name", "sess_*.mpd", "-mtime", "+1", "-exec", "rm", "-f", "{}", "\\;")
		err := cmd.Run()
		if err != nil {
			lg.LW(err, "failed to remove mpd file : [%s]", mpddir)
		}
		lg.LTM("execute command-[%s]", cmd.String())
		// 排他制御ファイルを削除
		if err := os.Remove(mpddir + "/.lock"); err != nil {
			lg.LW(err, "failed to remove lock file : [%s]", mpddir)
		}
		lg.LTM("execute remove lock file : [%s]", mpddir+"/.lock")
	}()

	return nil
}

//#########################################################################################
//#########################################################################################
//#########################################################################################

func RemoveMpdB(lg logger.Logger, rootdir, tenantId, companyId string) error {

	// MPDファイルのパスを取得
	mpddir := fmt.Sprintf("%s/asset/%s/%s", rootdir, tenantId, companyId)

	// ディレクトリが存在しない場合は処理を終了
	if _, err := os.Stat(mpddir); os.IsNotExist(err) {
		lg.LWM("Directory does not exist. mpddir: %s, err: %v", mpddir, err)
		return nil
	}

	// 排他制御ファイルをチェック
	lockFile := mpddir + "/.lock"
	fi, flerr := os.Stat(lockFile)
	if flerr == nil {
		if fi.ModTime().Before(time.Now().Add(-24 * time.Hour)) {
			// 排他制御ファイルが24時間以上前であれば削除
			if err := os.Remove(lockFile); err != nil {
				return lg.LEM("Failed to remove lock file. lockFile: %s, err: %v", lockFile, err)
			}
		} else {
			// ロックファイルが存在し24時間以内であれば処理を終了
			lg.LWM("Lock file remains within the last 24 hours. lockFile: %s", lockFile)
			return nil
		}
	} else if !os.IsNotExist(flerr) {
		// ロックファイルのステータス確認に失敗した場合
		return lg.LEM("Failed to check lock file. lockFile: %s, flerr: %v", lockFile, flerr)
	}

	// 排他制御ファイルを作成
	lockFileHandle, err := os.Create(lockFile)
	if err != nil {
		return lg.LEM("Failed to create lock file. lockFile: %s, err: %v", lockFile, err)
	}
	lockFileHandle.Close()

	// 関数終了時にロックファイルを削除
	defer func() {
		if err := os.Remove(lockFile); err != nil {
			lg.LEM("Failed to remove lock file. lockFile: %s, err: %v", lockFile, err)
		} else {
			lg.LTM("Executed remove lock file. lockFile: %s", lockFile)
		}
	}()

	// ファイル削除を並列処理で実行するためのチャネル
	type fileTask struct {
		path    string
		modTime time.Time
	}
	taskChan := make(chan fileTask, 300)

	// エラーチャネル
	errorChan := make(chan error, 1)

	// ワーカーを起動
	var wg sync.WaitGroup
	const numWorkers = 2
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskChan {
				// MPDファイルが24時間より前のものかどうかをチェック
				if task.modTime.Before(time.Now().Add(-24 * time.Hour)) {
					// ファイル削除実行
					if err := os.Remove(task.path); err != nil {
						lg.LWM("Failed to remove file. path: %s, err: %v", task.path, err)
						errorChan <- err
					} else {
						lg.LBM("Successfully removed file. path: %s", task.path)
					}
				}
			}
		}()
	}

	// ファイルを探索しタスクをチャネルに送る
	go func() {
		defer close(taskChan)
		err := filepath.Walk(mpddir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// ファイル名がパターンに一致するかどうかをチェック
			if matched, _ := filepath.Match("sess_*.mpd", info.Name()); matched {
				taskChan <- fileTask{path: path, modTime: info.ModTime()}
			}
			return nil
		})
		if err != nil {
			errorChan <- lg.LWM("Failed to walk the path. mpddir: %s, err: %v", mpddir, err)
		}
	}()

	// ワーカーが全て終了するのを待つ
	wg.Wait()
	close(errorChan)

	// エラーが発生していた場合最初のエラーを返す
	if err := <-errorChan; err != nil {
		return err
	}

	return nil
}
