package repofiles

import (
	"context"

	"code.gitea.io/gitea/models/db"
)

type FileMeta struct {
	ID     int64  `xorm:"pk autoincr"`
	Sha    string `xorm:"varchar(260) UNIQUE NOT NULL"`
	Sha256 string `xorm:"varchar(260) NOT NULL"`
}

func init() {
	db.RegisterModel(new(FileMeta))
}
func IsFileMetaExist(sha string) (bool, error) {
	if len(sha) == 0 {
		return false, nil
	}
	return db.GetEngine(db.DefaultContext).
		Where("sha = ?", sha).
		Get(&FileMeta{Sha: sha})

}

func UpdateFileMeta(Sha string, Sha256 string) error {
	isFileExist, err := IsFileMetaExist(Sha)
	if err != nil {
		return err
	}
	if !isFileExist {
		return CreateFileMeta(&FileMeta{
			Sha:    Sha,
			Sha256: Sha256,
		})
	}
	_, err = db.GetEngine(db.DefaultContext).Exec("UPDATE `file_meta` SET sha256 = ? WHERE sha = ?", Sha256, Sha)
	return err
}

func CreateFileMeta(f *FileMeta) (err error) {
	ctx, committer, err := db.TxContext(db.DefaultContext)
	if err != nil {
		return err
	}
	defer committer.Close()

	isExist, err := IsFileMetaExist(f.Sha)
	if err != nil {
		return err
	} else if isExist {
		return ErrFileMetaAlreadyExist{f.Sha}
	}
	if err = db.Insert(ctx, f); err != nil {
		return err
	}
	return committer.Commit()
}

func DeleteFileMeta(ctx context.Context, sha string) (int64, error) {
	affected, err := db.GetEngine(ctx).Delete(&FileMeta{Sha: sha})

	if err != nil {
		return -1, err
	}
	return affected, nil
}

func GetFileMeta(sha string) (*FileMeta, error) {
	f := &FileMeta{}
	isExist, err := db.GetEngine(db.DefaultContext).Where("sha = ?", sha).Get(f)
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, ErrFileMetaNotExist{sha}
	}
	return f, nil
}
