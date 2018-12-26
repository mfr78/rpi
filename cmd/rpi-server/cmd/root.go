package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	proto "github.com/gbbirkisson/rpi/proto"
	rpi "github.com/gbbirkisson/rpi/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "rpi-server",
	Short: "Raspberry PI IO server",
	Long:  `A gRPC server that allows you to do IO operations on the Raspberry PI`,
	RunE:  run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().IntP("port", "p", 8000, "server port")
	rootCmd.PersistentFlags().StringP("ip", "i", "0.0.0.0", "server ip")
	rootCmd.PersistentFlags().BoolP("gpio", "g", false, "gpio service enabled")

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("ip", rootCmd.PersistentFlags().Lookup("ip"))
	viper.BindPFlag("gpio", rootCmd.PersistentFlags().Lookup("gpio"))
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

func run(cmd *cobra.Command, args []string) error {
	log.Printf("rpi server started")
	ip := cmd.Flag("ip").Value.String()
	port := cmd.Flag("port").Value.String()
	address := ip + ":" + port

	srv := grpc.NewServer()

	proto.RegisterCommonServer(srv, &rpi.CommonServerImpl{})

	if cmd.Flag("gpio").Value.String() == "true" {
		log.Printf("adding gpio service")
		proto.RegisterGPIOServer(srv, rpi.GetGPIOServer())
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("failed starting server: %v\n", err)
		os.Exit(1)
	}

	log.Printf("listening to %s\n", address)
	log.Fatal(srv.Serve(lis))
	return nil
}
