package misoplug

import (
	"github.com/fxamacker/cbor/v2"
	"github.com/itsopenmiso/openmiso/pkg/plugsdk/builder"
	"github.com/itsopenmiso/openmiso/pkg/plugsdk/deployer"
	"github.com/itsopenmiso/openmiso/pkg/plugsdk/monitor"
	"github.com/itsopenmiso/openmiso/pkg/plugsdk/releaser"
	"github.com/mjwhodur/plugkit/plug"
)

type MisoPlug struct {
	Builder  builder.Builder
	Releaser releaser.Releaser
	Deployer deployer.Deployer
	Monitor  monitor.Monitor
	host     *plug.RawStreamPlug
}

// Handle implements logic handlers for specific Plug needs - ensures proper unmarshalling of messages and ensures
// that plug has correct interface.
func (m *MisoPlug) Handle(kind string, payload cbor.RawMessage) {
	switch kind {
	case "releaser":
		if m.Releaser != nil {

		} else {
			m.plugDoesNotSupport()
		}
	}
}

func (m *MisoPlug) Mount(c *plug.RawStreamPlug) {
	m.host = c
	// FIXME: In PlugKit likely unneeded, but tests depend on that...
}

func (m *MisoPlug) CloseSignal() {
	//TODO implement me
	// FIXME: Handle plugin exit by OS (SIGINT / SIGTERM)
	panic("implement me")
}

func NewMisoPlug() *MisoPlug {
	p := &MisoPlug{}
	h := plug.NewRawStreamPlug(p)
	p.host = h // FIXME: This is ugly
	return p
}

// Run runs the plug
func (m *MisoPlug) Run() {
	m.host = plug.NewRawStreamPlug(m)
	m.host.Main()
}

func (m *MisoPlug) Done() error {
	//TODO implement me
	// This function has to shutdown nicely the plug itself. Probably the interface is bad...
	// This means that plug ended its job. (probably can be handled externally... from the plug host.)
	panic("implement me")
	return nil
}

// WithBuilder ensures plug conforms builder interface
func (m *MisoPlug) WithBuilder(b builder.Builder) *MisoPlug {
	m.Builder = b
	return m
}

// WithReleaser ensures plug conforms releaser interface
func (m *MisoPlug) WithReleaser(r releaser.Releaser) *MisoPlug {
	m.Releaser = r
	return m
}

// WithDeployer ensures plug conforms deployer interface
func (m *MisoPlug) WithDeployer(d deployer.Deployer) *MisoPlug {
	m.Deployer = d
	return m
}

// WithMonitor ensures plug conforms monitor interface for monitoring the infrastructure
func (m *MisoPlug) WithMonitor(c monitor.Monitor) *MisoPlug {
	m.Monitor = c
	return m
}

// Throws error to plug hoster that this kind of message is not supported in this scenario
func (m *MisoPlug) plugDoesNotSupport() {
	panic("implement me")
}
