// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/kubelet/certificate/bootstrap"
)

var kubeConfig string
var bootstrapKubeconfig string
var certDir string
var nodeName string

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Call LoadClientCert to generate cert and kubeconfig",
	Long: `The kubelet authenticates to the apiserver using a client certificate issued by the apiserver 
	(controller-manager I believe). The kubelet is initially bootstrapped with a key and cert typically 
	via compute instance metadata on public cloud. Access to this metadata will provide you access to the 
	k8s API only to request a certificate. It won't be very useful for much else. 
	This tool uses LoadClientCert function from the kubelet bootstrap.go to generate valid cert and set 
	up a kubeconfig file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")

		bootstrap.LoadClientCert(kubeConfig, bootstrapKubeconfig, certDir, types.NodeName(nodeName))
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)

	bootstrapCmd.Flags().StringVar(&kubeConfig, "kubeconfig", "kubeconfig", "Location to write kubeconfig file")
	bootstrapCmd.Flags().StringVar(&bootstrapKubeconfig, "bootstrap-kubeconfig", "bootstrap-kubeconfig", "Location to read bootstrap-kubeconfig file")
	bootstrapCmd.Flags().StringVar(&certDir, "cert-dir", "pki", "Location to write new PEM to")
	bootstrapCmd.Flags().StringVar(&nodeName, "node-name", "", "Node name to use for CSR")
	bootstrapCmd.MarkFlagRequired("node-name")

}
