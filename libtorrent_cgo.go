// +build !android

package libtorrent

// #cgo pkg-config: --static libtorrent-rasterbar boost openssl
// #cgo CXXFLAGS: -I/opt/boost/1.60.0/include -ferror-limit=1
// #cgo darwin LDFLAGS: -lm -lstdc++
// #cgo linux LDFLAGS: -lm -lstdc++ -ldl -lrt
// #cgo windows CXXFLAGS: -DIPV6_TCLASS=39
// #cgo windows LDFLAGS: -static-libgcc -static-libstdc++
import "C"
