package main

import (
	"archive/tar"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/Benyam-S/go-tg-bot/log"
)

func main() {

	maxSize := 2097152                // 2MB
	interval := int64(time.Hour) * 24 // 1 day

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	logs := &log.LogContainer{
		ServerLogFile:  filepath.Join(wd, `../logs/server.log`),
		BotLogFile:     filepath.Join(wd, `../logs/bot.log`),
		ErrorLogFile:   filepath.Join(wd, `../logs/error.log`),
		ArchiveLogFile: filepath.Join(wd, `../logs/archive.log`),
	}
	logger := log.NewLogger(logs)

	// Checking the validity of the given log file
	logFiles := []string{logs.ServerLogFile, logs.BotLogFile, logs.ErrorLogFile}

	Archive(logger, logFiles, int64(maxSize), interval, filepath.Join(wd, `../archives`))
}

// Archive is a function or a process that archives log files when they reaches a certain size
func Archive(logger log.ILogger, logFiles []string, maxSize, interval int64, archiveLocation string) {

	for {

		time.Sleep(time.Duration(interval))

		for _, logFile := range logFiles {

			file, err := os.OpenFile(logFile, os.O_RDWR, 0644)
			if err != nil {
				logger.LogFileArchive(err.Error())
				continue
			}

			status, err := file.Stat()
			if err != nil {
				logger.LogFileArchive(err.Error())
				file.Close()
				continue
			}

			if maxSize <= status.Size() {

				// Logging entry point
				logger.LogFileArchive(fmt.Sprintf("--------------- Start archiving file %s", logFile))

				current_time := time.Now()
				timeStamp := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
					current_time.Year(), current_time.Month(), current_time.Day(),
					current_time.Hour(), current_time.Minute(), current_time.Second())

				archivedFileName := fmt.Sprintf("%s_%s.%s", filepath.Base(file.Name()), timeStamp, "tar")
				archivedFilePath := filepath.Join(archiveLocation, archivedFileName)
				archivedFile, err := os.Create(archivedFilePath)
				if err != nil {
					logger.LogFileArchive(err.Error())
					file.Close()
					continue
				}

				archiver := tar.NewWriter(archivedFile)
				hdr := &tar.Header{
					Name: filepath.Base(file.Name()),
					Mode: 0600,
					Size: status.Size(),
				}

				if err := archiver.WriteHeader(hdr); err != nil {
					logger.LogFileArchive(err.Error())
					archivedFile.Close()
					file.Close()
					continue
				}

				output, err := ioutil.ReadAll(file)
				if err != nil {
					logger.LogFileArchive(err.Error())
					archivedFile.Close()
					file.Close()
					continue
				}

				if _, err := archiver.Write(output); err != nil {
					logger.LogFileArchive(err.Error())
					archivedFile.Close()
					file.Close()
					continue
				}

				if err := archiver.Close(); err != nil {
					logger.LogFileArchive(err.Error())
					archivedFile.Close()
					file.Close()
					continue
				}

				err = archivedFile.Close()
				if err != nil {
					logger.LogFileArchive(err.Error())
					file.Close()
					continue
				}

				// cleaning the file
				err = file.Truncate(0)
				if err != nil {
					logger.LogFileArchive(err.Error())
					file.Close()
					continue
				}

				// Logging finishing point
				logger.LogFileArchive(fmt.Sprintf("--------------- Finished archiving file %s", logFile))
			}

			err = file.Close()
			if err != nil {
				logger.LogFileArchive(err.Error())
				continue
			}

		}
	}
}
