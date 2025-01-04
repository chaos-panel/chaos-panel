package migration

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/os/gres"
	"github.com/golang-migrate/migrate/v4/source"
)

// StringReadCloser 是一个结构体，实现了 io.ReadCloser 接口
type StringReadCloser struct {
	*strings.Reader
}

// Close 方法，实现 io.ReadCloser 接口的 Close 方法
func (src *StringReadCloser) Close() error {
	// 对于 StringReadCloser，Close 方法不需要做任何操作
	// 因为 strings.Reader 不需要关闭
	return nil
}

// NewStringReadCloser 返回一个新的 StringReadCloser 对象
func NewStringReadCloser(s string) *StringReadCloser {
	return &StringReadCloser{Reader: strings.NewReader(s)}
}

type driver struct {
	PartialDriver
}

// New returns a new Driver from io/fs#FS and a relative path.
func NewEmbededRes(path string) (source.Driver, error) {
	var i driver
	if err := i.Init(path); err != nil {
		return nil, fmt.Errorf("failed to init driver with path %s: %w", path, err)
	}
	return &i, nil
}

// Open is part of source.Driver interface implementation.
// Open cannot be called on the iofs passthrough driver.
func (d *driver) Open(url string) (source.Driver, error) {
	return nil, errors.New("Open() cannot be called on the iofs passthrough driver")
}

// PartialDriver is a helper service for creating new source drivers working with
// io/fs.FS instances. It implements all source.Driver interface methods
// except for Open(). New driver could embed this struct and add missing Open()
// method.
//
// To prepare PartialDriver for use Init() function.
type PartialDriver struct {
	migrations *source.Migrations
	path       string
}

// Init prepares not initialized IoFS instance to read migrations from a
// io/fs#FS instance and a relative path.
func (d *PartialDriver) Init(path string) error {
	entries := gres.ScanDirFile(path, "*.sql", true)
	if len(entries) == 0 {
		return errors.New("not found sql files in gres: " + path)
	}

	ms := source.NewMigrations()
	for _, e := range entries {
		m, err := source.DefaultParse(filepath.Base(e.Name()))
		if err != nil {
			continue
		}
		if !ms.Append(m) {
			return source.ErrDuplicateMigration{
				Migration: *m,
				FileInfo:  e.FileInfo(),
			}
		}
	}

	d.path = path
	d.migrations = ms
	return nil
}

// Close is part of source.Driver interface implementation.
// Closes the file system if possible.
func (d *PartialDriver) Close() error {
	return nil
}

// First is part of source.Driver interface implementation.
func (d *PartialDriver) First() (version uint, err error) {
	if version, ok := d.migrations.First(); ok {
		return version, nil
	}
	return 0, &fs.PathError{
		Op:   "first",
		Path: d.path,
		Err:  fs.ErrNotExist,
	}
}

// Prev is part of source.Driver interface implementation.
func (d *PartialDriver) Prev(version uint) (prevVersion uint, err error) {
	if version, ok := d.migrations.Prev(version); ok {
		return version, nil
	}
	return 0, &fs.PathError{
		Op:   "prev for version " + strconv.FormatUint(uint64(version), 10),
		Path: d.path,
		Err:  fs.ErrNotExist,
	}
}

// Next is part of source.Driver interface implementation.
func (d *PartialDriver) Next(version uint) (nextVersion uint, err error) {
	if version, ok := d.migrations.Next(version); ok {
		return version, nil
	}
	return 0, &fs.PathError{
		Op:   "next for version " + strconv.FormatUint(uint64(version), 10),
		Path: d.path,
		Err:  fs.ErrNotExist,
	}
}

// ReadUp is part of source.Driver interface implementation.
func (d *PartialDriver) ReadUp(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := d.migrations.Up(version); ok {
		content := gres.GetContent(path.Join(d.path, m.Raw))
		body := NewStringReadCloser(string(content))
		if err != nil {
			return nil, "", err
		}
		return body, m.Identifier, nil
	}
	return nil, "", &fs.PathError{
		Op:   "read up for version " + strconv.FormatUint(uint64(version), 10),
		Path: d.path,
		Err:  fs.ErrNotExist,
	}
}

// ReadDown is part of source.Driver interface implementation.
func (d *PartialDriver) ReadDown(version uint) (r io.ReadCloser, identifier string, err error) {
	if m, ok := d.migrations.Down(version); ok {
		content := gres.GetContent(path.Join(d.path, m.Raw))
		body := NewStringReadCloser(string(content))
		if err != nil {
			return nil, "", err
		}
		return body, m.Identifier, nil
	}
	return nil, "", &fs.PathError{
		Op:   "read down for version " + strconv.FormatUint(uint64(version), 10),
		Path: d.path,
		Err:  fs.ErrNotExist,
	}
}
