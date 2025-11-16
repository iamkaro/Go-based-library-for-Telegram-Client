/*!
 * I am Karo  😊👍
 *
 * Contact me:
 *     https://www.karo.link/
 *     https://github.com/iamkaro
 *     https://www.linkedin.com/in/iamkaro
 *
 * Go-based library for developing Telegram client applications
 * https://github.com/iamkaro/Go-based-library-for-Telegram-Client.git
 * Copyright © 2020 developed.
 */

package telegram

func (client *Client) newFiles() {
	client.Files = &files{
		client: client,
	}
}

type (
	files struct {
		client *Client
	}
)

func (it *files) GetFile(fileId int32) *File {
	var fileData = object{}
	if it.client.Load(&fileData, Object{"@type": "getFile", "file_id": fileId}) {
		return getFile(fileData)
	}
	return nil
}

func (it *files) GetFileMimeType(filePath string) string {
	var mimeType = object{}
	if it.client.Load(&mimeType, Object{"@type": "getFileMimeType", "file_name": filePath}) {
		return mimeType.string("text")
	}
	return ""
}

/*----------------------------------------/         items         /-----------*/
type (
	File struct {
		Id           int32
		Size         int32
		ExpectedSize int32
		Local        *localFile
		Remote       *remoteFile
	}
	localFile struct {
		Path                   string
		CanBeDownloaded        bool
		CanBeDeleted           bool
		IsDownloadingActive    bool
		IsDownloadingCompleted bool
		DownloadOffset         int32
		DownloadedPrefixSize   int32
		DownloadedSize         int32
	}
	remoteFile struct {
		Id                   string
		UniqueId             string
		IsUploadingActive    bool
		IsUploadingCompleted bool
		UploadedSize         int32
	}
)

func getLocalFile(value object) *localFile {
	if CheckType(value, "localFile") {
		return &localFile{
			Path:                   value.string("path"),
			CanBeDownloaded:        value.bool("can_be_downloaded"),
			CanBeDeleted:           value.bool("can_be_deleted"),
			IsDownloadingActive:    value.bool("is_downloading_active"),
			IsDownloadingCompleted: value.bool("is_downloading_completed"),
			DownloadOffset:         value.int32("download_offset"),
			DownloadedPrefixSize:   value.int32("downloaded_prefix_size"),
			DownloadedSize:         value.int32("downloaded_size"),
		}
	}
	return nil
}

func getRemoteFile(value object) *remoteFile {
	if CheckType(value, "remoteFile") {
		return &remoteFile{
			Id:                   value.string("id"),
			UniqueId:             value.string("unique_id"),
			IsUploadingActive:    value.bool("is_uploading_active"),
			IsUploadingCompleted: value.bool("is_uploading_completed"),
			UploadedSize:         value.int32("uploaded_size"),
		}
	}
	return nil
}

func getFile(value object) *File {
	if CheckType(value, "file") {
		return &File{
			Id:           value.int32("id"),
			Size:         value.int32("size"),
			ExpectedSize: value.int32("expected_size"),
			Local:        getLocalFile(value.Object("local")),
			Remote:       getRemoteFile(value.Object("remote")),
		}
	}
	return nil
}
