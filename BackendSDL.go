package giu

import (
	"fmt"
	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/sdlbackend"
)

// GIUBackendSDL is an abstraction layer between cimgui-go's Backends.
type GIUBackendSDL backend.Backend[MasterWindowFlags]

var _ GIUBackendSDL = &SDLBackend{}

// SDLBackend is an implementation of sdlbackend.SDLBackend cimgui-go backend with respect to
// giu's MasterWindowFlags.
type SDLBackend struct {
	*sdlbackend.SDLBackend
}

// NewSDLBackend creates a new instance of SDLBackend.
func NewSDLBackend() *SDLBackend {
	return &SDLBackend{
		SDLBackend: sdlbackend.NewSDLBackend(),
	}
}

// SetInputMode implements backend.Backend interface.
func (b *SDLBackend) SetInputMode(mode, _ MasterWindowFlags) {
	flag := b.parseFlag(mode)
	b.SDLBackend.SetInputMode(flag.flag, sdlbackend.SDLWindowFlags(flag.value))
}

// SetSwapInterval implements backend.Backend interface.
func (b *SDLBackend) SetSwapInterval(interval MasterWindowFlags) error {
	intervalV := b.parseFlag(interval).flag
	if err := b.SDLBackend.SetSwapInterval(intervalV); err != nil {
		return fmt.Errorf("giu.GLFWBackend got error while SwapInterval: %w", err)
	}

	return nil
}

// SetWindowFlags implements backend.Backend interface.
func (b *SDLBackend) SetWindowFlags(flags MasterWindowFlags, _ int) {
	flag := b.parseFlag(flags)
	b.SDLBackend.SetWindowFlags(flag.flag, flag.value)
}

func (b *SDLBackend) parseFlag(m MasterWindowFlags) flagValue[sdlbackend.SDLWindowFlags] {
	data := map[MasterWindowFlags]flagValue[sdlbackend.SDLWindowFlags]{
		MasterWindowFlagsNotResizable: {sdlbackend.SDLWindowFlagsResizable, 0},
		MasterWindowFlagsMaximized:    {sdlbackend.SDLWindowFlagsMaximized, 1},
		MasterWindowFlagsFrameless:    {sdlbackend.SDLWindowFlagsDecorated, 0},
		MasterWindowFlagsTransparent:  {sdlbackend.SDLWindowFlagsTransparent, 1},
		MasterWindowFlagsHidden:       {sdlbackend.SDLWindowFlagsVisible, 0},
	}

	d, ok := data[m]
	Assert(ok, "SDLBackend", "parseFlag", "Unknown MasterWindowFlags")

	return d
}
