package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
)

func main() {
	paths := []string{"config"}
	conf, err := ReadConfigPaths(paths)
	if err != nil {
		log.Fatalf("Error reading '%s': %s", paths, err)
	}

	conf = MergeConfig(DefaultConfig(), conf)

	fmt.Println(conf.Node)
	fmt.Println(conf.NodeGroup)
	fmt.Println(conf.Connection)
}

type Config struct {
	Node       NodeConfig       `mapstructure:"node"`
	NodeGroup  NodeGroupConfig  `mapstructure:"node-group"`
	Connection ConnectionConfig `mapstructure:"connection"`
}

type NodeConfig struct {
	SshName     string             `mapstructure:"ssh-name"`
	SshPassword string             `mapstructure:"ssh-password"`
	StaticNodes []StaticNodeConfig `mapstructure:"static-nodes"`
}

type StaticNodeConfig struct {
	Id       string `mapstructure:"id"`
	Hostname string `mapstructure:"hostname"`
	Port     int    `mapstructure:"port"`
}

type NodeGroupConfig struct {
	StaticNodeGroups []StaticNodeGroupConfig `mapstructure:"static-node-groups"`
}

type StaticNodeGroupConfig struct {
	Id      string   `mapstructure:"id"`
	NodeIds []string `mapstructure:"node-ids"`
}

type ConnectionConfig struct {
	ScriptConnections []ScriptConnectionConfig `mapstructure:"script-connections"`
}

type ScriptConnectionConfig struct {
	Id             string   `mapstructure:"id"`
	TargetGroupIds []string `mapstructure:"target-group-ids"`
	ScriptPaths    []string `mapstructure:"script-paths"`
}

type dirEnts []os.FileInfo

func DefaultConfig() *Config {
	return &Config{
		Node: NodeConfig{
			SshName:     "",
			SshPassword: "",
			StaticNodes: nil,
		},
		NodeGroup: NodeGroupConfig{
			StaticNodeGroups: nil,
		},
		Connection: ConnectionConfig{
			ScriptConnections: nil,
		},
	}
}

func DecodeConfig(r io.Reader) (*Config, error) {
	var raw interface{}
	var result Config
	dec := json.NewDecoder(r)
	if err := dec.Decode(&raw); err != nil {
		return nil, err
	}

	// Decode
	var md mapstructure.Metadata
	msdec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &result,
	})

	if err != nil {
		return nil, err
	}

	if err := msdec.Decode(raw); err != nil {
		return nil, err
	}

	return &result, nil
}

func MergeConfig(a, b *Config) *Config {
	var result Config = *a

	if b.Node.SshName != "" {
		result.Node.SshName = b.Node.SshName
	}
	if b.Node.SshPassword != "" {
		result.Node.SshPassword = b.Node.SshPassword
	}
	if b.Node.StaticNodes != nil {
		result.Node.StaticNodes = b.Node.StaticNodes
	}
	if b.NodeGroup.StaticNodeGroups != nil {
		result.NodeGroup.StaticNodeGroups = b.NodeGroup.StaticNodeGroups
	}
	if b.Connection.ScriptConnections != nil {
		result.Connection.ScriptConnections = b.Connection.ScriptConnections
	}

	return &result
}

func ReadConfigPaths(paths []string) (*Config, error) {
	result := new(Config)
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		fi, err := f.Stat()
		if err != nil {
			f.Close()
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		if !fi.IsDir() {
			config, err := DecodeConfig(f)
			f.Close()

			if err != nil {
				return nil, fmt.Errorf("Error decoding '%s': %s", path, err)
			}

			result = MergeConfig(result, config)
			continue
		}

		contents, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return nil, fmt.Errorf("Error reading '%s': %s", path, err)
		}

		// Sort the contents, ensures lexical order
		sort.Sort(dirEnts(contents))

		for _, fi := range contents {
			// Don't recursively read contents
			if fi.IsDir() {
				continue
			}

			// If it isn't a JSON file, ignore it
			if !strings.HasSuffix(fi.Name(), ".json") {
				continue
			}

			subpath := filepath.Join(path, fi.Name())
			f, err := os.Open(subpath)
			if err != nil {
				return nil, fmt.Errorf("Error reading '%s': %s", subpath, err)
			}

			config, err := DecodeConfig(f)
			f.Close()

			if err != nil {
				return nil, fmt.Errorf("Error decoding '%s': %s", subpath, err)
			}

			result = MergeConfig(result, config)
		}
	}

	return result, nil
}

// Implement the sort interface for dirEnts
func (d dirEnts) Len() int {
	return len(d)
}

func (d dirEnts) Less(i, j int) bool {
	return d[i].Name() < d[j].Name()
}

func (d dirEnts) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
