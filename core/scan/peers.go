package scan

import "github.com/ipfs/go-ipfs-api"

type PeerMapper struct {
	shell *shell.Shell
}

func NewPeerMapper(apiShell *shell.Shell) PeerMapper {
	return PeerMapper{shell: apiShell}
}

func (pm *PeerMapper) ScanForPeers() ([]byte, error) {
	data, err := pm.getPeerData()
	return data, err
}

func (pm *PeerMapper) getPeerData() ([]byte, error) {
	resp, err := pm.shell.DiagNet("d3")
	return resp, err
}
