// Code generated by goctl. DO NOT EDIT.
package types

type CreateDirectoryReq struct {
	ParentId int64  `json:"parentId"`
	Name     string `json:"name"`
}

type CreateDirectoryResp struct {
}

type DeleteDirectoryReq struct {
	DirId int64 `json:"dirId"`
	Force bool  `json:"force"`
}

type DeleteDirectoryResp struct {
}

type DeleteFileReq struct {
	FileId int64 `json:"fileId"`
}

type DeleteFileResp struct {
}

type Directory struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
}

type File struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Ext    string `json:"ext"`
	Hash   string `json:"hash"`
	Status string `json:"status"`
}

type MoveDirectoryReq struct {
	DirId    int64 `json:"dirId"`
	ParentId int64 `json:"parentId"`
}

type MoveDirectoryResp struct {
}

type MoveFileReq struct {
	FileId int64 `json:"fileId"`
	DirId  int64 `json:"dirId"`
}

type MoveFileResp struct {
}

type RenameDirectoryReq struct {
	DirId  int64  `json:"dirId"`
	Rename string `json:"rename"`
}

type RenameDirectoryResp struct {
}

type RenameFileReq struct {
	FileId int64  `json:"fileId"`
	Rename string `json:"rename"`
}

type RenameFileResp struct {
}

type UserDirectoryListReq struct {
	ParentId int64 `json:"parentId"`
}

type UserDirectoryListResp struct {
	FileList      []File      `json:"fileList"`
	DirectoryList []Directory `json:"directoryList"`
}

type UserFileDetailReq struct {
	FileId int64 `json:"fileId"`
}

type UserFileDetailResp struct {
	File File `json:"file"`
}
