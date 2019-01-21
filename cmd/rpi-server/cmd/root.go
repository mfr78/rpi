package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	gpio "github.com/gbbirkisson/rpi/pkg/gpio"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	servers "github.com/gbbirkisson/rpi/pkg/servers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var cfgFile string

func startGpio(srv *grpc.Server) func() {
	ctx := context.Background()
	log.Printf("adding gpio service")
	g := gpio.Gpio{}
	err := g.Open(ctx)
	helper.ExitOnError("unable to open gpio pins", err)
	proto.RegisterGpioServiceServer(srv, &gpio.GpioServer{})
	return func() {
		g.Close(ctx)
	}
}

func startCommon(ctx context.Context, srv *grpc.Server) func() error {
	common := &rpi.Common{}

	cam_modprobe := viper.GetString("picam_modprobe")

	if cam_modprobe != "" {
		err := common.Modprobe(ctx, cam_modprobe)
		helper.ExitOnError(fmt.Sprintf("unable to modprobe %s", cam_modprobe), err)
	}

	proto.RegisterCommonServiceServer(srv, &servers.CommonServer{Common: common})

	return func() error { return nil }
}

func startCamera(ctx context.Context, srv *grpc.Server) func() error {
	log.Printf("adding picam service\n")

	piCam := &rpi.PiCam{
		Width:    viper.GetInt32("picam_width"),
		Height:   viper.GetInt32("picam_height"),
		Rotation: viper.GetInt32("picam_rotation"),
	}

	log.Printf("picam parameters: %+v\n", piCam)

	err := piCam.Open(ctx)
	helper.ExitOnError("unable to create camera", err)

	proto.RegisterPiCamServiceServer(srv, &servers.PiCamServer{Camera: piCam})
	return func() error {
		return piCam.Close(ctx)
	}
}

type NgrokLogger struct{}

func (nl *NgrokLogger) Write(p []byte) (n int, err error) {
	log.Printf("ngrok: %s", p)
	return len(p), nil
}

func startNgrok() {
	log.Printf("starting ngrok\n")

	cmd := exec.Command(
		"ngrok",
		"tcp",
		viper.GetString("port"),
		"--authtoken",
		viper.GetString("ngrok_token"),
		"--log=stdout",
		"--region",
		viper.GetString("ngrok_region"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	helper.ExitOnError("unable to start ngrok", err)
}

var rootCmd = &cobra.Command{
	Use:   "rpi-server",
	Short: "Raspberry PI IO server",
	Long:  `A gRPC server that allows you to do IO operations on the Raspberry PI`,
	RunE: func(_ *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		log.Printf("starting rpi server")
		srv, lis, err := servers.GrpcServerInsecure(viper.GetString("host"), viper.GetString("port"))
		if err != nil {
			helper.ExitOnError("unable to create server", err)
		}

		startCommon(ctx, srv)

		if viper.GetBool("gpio") {
			close := startGpio(srv)
			defer close()
		}

		if viper.GetBool("picam") {
			close := startCamera(ctx, srv)
			defer close()
		}

		if viper.GetBool("ngrok") {
			startNgrok()
		}

		log.Fatal(srv.Serve(lis))
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.SetEnvPrefix("rpi")

	rootCmd.PersistentFlags().String("host", "0.0.0.0", "server ip")
	rootCmd.PersistentFlags().Int("port", 8000, "server port")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

	rootCmd.PersistentFlags().BoolP("gpio", "g", false, "gpio service enabled")
	viper.BindPFlag("gpio", rootCmd.PersistentFlags().Lookup("gpio"))

	rootCmd.PersistentFlags().BoolP("picam", "c", false, "picam service enabled")
	rootCmd.PersistentFlags().String("picam_modprobe", "", "modprobe on start")
	rootCmd.PersistentFlags().Int32("picam_width", 648, "Width of the image from pi camera")
	rootCmd.PersistentFlags().Int32("picam_height", 486, "Height of the image from pi camera")
	rootCmd.PersistentFlags().Int32("picam_rotation", 0, "Rotation of pi camera image")

	viper.BindPFlag("picam", rootCmd.PersistentFlags().Lookup("picam"))
	viper.BindPFlag("picam_modprobe", rootCmd.PersistentFlags().Lookup("picam_modprobe"))
	viper.BindPFlag("picam_width", rootCmd.PersistentFlags().Lookup("picam_width"))
	viper.BindPFlag("picam_height", rootCmd.PersistentFlags().Lookup("picam_height"))
	viper.BindPFlag("picam_rotation", rootCmd.PersistentFlags().Lookup("picam_rotation"))

	rootCmd.PersistentFlags().Bool("ngrok", false, "Start a ngrok tunnel")
	rootCmd.PersistentFlags().String("ngrok_token", "", "Ngrok authentication token")
	rootCmd.PersistentFlags().String("ngrok_region", "eu", "Ngrok region")

	viper.BindPFlag("ngrok", rootCmd.PersistentFlags().Lookup("ngrok"))
	viper.BindPFlag("ngrok_token", rootCmd.PersistentFlags().Lookup("ngrok_token"))
	viper.BindPFlag("ngrok_region", rootCmd.PersistentFlags().Lookup("ngrok_region"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	configPath := "/etc/rpi-server"
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		file := filepath.Join(configPath, "config.yaml")

		fmt.Fprintf(os.Stderr, "Config file not found, creating it: %s\n", file)

		err = os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed creating configuration folder: %v\n", err)
			return
		}

		err = viper.WriteConfigAs(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed creating configuration file: %v\n", err)
			return
		}
	}
}
