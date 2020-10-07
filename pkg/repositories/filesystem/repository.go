package filesystem

import (
    "os"
    "time"

    "github.com/fanie42/sansa/pkg/domains"
    "github.com/fanie42/sansa/pkg/repositories"
)

// implements Repository
type repository struct {
    config Config
    file   *os.File
}

// NewRepository TODO
func NewRepository(
    config Config,
) repositories.Repository {
    return &repository{
        config: config,
    }
}

func (repo *repository) getBase(t time.Time) string {
    pre := repo.config.Base.Prefix
    pattern := t.Format(repo.config.Base.Pattern)
    post := repo.config.Base.Postfix
    ext := repo.config.Base.Extension

    return pre + pattern + post + "." + ext
}

func (repo *repository) getPath(t time.Time) string {
    loc := repo.config.Path.Location
    pattern := t.Format(repo.config.Path.Pattern)

    // Add in some path logic here for slashes and such
    return loc + pattern
}

// Save TODO
func (repo *repository) Save(item domains.Model) error {
    timestamp := item.Timestamp()

    base := repo.getBase(timestamp)
    path := repo.getPath(timestamp)
    abs := path + "/" + base

    // info, err := repo.file.Stat()

    if repo.file != nil {
        // There is a file open - check if it's the correct file
        if repo.file.Name() != abs {
            // We don't have the correct file, open the right one.
            if _, err := os.Stat(abs); os.IsNotExist(err) {
                // File does not exist, create dir, file and open it
                err := os.MkdirAll(path, 0775)
                if err != nil {
                    return err
                }
                // -> LOG THAT WE CREATED A NEW FILE
            }
            err := repo.file.Close()
            if err != nil {
                return err
            }
            f, err := os.OpenFile(
                abs,
                os.O_WRONLY|os.O_APPEND|os.O_CREATE,
                0664,
            )
            if err != nil {
                return err
            }
            repo.file = f
        }
    } else {
        // We don't have a file - create/open the right one
        if _, err := os.Stat(abs); os.IsNotExist(err) {
            // File does not exist, create dir, file and open it
            err := os.MkdirAll(path, 0775)
            if err != nil {
                return err
            }
            // -> LOG THAT WE CREATED A NEW FILE
        }
        f, err := os.OpenFile(
            abs,
            os.O_WRONLY|os.O_APPEND|os.O_CREATE,
            0664,
        )
        if err != nil {
            return err
        }
        repo.file = f
    }

    _, err := repo.file.WriteString(item.String() + "\n")

    return err
}

// GetByID TODO - we need to implement this still - read from file
func (repo *repository) GetByID(id string) (domains.Model, error) {
    return nil, nil
}
