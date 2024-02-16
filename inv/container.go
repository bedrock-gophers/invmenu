package inv

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"sync"
)

var (
	menuMu            sync.Mutex
	openedMenus       = map[block.ContainerViewer]Menu{}
	fakeContainersPos = map[byte]cube.Pos{}
)

const (
	// ContainerTypeChest is a container type for a chest.
	ContainerTypeChest byte = iota
	// ContainerTypeBarrel is a container type for a barrel.
	ContainerTypeBarrel
)

func blockFromContainerKind(t byte) world.Block {
	switch t {
	case ContainerTypeChest:
		return block.NewChest()
	case ContainerTypeBarrel:
		return block.NewBarrel()
	default:
		panic("invalid container type")
	}
}

// PlaceFakeContainer places a fake container at the position and world passed.
func PlaceFakeContainer(w *world.World, pos cube.Pos) {
	// TODO: Add support for other container types.
	kind := ContainerTypeChest

	w.SetBlock(pos, blockFromContainerKind(kind), nil)
	fakeContainersPos[kind] = pos
}

// CloseContainer closes the container that the session passed is currently viewing.
func CloseContainer(p *player.Player) {
	menuMu.Lock()
	s := player_session(p)
	m, ok := openedMenus[s]
	if ok {
		if s != session.Nop {
			if closeable, ok := m.submittable.(Closer); ok {
				closeable.Close(p)
			}
			s.ViewBlockUpdate(m.pos, block.Air{}, 0)
		}
		delete(openedMenus, s)
	}
	menuMu.Unlock()

}
