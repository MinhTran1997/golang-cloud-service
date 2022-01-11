package drop_box

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

type DropboxService struct {
	Token	string
	Client	files.Client
}

type bodyRequestDropbox struct {
	Path		string			`json:"path"`
	Settings	settingsStruct	`json:"settings"`
}

type settingsStruct struct{
	Audience			string	`json:"audience"`
	Access				string	`json:"access"`
	RequestedVisibility	string	`json:"requested_visibility"`
	AllowDownload		bool	`json:"allow_download"`
}

func NewDropboxService(token string) (*DropboxService, error) {
	config := dropbox.Config{
		Token: token,
	}
	client := files.New(config)

	return &DropboxService{Token: token, Client: client}, nil
}

func (d DropboxService) Upload(ctx context.Context, directory string, filename string, data []byte, contentType string) (string, error) {
	file := bytes.NewReader(data)

	// create new client to access drop_box cloud with token generated in drop_box console
	client := d.Client
	if client == nil {
		config := dropbox.Config{
			Token: d.Token,
		}
		client = files.New(config)
	}

	// create new upload info
	filepath := fmt.Sprintf("/%s/%s", directory, filename)
	arg := files.NewCommitInfo(filepath)

	//upload file
	_, err2 := client.Upload(arg, file)
	if err2 != nil {
		panic(err2)
	}

	msg := fmt.Sprintf("uploaded file '%s' to dropbox successfully!!!", filename)
	return msg, err2
}

func (d DropboxService)  Delete(ctx context.Context, directory string, fileName string) (bool, error) {
	client := d.Client
	if client == nil {
		config := dropbox.Config{
			Token: d.Token,
		}
		client = files.New(config)
	}

	filepath := fmt.Sprintf("/%s/%s", directory, fileName)
	arg := files.NewDeleteArg(filepath)
	_, err := client.DeleteV2(arg)
	if err != nil {
		return false, err
	}

	return true, nil
}