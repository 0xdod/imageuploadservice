/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cli

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/0xdod/imageuploadservice/internal/grpc"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload an image to cloud storage",
	Long: `Upload an image image-uploader-cli

USAGE:
   image-uploader-cli upload ./image.png ./image-2.png
.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		locations := make(map[string]string)

		for _, arg := range args {
			loc, err := doUpload(arg)

			if err != nil {
				fmt.Fprintf(os.Stdout, "cannot upload %q: %v", arg, err)
			} else {
				locations[arg] = loc
			}
		}

		if len(locations) > 0 {
			fmt.Fprintln(os.Stdout, "\nUploaded images: ")
			fmt.Fprintln(os.Stdout)
			for k, v := range locations {
				fmt.Fprintf(os.Stdout, "%s: %s\n", k, v)
			}
		}

		return nil
	},
}

func doUpload(arg string) (string, error) {
	file, err := os.Open(arg)

	if err != nil {
		return "", fmt.Errorf("cannot open file: %v", err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return "", fmt.Errorf("cannot read file: %v", err)
	}

	c := grpc.NewClient()

	if err = c.DialServer(net.JoinHostPort("127.0.0.1", "50051")); err != nil {
		return "", fmt.Errorf("cannot connect to server: %v", err)
	}

	defer c.Close()

	fmt.Fprintf(os.Stdout, "Uploading image %q in progress...\n", file.Name())

	location, err := c.UploadImage(context.Background(), file.Name(), data)

	if err != nil {
		return "", fmt.Errorf("cannot upload image: %v", err)
	}

	fmt.Fprintf(os.Stdout, "Uploaded image %q to: %q\n", file.Name(), location)

	return location, nil
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
