package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kong/deck/dump"
	"github.com/kong/deck/state"
	"github.com/kong/deck/utils"
	"github.com/kong/go-kong/kong"
)

func getKongAddress() string {
	if os.Getenv("KONG_ADDRESS") != "" {
		return os.Getenv("KONG_ADDRESS")
	}
	return "http://localhost:8001"
}

func getTestClient() (*kong.Client, error) {
	return utils.GetKongClient(utils.KongClientConfig{
		Address: getKongAddress(),
	})
}

func cleanUpEnv(client *kong.Client) error {
	ctx := context.Background()
	currentState, err := fetchCurrentState(ctx, client, dumpConfig)
	if err != nil {
		return err
	}
	targetState, err := state.NewKongState()
	if err != nil {
		return err
	}
	_, err = performDiff(ctx, currentState, targetState, false, 10, 0, client)
	return err
}

func normalizeOutput(content string) string {
	lines := strings.Split(strings.TrimSuffix(content, "\n"), "\n")
	sort.Strings(lines)
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return strings.Join(lines, "\n")
}

func loadExpectedOutput(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func String(v string) *string { return &v }

func Int(i int) *int { return &i }

func Bool(b bool) *bool { return &b }

func sortServices(x, y *kong.Service) bool {
	xName := *x.Name
	yName := *y.Name
	return xName < yName
}

func Test_Sync(t *testing.T) {
	var (
		subcommand = "sync"
		flag       = "-s"
	)
	tests := []struct {
		name             string
		expectedServices []*kong.Service
		expectedPlugins  []*kong.Plugin
	}{
		{
			name: "create_service",
			expectedServices: []*kong.Service{
				{
					Name:           String("svc1"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
					Tags:           []*string{String("team-svc1")},
				},
				{
					Name:           String("svc2"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
				},
				{
					Name:           String("svc3"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
				},
			},
			expectedPlugins: nil,
		},
		{
			name: "create_service_no_change",
			expectedServices: []*kong.Service{
				{
					Name:           String("svc1"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
					Tags:           []*string{String("team-svc1")},
				},
				{
					Name:           String("svc2"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
				},
				{
					Name:           String("svc3"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
				},
			},
			expectedPlugins: nil,
		},
		{
			name: "create_plugin",
			expectedServices: []*kong.Service{
				{
					Name:           String("svc1"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
					Tags:           []*string{String("team-svc1")},
				},
				{
					Name:           String("svc2"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
				},
				{
					Name:           String("svc3"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
				},
			},
			expectedPlugins: []*kong.Plugin{
				{
					Name: String("rate-limiting"),
					Config: kong.Configuration{
						"day":                 nil,
						"fault_tolerant":      true,
						"header_name":         nil,
						"hide_client_headers": false,
						"hour":                float64(10),
						"limit_by":            "consumer",
						"minute":              nil,
						"month":               nil,
						"path":                nil,
						"policy":              string("cluster"),
						"redis_database":      float64(0),
						"redis_host":          nil,
						"redis_password":      nil,
						"redis_port":          float64(6379),
						"redis_server_name":   nil,
						"redis_ssl":           bool(false),
						"redis_ssl_verify":    bool(false),
						"redis_timeout":       float64(2000),
						"second":              nil,
						"year":                nil,
					},
					Enabled:   Bool(true),
					Protocols: []*string{String("grpc"), String("grpcs"), String("http"), String("https")},
				},
			},
		},
		{
			name: "create_plugin_without_defaults_no_change",
			expectedServices: []*kong.Service{
				{
					Name:           String("svc1"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
					Tags:           []*string{String("team-svc1")},
				},
				{
					Name:           String("svc2"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
				},
				{
					Name:           String("svc3"),
					ConnectTimeout: Int(60000),
					Host:           String("mockbin.org"),
					Port:           Int(80),
					Protocol:       String("http"),
					ReadTimeout:    Int(60000),
					Retries:        Int(5),
					WriteTimeout:   Int(60000),
				},
			},
			expectedPlugins: []*kong.Plugin{
				{
					Name: String("rate-limiting"),
					Config: kong.Configuration{
						"day":                 nil,
						"fault_tolerant":      true,
						"header_name":         nil,
						"hide_client_headers": false,
						"hour":                float64(10),
						"limit_by":            "consumer",
						"minute":              nil,
						"month":               nil,
						"path":                nil,
						"policy":              string("cluster"),
						"redis_database":      float64(0),
						"redis_host":          nil,
						"redis_password":      nil,
						"redis_port":          float64(6379),
						"redis_server_name":   nil,
						"redis_ssl":           bool(false),
						"redis_ssl_verify":    bool(false),
						"redis_timeout":       float64(2000),
						"second":              nil,
						"year":                nil,
					},
					Enabled:   Bool(true),
					Protocols: []*string{String("grpc"), String("grpcs"), String("http"), String("https")},
				},
			},
		},
	}
	client, err := getTestClient()
	if err != nil {
		t.Errorf(err.Error())
	}
	if err := cleanUpEnv(client); err != nil {
		t.Errorf(err.Error())
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			tcPath := fmt.Sprintf("testdata/%s/%s/", subcommand, tc.name)
			kongFile := fmt.Sprintf("%s/kong.yaml", tcPath)
			cmd := exec.Command(
				"deck",
				"--kong-addr",
				getKongAddress(),
				subcommand,
				flag,
				kongFile,
			) // #nosec G204

			var stdout bytes.Buffer
			cmd.Stdout = &stdout
			if err := cmd.Run(); err != nil {
				t.Errorf(err.Error())
			}

			expected, err := loadExpectedOutput(fmt.Sprintf("%s/output", tcPath))
			if err != nil {
				t.Errorf(err.Error())
			}

			expectedNormalized := normalizeOutput(expected)
			obtainedNormalized := normalizeOutput(stdout.String())
			if !reflect.DeepEqual(obtainedNormalized, expectedNormalized) {
				t.Errorf(cmp.Diff(obtainedNormalized, expectedNormalized))
			}

			// Get entities from Kong
			services, err := dump.GetAllServices(ctx, client, []string{})
			if err != nil {
				t.Errorf(err.Error())
			}
			plugins, err := dump.GetAllPlugins(ctx, client, []string{})
			if err != nil {
				t.Errorf(err.Error())
			}

			opt := []cmp.Option{
				cmpopts.IgnoreFields(kong.Service{}, "ID", "CreatedAt", "UpdatedAt"),
				cmpopts.IgnoreFields(kong.Plugin{}, "ID", "CreatedAt"),
				cmpopts.SortSlices(sortServices),
			}
			if diff := cmp.Diff(services, tc.expectedServices, opt...); diff != "" {
				t.Errorf(diff)
			}
			if diff := cmp.Diff(plugins, tc.expectedPlugins, opt...); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
