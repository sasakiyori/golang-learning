1. android supporting
2. reflect
3. cgo
   - add link flags like:
   ```go
    // #cgo CFLAGS: -DPNG_DEBUG=1
    // #cgo linux CFLAGS: -DLINUX=1
    // #cgo LDFLAGS: -lpng
    // #include <png.h>
    import "C"
   ```
   - or:
   ```go
    // #cgo pkg-config: png cairo
    // #include <png.h>
    import "C"
   ```