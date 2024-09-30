package util

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/mholt/archiver/v4"
)

func Extract(file *os.File, targetPath string) error {
	format, reader, err := archiver.Identify(file.Name(), file)
	if err != nil {
		return err
	}
	err = format.(archiver.Extractor).Extract(context.TODO(), reader, nil, func(ctx context.Context, f archiver.File) error {
		path := filepath.Join(targetPath, f.NameInArchive)
		if err = os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
			return err
		}
		if f.FileInfo.IsDir() {
			return os.MkdirAll(path, f.Mode())
		}
		v, err := f.Open()
		if err != nil {
			return err
		}
		defer func(v io.ReadCloser) {
			_ = v.Close()
		}(v)

		bytes, err := io.ReadAll(v)
		if err != nil {
			return err
		}

		return os.WriteFile(targetPath+f.NameInArchive, bytes, f.Mode())
	})
	return err
}

func ExtractExc(file *os.File, targetPath string, except []string) error {
	format, reader, err := archiver.Identify(file.Name(), file)
	if err != nil {
		return err
	}
	err = format.(archiver.Extractor).Extract(context.TODO(), reader, nil, func(ctx context.Context, f archiver.File) error {
		if slices.Contains(except, f.FileInfo.Name()) {
			return nil
		}
		path := filepath.Join(targetPath, f.NameInArchive)
		if err = os.MkdirAll(filepath.Dir(path), f.Mode()); err != nil {
			return err
		}
		if f.FileInfo.IsDir() {
			return os.MkdirAll(path, f.Mode())
		}
		v, err := f.Open()
		if err != nil {
			return err
		}
		defer func(v io.ReadCloser) {
			_ = v.Close()
		}(v)

		bytes, err := io.ReadAll(v)
		if err != nil {
			return err
		}

		return os.WriteFile(targetPath+f.NameInArchive, bytes, f.Mode())
	})
	return err
}

func CompressAndArchive(pathToArchive string, compressedPackageSuffix string) error {
	format := archiver.CompressedArchive{}
	switch compressedPackageSuffix {
	case ".tar.zst":
		format.Compression = archiver.Zstd{}
		format.Archival = archiver.Tar{}
	case "zip":
		format.Archival = archiver.Zip{}
	default:
		return errors.New("compress format not supported: " + compressedPackageSuffix)
	}
	tarFiles := map[string]string{}
	base := filepath.Base(pathToArchive)
	err := filepath.WalkDir(pathToArchive, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		tarFiles[path] = base + "/"
		return nil
	})
	files, err := archiver.FilesFromDisk(nil, tarFiles)
	if err != nil {
		return err
	}
	out, err := os.Create(filepath.Join(pathToArchive, "../") + "/" + base + compressedPackageSuffix)
	defer out.Close()
	if err != nil {
		return err
	}
	return format.Archive(context.TODO(), out, files)
}
