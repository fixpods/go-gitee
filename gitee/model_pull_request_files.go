/*
 * 码云 Open API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.3.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gitee

type PRFilePath struct {
	Diff       string `json:"diff,omitempty"`
	Newpath    string `json:"new_path,omitempty"`
	Oldpath    string `json:"old_path,omitempty"`
	Amode      string `json:"a_mode,omitempty"`
	Bmode      string `json:"b_mode,omitempty"`
	Newfile    bool   `json:"new_file,omitempty"`
	Renamefile bool   `json:"renamed_file,omitempty"`
	Deletefile bool   `json:"deleted_file,omitempty"`
	Toolarge   bool   `json:"too_large,omitempty"`
}

// Pull Request Commit文件列表。最多显示300条diff
type PullRequestFiles struct {
	Sha       string     `json:"sha,omitempty"`
	Filename  string     `json:"filename,omitempty"`
	Status    string     `json:"status,omitempty"`
	Additions string     `json:"additions,omitempty"`
	Deletions string     `json:"deletions,omitempty"`
	BlobUrl   string     `json:"blob_url,omitempty"`
	RawUrl    string     `json:"raw_url,omitempty"`
	Patch *PullRequestFilePath `json:"patch,omitempty"`
}
