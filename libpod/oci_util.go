package libpod

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/containers/common/libnetwork/types"
	"github.com/hanks177/podman/v4/libpod/define"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Timeout before declaring that runtime has failed to kill a given
// container
const killContainerTimeout = 5 * time.Second

// ociError is used to parse the OCI runtime JSON log.  It is not part of the
// OCI runtime specifications, it follows what runc does
type ociError struct {
	Level string `json:"level,omitempty"`
	Time  string `json:"time,omitempty"`
	Msg   string `json:"msg,omitempty"`
}

// Create systemd unit name for cgroup scopes
func createUnitName(prefix string, name string) string {
	return fmt.Sprintf("%s-%s.scope", prefix, name)
}

// Bind ports to keep them closed on the host
func bindPorts(ports []types.PortMapping) ([]*os.File, error) {
	var files []*os.File
	sctpWarning := true
	for _, port := range ports {
		isV6 := net.ParseIP(port.HostIP).To4() == nil
		if port.HostIP == "" {
			isV6 = false
		}
		protocols := strings.Split(port.Protocol, ",")
		for _, protocol := range protocols {
			for i := uint16(0); i < port.Range; i++ {
				f, err := bindPort(protocol, port.HostIP, port.HostPort+i, isV6, &sctpWarning)
				if err != nil {
					return files, err
				}
				if f != nil {
					files = append(files, f)
				}
			}
		}
	}
	return files, nil
}

func bindPort(protocol, hostIP string, port uint16, isV6 bool, sctpWarning *bool) (*os.File, error) {
	var file *os.File
	switch protocol {
	case "udp":
		var (
			addr *net.UDPAddr
			err  error
		)
		if isV6 {
			addr, err = net.ResolveUDPAddr("udp6", fmt.Sprintf("[%s]:%d", hostIP, port))
		} else {
			addr, err = net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", hostIP, port))
		}
		if err != nil {
			return nil, errors.Wrapf(err, "cannot resolve the UDP address")
		}

		proto := "udp4"
		if isV6 {
			proto = "udp6"
		}
		server, err := net.ListenUDP(proto, addr)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot listen on the UDP port")
		}
		file, err = server.File()
		if err != nil {
			return nil, errors.Wrapf(err, "cannot get file for UDP socket")
		}
		// close the listener
		// note that this does not affect the fd, see the godoc for server.File()
		err = server.Close()
		if err != nil {
			logrus.Warnf("Failed to close connection: %v", err)
		}

	case "tcp":
		var (
			addr *net.TCPAddr
			err  error
		)
		if isV6 {
			addr, err = net.ResolveTCPAddr("tcp6", fmt.Sprintf("[%s]:%d", hostIP, port))
		} else {
			addr, err = net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", hostIP, port))
		}
		if err != nil {
			return nil, errors.Wrapf(err, "cannot resolve the TCP address")
		}

		proto := "tcp4"
		if isV6 {
			proto = "tcp6"
		}
		server, err := net.ListenTCP(proto, addr)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot listen on the TCP port")
		}
		file, err = server.File()
		if err != nil {
			return nil, errors.Wrapf(err, "cannot get file for TCP socket")
		}
		// close the listener
		// note that this does not affect the fd, see the godoc for server.File()
		err = server.Close()
		if err != nil {
			logrus.Warnf("Failed to close connection: %v", err)
		}

	case "sctp":
		if *sctpWarning {
			logrus.Info("Port reservation for SCTP is not supported")
			*sctpWarning = false
		}
	default:
		return nil, fmt.Errorf("unknown protocol %s", protocol)
	}
	return file, nil
}

func getOCIRuntimeError(name, runtimeMsg string) error {
	includeFullOutput := logrus.GetLevel() == logrus.DebugLevel

	if match := regexp.MustCompile("(?i).*permission denied.*|.*operation not permitted.*").FindString(runtimeMsg); match != "" {
		errStr := match
		if includeFullOutput {
			errStr = runtimeMsg
		}
		return errors.Wrapf(define.ErrOCIRuntimePermissionDenied, "%s: %s", name, strings.Trim(errStr, "\n"))
	}
	if match := regexp.MustCompile("(?i).*executable file not found in.*|.*no such file or directory.*").FindString(runtimeMsg); match != "" {
		errStr := match
		if includeFullOutput {
			errStr = runtimeMsg
		}
		return errors.Wrapf(define.ErrOCIRuntimeNotFound, "%s: %s", name, strings.Trim(errStr, "\n"))
	}
	if match := regexp.MustCompile("`/proc/[a-z0-9-].+/attr.*`").FindString(runtimeMsg); match != "" {
		errStr := match
		if includeFullOutput {
			errStr = runtimeMsg
		}
		if strings.HasSuffix(match, "/exec`") {
			return errors.Wrapf(define.ErrSetSecurityAttribute, "%s: %s", name, strings.Trim(errStr, "\n"))
		} else if strings.HasSuffix(match, "/current`") {
			return errors.Wrapf(define.ErrGetSecurityAttribute, "%s: %s", name, strings.Trim(errStr, "\n"))
		}
		return errors.Wrapf(define.ErrSecurityAttribute, "%s: %s", name, strings.Trim(errStr, "\n"))
	}
	return errors.Wrapf(define.ErrOCIRuntime, "%s: %s", name, strings.Trim(runtimeMsg, "\n"))
}
