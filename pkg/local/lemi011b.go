package local

import (
    "fmt"
    "os"
    "time"

    "github.com/fanie42/sansa/pkg/lemi011b"
)

// Save TODO -- Repetitive code can just be moved to a
// utility file.
func (repo *lemi011bRepo) Save(
    d *lemi011b.Data,
) error {
    timestamp := d.Timestamp

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

    s := fmt.Sprintf(
        "%s, %d, %d, %d, %d\n",
        timestamp.Format("2006-01-02 15:04:05.000000"),
        d.X,
        d.Y,
        d.Z,
        d.Temperature,
    )

    _, err := repo.file.WriteString(s)

    return err
}

func (repo *lemi011bRepo) getBase(t time.Time) string {
    pre := repo.config.Base.Prefix
    pattern := t.Format(repo.config.Base.Format)
    post := repo.config.Base.Postfix
    ext := repo.config.Base.Extension

    return pre + pattern + post + "." + ext
}

func (repo *lemi011bRepo) getPath(t time.Time) string {
    loc := repo.config.Path.Location
    pattern := t.Format(repo.config.Path.Format)

    // Add in some path logic here for slashes and such
    return loc + pattern
}
